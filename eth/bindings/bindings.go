// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SkavengeMetaData contains all meta data concerning the Skavenge contract.
var SkavengeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialMinter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ClueNotForSale\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPointValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolvedClueCannotBeSold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolvedClueTransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferAlreadyInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedMinterUpdate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMinter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"AuthorizedMinterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"attemptedSolution\",\"type\":\"string\"}],\"name\":\"ClueAttemptFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"ClueMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"solution\",\"type\":\"string\"}],\"name\":\"ClueSolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rValueHash\",\"type\":\"bytes32\"}],\"name\":\"ProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"ProofVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SalePriceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SalePriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cancelledBy\",\"type\":\"address\"}],\"name\":\"TransferCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"name\":\"TransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TransferInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_TIMEOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_TIMEOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeTransferIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"solution\",\"type\":\"string\"}],\"name\":\"attemptSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorizedMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"cancelTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"clues\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"solutionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isSolved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"pointValue\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"solveReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cluesForSale\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"newEncryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"name\":\"completeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"generateTransferId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getClueContents\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getCluesForSale\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"solvedStatus\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getPointValue\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getSolveReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalCluesForSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"initiatePurchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"solutionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"pointValue\",\"type\":\"uint8\"}],\"name\":\"mintClue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"}],\"name\":\"provideProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"setSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initiatedAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rValueHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"proofVerified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"proofProvidedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"updateAuthorizedMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516170b83803806170b883398181016040528101906100319190610172565b6040518060400160405280600881526020017f536b6176656e67650000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f534b564700000000000000000000000000000000000000000000000000000000815250815f90816100ab91906103da565b5080600190816100bb91906103da565b5050506001600a819055506001600b8190555080600c5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506104a9565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61014182610118565b9050919050565b61015181610137565b811461015b575f5ffd5b50565b5f8151905061016c81610148565b92915050565b5f6020828403121561018757610186610114565b5b5f6101948482850161015e565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061021857607f821691505b60208210810361022b5761022a6101d4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261028d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610252565b6102978683610252565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102db6102d66102d1846102af565b6102b8565b6102af565b9050919050565b5f819050919050565b6102f4836102c1565b610308610300826102e2565b84845461025e565b825550505050565b5f5f905090565b61031f610310565b61032a8184846102eb565b505050565b5b8181101561034d576103425f82610317565b600181019050610330565b5050565b601f8211156103925761036381610231565b61036c84610243565b8101602085101561037b578190505b61038f61038785610243565b83018261032f565b50505b505050565b5f82821c905092915050565b5f6103b25f1984600802610397565b1980831691505092915050565b5f6103ca83836103a3565b9150826002028217905092915050565b6103e38261019d565b67ffffffffffffffff8111156103fc576103fb6101a7565b5b6104068254610201565b610411828285610351565b5f60209050601f831160018114610442575f8415610430578287015190505b61043a85826103bf565b8655506104a1565b601f19841661045086610231565b5f5b8281101561047757848901518255600182019150602085019450602081019050610452565b868310156104945784890151610490601f8916826103a3565b8355505b6001600288020188555050505b505050505050565b616c02806104b65f395ff3fe608060405260043610610250575f3560e01c806374b19a0711610138578063c2d554ae116100b5578063de38eb3a11610079578063de38eb3a1461094e578063e985e9c514610978578063eb927a83146109b4578063f12b72ba146109f0578063f8f5a54414610a2f578063fae5380c14610a5757610250565b8063c2d554ae14610862578063c87b56dd1461088a578063d32d5790146108c6578063d4098ca814610902578063dd142be01461093257610250565b8063a6cd5ff5116100fc578063a6cd5ff514610786578063aff202b4146107c2578063b142b4ec146107ea578063b329bf5c14610812578063b88d4fde1461083a57610250565b806374b19a07146106a657806379096ee8146106d05780638d7cf3e41461070c57806395d89b4114610734578063a22cb4651461075e57610250565b806334499fff116101d1578063543ad1df11610195578063543ad1df1461057657806356189236146105a05780636352211e146105ca5780636c39cc341461060657806370a082311461062e57806373c2c7e11461066a57610250565b806334499fff146104555780633c64f04b1461049157806342842e0e146104d6578063437fdc33146104fe5780634f6ccce71461053a57610250565b80631ba538cd116102185780631ba538cd1461034857806323b872dd146103725780632f745c591461039a57806330f37c7f146103d65780633427ee941461041957610250565b806301ffc9a71461025457806306fdde0314610290578063081812fc146102ba578063095ea7b3146102f657806318160ddd1461031e575b5f5ffd5b34801561025f575f5ffd5b5061027a60048036038101906102759190614e38565b610a7f565b6040516102879190614e7d565b60405180910390f35b34801561029b575f5ffd5b506102a4610af8565b6040516102b19190614f06565b60405180910390f35b3480156102c5575f5ffd5b506102e060048036038101906102db9190614f59565b610b87565b6040516102ed9190614fc3565b60405180910390f35b348015610301575f5ffd5b5061031c60048036038101906103179190615006565b610ba2565b005b348015610329575f5ffd5b50610332610bb8565b60405161033f9190615053565b60405180910390f35b348015610353575f5ffd5b5061035c610bc4565b6040516103699190614fc3565b60405180910390f35b34801561037d575f5ffd5b506103986004803603810190610393919061506c565b610be9565b005b3480156103a5575f5ffd5b506103c060048036038101906103bb9190615006565b610ce8565b6040516103cd9190615053565b60405180910390f35b3480156103e1575f5ffd5b506103fc60048036038101906103f79190614f59565b610d8c565b604051610410989796959493929190615141565b60405180910390f35b348015610424575f5ffd5b5061043f600480360381019061043a9190614f59565b610e6e565b60405161044c9190614e7d565b60405180910390f35b348015610460575f5ffd5b5061047b60048036038101906104769190614f59565b610e8b565b6040516104889190614e7d565b60405180910390f35b34801561049c575f5ffd5b506104b760048036038101906104b291906151ee565b610ea8565b6040516104cd9a99989796959493929190615219565b60405180910390f35b3480156104e1575f5ffd5b506104fc60048036038101906104f7919061506c565b610fa9565b005b348015610509575f5ffd5b50610524600480360381019061051f9190614f59565b610fc8565b60405161053191906152ba565b60405180910390f35b348015610545575f5ffd5b50610560600480360381019061055b9190614f59565b610ffb565b60405161056d9190615053565b60405180910390f35b348015610581575f5ffd5b5061058a61106d565b6040516105979190615053565b60405180910390f35b3480156105ab575f5ffd5b506105b4611071565b6040516105c19190615053565b60405180910390f35b3480156105d5575f5ffd5b506105f060048036038101906105eb9190614f59565b61107a565b6040516105fd9190614fc3565b60405180910390f35b348015610611575f5ffd5b5061062c600480360381019061062791906152d3565b61108b565b005b348015610639575f5ffd5b50610654600480360381019061064f9190615323565b6112df565b6040516106619190615053565b60405180910390f35b348015610675575f5ffd5b50610690600480360381019061068b9190614f59565b611395565b60405161069d9190615053565b60405180910390f35b3480156106b1575f5ffd5b506106ba6113bc565b6040516106c79190615053565b60405180910390f35b3480156106db575f5ffd5b506106f660048036038101906106f19190614f59565b6113c8565b604051610703919061534e565b60405180910390f35b348015610717575f5ffd5b50610732600480360381019061072d9190614f59565b6113dd565b005b34801561073f575f5ffd5b506107486114c9565b6040516107559190614f06565b60405180910390f35b348015610769575f5ffd5b50610784600480360381019061077f9190615391565b611559565b005b348015610791575f5ffd5b506107ac60048036038101906107a79190615006565b61156f565b6040516107b9919061534e565b60405180910390f35b3480156107cd575f5ffd5b506107e860048036038101906107e39190615430565b6115a1565b005b3480156107f5575f5ffd5b50610810600480360381019061080b91906151ee565b6118bc565b005b34801561081d575f5ffd5b50610838600480360381019061083391906151ee565b611a76565b005b348015610845575f5ffd5b50610860600480360381019061085b91906155b5565b611f1d565b005b34801561086d575f5ffd5b506108886004803603810190610883919061568a565b611f42565b005b348015610895575f5ffd5b506108b060048036038101906108ab9190614f59565b6122fe565b6040516108bd9190614f06565b60405180910390f35b3480156108d1575f5ffd5b506108ec60048036038101906108e79190614f59565b612364565b6040516108f99190615053565b60405180910390f35b61091c60048036038101906109179190615725565b612381565b6040516109299190615053565b60405180910390f35b61094c60048036038101906109479190614f59565b6125d2565b005b348015610959575f5ffd5b50610962612a13565b60405161096f9190615053565b60405180910390f35b348015610983575f5ffd5b5061099e600480360381019061099991906157a9565b612a1a565b6040516109ab9190614e7d565b60405180910390f35b3480156109bf575f5ffd5b506109da60048036038101906109d59190614f59565b612aa8565b6040516109e791906157e7565b60405180910390f35b3480156109fb575f5ffd5b50610a166004803603810190610a119190615807565b612b55565b604051610a269493929190615a6a565b60405180910390f35b348015610a3a575f5ffd5b50610a556004803603810190610a509190615323565b612e18565b005b348015610a62575f5ffd5b50610a7d6004803603810190610a789190615ac9565b612f61565b005b5f7f780e9d63000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610af15750610af0826135f8565b5b9050919050565b60605f8054610b0690615b67565b80601f0160208091040260200160405190810160405280929190818152602001828054610b3290615b67565b8015610b7d5780601f10610b5457610100808354040283529160200191610b7d565b820191905f5260205f20905b815481529060010190602001808311610b6057829003601f168201915b5050505050905090565b5f610b91826136d9565b50610b9b8261375f565b9050919050565b610bb48282610baf613798565b61379f565b5050565b5f600880549050905090565b600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c59575f6040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610c509190614fc3565b60405180910390fd5b5f610c6c8383610c67613798565b6137b1565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610ce2578382826040517f64283d7b000000000000000000000000000000000000000000000000000000008152600401610cd993929190615b97565b60405180910390fd5b50505050565b5f610cf2836112df565b8210610d375782826040517fa57d13dc000000000000000000000000000000000000000000000000000000008152600401610d2e929190615bcc565b60405180910390fd5b60065f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f2054905092915050565b600d602052805f5260405f205f91509050805f018054610dab90615b67565b80601f0160208091040260200160405190810160405280929190818152602001828054610dd790615b67565b8015610e225780601f10610df957610100808354040283529160200191610e22565b820191905f5260205f20905b815481529060010190602001808311610e0557829003601f168201915b505050505090806001015490806002015f9054906101000a900460ff1690806003015490806004015490806005015490806006015f9054906101000a900460ff16908060070154905088565b600f602052805f5260405f205f915054906101000a900460ff1681565b6011602052805f5260405f205f915054906101000a900460ff1681565b600e602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002015490806003015490806004018054610efe90615b67565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2a90615b67565b8015610f755780601f10610f4c57610100808354040283529160200191610f75565b820191905f5260205f20905b815481529060010190602001808311610f5857829003601f168201915b505050505090806005015490806006015490806007015f9054906101000a900460ff1690806008015490806009015490508a565b610fc383838360405180602001604052805f815250611f1d565b505050565b5f610fd28261107a565b50600d5f8381526020019081526020015f206006015f9054906101000a900460ff169050919050565b5f611004610bb8565b8210611049575f826040517fa57d13dc000000000000000000000000000000000000000000000000000000008152600401611040929190615bcc565b60405180910390fd5b6008828154811061105d5761105c615bf3565b5b905f5260205f2001549050919050565b5f81565b5f600b54905090565b5f611084826136d9565b9050919050565b3373ffffffffffffffffffffffffffffffffffffffff166110ab8461107a565b73ffffffffffffffffffffffffffffffffffffffff1614611101576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110f890615c6a565b60405180910390fd5b600d5f8481526020019081526020015f206002015f9054906101000a900460ff1615611159576040517fff1e4dda00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8211156111cd575f81101580156111745750620151808111155b6111b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111aa90615cd2565b60405180910390fd5b80600d5f8581526020019081526020015f20600501819055505b81600d5f8581526020019081526020015f2060030181905550600f5f8481526020019081526020015f205f9054906101000a900460ff1615801561121057505f82115b15611269576001600f5f8581526020019081526020015f205f6101000a81548160ff021916908315150217905550601083908060018154018082558091505060019003905f5260205f20015f90919091909150556112a2565b600f5f8481526020019081526020015f205f9054906101000a900460ff16801561129257505f82145b156112a1576112a08361387f565b5b5b827fe23ea816dce6d7f5c0b85cbd597e7c3b97b2453791152c0b94e5e5c5f314d2f0836040516112d29190615053565b60405180910390a2505050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611350575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016113479190614fc3565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f61139f8261107a565b50600d5f8381526020019081526020015f20600701549050919050565b5f601080549050905090565b6012602052805f5260405f205f915090505481565b3373ffffffffffffffffffffffffffffffffffffffff166113fd8261107a565b73ffffffffffffffffffffffffffffffffffffffff1614611453576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161144a90615c6a565b60405180910390fd5b600f5f8281526020019081526020015f205f9054906101000a900460ff16156114805761147f8161387f565b5b5f600d5f8381526020019081526020015f2060030181905550807f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a250565b6060600180546114d890615b67565b80601f016020809104026020016040519081016040528092919081815260200182805461150490615b67565b801561154f5780601f106115265761010080835404028352916020019161154f565b820191905f5260205f20905b81548152906001019060200180831161153257829003601f168201915b5050505050905090565b61156b611564613798565b8383613bfc565b5050565b5f8282604051602001611583929190615d55565b60405160208183030381529060405280519060200120905092915050565b6115a9613d65565b3373ffffffffffffffffffffffffffffffffffffffff166115c98461107a565b73ffffffffffffffffffffffffffffffffffffffff161461161f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161161690615c6a565b60405180910390fd5b600d5f8481526020019081526020015f206002015f9054906101000a900460ff1615611680576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161167790615dca565b60405180910390fd5b600d5f8481526020019081526020015f206001015482826040516116a5929190615e16565b604051809103902003611874576001600d5f8581526020019081526020015f206002015f6101000a81548160ff021916908315150217905550600f5f8481526020019081526020015f205f9054906101000a900460ff16156117515761170a8361387f565b5f600d5f8581526020019081526020015f2060030181905550827f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a25b5f600d5f8581526020019081526020015f206007015490505f811115611834575f600d5f8681526020019081526020015f20600701819055505f3373ffffffffffffffffffffffffffffffffffffffff16826040516117af90615e51565b5f6040518083038185875af1925050503d805f81146117e9576040519150601f19603f3d011682016040523d82523d5f602084013e6117ee565b606091505b5050905080611832576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161182990615eaf565b60405180910390fd5b505b837f3138eb607d845be3efb1a7ea147da7816c8a05f683313c459e6bf953ea4f988e8484604051611866929190615ef9565b60405180910390a2506118af565b827f65cc0e9121123eab4b9d9814a9160e5954b2f7ce53d78b9cdbdd055af308b9f583836040516118a6929190615ef9565b60405180910390a25b6118b7613dab565b505050565b6118c4613d65565b5f600e5f8381526020019081526020015f2090503373ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611968576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161195f90615f65565b60405180910390fd5b5f8160080154116119ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119a590615fcd565b60405180910390fd5b600d5f826001015481526020019081526020015f20600501548160080154426119d79190616018565b1115611a18576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a0f90616095565b60405180910390fd5b6001816007015f6101000a81548160ff021916908315150217905550428160090181905550817f543093db8d78fd8619586d3a0be12a5736836393feede0888f262888c81ce4c360405160405180910390a250611a73613dab565b50565b611a7e613d65565b5f600e5f8381526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611b22576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b19906160fd565b60405180910390fd5b5f3373ffffffffffffffffffffffffffffffffffffffff16825f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161490505f3373ffffffffffffffffffffffffffffffffffffffff16611b9b846001015461107a565b73ffffffffffffffffffffffffffffffffffffffff161490505f5f90508215611c7a575f8460080154148015611bf75750600d5f856001015481526020019081526020015f2060050154846003015442611bf59190616018565b115b15611c055760019050611c79565b5f8460080154118015611c265750836007015f9054906101000a900460ff16155b15611c345760019050611c78565b5f8460090154118015611c6d5750600d5f856001015481526020019081526020015f2060050154846009015442611c6b9190616018565b115b15611c7757600190505b5b5b5b8115611cde575f8460080154118015611ca15750836007015f9054906101000a900460ff16155b8015611cd35750600d5f856001015481526020019081526020015f2060050154846008015442611cd19190616018565b115b15611cdd57600190505b5b80611d1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d1590616165565b60405180910390fd5b5f84600201541115611dfa575f845f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168560020154604051611d7590615e51565b5f6040518083038185875af1925050503d805f8114611daf576040519150601f19603f3d011682016040523d82523d5f602084013e611db4565b606091505b5050905080611df8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611def906161cd565b60405180910390fd5b505b5f60115f866001015481526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f856001015481526020019081526020015f205f90553373ffffffffffffffffffffffffffffffffffffffff16857f1ed784ea0b4551753ccb1bbf1711421d8a07aff605d39bb9d770c25943aea48560405160405180910390a3600e5f8681526020019081526020015f205f5f82015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f611edd9190614d7a565b600582015f9055600682015f9055600782015f6101000a81549060ff0219169055600882015f9055600982015f9055505050505050611f1a613dab565b50565b611f28848484610be9565b611f3c611f33613798565b85858585613db5565b50505050565b611f4a613d65565b5f600e5f8681526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611fee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fe5906160fd565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16612012826001015461107a565b73ffffffffffffffffffffffffffffffffffffffff1614612068576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161205f90615c6a565b60405180910390fd5b600d5f826001015481526020019081526020015f20600501548160030154426120919190616018565b11156120d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120c990616235565b60405180910390fd5b6024848490501015612119576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121109061629d565b60405180910390fd5b5f8484600381811061212e5761212d615bf3565b5b9050013560f81c60f81b60f81c60ff1660088686600281811061215457612153615bf3565b5b9050013560f81c60f81b60f81c60ff1663ffffffff16901b60108787600181811061218257612181615bf3565b5b9050013560f81c60f81b60f81c60ff1663ffffffff16901b601888885f8181106121af576121ae615bf3565b5b9050013560f81c60f81b60f81c60ff1663ffffffff16901b171717905060208160046121db91906162ca565b6121e591906162ca565b63ffffffff16858590501015612230576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122279061634b565b60405180910390fd5b5f8585602084600461224291906162ca565b61224c9190616369565b63ffffffff169084600461226091906162ca565b63ffffffff1692612273939291906163a8565b9061227e91906163f8565b905085858460040191826122939291906165ea565b50838360050181905550808360060181905550428360080181905550867f319414a72bfc3d93a989d08f1055fd74a1b953a652be46d0dff852ac157c12f2878787856040516122e594939291906166e3565b60405180910390a25050506122f8613dab565b50505050565b6060612309826136d9565b505f612313613f61565b90505f8151116123315760405180602001604052805f81525061235c565b8061233b84613f77565b60405160200161234c92919061675b565b6040516020818303038152906040525b915050919050565b5f600d5f8381526020019081526020015f20600401549050919050565b5f600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612408576040517f955c501b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018260ff16108061241d575060058260ff16115b15612454576040517ff5aab1a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b5f8154809291906124669061677e565b91905055905060405180610100016040528087878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018581526020015f151581526020015f81526020018481526020015f81526020018360ff16815260200134815250600d5f8381526020019081526020015f205f820151815f01908161251291906167c5565b50602082015181600101556040820151816002015f6101000a81548160ff021916908315150217905550606082015181600301556080820151816004015560a0820151816005015560c0820151816006015f6101000a81548160ff021916908360ff16021790555060e082015181600701559050506125913382614041565b807fa90e59f66e7533243b5959b6498caf4949957dbf8ccaa6b6534177c10041ea54336040516125c19190614fc3565b60405180910390a295945050505050565b6125da613d65565b6125e38161107a565b50600f5f8281526020019081526020015f205f9054906101000a900460ff16158061262157505f600d5f8381526020019081526020015f2060030154145b15612658576040517fa7d67ebb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5f8281526020019081526020015f20600301543410156126a6576040517f356680b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5f8281526020019081526020015f206002015f9054906101000a900460ff16156126fe576040517f6e40ff0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60115f8281526020019081526020015f205f9054906101000a900460ff1615612753576040517f74ed79ae00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61275e338361156f565b90505f73ffffffffffffffffffffffffffffffffffffffff16600e5f8381526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612800576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127f7906168de565b60405180910390fd5b600160115f8481526020019081526020015f205f6101000a81548160ff0219169083151502179055508060125f8481526020019081526020015f20819055506040518061014001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018381526020013481526020014281526020015f67ffffffffffffffff81111561289357612892615491565b5b6040519080825280601f01601f1916602001820160405280156128c55781602001600182028036833780820191505090505b5081526020015f5f1b81526020015f5f1b81526020015f151581526020015f81526020015f815250600e5f8381526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401908161297591906167c5565b5060a0820151816005015560c0820151816006015560e0820151816007015f6101000a81548160ff02191690831515021790555061010082015181600801556101208201518160090155905050813373ffffffffffffffffffffffffffffffffffffffff16827f2d18295f817f7e46b8d3401af48ee043761aba21f602005110a282939c3c4c7260405160405180910390a450612a10613dab565b50565b6201518081565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b6060612ab38261107a565b50600d5f8381526020019081526020015f205f018054612ad290615b67565b80601f0160208091040260200160405190810160405280929190818152602001828054612afe90615b67565b8015612b495780601f10612b2057610100808354040283529160200191612b49565b820191905f5260205f20905b815481529060010190602001808311612b2c57829003601f168201915b50505050509050919050565b6060806060805f60108054905090505f818789612b7291906168fc565b11612b88578688612b8391906168fc565b612b8a565b815b90505f888211612b9a575f612ba7565b8882612ba69190616018565b5b90508067ffffffffffffffff811115612bc357612bc2615491565b5b604051908082528060200260200182016040528015612bf15781602001602082028036833780820191505090505b5096508067ffffffffffffffff811115612c0e57612c0d615491565b5b604051908082528060200260200182016040528015612c3c5781602001602082028036833780820191505090505b5095508067ffffffffffffffff811115612c5957612c58615491565b5b604051908082528060200260200182016040528015612c875781602001602082028036833780820191505090505b5094508067ffffffffffffffff811115612ca457612ca3615491565b5b604051908082528060200260200182016040528015612cd25781602001602082028036833780820191505090505b5093505f5f90505b81811015612e0b575f6010828c612cf191906168fc565b81548110612d0257612d01615bf3565b5b905f5260205f200154905080898381518110612d2157612d20615bf3565b5b602002602001018181525050612d368161107a565b888381518110612d4957612d48615bf3565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600d5f8281526020019081526020015f2060030154878381518110612dab57612daa615bf3565b5b602002602001018181525050600d5f8281526020019081526020015f206002015f9054906101000a900460ff16868381518110612deb57612dea615bf3565b5b602002602001019015159081151581525050508080600101915050612cda565b5050505092959194509250565b600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612e9e576040517f7efb568f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600c5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f808ec13129987deb49ec337ab895a2cf7af16a4d0d55a51ddc054e2c7fb2515b60405160405180910390a35050565b612f69613d65565b5f600e5f8681526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361300d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613004906160fd565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16613031826001015461107a565b73ffffffffffffffffffffffffffffffffffffffff1614613087576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161307e90615c6a565b60405180910390fd5b5f8160090154116130cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130c490616979565b60405180910390fd5b600d5f826001015481526020019081526020015f20600501548160090154426130f69190616018565b1115613137576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161312e906169e1565b60405180910390fd5b8060050154848460405161314c929190615e16565b604051809103902014613194576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161318b90616a49565b60405180910390fd5b8060060154826040516020016131aa9190616a67565b6040516020818303038152906040528051906020012014613200576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131f790616acb565b60405180910390fd5b600d5f826001015481526020019081526020015f206002015f9054906101000a900460ff161561325c576040517f6e40ff0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8383600d5f846001015481526020019081526020015f205f0191826132829291906165ea565b5081600d5f836001015481526020019081526020015f20600401819055505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f826002015490505f83600101549050600f5f8281526020019081526020015f205f9054906101000a900460ff1615613423575f600f5f8381526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f90505b6010805490508110156133db57816010828154811061334857613347615bf3565b5b905f5260205f200154036133ce576010600160108054905061336a9190616018565b8154811061337b5761337a615bf3565b5b905f5260205f2001546010828154811061339857613397615bf3565b5b905f5260205f20018190555060108054806133b6576133b5616ae9565b5b600190038181905f5260205f20015f905590556133db565b8080600101915050613326565b505f600d5f8381526020019081526020015f2060030181905550807f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a25b5f60115f8381526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f8281526020019081526020015f205f9055600e5f8981526020019081526020015f205f5f82015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f6134ba9190614d7a565b600582015f9055600682015f9055600782015f6101000a81549060ff0219169055600882015f9055600982015f9055505061350533848360405180602001604052805f815250614134565b5f3373ffffffffffffffffffffffffffffffffffffffff168360405161352a90615e51565b5f6040518083038185875af1925050503d805f8114613564576040519150601f19603f3d011682016040523d82523d5f602084013e613569565b606091505b50509050806135ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135a4906161cd565b60405180910390fd5b887f062fb96142a4ea35fc5c48049c3a7d7a418829dea520220e03d76440bbe275c0876040516135dd9190615053565b60405180910390a250505050506135f2613dab565b50505050565b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806136c257507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806136d257506136d182614159565b5b9050919050565b5f5f6136e4836141c2565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361375657826040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161374d9190615053565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b6137ac83838360016141fb565b505050565b5f5f6137be8585856143ba565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561382857505f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614155b1561387457600f5f8581526020019081526020015f205f9054906101000a900460ff161561385a576138598461387f565b5b5f600d5f8681526020019081526020015f20600301819055505b809150509392505050565b600f5f8281526020019081526020015f205f9054906101000a900460ff1615613bf9575f60125f8381526020019081526020015f205490505f5f1b8114613b14575f600e5f8381526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613b12575f81600201541115613a05575f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826002015460405161398090615e51565b5f6040518083038185875af1925050503d805f81146139ba576040519150601f19603f3d011682016040523d82523d5f602084013e6139bf565b606091505b5050905080613a03576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139fa906161cd565b60405180910390fd5b505b5f60115f8581526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f8481526020019081526020015f205f90553373ffffffffffffffffffffffffffffffffffffffff16827f1ed784ea0b4551753ccb1bbf1711421d8a07aff605d39bb9d770c25943aea48560405160405180910390a3600e5f8381526020019081526020015f205f5f82015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f613ae09190614d7a565b600582015f9055600682015f9055600782015f6101000a81549060ff0219169055600882015f9055600982015f905550505b505b5f600f5f8481526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f90505b601080549050811015613bf6578260108281548110613b6357613b62615bf3565b5b905f5260205f20015403613be95760106001601080549050613b859190616018565b81548110613b9657613b95615bf3565b5b905f5260205f20015460108281548110613bb357613bb2615bf3565b5b905f5260205f2001819055506010805480613bd157613bd0616ae9565b5b600190038181905f5260205f20015f90559055613bf6565b8080600101915050613b41565b50505b50565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603613c6c57816040517f5b08ba18000000000000000000000000000000000000000000000000000000008152600401613c639190614fc3565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051613d589190614e7d565b60405180910390a3505050565b6002600a5403613da1576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600a81905550565b6001600a81905550565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115613f5a578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b8152600401613e139493929190616b16565b6020604051808303815f875af1925050508015613e4e57506040513d601f19601f82011682018060405250810190613e4b9190616b74565b60015b613ecf573d805f8114613e7c576040519150601f19603f3d011682016040523d82523d5f602084013e613e81565b606091505b505f815103613ec757836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401613ebe9190614fc3565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614613f5857836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401613f4f9190614fc3565b60405180910390fd5b505b5050505050565b606060405180602001604052805f815250905090565b60605f6001613f85846144d4565b0190505f8167ffffffffffffffff811115613fa357613fa2615491565b5b6040519080825280601f01601f191660200182016040528015613fd55781602001600182028036833780820191505090505b5090505f82602001820190505b600115614036578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a858161402b5761402a616b9f565b5b0494505f8503613fe2575b819350505050919050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036140b1575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016140a89190614fc3565b60405180910390fd5b5f6140bd83835f6137b1565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461412f575f6040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081526004016141269190614fc3565b60405180910390fd5b505050565b61413f848484614625565b61415361414a613798565b85858585613db5565b50505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061423357505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15614365575f614242846136d9565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156142ac57508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156142bf57506142bd8184612a1a565b155b1561430157826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016142f89190614fc3565b60405180910390fd5b811561436357838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b5f5f6143c785858561478d565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361440a5761440584614998565b614449565b8473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146144485761444781856149dc565b5b5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff160361448a5761448584614ab3565b6144c9565b8473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146144c8576144c78585614b73565b5b5b809150509392505050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310614530577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000838161452657614525616b9f565b5b0492506040810190505b6d04ee2d6d415b85acef8100000000831061456d576d04ee2d6d415b85acef8100000000838161456357614562616b9f565b5b0492506020810190505b662386f26fc10000831061459c57662386f26fc10000838161459257614591616b9f565b5b0492506010810190505b6305f5e10083106145c5576305f5e10083816145bb576145ba616b9f565b5b0492506008810190505b61271083106145ea5761271083816145e0576145df616b9f565b5b0492506004810190505b6064831061460d576064838161460357614602616b9f565b5b0492506002810190505b600a831061461c576001810190505b80915050919050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603614695575f6040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161468c9190614fc3565b60405180910390fd5b5f6146a183835f6137b1565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361471357816040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161470a9190615053565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614614787578382826040517f64283d7b00000000000000000000000000000000000000000000000000000000815260040161477e93929190615b97565b60405180910390fd5b50505050565b5f5f614798846141c2565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146147d9576147d8818486614bf7565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614614864576148185f855f5f6141fb565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16146148e357600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b60088054905060095f8381526020019081526020015f2081905550600881908060018154018082558091505060019003905f5260205f20015f909190919091505550565b5f6149e6836112df565b90505f60075f8481526020019081526020015f205490505f60065f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050828214614a85575f815f8581526020019081526020015f2054905080825f8581526020019081526020015f20819055508260075f8381526020019081526020015f2081905550505b60075f8581526020019081526020015f205f9055805f8481526020019081526020015f205f90555050505050565b5f6001600880549050614ac69190616018565b90505f60095f8481526020019081526020015f205490505f60088381548110614af257614af1615bf3565b5b905f5260205f20015490508060088381548110614b1257614b11615bf3565b5b905f5260205f2001819055508160095f8381526020019081526020015f208190555060095f8581526020019081526020015f205f90556008805480614b5a57614b59616ae9565b5b600190038181905f5260205f20015f9055905550505050565b5f6001614b7f846112df565b614b899190616018565b90508160065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f20819055508060075f8481526020019081526020015f2081905550505050565b614c02838383614cba565b614cb5575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603614c7657806040517f7e273289000000000000000000000000000000000000000000000000000000008152600401614c6d9190615053565b60405180910390fd5b81816040517f177e802f000000000000000000000000000000000000000000000000000000008152600401614cac929190615bcc565b60405180910390fd5b505050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015614d7157508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480614d325750614d318484612a1a565b5b80614d7057508273ffffffffffffffffffffffffffffffffffffffff16614d588361375f565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b508054614d8690615b67565b5f825580601f10614d975750614db4565b601f0160209004905f5260205f2090810190614db39190614db7565b5b50565b5b80821115614dce575f815f905550600101614db8565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b614e1781614de3565b8114614e21575f5ffd5b50565b5f81359050614e3281614e0e565b92915050565b5f60208284031215614e4d57614e4c614ddb565b5b5f614e5a84828501614e24565b91505092915050565b5f8115159050919050565b614e7781614e63565b82525050565b5f602082019050614e905f830184614e6e565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f614ed882614e96565b614ee28185614ea0565b9350614ef2818560208601614eb0565b614efb81614ebe565b840191505092915050565b5f6020820190508181035f830152614f1e8184614ece565b905092915050565b5f819050919050565b614f3881614f26565b8114614f42575f5ffd5b50565b5f81359050614f5381614f2f565b92915050565b5f60208284031215614f6e57614f6d614ddb565b5b5f614f7b84828501614f45565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f614fad82614f84565b9050919050565b614fbd81614fa3565b82525050565b5f602082019050614fd65f830184614fb4565b92915050565b614fe581614fa3565b8114614fef575f5ffd5b50565b5f8135905061500081614fdc565b92915050565b5f5f6040838503121561501c5761501b614ddb565b5b5f61502985828601614ff2565b925050602061503a85828601614f45565b9150509250929050565b61504d81614f26565b82525050565b5f6020820190506150665f830184615044565b92915050565b5f5f5f6060848603121561508357615082614ddb565b5b5f61509086828701614ff2565b93505060206150a186828701614ff2565b92505060406150b286828701614f45565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f6150e0826150bc565b6150ea81856150c6565b93506150fa818560208601614eb0565b61510381614ebe565b840191505092915050565b5f819050919050565b6151208161510e565b82525050565b5f60ff82169050919050565b61513b81615126565b82525050565b5f610100820190508181035f83015261515a818b6150d6565b9050615169602083018a615117565b6151766040830189614e6e565b6151836060830188615044565b6151906080830187615044565b61519d60a0830186615044565b6151aa60c0830185615132565b6151b760e0830184615044565b9998505050505050505050565b6151cd8161510e565b81146151d7575f5ffd5b50565b5f813590506151e8816151c4565b92915050565b5f6020828403121561520357615202614ddb565b5b5f615210848285016151da565b91505092915050565b5f6101408201905061522d5f83018d614fb4565b61523a602083018c615044565b615247604083018b615044565b615254606083018a615044565b818103608083015261526681896150d6565b905061527560a0830188615117565b61528260c0830187615117565b61528f60e0830186614e6e565b61529d610100830185615044565b6152ab610120830184615044565b9b9a5050505050505050505050565b5f6020820190506152cd5f830184615132565b92915050565b5f5f5f606084860312156152ea576152e9614ddb565b5b5f6152f786828701614f45565b935050602061530886828701614f45565b925050604061531986828701614f45565b9150509250925092565b5f6020828403121561533857615337614ddb565b5b5f61534584828501614ff2565b91505092915050565b5f6020820190506153615f830184615117565b92915050565b61537081614e63565b811461537a575f5ffd5b50565b5f8135905061538b81615367565b92915050565b5f5f604083850312156153a7576153a6614ddb565b5b5f6153b485828601614ff2565b92505060206153c58582860161537d565b9150509250929050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126153f0576153ef6153cf565b5b8235905067ffffffffffffffff81111561540d5761540c6153d3565b5b602083019150836001820283011115615429576154286153d7565b5b9250929050565b5f5f5f6040848603121561544757615446614ddb565b5b5f61545486828701614f45565b935050602084013567ffffffffffffffff81111561547557615474614ddf565b5b615481868287016153db565b92509250509250925092565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6154c782614ebe565b810181811067ffffffffffffffff821117156154e6576154e5615491565b5b80604052505050565b5f6154f8614dd2565b905061550482826154be565b919050565b5f67ffffffffffffffff82111561552357615522615491565b5b61552c82614ebe565b9050602081019050919050565b828183375f83830152505050565b5f61555961555484615509565b6154ef565b9050828152602081018484840111156155755761557461548d565b5b615580848285615539565b509392505050565b5f82601f83011261559c5761559b6153cf565b5b81356155ac848260208601615547565b91505092915050565b5f5f5f5f608085870312156155cd576155cc614ddb565b5b5f6155da87828801614ff2565b94505060206155eb87828801614ff2565b93505060406155fc87828801614f45565b925050606085013567ffffffffffffffff81111561561d5761561c614ddf565b5b61562987828801615588565b91505092959194509250565b5f5f83601f84011261564a576156496153cf565b5b8235905067ffffffffffffffff811115615667576156666153d3565b5b602083019150836001820283011115615683576156826153d7565b5b9250929050565b5f5f5f5f606085870312156156a2576156a1614ddb565b5b5f6156af878288016151da565b945050602085013567ffffffffffffffff8111156156d0576156cf614ddf565b5b6156dc87828801615635565b935093505060406156ef878288016151da565b91505092959194509250565b61570481615126565b811461570e575f5ffd5b50565b5f8135905061571f816156fb565b92915050565b5f5f5f5f5f6080868803121561573e5761573d614ddb565b5b5f86013567ffffffffffffffff81111561575b5761575a614ddf565b5b61576788828901615635565b9550955050602061577a888289016151da565b935050604061578b88828901614f45565b925050606061579c88828901615711565b9150509295509295909350565b5f5f604083850312156157bf576157be614ddb565b5b5f6157cc85828601614ff2565b92505060206157dd85828601614ff2565b9150509250929050565b5f6020820190508181035f8301526157ff81846150d6565b905092915050565b5f5f6040838503121561581d5761581c614ddb565b5b5f61582a85828601614f45565b925050602061583b85828601614f45565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61587781614f26565b82525050565b5f615888838361586e565b60208301905092915050565b5f602082019050919050565b5f6158aa82615845565b6158b4818561584f565b93506158bf8361585f565b805f5b838110156158ef5781516158d6888261587d565b97506158e183615894565b9250506001810190506158c2565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61592e81614fa3565b82525050565b5f61593f8383615925565b60208301905092915050565b5f602082019050919050565b5f615961826158fc565b61596b8185615906565b935061597683615916565b805f5b838110156159a657815161598d8882615934565b97506159988361594b565b925050600181019050615979565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6159e581614e63565b82525050565b5f6159f683836159dc565b60208301905092915050565b5f602082019050919050565b5f615a18826159b3565b615a2281856159bd565b9350615a2d836159cd565b805f5b83811015615a5d578151615a4488826159eb565b9750615a4f83615a02565b925050600181019050615a30565b5085935050505092915050565b5f6080820190508181035f830152615a8281876158a0565b90508181036020830152615a968186615957565b90508181036040830152615aaa81856158a0565b90508181036060830152615abe8184615a0e565b905095945050505050565b5f5f5f5f60608587031215615ae157615ae0614ddb565b5b5f615aee878288016151da565b945050602085013567ffffffffffffffff811115615b0f57615b0e614ddf565b5b615b1b87828801615635565b93509350506040615b2e87828801614f45565b91505092959194509250565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680615b7e57607f821691505b602082108103615b9157615b90615b3a565b5b50919050565b5f606082019050615baa5f830186614fb4565b615bb76020830185615044565b615bc46040830184614fb4565b949350505050565b5f604082019050615bdf5f830185614fb4565b615bec6020830184615044565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e6f7420746f6b656e206f776e657200000000000000000000000000000000005f82015250565b5f615c54600f83614ea0565b9150615c5f82615c20565b602082019050919050565b5f6020820190508181035f830152615c8181615c48565b9050919050565b7f496e76616c69642074696d656f757400000000000000000000000000000000005f82015250565b5f615cbc600f83614ea0565b9150615cc782615c88565b602082019050919050565b5f6020820190508181035f830152615ce981615cb0565b9050919050565b5f8160601b9050919050565b5f615d0682615cf0565b9050919050565b5f615d1782615cfc565b9050919050565b615d2f615d2a82614fa3565b615d0d565b82525050565b5f819050919050565b615d4f615d4a82614f26565b615d35565b82525050565b5f615d608285615d1e565b601482019150615d708284615d3e565b6020820191508190509392505050565b7f436c756520616c726561647920736f6c766564000000000000000000000000005f82015250565b5f615db4601383614ea0565b9150615dbf82615d80565b602082019050919050565b5f6020820190508181035f830152615de181615da8565b9050919050565b5f81905092915050565b5f615dfd8385615de8565b9350615e0a838584615539565b82840190509392505050565b5f615e22828486615df2565b91508190509392505050565b50565b5f615e3c5f83615de8565b9150615e4782615e2e565b5f82019050919050565b5f615e5b82615e31565b9150819050919050565b7f4661696c656420746f2073656e6420736f6c76652072657761726400000000005f82015250565b5f615e99601b83614ea0565b9150615ea482615e65565b602082019050919050565b5f6020820190508181035f830152615ec681615e8d565b9050919050565b5f615ed88385614ea0565b9350615ee5838584615539565b615eee83614ebe565b840190509392505050565b5f6020820190508181035f830152615f12818486615ecd565b90509392505050565b7f4e6f7420746865206275796572000000000000000000000000000000000000005f82015250565b5f615f4f600d83614ea0565b9150615f5a82615f1b565b602082019050919050565b5f6020820190508181035f830152615f7c81615f43565b9050919050565b7f50726f6f66206e6f74207965742070726f7669646564000000000000000000005f82015250565b5f615fb7601683614ea0565b9150615fc282615f83565b602082019050919050565b5f6020820190508181035f830152615fe481615fab565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61602282614f26565b915061602d83614f26565b925082820390508181111561604557616044615feb565b5b92915050565b7f50726f6f6620766572696669636174696f6e20657870697265640000000000005f82015250565b5f61607f601a83614ea0565b915061608a8261604b565b602082019050919050565b5f6020820190508181035f8301526160ac81616073565b9050919050565b7f5472616e7366657220646f6573206e6f742065786973740000000000000000005f82015250565b5f6160e7601783614ea0565b91506160f2826160b3565b602082019050919050565b5f6020820190508181035f830152616114816160db565b9050919050565b7f4e6f7420617574686f72697a656420746f2063616e63656c00000000000000005f82015250565b5f61614f601883614ea0565b915061615a8261611b565b602082019050919050565b5f6020820190508181035f83015261617c81616143565b9050919050565b7f4661696c656420746f2073656e642045746865720000000000000000000000005f82015250565b5f6161b7601483614ea0565b91506161c282616183565b602082019050919050565b5f6020820190508181035f8301526161e4816161ab565b9050919050565b7f5472616e736665722065787069726564000000000000000000000000000000005f82015250565b5f61621f601083614ea0565b915061622a826161eb565b602082019050919050565b5f6020820190508181035f83015261624c81616213565b9050919050565b7f50726f6f6620746f6f2073686f727400000000000000000000000000000000005f82015250565b5f616287600f83614ea0565b915061629282616253565b602082019050919050565b5f6020820190508181035f8301526162b48161627b565b9050919050565b5f63ffffffff82169050919050565b5f6162d4826162bb565b91506162df836162bb565b9250828201905063ffffffff8111156162fb576162fa615feb565b5b92915050565b7f496e76616c69642070726f6f66207374727563747572650000000000000000005f82015250565b5f616335601783614ea0565b915061634082616301565b602082019050919050565b5f6020820190508181035f83015261636281616329565b9050919050565b5f616373826162bb565b915061637e836162bb565b9250828203905063ffffffff81111561639a57616399615feb565b5b92915050565b5f5ffd5b5f5ffd5b5f5f858511156163bb576163ba6163a0565b5b838611156163cc576163cb6163a4565b5b6001850283019150848603905094509492505050565b5f82905092915050565b5f82821b905092915050565b5f61640383836163e2565b8261640e813561510e565b9250602082101561644e576164497fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff836020036008026163ec565b831692505b505092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f600883026164a67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826163ec565b6164b086836163ec565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6164eb6164e66164e184614f26565b6164c8565b614f26565b9050919050565b5f819050919050565b616504836164d1565b616518616510826164f2565b848454616477565b825550505050565b5f5f905090565b61652f616520565b61653a8184846164fb565b505050565b5b8181101561655d576165525f82616527565b600181019050616540565b5050565b601f8211156165a25761657381616456565b61657c84616468565b8101602085101561658b578190505b61659f61659785616468565b83018261653f565b50505b505050565b5f82821c905092915050565b5f6165c25f19846008026165a7565b1980831691505092915050565b5f6165da83836165b3565b9150826002028217905092915050565b6165f483836163e2565b67ffffffffffffffff81111561660d5761660c615491565b5b6166178254615b67565b616622828285616561565b5f601f83116001811461664f575f841561663d578287013590505b61664785826165cf565b8655506166ae565b601f19841661665d86616456565b5f5b828110156166845784890135825560018201915060208501945060208101905061665f565b868310156166a1578489013561669d601f8916826165b3565b8355505b6001600288020188555050505b50505050505050565b5f6166c283856150c6565b93506166cf838584615539565b6166d883614ebe565b840190509392505050565b5f6060820190508181035f8301526166fc8186886166b7565b905061670b6020830185615117565b6167186040830184615117565b95945050505050565b5f81905092915050565b5f61673582614e96565b61673f8185616721565b935061674f818560208601614eb0565b80840191505092915050565b5f616766828561672b565b9150616772828461672b565b91508190509392505050565b5f61678882614f26565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036167ba576167b9615feb565b5b600182019050919050565b6167ce826150bc565b67ffffffffffffffff8111156167e7576167e6615491565b5b6167f18254615b67565b6167fc828285616561565b5f60209050601f83116001811461682d575f841561681b578287015190505b61682585826165cf565b86555061688c565b601f19841661683b86616456565b5f5b828110156168625784890151825560018201915060208501945060208101905061683d565b8683101561687f578489015161687b601f8916826165b3565b8355505b6001600288020188555050505b505050505050565b7f5472616e7366657220616c726561647920696e697469617465640000000000005f82015250565b5f6168c8601a83614ea0565b91506168d382616894565b602082019050919050565b5f6020820190508181035f8301526168f5816168bc565b9050919050565b5f61690682614f26565b915061691183614f26565b925082820190508082111561692957616928615feb565b5b92915050565b7f50726f6f66206e6f7420766572696669656400000000000000000000000000005f82015250565b5f616963601283614ea0565b915061696e8261692f565b602082019050919050565b5f6020820190508181035f83015261699081616957565b9050919050565b7f5472616e7366657220636f6d706c6574696f6e206578706972656400000000005f82015250565b5f6169cb601b83614ea0565b91506169d682616997565b602082019050919050565b5f6020820190508181035f8301526169f8816169bf565b9050919050565b7f436f6e74656e742068617368206d69736d6174636800000000000000000000005f82015250565b5f616a33601583614ea0565b9150616a3e826169ff565b602082019050919050565b5f6020820190508181035f830152616a6081616a27565b9050919050565b5f616a728284615d3e565b60208201915081905092915050565b7f522076616c75652068617368206d69736d6174636800000000000000000000005f82015250565b5f616ab5601583614ea0565b9150616ac082616a81565b602082019050919050565b5f6020820190508181035f830152616ae281616aa9565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f608082019050616b295f830187614fb4565b616b366020830186614fb4565b616b436040830185615044565b8181036060830152616b5581846150d6565b905095945050505050565b5f81519050616b6e81614e0e565b92915050565b5f60208284031215616b8957616b88614ddb565b5b5f616b9684828501616b60565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffdfea2646970667358221220b9ccfe647e71ec030d6e63a87f8025a2002971f2c8435078dba0e6068914a62c64736f6c634300081c0033",
}

