// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title Skavenge
 * @dev A scavenger hunt NFT that stores encrypted clues which users can solve and trade
 */
contract Skavenge is ERC721Enumerable, ReentrancyGuard {
    // Clue structure
    struct Clue {
        bytes encryptedContents; // ElGamal encrypted content of the clue
        bytes32 solutionHash; // Hash of the solution
        bool isSolved; // Whether the clue has been solved
        uint256 salePrice; // Price in wei for which the clue is for sale
        uint256 rValue; // ElGamal encryption r value (needed for decryption)
        uint256 timeout; // Transfer timeout in seconds (set by seller)
        uint8 pointValue; // Point value of the clue (1-5)
        uint256 solveReward; // ETH reward for solving the clue
    }

    // TokenTransfer structure
    struct TokenTransfer {
        address buyer; // Address of the buyer
        uint256 tokenId; // Token ID being transferred
        uint256 value; // Value sent with the transfer
        uint256 initiatedAt; // Timestamp when transfer was initiated
        bytes proof; // DLEQ proof
        bytes32 newClueHash; // Hash of the new encrypted clue for the buyer
        bytes32 rValueHash; // Hash commitment to r value
        bool proofVerified; // Whether the proof has been verified
        uint256 proofProvidedAt; // Timestamp when proof was provided
        uint256 verifiedAt; // Timestamp when proof was verified
    }

    // Transfer timeout bounds in seconds
    // uint256 public constant MIN_TIMEOUT = 3600; // 1 hour
    uint256 public constant MIN_TIMEOUT = 0; // no minimum for testing
    uint256 public constant MAX_TIMEOUT = 86400; // 24 hours

    // Current token ID counter
    uint256 private _tokenIdCounter;

    // Address authorized to mint new clues
    address public authorizedMinter;

    // Mapping from token ID to Clue struct
    mapping(uint256 => Clue) public clues;

    // Mapping from transfer ID to Transfer struct
    mapping(bytes32 => TokenTransfer) public transfers;

    // Mapping to track which token IDs are for sale
    mapping(uint256 => bool) public cluesForSale;

    // Array to store token IDs that are for sale (for pagination)
    uint256[] private _cluesForSaleList;

    // Mapping to track tokens with active transfers (prevents concurrent purchases)
    mapping(uint256 => bool) public transferInProgress;

    // Mapping from token ID to active transfer ID (for cancelling transfers when sale is removed)
    mapping(uint256 => bytes32) public activeTransferIds;

    // Error for attempting to transfer a solved clue
    error SolvedClueTransferNotAllowed();

    // Error for attempting to buy a clue that is not for sale
    error ClueNotForSale();

    // Error for attempting to buy a clue with insufficient funds
    error InsufficientFunds();

    // Error for attempting to set a sale price on a solved clue
    error SolvedClueCannotBeSold();

    // Error for unauthorized minting
    error UnauthorizedMinter();

    // Error for unauthorized minter update
    error UnauthorizedMinterUpdate();

    // Error for attempting to purchase a clue with an active transfer
    error TransferAlreadyInProgress();

    // Error for invalid point value
    error InvalidPointValue();

    // Events
    event ClueMinted(uint256 indexed tokenId, address indexed minter, address indexed recipient);
    event ClueSolved(uint256 indexed tokenId, string solution);
    event ClueAttemptFailed(uint256 indexed tokenId, string attemptedSolution);
    event SalePriceSet(uint256 indexed tokenId, uint256 price);
    event SalePriceRemoved(uint256 indexed tokenId);

    event TransferInitiated(
        bytes32 indexed transferId,
        address indexed buyer,
        uint256 indexed tokenId
    );
    event ProofProvided(
        bytes32 indexed transferId,
        bytes proof,
        bytes32 newClueHash,
        bytes32 rValueHash
    );
    event ProofVerified(bytes32 indexed transferId);
    event TransferCompleted(bytes32 indexed transferId, uint256 rValue);
    event TransferCancelled(
        bytes32 indexed transferId,
        address indexed cancelledBy
    );

    // Event emitted when authorized minter is updated
    event AuthorizedMinterUpdated(
        address indexed oldMinter,
        address indexed newMinter
    );

    /**
     * @dev Constructor for the Skavenge contract
     * @param initialMinter Address authorized to mint new clues
     */
    constructor(address initialMinter) ERC721("Skavenge", "SKVG") {
        _tokenIdCounter = 1; // Start token IDs at 1
        authorizedMinter = initialMinter;
    }

    /**
     * @dev Get the current token ID counter
     */
    function getCurrentTokenId() external view returns (uint256) {
        return _tokenIdCounter;
    }

    /**
     * @dev Mint a new clue
     * @param encryptedContents ElGamal encrypted content of the clue
     * @param solutionHash Hash of the solution
     * @param rValue ElGamal encryption r value
     * @param pointValue Point value of the clue (1-5)
     * @param recipient Address to mint the clue to
     * @notice Can optionally send ETH to set a solve reward for the clue
     */
    function mintClue(
        bytes calldata encryptedContents,
        bytes32 solutionHash,
        uint256 rValue,
        uint8 pointValue,
        address recipient
    ) external payable returns (uint256 tokenId) {
        if (msg.sender != authorizedMinter) {
            revert UnauthorizedMinter();
        }

        // Validate point value is between 1 and 5 (inclusive)
        if (pointValue < 1 || pointValue > 5) {
            revert InvalidPointValue();
        }

        tokenId = _tokenIdCounter++;

        clues[tokenId] = Clue({
            encryptedContents: encryptedContents,
            solutionHash: solutionHash,
            isSolved: false,
            salePrice: 0,
            rValue: rValue,
            timeout: 0,
            pointValue: pointValue,
            solveReward: msg.value
        });

        _mint(recipient, tokenId);

        emit ClueMinted(tokenId, msg.sender, recipient);

        return tokenId;
    }

    /**
     * @dev Get the encrypted contents of a clue
     * @param tokenId Token ID of the clue
     */
    function getClueContents(
        uint256 tokenId
    ) external view returns (bytes memory) {
        ownerOf(tokenId); // Will revert if token doesn't exist
        return clues[tokenId].encryptedContents;
    }

    /**
     * @dev Get the r value for a clue (needed for decryption along with private key)
     * @param tokenId Token ID of the clue
     */
    function getRValue(uint256 tokenId) external view returns (uint256) {
        return clues[tokenId].rValue;
    }

    /**
     * @dev Get the point value for a clue
     * @param tokenId Token ID of the clue
     */
    function getPointValue(uint256 tokenId) external view returns (uint8) {
        ownerOf(tokenId); // Will revert if token doesn't exist
        return clues[tokenId].pointValue;
    }

    /**
     * @dev Get the solve reward for a clue
     * @param tokenId Token ID of the clue
     */
    function getSolveReward(uint256 tokenId) external view returns (uint256) {
        ownerOf(tokenId); // Will revert if token doesn't exist
        return clues[tokenId].solveReward;
    }

    /**
     * @dev Set a sale price for a clue
     * @param tokenId Token ID of the clue
     * @param price Price in wei
     * @param timeout Transfer timeout in seconds (must be between MIN_TIMEOUT and MAX_TIMEOUT)
     */
    function setSalePrice(
        uint256 tokenId,
        uint256 price,
        uint256 timeout
    ) external {
        require(ownerOf(tokenId) == msg.sender, "Not token owner");

        if (clues[tokenId].isSolved) {
            revert SolvedClueCannotBeSold();
        }

        // Validate timeout only if listing for sale (price > 0)
        if (price > 0) {
            require(
                timeout >= MIN_TIMEOUT && timeout <= MAX_TIMEOUT,
                "Invalid timeout"
            );
            clues[tokenId].timeout = timeout;
        }

        clues[tokenId].salePrice = price;

        // Add to the for-sale list if not already there
        if (!cluesForSale[tokenId] && price > 0) {
            cluesForSale[tokenId] = true;
            _cluesForSaleList.push(tokenId);
        }
        // Remove from the for-sale list if price is set to 0
        else if (cluesForSale[tokenId] && price == 0) {
            _removeFromForSaleList(tokenId);
        }

        emit SalePriceSet(tokenId, price);
    }

    /**
     * @dev Remove a clue from sale
     * @param tokenId Token ID of the clue
     */
    function removeSalePrice(uint256 tokenId) external {
        require(ownerOf(tokenId) == msg.sender, "Not token owner");

        if (cluesForSale[tokenId]) {
            _removeFromForSaleList(tokenId);
        }

        clues[tokenId].salePrice = 0;
        emit SalePriceRemoved(tokenId);
    }

    /**
     * @dev Get clues for sale with pagination
     * @param offset Starting index
     * @param limit Maximum number of items to return
     */
    function getCluesForSale(
        uint256 offset,
        uint256 limit
    )
        external
        view
        returns (
            uint256[] memory tokenIds,
            address[] memory owners,
            uint256[] memory prices,
            bool[] memory solvedStatus
        )
    {
        uint256 total = _cluesForSaleList.length;
        uint256 end = offset + limit > total ? total : offset + limit;
        uint256 resultSize = end > offset ? end - offset : 0;

        tokenIds = new uint256[](resultSize);
        owners = new address[](resultSize);
        prices = new uint256[](resultSize);
        solvedStatus = new bool[](resultSize);

        for (uint256 i = 0; i < resultSize; i++) {
            uint256 tokenId = _cluesForSaleList[offset + i];
            tokenIds[i] = tokenId;
            owners[i] = ownerOf(tokenId);
            prices[i] = clues[tokenId].salePrice;
            solvedStatus[i] = clues[tokenId].isSolved;
        }

        return (tokenIds, owners, prices, solvedStatus);
    }

    /**
     * @dev Get the total number of clues for sale
     */
    function getTotalCluesForSale() external view returns (uint256) {
        return _cluesForSaleList.length;
    }

    /**
     * @dev Attempt to solve a clue
     * @param tokenId Token ID of the clue
     * @param solution Proposed solution
     */
    function attemptSolution(
        uint256 tokenId,
        string calldata solution
    ) external nonReentrant {
        require(ownerOf(tokenId) == msg.sender, "Not token owner");
        require(!clues[tokenId].isSolved, "Clue already solved");

        if (keccak256(bytes(solution)) == clues[tokenId].solutionHash) {
            clues[tokenId].isSolved = true;

            // Remove from sale if it was for sale
            if (cluesForSale[tokenId]) {
                _removeFromForSaleList(tokenId);
                clues[tokenId].salePrice = 0;
                emit SalePriceRemoved(tokenId);
            }

            // Transfer solve reward if one exists
            uint256 reward = clues[tokenId].solveReward;
            if (reward > 0) {
                clues[tokenId].solveReward = 0; // Clear reward before transfer
                (bool sent, ) = payable(msg.sender).call{value: reward}("");
                require(sent, "Failed to send solve reward");
            }

            emit ClueSolved(tokenId, solution);
        } else {
            emit ClueAttemptFailed(tokenId, solution);
        }
    }

    /**
     * @dev Initiate purchase of a clue
     * @param tokenId Token ID of the clue
     */
    function initiatePurchase(uint256 tokenId) external payable nonReentrant {
        ownerOf(tokenId); // Verify token exists

        // Check if the clue is for sale
        if (!cluesForSale[tokenId] || clues[tokenId].salePrice == 0) {
            revert ClueNotForSale();
        }

        // Check if sent value is at least the sale price
        if (msg.value < clues[tokenId].salePrice) {
            revert InsufficientFunds();
        }

        // Check if the clue is solved
        if (clues[tokenId].isSolved) {
            revert SolvedClueTransferNotAllowed();
        }

        // Check if token already has an active transfer
        // Prevents multiple buyers from initiating concurrent purchases for same token
        if (transferInProgress[tokenId]) {
            revert TransferAlreadyInProgress();
        }

        // Check if transfer already initiated by this buyer
        bytes32 transferId = generateTransferId(msg.sender, tokenId);
        require(
            transfers[transferId].buyer == address(0),
            "Transfer already initiated"
        );

        // Mark token as having an active transfer
        transferInProgress[tokenId] = true;

        // Store the active transfer ID for this token
        activeTransferIds[tokenId] = transferId;

        // Create transfer record
        transfers[transferId] = TokenTransfer({
            buyer: msg.sender,
            tokenId: tokenId,
            value: msg.value,
            initiatedAt: block.timestamp,
            proof: new bytes(0),
            newClueHash: bytes32(0),
            rValueHash: bytes32(0),
            proofVerified: false,
            proofProvidedAt: 0,
            verifiedAt: 0
        });

        emit TransferInitiated(transferId, msg.sender, tokenId);
    }

    /**
     * @dev Generate a transfer ID
     * @param buyer Address of the buyer
     * @param tokenId Token ID being transferred
     */
    function generateTransferId(
        address buyer,
        uint256 tokenId
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(buyer, tokenId));
    }

    /**
     * @dev Provide proof for a clue transfer
     * @param transferId ID of the transfer
     * @param proof Complete TransferProof (DLEQ + Plaintext proofs)
     * @param newClueHash Hash of the new encrypted clue for the buyer
     */
    function provideProof(
        bytes32 transferId,
        bytes calldata proof,
        bytes32 newClueHash
    ) external nonReentrant {
        TokenTransfer storage transfer = transfers[transferId];
        require(transfer.buyer != address(0), "Transfer does not exist");
        require(ownerOf(transfer.tokenId) == msg.sender, "Not token owner");

        // Check if transfer has timed out
        require(
            block.timestamp - transfer.initiatedAt <=
                clues[transfer.tokenId].timeout,
            "Transfer expired"
        );

        // Extract rHash from the TransferProof structure
        // Format: len(DLEQ) | DLEQ_bytes | Plaintext_bytes(32)
        // The rHash is at the end of the DLEQ portion
        require(proof.length >= 36, "Proof too short"); // At least 4 bytes for length + 32 for rHash

        // Read DLEQ length from first 4 bytes (big-endian)
        uint32 dleqLen = (uint32(uint8(proof[0])) << 24) |
            (uint32(uint8(proof[1])) << 16) |
            (uint32(uint8(proof[2])) << 8) |
            uint32(uint8(proof[3]));

        require(proof.length >= 4 + dleqLen + 32, "Invalid proof structure");

        // rHash is the last 32 bytes of the DLEQ proof
        // Position: 4 + dleqLen - 32 to 4 + dleqLen
        bytes32 extractedRHash = bytes32(proof[4 + dleqLen - 32:4 + dleqLen]);

        transfer.proof = proof;
        transfer.newClueHash = newClueHash;
        transfer.rValueHash = extractedRHash;
        transfer.proofProvidedAt = block.timestamp;

        emit ProofProvided(transferId, proof, newClueHash, extractedRHash);
    }

    /**
     * @dev Verify proof for a clue transfer
     * @param transferId ID of the transfer
     */
    function verifyProof(bytes32 transferId) external nonReentrant {
        TokenTransfer storage transfer = transfers[transferId];
        require(transfer.buyer == msg.sender, "Not the buyer");
        require(transfer.proofProvidedAt > 0, "Proof not yet provided");

        // Check if proof provision has timed out
        require(
            block.timestamp - transfer.proofProvidedAt <=
                clues[transfer.tokenId].timeout,
            "Proof verification expired"
        );

        transfer.proofVerified = true;
        transfer.verifiedAt = block.timestamp;

        emit ProofVerified(transferId);
    }

    /**
     * @dev Complete a clue transfer with new encrypted contents and r value
     * @param transferId ID of the transfer
     * @param newEncryptedContents New encrypted contents of the clue for the buyer
     * @param rValue The r value used in ElGamal encryption
     */
    function completeTransfer(
        bytes32 transferId,
        bytes calldata newEncryptedContents,
        uint256 rValue
    ) external nonReentrant {
        TokenTransfer storage transfer = transfers[transferId];
        require(transfer.buyer != address(0), "Transfer does not exist");
        require(ownerOf(transfer.tokenId) == msg.sender, "Not token owner");

        // Check if verification has timed out
        require(transfer.verifiedAt > 0, "Proof not verified");
        require(
            block.timestamp - transfer.verifiedAt <=
                clues[transfer.tokenId].timeout,
            "Transfer completion expired"
        );

        // Verify the hash of the new encrypted contents
        require(
            keccak256(newEncryptedContents) == transfer.newClueHash,
            "Content hash mismatch"
        );

        // Verify r value matches commitment
        require(
            keccak256(abi.encodePacked(rValue)) == transfer.rValueHash,
            "R value hash mismatch"
        );

        // Note: We don't verify g^r == C1 on-chain because:
        // 1. The ecmul precompile (0x07) only supports alt_bn128, not secp256k1
        // 2. The hash commitment already prevents r value tampering
        // 3. The buyer verified the DLEQ proof off-chain before calling verifyProof()
        // 4. Implementing secp256k1 multiplication in Solidity is prohibitively expensive

        // Check if the clue is solved
        if (clues[transfer.tokenId].isSolved) {
            revert SolvedClueTransferNotAllowed();
        }

        // Update the clue contents and r value
        clues[transfer.tokenId].encryptedContents = newEncryptedContents;
        clues[transfer.tokenId].rValue = rValue;

        // Store values we'll need after clearing transfer state
        address buyer = transfer.buyer;
        uint256 value = transfer.value;
        uint256 tokenId = transfer.tokenId;

        // Remove from sale list BEFORE transfer (must happen before _safeTransfer)
        if (cluesForSale[tokenId]) {
            cluesForSale[tokenId] = false;
            // Find and remove the token ID from the array
            for (uint256 i = 0; i < _cluesForSaleList.length; i++) {
                if (_cluesForSaleList[i] == tokenId) {
                    // Replace the item with the last item in the array
                    _cluesForSaleList[i] = _cluesForSaleList[
                        _cluesForSaleList.length - 1
                    ];
                    // Remove the last item
                    _cluesForSaleList.pop();
                    break;
                }
            }
            clues[tokenId].salePrice = 0;
            emit SalePriceRemoved(tokenId);
        }

        // Clear the transfer in progress flag
        transferInProgress[tokenId] = false;

        // Clear the active transfer ID
        delete activeTransferIds[tokenId];

        // Clear the transfer BEFORE _safeTransfer
        delete transfers[transferId];

        // Transfer ownership
        _safeTransfer(msg.sender, buyer, tokenId, "");

        // Send payment to seller
        (bool sent, ) = payable(msg.sender).call{value: value}("");
        require(sent, "Failed to send Ether");

        emit TransferCompleted(transferId, rValue);
    }

    /**
     * @dev Cancel a clue transfer
     * @param transferId ID of the transfer
     */
    function cancelTransfer(bytes32 transferId) external nonReentrant {
        TokenTransfer storage transfer = transfers[transferId];
        require(transfer.buyer != address(0), "Transfer does not exist");

        bool isBuyer = transfer.buyer == msg.sender;
        bool isSeller = ownerOf(transfer.tokenId) == msg.sender;

        // Check cancellation conditions
        bool canCancel = false;

        // Buyer can cancel if:
        // 1. Timeout elapsed since initiation and waiting for seller to provide proof
        // 2. Seller has provided proof (can cancel anytime, regardless of timeout)
        // 3. Timeout elapsed since verification and waiting for seller to complete
        if (isBuyer) {
            // State 1: Waiting for seller to provide proof
            if (
                transfer.proofProvidedAt == 0 &&
                block.timestamp - transfer.initiatedAt >
                clues[transfer.tokenId].timeout
            ) {
                canCancel = true;
            }
            // State 2: Seller has provided proof (buyer can cancel anytime)
            else if (transfer.proofProvidedAt > 0 && !transfer.proofVerified) {
                canCancel = true;
            }
            // State 3: Proof verified, waiting for seller to complete
            else if (
                transfer.verifiedAt > 0 &&
                block.timestamp - transfer.verifiedAt >
                clues[transfer.tokenId].timeout
            ) {
                canCancel = true;
            }
        }

        // Seller can cancel if:
        // Proof provided but not verified and timeout elapsed (waiting for buyer)
        if (isSeller) {
            if (
                transfer.proofProvidedAt > 0 &&
                !transfer.proofVerified &&
                block.timestamp - transfer.proofProvidedAt >
                clues[transfer.tokenId].timeout
            ) {
                canCancel = true;
            }
        }

        require(canCancel, "Not authorized to cancel");

        // Refund the buyer
        if (transfer.value > 0) {
            (bool sent, ) = payable(transfer.buyer).call{value: transfer.value}(
                ""
            );
            require(sent, "Failed to send Ether");
        }

        // Clear the transfer in progress flag
        transferInProgress[transfer.tokenId] = false;

        // Clear the active transfer ID
        delete activeTransferIds[transfer.tokenId];

        emit TransferCancelled(transferId, msg.sender);

        // Clear the transfer
        delete transfers[transferId];
    }

    /**
     * @dev Remove a token ID from the for-sale list
     * @param tokenId Token ID to remove
     */
    function _removeFromForSaleList(uint256 tokenId) private {
        if (!cluesForSale[tokenId]) {
            return;
        }

        // Cancel any pending transfer for this token
        bytes32 transferId = activeTransferIds[tokenId];
        if (transferId != bytes32(0)) {
            TokenTransfer storage transfer = transfers[transferId];
            if (transfer.buyer != address(0)) {
                // Refund the buyer
                if (transfer.value > 0) {
                    (bool sent, ) = payable(transfer.buyer).call{
                        value: transfer.value
                    }("");
                    require(sent, "Failed to send Ether");
                }

                // Clear the transfer in progress flag
                transferInProgress[tokenId] = false;

                // Clear the active transfer ID
                delete activeTransferIds[tokenId];

                emit TransferCancelled(transferId, msg.sender);

                // Clear the transfer
                delete transfers[transferId];
            }
        }

        cluesForSale[tokenId] = false;

        // Find and remove the token ID from the array
        for (uint256 i = 0; i < _cluesForSaleList.length; i++) {
            if (_cluesForSaleList[i] == tokenId) {
                // Replace the item with the last item in the array
                _cluesForSaleList[i] = _cluesForSaleList[
                    _cluesForSaleList.length - 1
                ];
                // Remove the last item
                _cluesForSaleList.pop();
                break;
            }
        }
    }

    /**
     * @dev Update the authorized minter address
     * @param newMinter New address authorized to mint clues
     */
    function updateAuthorizedMinter(address newMinter) external {
        if (msg.sender != authorizedMinter) {
            revert UnauthorizedMinterUpdate();
        }
        address oldMinter = authorizedMinter;
        authorizedMinter = newMinter;
        emit AuthorizedMinterUpdated(oldMinter, newMinter);
    }

    /**
     * @dev Override _update to handle transfers and maintain enumerable functionality
     */
    function _update(
        address to,
        uint256 tokenId,
        address auth
    ) internal virtual override returns (address) {
        address from = super._update(to, tokenId, auth);

        // If this is a transfer (not a mint or burn), remove from sale list
        if (from != address(0) && to != address(0)) {
            if (cluesForSale[tokenId]) {
                // Remove from sale list and cancel any pending transfer
                // Note: During completeTransfer, the transfer is already deleted
                // before _safeTransfer is called, so this won't double-refund
                _removeFromForSaleList(tokenId);
            }
            clues[tokenId].salePrice = 0;
        }

        return from;
    }
}
