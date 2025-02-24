// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title Skavenge
 * @dev Implementation of a scavenger hunt game using ERC721 tokens.
 * Each token represents a clue that can be solved and traded.
 */
contract Skavenge is ERC721, ReentrancyGuard {
    // Constants
    uint256 public constant TRANSFER_TIMEOUT = 3 minutes;
    uint256 public constant MAX_SOLVE_ATTEMPTS = 3;

    // Simple counter to replace Counters.sol
    uint256 private _currentTokenId;

    // Structs
    struct Clue {
        bytes encryptedContents; // Encrypted clue content, readable by owner
        bytes32 solutionHash; // Hash of the correct solution
        bool isSolved; // Whether the clue has been solved
        uint256 solveAttempts; // Number of attempts made to solve
    }

    // Renamed from Transfer to TokenTransfer to avoid conflict with ERC721 event
    struct TokenTransfer {
        address buyer; // Address of the potential buyer
        uint256 tokenId; // ID of the token being transferred
        uint256 value; // Amount of ETH offered
        uint256 initiatedAt; // Timestamp when transfer was initiated
        bytes proof; // Proof provided by seller
        bytes32 newClueHash; // Hash of the new encrypted clue
        bool proofVerified; // Whether buyer has verified the proof
        uint256 proofProvidedAt; // Timestamp when proof was provided
        uint256 verifiedAt; // Timestamp when proof was verified
    }

    // State variables
    mapping(uint256 => Clue) public clues;
    mapping(bytes32 => TokenTransfer) public transfers; // transferId => TokenTransfer

    // Events
    event TransferInitiated(
        bytes32 indexed transferId,
        address indexed buyer,
        uint256 indexed tokenId
    );
    event ProofProvided(
        bytes32 indexed transferId,
        bytes proof,
        bytes32 newClueHash
    );
    event ProofVerified(bytes32 indexed transferId);
    event TransferCompleted(bytes32 indexed transferId);
    event TransferCancelled(bytes32 indexed transferId);
    event ClueSolved(uint256 indexed tokenId);
    event ClueAttempted(uint256 indexed tokenId, uint256 remainingAttempts);

    // Custom error to prevent solved clues from being transferred
    error SolvedClueTransferNotAllowed();

    constructor() ERC721("Skavenge", "SKVG") {}

    /**
     * @dev Generates a unique transfer ID from buyer address and token ID
     * @param buyer Address of the buyer
     * @param tokenId ID of the token being transferred
     */
    function generateTransferId(
        address buyer,
        uint256 tokenId
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(buyer, tokenId));
    }

    /**
     * @dev Initiates the purchase of a clue token
     * @param tokenId The ID of the token to purchase
     */
    function initiatePurchase(uint256 tokenId) external payable nonReentrant {
        // Use ownerOf to check if token exists instead of _exists
        // This will revert if token doesn't exist
        ownerOf(tokenId);

        // Check if the clue is solved
        if (clues[tokenId].isSolved) {
            revert SolvedClueTransferNotAllowed();
        }

        bytes32 transferId = generateTransferId(msg.sender, tokenId);
        require(
            transfers[transferId].buyer == address(0),
            "Transfer already initiated"
        );

        transfers[transferId] = TokenTransfer({
            buyer: msg.sender,
            tokenId: tokenId,
            value: msg.value,
            initiatedAt: block.timestamp,
            proof: "",
            newClueHash: bytes32(0),
            proofVerified: false,
            proofProvidedAt: 0,
            verifiedAt: 0
        });

        emit TransferInitiated(transferId, msg.sender, tokenId);
    }

    /**
     * @dev Called by the seller to provide proof and the hash of the new encrypted clue
     * @param transferId The ID of the transfer
     * @param proof The proof data
     * @param newClueHash Hash of the newly encrypted clue content
     */
    function provideProof(
        bytes32 transferId,
        bytes calldata proof,
        bytes32 newClueHash
    ) external {
        TokenTransfer storage t = transfers[transferId];
        require(t.buyer != address(0), "Transfer does not exist");
        require(
            block.timestamp <= t.initiatedAt + TRANSFER_TIMEOUT,
            "Transfer expired"
        );
        require(ownerOf(t.tokenId) == msg.sender, "Not token owner");

        t.proof = proof;
        t.newClueHash = newClueHash;
        t.proofProvidedAt = block.timestamp;

        emit ProofProvided(transferId, proof, newClueHash);
    }

    /**
     * @dev Called by the buyer to verify they accept the proof
     * @param transferId The ID of the transfer
     */
    function verifyProof(bytes32 transferId) external {
        TokenTransfer storage t = transfers[transferId];
        require(t.buyer == msg.sender, "Not the buyer");
        require(t.proofProvidedAt > 0, "Proof not yet provided");
        require(
            block.timestamp <= t.proofProvidedAt + TRANSFER_TIMEOUT,
            "Proof verification expired"
        );

        t.proofVerified = true;
        t.verifiedAt = block.timestamp;

        emit ProofVerified(transferId);
    }

    /**
     * @dev Completes the transfer process
     * @param transferId The ID of the transfer
     * @param newEncryptedContents The newly encrypted clue contents
     */
    function completeTransfer(
        bytes32 transferId,
        bytes calldata newEncryptedContents
    ) external nonReentrant {
        TokenTransfer storage t = transfers[transferId];
        require(ownerOf(t.tokenId) == msg.sender, "Not token owner");
        require(t.proofVerified, "Proof not verified");
        require(
            block.timestamp <= t.verifiedAt + TRANSFER_TIMEOUT,
            "Transfer completion expired"
        );

        // Verify the new encrypted contents match the previously provided hash
        require(
            keccak256(newEncryptedContents) == t.newClueHash,
            "Content hash mismatch"
        );

        // Check if the clue is solved
        if (clues[t.tokenId].isSolved) {
            revert SolvedClueTransferNotAllowed();
        }

        // Update clue contents and reset solve attempts
        Clue storage clue = clues[t.tokenId];
        clue.encryptedContents = newEncryptedContents;
        clue.solveAttempts = 0;

        // Transfer token and payment
        _transfer(msg.sender, t.buyer, t.tokenId);
        payable(msg.sender).transfer(t.value);

        // Cleanup transfer data
        emit TransferCompleted(transferId);
        delete transfers[transferId];
    }

    /**
     * @dev Cancels a stalled transfer
     * @param transferId The ID of the transfer to cancel
     */
    function cancelTransfer(bytes32 transferId) external {
        TokenTransfer storage t = transfers[transferId];
        require(t.buyer != address(0), "Transfer does not exist");

        bool isExpired = (t.proofProvidedAt == 0 &&
            block.timestamp > t.initiatedAt + TRANSFER_TIMEOUT) ||
            (t.verifiedAt == 0 &&
                t.proofProvidedAt > 0 &&
                block.timestamp > t.proofProvidedAt + TRANSFER_TIMEOUT) ||
            (t.verifiedAt > 0 &&
                block.timestamp > t.verifiedAt + TRANSFER_TIMEOUT);

        require(isExpired, "Transfer not expired");

        // Refund buyer
        if (t.value > 0) {
            payable(t.buyer).transfer(t.value);
        }

        emit TransferCancelled(transferId);
        delete transfers[transferId];
    }

    /**
     * @dev Attempts to solve a clue
     * @param tokenId The ID of the clue token
     * @param solution The proposed solution
     */
    function attemptSolution(
        uint256 tokenId,
        string calldata solution
    ) external {
        require(ownerOf(tokenId) == msg.sender, "Not token owner");
        require(!clues[tokenId].isSolved, "Clue already solved");
        require(
            clues[tokenId].solveAttempts < MAX_SOLVE_ATTEMPTS,
            "No attempts remaining"
        );

        clues[tokenId].solveAttempts++;
        emit ClueAttempted(
            tokenId,
            MAX_SOLVE_ATTEMPTS - clues[tokenId].solveAttempts
        );

        if (
            keccak256(abi.encodePacked(solution)) == clues[tokenId].solutionHash
        ) {
            clues[tokenId].isSolved = true;
            emit ClueSolved(tokenId);
        }
    }

    /**
     * @dev Creates a new clue token
     * @param encryptedContents The encrypted contents of the clue
     * @param solutionHash The hash of the correct solution
     */
    function mintClue(
        bytes calldata encryptedContents,
        bytes32 solutionHash
    ) external {
        // Increment token ID (replacing Counters functionality)
        _currentTokenId += 1;
        uint256 newTokenId = _currentTokenId;

        clues[newTokenId] = Clue({
            encryptedContents: encryptedContents,
            solutionHash: solutionHash,
            isSolved: false,
            solveAttempts: 0
        });

        _mint(msg.sender, newTokenId);
    }

    /**
     * @dev Returns the encrypted contents of a clue
     * @param tokenId The ID of the clue token
     */
    function getClueContents(
        uint256 tokenId
    ) external view returns (bytes memory) {
        // Using ownerOf instead of _exists
        // This will revert if the token doesn't exist
        ownerOf(tokenId); // Will revert if token doesn't exist
        return clues[tokenId].encryptedContents;
    }

    /**
     * @dev Hook that is called before any token transfer
     * This includes minting and burning.
     */
    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 tokenId
    ) internal virtual {
        // Check if the token exists (will fail for non-existent tokens)
        if (from != address(0) && to != address(0)) {
            // Only check for transfers (not minting or burning)
            if (clues[tokenId].isSolved) {
                revert SolvedClueTransferNotAllowed();
            }
        }
    }

    /**
     * @dev Returns the current token ID
     */
    function getCurrentTokenId() external view returns (uint256) {
        return _currentTokenId;
    }
}