// SkavengeABI is the input ABI used to generate the binding from.
// Deprecated: Use SkavengeMetaData.ABI instead.
var SkavengeABI = SkavengeMetaData.ABI

// SkavengeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SkavengeMetaData.Bin instead.
var SkavengeBin = SkavengeMetaData.Bin

// DeploySkavenge deploys a new Ethereum contract, binding an instance of Skavenge to it.
func DeploySkavenge(auth *bind.TransactOpts, backend bind.ContractBackend, initialMinter common.Address) (common.Address, *types.Transaction, *Skavenge, error) {
	parsed, err := SkavengeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SkavengeBin), backend, initialMinter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Skavenge{SkavengeCaller: SkavengeCaller{contract: contract}, SkavengeTransactor: SkavengeTransactor{contract: contract}, SkavengeFilterer: SkavengeFilterer{contract: contract}}, nil
}

// Skavenge is an auto generated Go binding around an Ethereum contract.
type Skavenge struct {
	SkavengeCaller     // Read-only binding to the contract
	SkavengeTransactor // Write-only binding to the contract
	SkavengeFilterer   // Log filterer for contract events
}

// SkavengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SkavengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SkavengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SkavengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SkavengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SkavengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SkavengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SkavengeSession struct {
	Contract     *Skavenge         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SkavengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SkavengeCallerSession struct {
	Contract *SkavengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SkavengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SkavengeTransactorSession struct {
	Contract     *SkavengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SkavengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SkavengeRaw struct {
	Contract *Skavenge // Generic contract binding to access the raw methods on
}

// SkavengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SkavengeCallerRaw struct {
	Contract *SkavengeCaller // Generic read-only contract binding to access the raw methods on
}

// SkavengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SkavengeTransactorRaw struct {
	Contract *SkavengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSkavenge creates a new instance of Skavenge, bound to a specific deployed contract.
func NewSkavenge(address common.Address, backend bind.ContractBackend) (*Skavenge, error) {
	contract, err := bindSkavenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Skavenge{SkavengeCaller: SkavengeCaller{contract: contract}, SkavengeTransactor: SkavengeTransactor{contract: contract}, SkavengeFilterer: SkavengeFilterer{contract: contract}}, nil
}

// NewSkavengeCaller creates a new read-only instance of Skavenge, bound to a specific deployed contract.
func NewSkavengeCaller(address common.Address, caller bind.ContractCaller) (*SkavengeCaller, error) {
	contract, err := bindSkavenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SkavengeCaller{contract: contract}, nil
}

// NewSkavengeTransactor creates a new write-only instance of Skavenge, bound to a specific deployed contract.
func NewSkavengeTransactor(address common.Address, transactor bind.ContractTransactor) (*SkavengeTransactor, error) {
	contract, err := bindSkavenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SkavengeTransactor{contract: contract}, nil
}

// NewSkavengeFilterer creates a new log filterer instance of Skavenge, bound to a specific deployed contract.
func NewSkavengeFilterer(address common.Address, filterer bind.ContractFilterer) (*SkavengeFilterer, error) {
	contract, err := bindSkavenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SkavengeFilterer{contract: contract}, nil
}

// bindSkavenge binds a generic wrapper to an already deployed contract.
func bindSkavenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SkavengeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Skavenge *SkavengeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Skavenge.Contract.SkavengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Skavenge *SkavengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Skavenge.Contract.SkavengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Skavenge *SkavengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Skavenge.Contract.SkavengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Skavenge *SkavengeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Skavenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Skavenge *SkavengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Skavenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Skavenge *SkavengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Skavenge.Contract.contract.Transact(opts, method, params...)
}

// MAXTIMEOUT is a free data retrieval call binding the contract method 0xde38eb3a.
//
// Solidity: function MAX_TIMEOUT() view returns(uint256)
func (_Skavenge *SkavengeCaller) MAXTIMEOUT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "MAX_TIMEOUT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTIMEOUT is a free data retrieval call binding the contract method 0xde38eb3a.
//
// Solidity: function MAX_TIMEOUT() view returns(uint256)
func (_Skavenge *SkavengeSession) MAXTIMEOUT() (*big.Int, error) {
	return _Skavenge.Contract.MAXTIMEOUT(&_Skavenge.CallOpts)
}

// MAXTIMEOUT is a free data retrieval call binding the contract method 0xde38eb3a.
//
// Solidity: function MAX_TIMEOUT() view returns(uint256)
func (_Skavenge *SkavengeCallerSession) MAXTIMEOUT() (*big.Int, error) {
	return _Skavenge.Contract.MAXTIMEOUT(&_Skavenge.CallOpts)
}

// MINTIMEOUT is a free data retrieval call binding the contract method 0x543ad1df.
//
// Solidity: function MIN_TIMEOUT() view returns(uint256)
func (_Skavenge *SkavengeCaller) MINTIMEOUT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "MIN_TIMEOUT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTIMEOUT is a free data retrieval call binding the contract method 0x543ad1df.
//
// Solidity: function MIN_TIMEOUT() view returns(uint256)
func (_Skavenge *SkavengeSession) MINTIMEOUT() (*big.Int, error) {
	return _Skavenge.Contract.MINTIMEOUT(&_Skavenge.CallOpts)
}

// MINTIMEOUT is a free data retrieval call binding the contract method 0x543ad1df.
//
// Solidity: function MIN_TIMEOUT() view returns(uint256)
func (_Skavenge *SkavengeCallerSession) MINTIMEOUT() (*big.Int, error) {
	return _Skavenge.Contract.MINTIMEOUT(&_Skavenge.CallOpts)
}

// ActiveTransferIds is a free data retrieval call binding the contract method 0x79096ee8.
//
// Solidity: function activeTransferIds(uint256 ) view returns(bytes32)
func (_Skavenge *SkavengeCaller) ActiveTransferIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "activeTransferIds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActiveTransferIds is a free data retrieval call binding the contract method 0x79096ee8.
//
// Solidity: function activeTransferIds(uint256 ) view returns(bytes32)
func (_Skavenge *SkavengeSession) ActiveTransferIds(arg0 *big.Int) ([32]byte, error) {
	return _Skavenge.Contract.ActiveTransferIds(&_Skavenge.CallOpts, arg0)
}

// ActiveTransferIds is a free data retrieval call binding the contract method 0x79096ee8.
//
// Solidity: function activeTransferIds(uint256 ) view returns(bytes32)
func (_Skavenge *SkavengeCallerSession) ActiveTransferIds(arg0 *big.Int) ([32]byte, error) {
	return _Skavenge.Contract.ActiveTransferIds(&_Skavenge.CallOpts, arg0)
}

// AuthorizedMinter is a free data retrieval call binding the contract method 0x1ba538cd.
//
// Solidity: function authorizedMinter() view returns(address)
func (_Skavenge *SkavengeCaller) AuthorizedMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "authorizedMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthorizedMinter is a free data retrieval call binding the contract method 0x1ba538cd.
//
// Solidity: function authorizedMinter() view returns(address)
func (_Skavenge *SkavengeSession) AuthorizedMinter() (common.Address, error) {
	return _Skavenge.Contract.AuthorizedMinter(&_Skavenge.CallOpts)
}

// AuthorizedMinter is a free data retrieval call binding the contract method 0x1ba538cd.
//
// Solidity: function authorizedMinter() view returns(address)
func (_Skavenge *SkavengeCallerSession) AuthorizedMinter() (common.Address, error) {
	return _Skavenge.Contract.AuthorizedMinter(&_Skavenge.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Skavenge *SkavengeCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Skavenge *SkavengeSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Skavenge.Contract.BalanceOf(&_Skavenge.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Skavenge *SkavengeCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Skavenge.Contract.BalanceOf(&_Skavenge.CallOpts, owner)
}

// Clues is a free data retrieval call binding the contract method 0x30f37c7f.
//
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 salePrice, uint256 rValue, uint256 timeout, uint8 pointValue, uint256 solveReward)
func (_Skavenge *SkavengeCaller) Clues(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
	SalePrice         *big.Int
	RValue            *big.Int
	Timeout           *big.Int
	PointValue        uint8
	SolveReward       *big.Int
}, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "clues", arg0)

	outstruct := new(struct {
		EncryptedContents []byte
		SolutionHash      [32]byte
		IsSolved          bool
		SalePrice         *big.Int
		RValue            *big.Int
		Timeout           *big.Int
		PointValue        uint8
		SolveReward       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EncryptedContents = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.SolutionHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.IsSolved = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.SalePrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Timeout = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PointValue = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.SolveReward = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Clues is a free data retrieval call binding the contract method 0x30f37c7f.
//
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 salePrice, uint256 rValue, uint256 timeout, uint8 pointValue, uint256 solveReward)
func (_Skavenge *SkavengeSession) Clues(arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
	SalePrice         *big.Int
	RValue            *big.Int
	Timeout           *big.Int
	PointValue        uint8
	SolveReward       *big.Int
}, error) {
	return _Skavenge.Contract.Clues(&_Skavenge.CallOpts, arg0)
}

// Clues is a free data retrieval call binding the contract method 0x30f37c7f.
//
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 salePrice, uint256 rValue, uint256 timeout, uint8 pointValue, uint256 solveReward)
func (_Skavenge *SkavengeCallerSession) Clues(arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
	SalePrice         *big.Int
	RValue            *big.Int
	Timeout           *big.Int
	PointValue        uint8
	SolveReward       *big.Int
}, error) {
	return _Skavenge.Contract.Clues(&_Skavenge.CallOpts, arg0)
}

// CluesForSale is a free data retrieval call binding the contract method 0x3427ee94.
//
// Solidity: function cluesForSale(uint256 ) view returns(bool)
func (_Skavenge *SkavengeCaller) CluesForSale(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "cluesForSale", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CluesForSale is a free data retrieval call binding the contract method 0x3427ee94.
//
// Solidity: function cluesForSale(uint256 ) view returns(bool)
func (_Skavenge *SkavengeSession) CluesForSale(arg0 *big.Int) (bool, error) {
	return _Skavenge.Contract.CluesForSale(&_Skavenge.CallOpts, arg0)
}

// CluesForSale is a free data retrieval call binding the contract method 0x3427ee94.
//
// Solidity: function cluesForSale(uint256 ) view returns(bool)
func (_Skavenge *SkavengeCallerSession) CluesForSale(arg0 *big.Int) (bool, error) {
	return _Skavenge.Contract.CluesForSale(&_Skavenge.CallOpts, arg0)
}

// GenerateTransferId is a free data retrieval call binding the contract method 0xa6cd5ff5.
//
// Solidity: function generateTransferId(address buyer, uint256 tokenId) pure returns(bytes32)
func (_Skavenge *SkavengeCaller) GenerateTransferId(opts *bind.CallOpts, buyer common.Address, tokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "generateTransferId", buyer, tokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GenerateTransferId is a free data retrieval call binding the contract method 0xa6cd5ff5.
//
// Solidity: function generateTransferId(address buyer, uint256 tokenId) pure returns(bytes32)
func (_Skavenge *SkavengeSession) GenerateTransferId(buyer common.Address, tokenId *big.Int) ([32]byte, error) {
	return _Skavenge.Contract.GenerateTransferId(&_Skavenge.CallOpts, buyer, tokenId)
}

// GenerateTransferId is a free data retrieval call binding the contract method 0xa6cd5ff5.
//
// Solidity: function generateTransferId(address buyer, uint256 tokenId) pure returns(bytes32)
func (_Skavenge *SkavengeCallerSession) GenerateTransferId(buyer common.Address, tokenId *big.Int) ([32]byte, error) {
	return _Skavenge.Contract.GenerateTransferId(&_Skavenge.CallOpts, buyer, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Skavenge *SkavengeCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Skavenge *SkavengeSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Skavenge.Contract.GetApproved(&_Skavenge.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Skavenge *SkavengeCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Skavenge.Contract.GetApproved(&_Skavenge.CallOpts, tokenId)
}

// GetClueContents is a free data retrieval call binding the contract method 0xeb927a83.
//
// Solidity: function getClueContents(uint256 tokenId) view returns(bytes)
func (_Skavenge *SkavengeCaller) GetClueContents(opts *bind.CallOpts, tokenId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "getClueContents", tokenId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetClueContents is a free data retrieval call binding the contract method 0xeb927a83.
//
// Solidity: function getClueContents(uint256 tokenId) view returns(bytes)
func (_Skavenge *SkavengeSession) GetClueContents(tokenId *big.Int) ([]byte, error) {
	return _Skavenge.Contract.GetClueContents(&_Skavenge.CallOpts, tokenId)
}

// GetClueContents is a free data retrieval call binding the contract method 0xeb927a83.
//
// Solidity: function getClueContents(uint256 tokenId) view returns(bytes)
func (_Skavenge *SkavengeCallerSession) GetClueContents(tokenId *big.Int) ([]byte, error) {
	return _Skavenge.Contract.GetClueContents(&_Skavenge.CallOpts, tokenId)
}

// GetCluesForSale is a free data retrieval call binding the contract method 0xf12b72ba.
//
// Solidity: function getCluesForSale(uint256 offset, uint256 limit) view returns(uint256[] tokenIds, address[] owners, uint256[] prices, bool[] solvedStatus)
func (_Skavenge *SkavengeCaller) GetCluesForSale(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	TokenIds     []*big.Int
	Owners       []common.Address
	Prices       []*big.Int
	SolvedStatus []bool
}, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "getCluesForSale", offset, limit)

	outstruct := new(struct {
		TokenIds     []*big.Int
		Owners       []common.Address
		Prices       []*big.Int
		SolvedStatus []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenIds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Owners = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Prices = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.SolvedStatus = *abi.ConvertType(out[3], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetCluesForSale is a free data retrieval call binding the contract method 0xf12b72ba.
//
// Solidity: function getCluesForSale(uint256 offset, uint256 limit) view returns(uint256[] tokenIds, address[] owners, uint256[] prices, bool[] solvedStatus)
func (_Skavenge *SkavengeSession) GetCluesForSale(offset *big.Int, limit *big.Int) (struct {
	TokenIds     []*big.Int
	Owners       []common.Address
	Prices       []*big.Int
	SolvedStatus []bool
}, error) {
	return _Skavenge.Contract.GetCluesForSale(&_Skavenge.CallOpts, offset, limit)
}

// GetCluesForSale is a free data retrieval call binding the contract method 0xf12b72ba.
//
// Solidity: function getCluesForSale(uint256 offset, uint256 limit) view returns(uint256[] tokenIds, address[] owners, uint256[] prices, bool[] solvedStatus)
func (_Skavenge *SkavengeCallerSession) GetCluesForSale(offset *big.Int, limit *big.Int) (struct {
	TokenIds     []*big.Int
	Owners       []common.Address
	Prices       []*big.Int
	SolvedStatus []bool
}, error) {
	return _Skavenge.Contract.GetCluesForSale(&_Skavenge.CallOpts, offset, limit)
}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_Skavenge *SkavengeCaller) GetCurrentTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "getCurrentTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_Skavenge *SkavengeSession) GetCurrentTokenId() (*big.Int, error) {
	return _Skavenge.Contract.GetCurrentTokenId(&_Skavenge.CallOpts)
}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_Skavenge *SkavengeCallerSession) GetCurrentTokenId() (*big.Int, error) {
	return _Skavenge.Contract.GetCurrentTokenId(&_Skavenge.CallOpts)
}

// GetPointValue is a free data retrieval call binding the contract method 0x437fdc33.
//
// Solidity: function getPointValue(uint256 tokenId) view returns(uint8)
func (_Skavenge *SkavengeCaller) GetPointValue(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "getPointValue", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPointValue is a free data retrieval call binding the contract method 0x437fdc33.
//
// Solidity: function getPointValue(uint256 tokenId) view returns(uint8)
func (_Skavenge *SkavengeSession) GetPointValue(tokenId *big.Int) (uint8, error) {
	return _Skavenge.Contract.GetPointValue(&_Skavenge.CallOpts, tokenId)
}

// GetPointValue is a free data retrieval call binding the contract method 0x437fdc33.
//
// Solidity: function getPointValue(uint256 tokenId) view returns(uint8)
func (_Skavenge *SkavengeCallerSession) GetPointValue(tokenId *big.Int) (uint8, error) {
	return _Skavenge.Contract.GetPointValue(&_Skavenge.CallOpts, tokenId)
}

// GetRValue is a free data retrieval call binding the contract method 0xd32d5790.
//
// Solidity: function getRValue(uint256 tokenId) view returns(uint256)
func (_Skavenge *SkavengeCaller) GetRValue(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "getRValue", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRValue is a free data retrieval call binding the contract method 0xd32d5790.
//
// Solidity: function getRValue(uint256 tokenId) view returns(uint256)
func (_Skavenge *SkavengeSession) GetRValue(tokenId *big.Int) (*big.Int, error) {
	return _Skavenge.Contract.GetRValue(&_Skavenge.CallOpts, tokenId)
}

// GetRValue is a free data retrieval call binding the contract method 0xd32d5790.
//
// Solidity: function getRValue(uint256 tokenId) view returns(uint256)
func (_Skavenge *SkavengeCallerSession) GetRValue(tokenId *big.Int) (*big.Int, error) {
	return _Skavenge.Contract.GetRValue(&_Skavenge.CallOpts, tokenId)
}

// GetSolveReward is a free data retrieval call binding the contract method 0x73c2c7e1.
//
// Solidity: function getSolveReward(uint256 tokenId) view returns(uint256)
func (_Skavenge *SkavengeCaller) GetSolveReward(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "getSolveReward", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSolveReward is a free data retrieval call binding the contract method 0x73c2c7e1.
//
// Solidity: function getSolveReward(uint256 tokenId) view returns(uint256)
func (_Skavenge *SkavengeSession) GetSolveReward(tokenId *big.Int) (*big.Int, error) {
	return _Skavenge.Contract.GetSolveReward(&_Skavenge.CallOpts, tokenId)
}

// GetSolveReward is a free data retrieval call binding the contract method 0x73c2c7e1.
//
// Solidity: function getSolveReward(uint256 tokenId) view returns(uint256)
func (_Skavenge *SkavengeCallerSession) GetSolveReward(tokenId *big.Int) (*big.Int, error) {
	return _Skavenge.Contract.GetSolveReward(&_Skavenge.CallOpts, tokenId)
}

// GetTotalCluesForSale is a free data retrieval call binding the contract method 0x74b19a07.
//
// Solidity: function getTotalCluesForSale() view returns(uint256)
func (_Skavenge *SkavengeCaller) GetTotalCluesForSale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "getTotalCluesForSale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalCluesForSale is a free data retrieval call binding the contract method 0x74b19a07.
//
// Solidity: function getTotalCluesForSale() view returns(uint256)
func (_Skavenge *SkavengeSession) GetTotalCluesForSale() (*big.Int, error) {
	return _Skavenge.Contract.GetTotalCluesForSale(&_Skavenge.CallOpts)
}

// GetTotalCluesForSale is a free data retrieval call binding the contract method 0x74b19a07.
//
// Solidity: function getTotalCluesForSale() view returns(uint256)
func (_Skavenge *SkavengeCallerSession) GetTotalCluesForSale() (*big.Int, error) {
	return _Skavenge.Contract.GetTotalCluesForSale(&_Skavenge.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Skavenge *SkavengeCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Skavenge *SkavengeSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Skavenge.Contract.IsApprovedForAll(&_Skavenge.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Skavenge *SkavengeCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Skavenge.Contract.IsApprovedForAll(&_Skavenge.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Skavenge *SkavengeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Skavenge *SkavengeSession) Name() (string, error) {
	return _Skavenge.Contract.Name(&_Skavenge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Skavenge *SkavengeCallerSession) Name() (string, error) {
	return _Skavenge.Contract.Name(&_Skavenge.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Skavenge *SkavengeCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Skavenge *SkavengeSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Skavenge.Contract.OwnerOf(&_Skavenge.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Skavenge *SkavengeCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Skavenge.Contract.OwnerOf(&_Skavenge.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Skavenge *SkavengeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Skavenge *SkavengeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Skavenge.Contract.SupportsInterface(&_Skavenge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Skavenge *SkavengeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Skavenge.Contract.SupportsInterface(&_Skavenge.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Skavenge *SkavengeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Skavenge *SkavengeSession) Symbol() (string, error) {
	return _Skavenge.Contract.Symbol(&_Skavenge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Skavenge *SkavengeCallerSession) Symbol() (string, error) {
	return _Skavenge.Contract.Symbol(&_Skavenge.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Skavenge *SkavengeCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Skavenge *SkavengeSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Skavenge.Contract.TokenByIndex(&_Skavenge.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Skavenge *SkavengeCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Skavenge.Contract.TokenByIndex(&_Skavenge.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Skavenge *SkavengeCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Skavenge *SkavengeSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Skavenge.Contract.TokenOfOwnerByIndex(&_Skavenge.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Skavenge *SkavengeCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Skavenge.Contract.TokenOfOwnerByIndex(&_Skavenge.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Skavenge *SkavengeCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Skavenge *SkavengeSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Skavenge.Contract.TokenURI(&_Skavenge.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Skavenge *SkavengeCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Skavenge.Contract.TokenURI(&_Skavenge.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Skavenge *SkavengeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Skavenge *SkavengeSession) TotalSupply() (*big.Int, error) {
	return _Skavenge.Contract.TotalSupply(&_Skavenge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Skavenge *SkavengeCallerSession) TotalSupply() (*big.Int, error) {
	return _Skavenge.Contract.TotalSupply(&_Skavenge.CallOpts)
}

// TransferInProgress is a free data retrieval call binding the contract method 0x34499fff.
//
// Solidity: function transferInProgress(uint256 ) view returns(bool)
func (_Skavenge *SkavengeCaller) TransferInProgress(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "transferInProgress", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferInProgress is a free data retrieval call binding the contract method 0x34499fff.
//
// Solidity: function transferInProgress(uint256 ) view returns(bool)
func (_Skavenge *SkavengeSession) TransferInProgress(arg0 *big.Int) (bool, error) {
	return _Skavenge.Contract.TransferInProgress(&_Skavenge.CallOpts, arg0)
}

// TransferInProgress is a free data retrieval call binding the contract method 0x34499fff.
//
// Solidity: function transferInProgress(uint256 ) view returns(bool)
func (_Skavenge *SkavengeCallerSession) TransferInProgress(arg0 *big.Int) (bool, error) {
	return _Skavenge.Contract.TransferInProgress(&_Skavenge.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(address buyer, uint256 tokenId, uint256 value, uint256 initiatedAt, bytes proof, bytes32 newClueHash, bytes32 rValueHash, bool proofVerified, uint256 proofProvidedAt, uint256 verifiedAt)
func (_Skavenge *SkavengeCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Buyer           common.Address
	TokenId         *big.Int
	Value           *big.Int
	InitiatedAt     *big.Int
	Proof           []byte
	NewClueHash     [32]byte
	RValueHash      [32]byte
	ProofVerified   bool
	ProofProvidedAt *big.Int
	VerifiedAt      *big.Int
}, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "transfers", arg0)

	outstruct := new(struct {
		Buyer           common.Address
		TokenId         *big.Int
		Value           *big.Int
		InitiatedAt     *big.Int
		Proof           []byte
		NewClueHash     [32]byte
		RValueHash      [32]byte
		ProofVerified   bool
		ProofProvidedAt *big.Int
		VerifiedAt      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Buyer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InitiatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Proof = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.NewClueHash = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.RValueHash = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.ProofVerified = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.ProofProvidedAt = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.VerifiedAt = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(address buyer, uint256 tokenId, uint256 value, uint256 initiatedAt, bytes proof, bytes32 newClueHash, bytes32 rValueHash, bool proofVerified, uint256 proofProvidedAt, uint256 verifiedAt)
func (_Skavenge *SkavengeSession) Transfers(arg0 [32]byte) (struct {
	Buyer           common.Address
	TokenId         *big.Int
	Value           *big.Int
	InitiatedAt     *big.Int
	Proof           []byte
	NewClueHash     [32]byte
	RValueHash      [32]byte
	ProofVerified   bool
	ProofProvidedAt *big.Int
	VerifiedAt      *big.Int
}, error) {
	return _Skavenge.Contract.Transfers(&_Skavenge.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(address buyer, uint256 tokenId, uint256 value, uint256 initiatedAt, bytes proof, bytes32 newClueHash, bytes32 rValueHash, bool proofVerified, uint256 proofProvidedAt, uint256 verifiedAt)
func (_Skavenge *SkavengeCallerSession) Transfers(arg0 [32]byte) (struct {
	Buyer           common.Address
	TokenId         *big.Int
	Value           *big.Int
	InitiatedAt     *big.Int
	Proof           []byte
	NewClueHash     [32]byte
	RValueHash      [32]byte
	ProofVerified   bool
	ProofProvidedAt *big.Int
	VerifiedAt      *big.Int
}, error) {
	return _Skavenge.Contract.Transfers(&_Skavenge.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Skavenge *SkavengeTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Skavenge *SkavengeSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.Approve(&_Skavenge.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Skavenge *SkavengeTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.Approve(&_Skavenge.TransactOpts, to, tokenId)
}

// AttemptSolution is a paid mutator transaction binding the contract method 0xaff202b4.
//
// Solidity: function attemptSolution(uint256 tokenId, string solution) returns()
func (_Skavenge *SkavengeTransactor) AttemptSolution(opts *bind.TransactOpts, tokenId *big.Int, solution string) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "attemptSolution", tokenId, solution)
}

// AttemptSolution is a paid mutator transaction binding the contract method 0xaff202b4.
//
// Solidity: function attemptSolution(uint256 tokenId, string solution) returns()
func (_Skavenge *SkavengeSession) AttemptSolution(tokenId *big.Int, solution string) (*types.Transaction, error) {
	return _Skavenge.Contract.AttemptSolution(&_Skavenge.TransactOpts, tokenId, solution)
}

// AttemptSolution is a paid mutator transaction binding the contract method 0xaff202b4.
//
// Solidity: function attemptSolution(uint256 tokenId, string solution) returns()
func (_Skavenge *SkavengeTransactorSession) AttemptSolution(tokenId *big.Int, solution string) (*types.Transaction, error) {
	return _Skavenge.Contract.AttemptSolution(&_Skavenge.TransactOpts, tokenId, solution)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0xb329bf5c.
//
// Solidity: function cancelTransfer(bytes32 transferId) returns()
func (_Skavenge *SkavengeTransactor) CancelTransfer(opts *bind.TransactOpts, transferId [32]byte) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "cancelTransfer", transferId)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0xb329bf5c.
//
// Solidity: function cancelTransfer(bytes32 transferId) returns()
func (_Skavenge *SkavengeSession) CancelTransfer(transferId [32]byte) (*types.Transaction, error) {
	return _Skavenge.Contract.CancelTransfer(&_Skavenge.TransactOpts, transferId)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0xb329bf5c.
//
// Solidity: function cancelTransfer(bytes32 transferId) returns()
func (_Skavenge *SkavengeTransactorSession) CancelTransfer(transferId [32]byte) (*types.Transaction, error) {
	return _Skavenge.Contract.CancelTransfer(&_Skavenge.TransactOpts, transferId)
}

// CompleteTransfer is a paid mutator transaction binding the contract method 0xfae5380c.
//
// Solidity: function completeTransfer(bytes32 transferId, bytes newEncryptedContents, uint256 rValue) returns()
func (_Skavenge *SkavengeTransactor) CompleteTransfer(opts *bind.TransactOpts, transferId [32]byte, newEncryptedContents []byte, rValue *big.Int) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "completeTransfer", transferId, newEncryptedContents, rValue)
}

// CompleteTransfer is a paid mutator transaction binding the contract method 0xfae5380c.
//
// Solidity: function completeTransfer(bytes32 transferId, bytes newEncryptedContents, uint256 rValue) returns()
func (_Skavenge *SkavengeSession) CompleteTransfer(transferId [32]byte, newEncryptedContents []byte, rValue *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.CompleteTransfer(&_Skavenge.TransactOpts, transferId, newEncryptedContents, rValue)
}

// CompleteTransfer is a paid mutator transaction binding the contract method 0xfae5380c.
//
// Solidity: function completeTransfer(bytes32 transferId, bytes newEncryptedContents, uint256 rValue) returns()
func (_Skavenge *SkavengeTransactorSession) CompleteTransfer(transferId [32]byte, newEncryptedContents []byte, rValue *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.CompleteTransfer(&_Skavenge.TransactOpts, transferId, newEncryptedContents, rValue)
}

// InitiatePurchase is a paid mutator transaction binding the contract method 0xdd142be0.
//
// Solidity: function initiatePurchase(uint256 tokenId) payable returns()
func (_Skavenge *SkavengeTransactor) InitiatePurchase(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "initiatePurchase", tokenId)
}

// InitiatePurchase is a paid mutator transaction binding the contract method 0xdd142be0.
//
// Solidity: function initiatePurchase(uint256 tokenId) payable returns()
func (_Skavenge *SkavengeSession) InitiatePurchase(tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.InitiatePurchase(&_Skavenge.TransactOpts, tokenId)
}

// InitiatePurchase is a paid mutator transaction binding the contract method 0xdd142be0.
//
// Solidity: function initiatePurchase(uint256 tokenId) payable returns()
func (_Skavenge *SkavengeTransactorSession) InitiatePurchase(tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.InitiatePurchase(&_Skavenge.TransactOpts, tokenId)
}

// MintClue is a paid mutator transaction binding the contract method 0xd4098ca8.
//
// Solidity: function mintClue(bytes encryptedContents, bytes32 solutionHash, uint256 rValue, uint8 pointValue) payable returns(uint256 tokenId)
func (_Skavenge *SkavengeTransactor) MintClue(opts *bind.TransactOpts, encryptedContents []byte, solutionHash [32]byte, rValue *big.Int, pointValue uint8) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "mintClue", encryptedContents, solutionHash, rValue, pointValue)
}

// MintClue is a paid mutator transaction binding the contract method 0xd4098ca8.
//
// Solidity: function mintClue(bytes encryptedContents, bytes32 solutionHash, uint256 rValue, uint8 pointValue) payable returns(uint256 tokenId)
func (_Skavenge *SkavengeSession) MintClue(encryptedContents []byte, solutionHash [32]byte, rValue *big.Int, pointValue uint8) (*types.Transaction, error) {
	return _Skavenge.Contract.MintClue(&_Skavenge.TransactOpts, encryptedContents, solutionHash, rValue, pointValue)
}

// MintClue is a paid mutator transaction binding the contract method 0xd4098ca8.
//
// Solidity: function mintClue(bytes encryptedContents, bytes32 solutionHash, uint256 rValue, uint8 pointValue) payable returns(uint256 tokenId)
func (_Skavenge *SkavengeTransactorSession) MintClue(encryptedContents []byte, solutionHash [32]byte, rValue *big.Int, pointValue uint8) (*types.Transaction, error) {
	return _Skavenge.Contract.MintClue(&_Skavenge.TransactOpts, encryptedContents, solutionHash, rValue, pointValue)
}

// ProvideProof is a paid mutator transaction binding the contract method 0xc2d554ae.
//
// Solidity: function provideProof(bytes32 transferId, bytes proof, bytes32 newClueHash) returns()
func (_Skavenge *SkavengeTransactor) ProvideProof(opts *bind.TransactOpts, transferId [32]byte, proof []byte, newClueHash [32]byte) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "provideProof", transferId, proof, newClueHash)
}

// ProvideProof is a paid mutator transaction binding the contract method 0xc2d554ae.
//
// Solidity: function provideProof(bytes32 transferId, bytes proof, bytes32 newClueHash) returns()
func (_Skavenge *SkavengeSession) ProvideProof(transferId [32]byte, proof []byte, newClueHash [32]byte) (*types.Transaction, error) {
	return _Skavenge.Contract.ProvideProof(&_Skavenge.TransactOpts, transferId, proof, newClueHash)
}

// ProvideProof is a paid mutator transaction binding the contract method 0xc2d554ae.
//
// Solidity: function provideProof(bytes32 transferId, bytes proof, bytes32 newClueHash) returns()
func (_Skavenge *SkavengeTransactorSession) ProvideProof(transferId [32]byte, proof []byte, newClueHash [32]byte) (*types.Transaction, error) {
	return _Skavenge.Contract.ProvideProof(&_Skavenge.TransactOpts, transferId, proof, newClueHash)
}

// RemoveSalePrice is a paid mutator transaction binding the contract method 0x8d7cf3e4.
//
// Solidity: function removeSalePrice(uint256 tokenId) returns()
func (_Skavenge *SkavengeTransactor) RemoveSalePrice(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "removeSalePrice", tokenId)
}

// RemoveSalePrice is a paid mutator transaction binding the contract method 0x8d7cf3e4.
//
// Solidity: function removeSalePrice(uint256 tokenId) returns()
func (_Skavenge *SkavengeSession) RemoveSalePrice(tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.RemoveSalePrice(&_Skavenge.TransactOpts, tokenId)
}

// RemoveSalePrice is a paid mutator transaction binding the contract method 0x8d7cf3e4.
//
// Solidity: function removeSalePrice(uint256 tokenId) returns()
func (_Skavenge *SkavengeTransactorSession) RemoveSalePrice(tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.RemoveSalePrice(&_Skavenge.TransactOpts, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Skavenge *SkavengeTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Skavenge *SkavengeSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.SafeTransferFrom(&_Skavenge.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Skavenge *SkavengeTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.SafeTransferFrom(&_Skavenge.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Skavenge *SkavengeTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Skavenge *SkavengeSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Skavenge.Contract.SafeTransferFrom0(&_Skavenge.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Skavenge *SkavengeTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Skavenge.Contract.SafeTransferFrom0(&_Skavenge.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Skavenge *SkavengeTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Skavenge *SkavengeSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Skavenge.Contract.SetApprovalForAll(&_Skavenge.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Skavenge *SkavengeTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Skavenge.Contract.SetApprovalForAll(&_Skavenge.TransactOpts, operator, approved)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0x6c39cc34.
//
// Solidity: function setSalePrice(uint256 tokenId, uint256 price, uint256 timeout) returns()
func (_Skavenge *SkavengeTransactor) SetSalePrice(opts *bind.TransactOpts, tokenId *big.Int, price *big.Int, timeout *big.Int) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "setSalePrice", tokenId, price, timeout)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0x6c39cc34.
//
// Solidity: function setSalePrice(uint256 tokenId, uint256 price, uint256 timeout) returns()
func (_Skavenge *SkavengeSession) SetSalePrice(tokenId *big.Int, price *big.Int, timeout *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.SetSalePrice(&_Skavenge.TransactOpts, tokenId, price, timeout)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0x6c39cc34.
//
// Solidity: function setSalePrice(uint256 tokenId, uint256 price, uint256 timeout) returns()
func (_Skavenge *SkavengeTransactorSession) SetSalePrice(tokenId *big.Int, price *big.Int, timeout *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.SetSalePrice(&_Skavenge.TransactOpts, tokenId, price, timeout)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Skavenge *SkavengeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Skavenge *SkavengeSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.TransferFrom(&_Skavenge.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Skavenge *SkavengeTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.TransferFrom(&_Skavenge.TransactOpts, from, to, tokenId)
}

// UpdateAuthorizedMinter is a paid mutator transaction binding the contract method 0xf8f5a544.
//
// Solidity: function updateAuthorizedMinter(address newMinter) returns()
func (_Skavenge *SkavengeTransactor) UpdateAuthorizedMinter(opts *bind.TransactOpts, newMinter common.Address) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "updateAuthorizedMinter", newMinter)
}

// UpdateAuthorizedMinter is a paid mutator transaction binding the contract method 0xf8f5a544.
//
// Solidity: function updateAuthorizedMinter(address newMinter) returns()
func (_Skavenge *SkavengeSession) UpdateAuthorizedMinter(newMinter common.Address) (*types.Transaction, error) {
	return _Skavenge.Contract.UpdateAuthorizedMinter(&_Skavenge.TransactOpts, newMinter)
}

// UpdateAuthorizedMinter is a paid mutator transaction binding the contract method 0xf8f5a544.
//
// Solidity: function updateAuthorizedMinter(address newMinter) returns()
func (_Skavenge *SkavengeTransactorSession) UpdateAuthorizedMinter(newMinter common.Address) (*types.Transaction, error) {
	return _Skavenge.Contract.UpdateAuthorizedMinter(&_Skavenge.TransactOpts, newMinter)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xb142b4ec.
//
// Solidity: function verifyProof(bytes32 transferId) returns()
func (_Skavenge *SkavengeTransactor) VerifyProof(opts *bind.TransactOpts, transferId [32]byte) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "verifyProof", transferId)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xb142b4ec.
//
// Solidity: function verifyProof(bytes32 transferId) returns()
func (_Skavenge *SkavengeSession) VerifyProof(transferId [32]byte) (*types.Transaction, error) {
	return _Skavenge.Contract.VerifyProof(&_Skavenge.TransactOpts, transferId)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xb142b4ec.
//
// Solidity: function verifyProof(bytes32 transferId) returns()
func (_Skavenge *SkavengeTransactorSession) VerifyProof(transferId [32]byte) (*types.Transaction, error) {
	return _Skavenge.Contract.VerifyProof(&_Skavenge.TransactOpts, transferId)
}

// SkavengeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Skavenge contract.
type SkavengeApprovalIterator struct {
	Event *SkavengeApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeApproval represents a Approval event raised by the Skavenge contract.
type SkavengeApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SkavengeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeApprovalIterator{contract: _Skavenge.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SkavengeApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeApproval)
				if err := _Skavenge.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) ParseApproval(log types.Log) (*SkavengeApproval, error) {
	event := new(SkavengeApproval)
	if err := _Skavenge.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Skavenge contract.
type SkavengeApprovalForAllIterator struct {
	Event *SkavengeApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeApprovalForAll represents a ApprovalForAll event raised by the Skavenge contract.
type SkavengeApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Skavenge *SkavengeFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SkavengeApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeApprovalForAllIterator{contract: _Skavenge.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Skavenge *SkavengeFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SkavengeApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeApprovalForAll)
				if err := _Skavenge.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Skavenge *SkavengeFilterer) ParseApprovalForAll(log types.Log) (*SkavengeApprovalForAll, error) {
	event := new(SkavengeApprovalForAll)
	if err := _Skavenge.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeAuthorizedMinterUpdatedIterator is returned from FilterAuthorizedMinterUpdated and is used to iterate over the raw logs and unpacked data for AuthorizedMinterUpdated events raised by the Skavenge contract.
type SkavengeAuthorizedMinterUpdatedIterator struct {
	Event *SkavengeAuthorizedMinterUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeAuthorizedMinterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeAuthorizedMinterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeAuthorizedMinterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeAuthorizedMinterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeAuthorizedMinterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeAuthorizedMinterUpdated represents a AuthorizedMinterUpdated event raised by the Skavenge contract.
type SkavengeAuthorizedMinterUpdated struct {
	OldMinter common.Address
	NewMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedMinterUpdated is a free log retrieval operation binding the contract event 0x808ec13129987deb49ec337ab895a2cf7af16a4d0d55a51ddc054e2c7fb2515b.
//
// Solidity: event AuthorizedMinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_Skavenge *SkavengeFilterer) FilterAuthorizedMinterUpdated(opts *bind.FilterOpts, oldMinter []common.Address, newMinter []common.Address) (*SkavengeAuthorizedMinterUpdatedIterator, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "AuthorizedMinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeAuthorizedMinterUpdatedIterator{contract: _Skavenge.contract, event: "AuthorizedMinterUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorizedMinterUpdated is a free log subscription operation binding the contract event 0x808ec13129987deb49ec337ab895a2cf7af16a4d0d55a51ddc054e2c7fb2515b.
//
// Solidity: event AuthorizedMinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_Skavenge *SkavengeFilterer) WatchAuthorizedMinterUpdated(opts *bind.WatchOpts, sink chan<- *SkavengeAuthorizedMinterUpdated, oldMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "AuthorizedMinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeAuthorizedMinterUpdated)
				if err := _Skavenge.contract.UnpackLog(event, "AuthorizedMinterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorizedMinterUpdated is a log parse operation binding the contract event 0x808ec13129987deb49ec337ab895a2cf7af16a4d0d55a51ddc054e2c7fb2515b.
//
// Solidity: event AuthorizedMinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_Skavenge *SkavengeFilterer) ParseAuthorizedMinterUpdated(log types.Log) (*SkavengeAuthorizedMinterUpdated, error) {
	event := new(SkavengeAuthorizedMinterUpdated)
	if err := _Skavenge.contract.UnpackLog(event, "AuthorizedMinterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeClueAttemptFailedIterator is returned from FilterClueAttemptFailed and is used to iterate over the raw logs and unpacked data for ClueAttemptFailed events raised by the Skavenge contract.
type SkavengeClueAttemptFailedIterator struct {
	Event *SkavengeClueAttemptFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeClueAttemptFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeClueAttemptFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeClueAttemptFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeClueAttemptFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeClueAttemptFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeClueAttemptFailed represents a ClueAttemptFailed event raised by the Skavenge contract.
type SkavengeClueAttemptFailed struct {
	TokenId           *big.Int
	AttemptedSolution string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClueAttemptFailed is a free log retrieval operation binding the contract event 0x65cc0e9121123eab4b9d9814a9160e5954b2f7ce53d78b9cdbdd055af308b9f5.
//
// Solidity: event ClueAttemptFailed(uint256 indexed tokenId, string attemptedSolution)
func (_Skavenge *SkavengeFilterer) FilterClueAttemptFailed(opts *bind.FilterOpts, tokenId []*big.Int) (*SkavengeClueAttemptFailedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "ClueAttemptFailed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeClueAttemptFailedIterator{contract: _Skavenge.contract, event: "ClueAttemptFailed", logs: logs, sub: sub}, nil
}

// WatchClueAttemptFailed is a free log subscription operation binding the contract event 0x65cc0e9121123eab4b9d9814a9160e5954b2f7ce53d78b9cdbdd055af308b9f5.
//
// Solidity: event ClueAttemptFailed(uint256 indexed tokenId, string attemptedSolution)
func (_Skavenge *SkavengeFilterer) WatchClueAttemptFailed(opts *bind.WatchOpts, sink chan<- *SkavengeClueAttemptFailed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "ClueAttemptFailed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeClueAttemptFailed)
				if err := _Skavenge.contract.UnpackLog(event, "ClueAttemptFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClueAttemptFailed is a log parse operation binding the contract event 0x65cc0e9121123eab4b9d9814a9160e5954b2f7ce53d78b9cdbdd055af308b9f5.
//
// Solidity: event ClueAttemptFailed(uint256 indexed tokenId, string attemptedSolution)
func (_Skavenge *SkavengeFilterer) ParseClueAttemptFailed(log types.Log) (*SkavengeClueAttemptFailed, error) {
	event := new(SkavengeClueAttemptFailed)
	if err := _Skavenge.contract.UnpackLog(event, "ClueAttemptFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeClueMintedIterator is returned from FilterClueMinted and is used to iterate over the raw logs and unpacked data for ClueMinted events raised by the Skavenge contract.
type SkavengeClueMintedIterator struct {
	Event *SkavengeClueMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeClueMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeClueMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeClueMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeClueMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeClueMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeClueMinted represents a ClueMinted event raised by the Skavenge contract.
type SkavengeClueMinted struct {
	TokenId *big.Int
	Minter  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClueMinted is a free log retrieval operation binding the contract event 0xa90e59f66e7533243b5959b6498caf4949957dbf8ccaa6b6534177c10041ea54.
//
// Solidity: event ClueMinted(uint256 indexed tokenId, address minter)
func (_Skavenge *SkavengeFilterer) FilterClueMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*SkavengeClueMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "ClueMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeClueMintedIterator{contract: _Skavenge.contract, event: "ClueMinted", logs: logs, sub: sub}, nil
}

// WatchClueMinted is a free log subscription operation binding the contract event 0xa90e59f66e7533243b5959b6498caf4949957dbf8ccaa6b6534177c10041ea54.
//
// Solidity: event ClueMinted(uint256 indexed tokenId, address minter)
func (_Skavenge *SkavengeFilterer) WatchClueMinted(opts *bind.WatchOpts, sink chan<- *SkavengeClueMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "ClueMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeClueMinted)
				if err := _Skavenge.contract.UnpackLog(event, "ClueMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClueMinted is a log parse operation binding the contract event 0xa90e59f66e7533243b5959b6498caf4949957dbf8ccaa6b6534177c10041ea54.
//
// Solidity: event ClueMinted(uint256 indexed tokenId, address minter)
func (_Skavenge *SkavengeFilterer) ParseClueMinted(log types.Log) (*SkavengeClueMinted, error) {
	event := new(SkavengeClueMinted)
	if err := _Skavenge.contract.UnpackLog(event, "ClueMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeClueSolvedIterator is returned from FilterClueSolved and is used to iterate over the raw logs and unpacked data for ClueSolved events raised by the Skavenge contract.
type SkavengeClueSolvedIterator struct {
	Event *SkavengeClueSolved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeClueSolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeClueSolved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeClueSolved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeClueSolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeClueSolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeClueSolved represents a ClueSolved event raised by the Skavenge contract.
type SkavengeClueSolved struct {
	TokenId  *big.Int
	Solution string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClueSolved is a free log retrieval operation binding the contract event 0x3138eb607d845be3efb1a7ea147da7816c8a05f683313c459e6bf953ea4f988e.
//
// Solidity: event ClueSolved(uint256 indexed tokenId, string solution)
func (_Skavenge *SkavengeFilterer) FilterClueSolved(opts *bind.FilterOpts, tokenId []*big.Int) (*SkavengeClueSolvedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "ClueSolved", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeClueSolvedIterator{contract: _Skavenge.contract, event: "ClueSolved", logs: logs, sub: sub}, nil
}

// WatchClueSolved is a free log subscription operation binding the contract event 0x3138eb607d845be3efb1a7ea147da7816c8a05f683313c459e6bf953ea4f988e.
//
// Solidity: event ClueSolved(uint256 indexed tokenId, string solution)
func (_Skavenge *SkavengeFilterer) WatchClueSolved(opts *bind.WatchOpts, sink chan<- *SkavengeClueSolved, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "ClueSolved", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeClueSolved)
				if err := _Skavenge.contract.UnpackLog(event, "ClueSolved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClueSolved is a log parse operation binding the contract event 0x3138eb607d845be3efb1a7ea147da7816c8a05f683313c459e6bf953ea4f988e.
//
// Solidity: event ClueSolved(uint256 indexed tokenId, string solution)
func (_Skavenge *SkavengeFilterer) ParseClueSolved(log types.Log) (*SkavengeClueSolved, error) {
	event := new(SkavengeClueSolved)
	if err := _Skavenge.contract.UnpackLog(event, "ClueSolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeProofProvidedIterator is returned from FilterProofProvided and is used to iterate over the raw logs and unpacked data for ProofProvided events raised by the Skavenge contract.
type SkavengeProofProvidedIterator struct {
	Event *SkavengeProofProvided // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeProofProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeProofProvided)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeProofProvided)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeProofProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeProofProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeProofProvided represents a ProofProvided event raised by the Skavenge contract.
type SkavengeProofProvided struct {
	TransferId  [32]byte
	Proof       []byte
	NewClueHash [32]byte
	RValueHash  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProofProvided is a free log retrieval operation binding the contract event 0x319414a72bfc3d93a989d08f1055fd74a1b953a652be46d0dff852ac157c12f2.
//
// Solidity: event ProofProvided(bytes32 indexed transferId, bytes proof, bytes32 newClueHash, bytes32 rValueHash)
func (_Skavenge *SkavengeFilterer) FilterProofProvided(opts *bind.FilterOpts, transferId [][32]byte) (*SkavengeProofProvidedIterator, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "ProofProvided", transferIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeProofProvidedIterator{contract: _Skavenge.contract, event: "ProofProvided", logs: logs, sub: sub}, nil
}

// WatchProofProvided is a free log subscription operation binding the contract event 0x319414a72bfc3d93a989d08f1055fd74a1b953a652be46d0dff852ac157c12f2.
//
// Solidity: event ProofProvided(bytes32 indexed transferId, bytes proof, bytes32 newClueHash, bytes32 rValueHash)
func (_Skavenge *SkavengeFilterer) WatchProofProvided(opts *bind.WatchOpts, sink chan<- *SkavengeProofProvided, transferId [][32]byte) (event.Subscription, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "ProofProvided", transferIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeProofProvided)
				if err := _Skavenge.contract.UnpackLog(event, "ProofProvided", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProofProvided is a log parse operation binding the contract event 0x319414a72bfc3d93a989d08f1055fd74a1b953a652be46d0dff852ac157c12f2.
//
// Solidity: event ProofProvided(bytes32 indexed transferId, bytes proof, bytes32 newClueHash, bytes32 rValueHash)
func (_Skavenge *SkavengeFilterer) ParseProofProvided(log types.Log) (*SkavengeProofProvided, error) {
	event := new(SkavengeProofProvided)
	if err := _Skavenge.contract.UnpackLog(event, "ProofProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeProofVerifiedIterator is returned from FilterProofVerified and is used to iterate over the raw logs and unpacked data for ProofVerified events raised by the Skavenge contract.
type SkavengeProofVerifiedIterator struct {
	Event *SkavengeProofVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeProofVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeProofVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeProofVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeProofVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeProofVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeProofVerified represents a ProofVerified event raised by the Skavenge contract.
type SkavengeProofVerified struct {
	TransferId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProofVerified is a free log retrieval operation binding the contract event 0x543093db8d78fd8619586d3a0be12a5736836393feede0888f262888c81ce4c3.
//
// Solidity: event ProofVerified(bytes32 indexed transferId)
func (_Skavenge *SkavengeFilterer) FilterProofVerified(opts *bind.FilterOpts, transferId [][32]byte) (*SkavengeProofVerifiedIterator, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "ProofVerified", transferIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeProofVerifiedIterator{contract: _Skavenge.contract, event: "ProofVerified", logs: logs, sub: sub}, nil
}

// WatchProofVerified is a free log subscription operation binding the contract event 0x543093db8d78fd8619586d3a0be12a5736836393feede0888f262888c81ce4c3.
//
// Solidity: event ProofVerified(bytes32 indexed transferId)
func (_Skavenge *SkavengeFilterer) WatchProofVerified(opts *bind.WatchOpts, sink chan<- *SkavengeProofVerified, transferId [][32]byte) (event.Subscription, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "ProofVerified", transferIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeProofVerified)
				if err := _Skavenge.contract.UnpackLog(event, "ProofVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProofVerified is a log parse operation binding the contract event 0x543093db8d78fd8619586d3a0be12a5736836393feede0888f262888c81ce4c3.
//
// Solidity: event ProofVerified(bytes32 indexed transferId)
func (_Skavenge *SkavengeFilterer) ParseProofVerified(log types.Log) (*SkavengeProofVerified, error) {
	event := new(SkavengeProofVerified)
	if err := _Skavenge.contract.UnpackLog(event, "ProofVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeSalePriceRemovedIterator is returned from FilterSalePriceRemoved and is used to iterate over the raw logs and unpacked data for SalePriceRemoved events raised by the Skavenge contract.
type SkavengeSalePriceRemovedIterator struct {
	Event *SkavengeSalePriceRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeSalePriceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeSalePriceRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeSalePriceRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeSalePriceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeSalePriceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeSalePriceRemoved represents a SalePriceRemoved event raised by the Skavenge contract.
type SkavengeSalePriceRemoved struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSalePriceRemoved is a free log retrieval operation binding the contract event 0x06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f268989.
//
// Solidity: event SalePriceRemoved(uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) FilterSalePriceRemoved(opts *bind.FilterOpts, tokenId []*big.Int) (*SkavengeSalePriceRemovedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "SalePriceRemoved", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeSalePriceRemovedIterator{contract: _Skavenge.contract, event: "SalePriceRemoved", logs: logs, sub: sub}, nil
}

// WatchSalePriceRemoved is a free log subscription operation binding the contract event 0x06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f268989.
//
// Solidity: event SalePriceRemoved(uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) WatchSalePriceRemoved(opts *bind.WatchOpts, sink chan<- *SkavengeSalePriceRemoved, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "SalePriceRemoved", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeSalePriceRemoved)
				if err := _Skavenge.contract.UnpackLog(event, "SalePriceRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSalePriceRemoved is a log parse operation binding the contract event 0x06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f268989.
//
// Solidity: event SalePriceRemoved(uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) ParseSalePriceRemoved(log types.Log) (*SkavengeSalePriceRemoved, error) {
	event := new(SkavengeSalePriceRemoved)
	if err := _Skavenge.contract.UnpackLog(event, "SalePriceRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeSalePriceSetIterator is returned from FilterSalePriceSet and is used to iterate over the raw logs and unpacked data for SalePriceSet events raised by the Skavenge contract.
type SkavengeSalePriceSetIterator struct {
	Event *SkavengeSalePriceSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeSalePriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeSalePriceSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeSalePriceSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeSalePriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeSalePriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeSalePriceSet represents a SalePriceSet event raised by the Skavenge contract.
type SkavengeSalePriceSet struct {
	TokenId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSalePriceSet is a free log retrieval operation binding the contract event 0xe23ea816dce6d7f5c0b85cbd597e7c3b97b2453791152c0b94e5e5c5f314d2f0.
//
// Solidity: event SalePriceSet(uint256 indexed tokenId, uint256 price)
func (_Skavenge *SkavengeFilterer) FilterSalePriceSet(opts *bind.FilterOpts, tokenId []*big.Int) (*SkavengeSalePriceSetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "SalePriceSet", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeSalePriceSetIterator{contract: _Skavenge.contract, event: "SalePriceSet", logs: logs, sub: sub}, nil
}

// WatchSalePriceSet is a free log subscription operation binding the contract event 0xe23ea816dce6d7f5c0b85cbd597e7c3b97b2453791152c0b94e5e5c5f314d2f0.
//
// Solidity: event SalePriceSet(uint256 indexed tokenId, uint256 price)
func (_Skavenge *SkavengeFilterer) WatchSalePriceSet(opts *bind.WatchOpts, sink chan<- *SkavengeSalePriceSet, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "SalePriceSet", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeSalePriceSet)
				if err := _Skavenge.contract.UnpackLog(event, "SalePriceSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSalePriceSet is a log parse operation binding the contract event 0xe23ea816dce6d7f5c0b85cbd597e7c3b97b2453791152c0b94e5e5c5f314d2f0.
//
// Solidity: event SalePriceSet(uint256 indexed tokenId, uint256 price)
func (_Skavenge *SkavengeFilterer) ParseSalePriceSet(log types.Log) (*SkavengeSalePriceSet, error) {
	event := new(SkavengeSalePriceSet)
	if err := _Skavenge.contract.UnpackLog(event, "SalePriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Skavenge contract.
type SkavengeTransferIterator struct {
	Event *SkavengeTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeTransfer represents a Transfer event raised by the Skavenge contract.
type SkavengeTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SkavengeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeTransferIterator{contract: _Skavenge.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SkavengeTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeTransfer)
				if err := _Skavenge.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) ParseTransfer(log types.Log) (*SkavengeTransfer, error) {
	event := new(SkavengeTransfer)
	if err := _Skavenge.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeTransferCancelledIterator is returned from FilterTransferCancelled and is used to iterate over the raw logs and unpacked data for TransferCancelled events raised by the Skavenge contract.
type SkavengeTransferCancelledIterator struct {
	Event *SkavengeTransferCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeTransferCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeTransferCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeTransferCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeTransferCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeTransferCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeTransferCancelled represents a TransferCancelled event raised by the Skavenge contract.
type SkavengeTransferCancelled struct {
	TransferId  [32]byte
	CancelledBy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferCancelled is a free log retrieval operation binding the contract event 0x1ed784ea0b4551753ccb1bbf1711421d8a07aff605d39bb9d770c25943aea485.
//
// Solidity: event TransferCancelled(bytes32 indexed transferId, address indexed cancelledBy)
func (_Skavenge *SkavengeFilterer) FilterTransferCancelled(opts *bind.FilterOpts, transferId [][32]byte, cancelledBy []common.Address) (*SkavengeTransferCancelledIterator, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}
	var cancelledByRule []interface{}
	for _, cancelledByItem := range cancelledBy {
		cancelledByRule = append(cancelledByRule, cancelledByItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "TransferCancelled", transferIdRule, cancelledByRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeTransferCancelledIterator{contract: _Skavenge.contract, event: "TransferCancelled", logs: logs, sub: sub}, nil
}

// WatchTransferCancelled is a free log subscription operation binding the contract event 0x1ed784ea0b4551753ccb1bbf1711421d8a07aff605d39bb9d770c25943aea485.
//
// Solidity: event TransferCancelled(bytes32 indexed transferId, address indexed cancelledBy)
func (_Skavenge *SkavengeFilterer) WatchTransferCancelled(opts *bind.WatchOpts, sink chan<- *SkavengeTransferCancelled, transferId [][32]byte, cancelledBy []common.Address) (event.Subscription, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}
	var cancelledByRule []interface{}
	for _, cancelledByItem := range cancelledBy {
		cancelledByRule = append(cancelledByRule, cancelledByItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "TransferCancelled", transferIdRule, cancelledByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeTransferCancelled)
				if err := _Skavenge.contract.UnpackLog(event, "TransferCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferCancelled is a log parse operation binding the contract event 0x1ed784ea0b4551753ccb1bbf1711421d8a07aff605d39bb9d770c25943aea485.
//
// Solidity: event TransferCancelled(bytes32 indexed transferId, address indexed cancelledBy)
func (_Skavenge *SkavengeFilterer) ParseTransferCancelled(log types.Log) (*SkavengeTransferCancelled, error) {
	event := new(SkavengeTransferCancelled)
	if err := _Skavenge.contract.UnpackLog(event, "TransferCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeTransferCompletedIterator is returned from FilterTransferCompleted and is used to iterate over the raw logs and unpacked data for TransferCompleted events raised by the Skavenge contract.
type SkavengeTransferCompletedIterator struct {
	Event *SkavengeTransferCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeTransferCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeTransferCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeTransferCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeTransferCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeTransferCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeTransferCompleted represents a TransferCompleted event raised by the Skavenge contract.
type SkavengeTransferCompleted struct {
	TransferId [32]byte
	RValue     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferCompleted is a free log retrieval operation binding the contract event 0x062fb96142a4ea35fc5c48049c3a7d7a418829dea520220e03d76440bbe275c0.
//
// Solidity: event TransferCompleted(bytes32 indexed transferId, uint256 rValue)
func (_Skavenge *SkavengeFilterer) FilterTransferCompleted(opts *bind.FilterOpts, transferId [][32]byte) (*SkavengeTransferCompletedIterator, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "TransferCompleted", transferIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeTransferCompletedIterator{contract: _Skavenge.contract, event: "TransferCompleted", logs: logs, sub: sub}, nil
}

// WatchTransferCompleted is a free log subscription operation binding the contract event 0x062fb96142a4ea35fc5c48049c3a7d7a418829dea520220e03d76440bbe275c0.
//
// Solidity: event TransferCompleted(bytes32 indexed transferId, uint256 rValue)
func (_Skavenge *SkavengeFilterer) WatchTransferCompleted(opts *bind.WatchOpts, sink chan<- *SkavengeTransferCompleted, transferId [][32]byte) (event.Subscription, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "TransferCompleted", transferIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeTransferCompleted)
				if err := _Skavenge.contract.UnpackLog(event, "TransferCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferCompleted is a log parse operation binding the contract event 0x062fb96142a4ea35fc5c48049c3a7d7a418829dea520220e03d76440bbe275c0.
//
// Solidity: event TransferCompleted(bytes32 indexed transferId, uint256 rValue)
func (_Skavenge *SkavengeFilterer) ParseTransferCompleted(log types.Log) (*SkavengeTransferCompleted, error) {
	event := new(SkavengeTransferCompleted)
	if err := _Skavenge.contract.UnpackLog(event, "TransferCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SkavengeTransferInitiatedIterator is returned from FilterTransferInitiated and is used to iterate over the raw logs and unpacked data for TransferInitiated events raised by the Skavenge contract.
type SkavengeTransferInitiatedIterator struct {
	Event *SkavengeTransferInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SkavengeTransferInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeTransferInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SkavengeTransferInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SkavengeTransferInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeTransferInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeTransferInitiated represents a TransferInitiated event raised by the Skavenge contract.
type SkavengeTransferInitiated struct {
	TransferId [32]byte
	Buyer      common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferInitiated is a free log retrieval operation binding the contract event 0x2d18295f817f7e46b8d3401af48ee043761aba21f602005110a282939c3c4c72.
//
// Solidity: event TransferInitiated(bytes32 indexed transferId, address indexed buyer, uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) FilterTransferInitiated(opts *bind.FilterOpts, transferId [][32]byte, buyer []common.Address, tokenId []*big.Int) (*SkavengeTransferInitiatedIterator, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "TransferInitiated", transferIdRule, buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeTransferInitiatedIterator{contract: _Skavenge.contract, event: "TransferInitiated", logs: logs, sub: sub}, nil
}

// WatchTransferInitiated is a free log subscription operation binding the contract event 0x2d18295f817f7e46b8d3401af48ee043761aba21f602005110a282939c3c4c72.
//
// Solidity: event TransferInitiated(bytes32 indexed transferId, address indexed buyer, uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) WatchTransferInitiated(opts *bind.WatchOpts, sink chan<- *SkavengeTransferInitiated, transferId [][32]byte, buyer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "TransferInitiated", transferIdRule, buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeTransferInitiated)
				if err := _Skavenge.contract.UnpackLog(event, "TransferInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferInitiated is a log parse operation binding the contract event 0x2d18295f817f7e46b8d3401af48ee043761aba21f602005110a282939c3c4c72.
//
// Solidity: event TransferInitiated(bytes32 indexed transferId, address indexed buyer, uint256 indexed tokenId)
func (_Skavenge *SkavengeFilterer) ParseTransferInitiated(log types.Log) (*SkavengeTransferInitiated, error) {
	event := new(SkavengeTransferInitiated)
	if err := _Skavenge.contract.UnpackLog(event, "TransferInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
