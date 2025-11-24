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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialMinter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ClueNotForSale\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolvedClueCannotBeSold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolvedClueTransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferAlreadyInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedMinterUpdate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMinter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"AuthorizedMinterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingAttempts\",\"type\":\"uint256\"}],\"name\":\"ClueAttempted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"ClueMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"solution\",\"type\":\"string\"}],\"name\":\"ClueSolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rValueHash\",\"type\":\"bytes32\"}],\"name\":\"ProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"ProofVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SalePriceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SalePriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"TransferCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"name\":\"TransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TransferInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SOLVE_ATTEMPTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_TIMEOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"solution\",\"type\":\"string\"}],\"name\":\"attemptSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorizedMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"cancelTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"clues\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"solutionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isSolved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"solveAttempts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cluesForSale\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"newEncryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"name\":\"completeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"generateTransferId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getClueContents\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getCluesForSale\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"solvedStatus\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalCluesForSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"initiatePurchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"solutionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"name\":\"mintClue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"}],\"name\":\"provideProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initiatedAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rValueHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"proofVerified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"proofProvidedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"updateAuthorizedMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516167bf3803806167bf83398181016040528101906100319190610172565b6040518060400160405280600881526020017f536b6176656e67650000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f534b564700000000000000000000000000000000000000000000000000000000815250815f90816100ab91906103da565b5080600190816100bb91906103da565b5050506001600a819055506001600b8190555080600c5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506104a9565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61014182610118565b9050919050565b61015181610137565b811461015b575f5ffd5b50565b5f8151905061016c81610148565b92915050565b5f6020828403121561018757610186610114565b5b5f6101948482850161015e565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061021857607f821691505b60208210810361022b5761022a6101d4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261028d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610252565b6102978683610252565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102db6102d66102d1846102af565b6102b8565b6102af565b9050919050565b5f819050919050565b6102f4836102c1565b610308610300826102e2565b84845461025e565b825550505050565b5f5f905090565b61031f610310565b61032a8184846102eb565b505050565b5b8181101561034d576103425f82610317565b600181019050610330565b5050565b601f8211156103925761036381610231565b61036c84610243565b8101602085101561037b578190505b61038f61038785610243565b83018261032f565b50505b505050565b5f82821c905092915050565b5f6103b25f1984600802610397565b1980831691505092915050565b5f6103ca83836103a3565b9150826002028217905092915050565b6103e38261019d565b67ffffffffffffffff8111156103fc576103fb6101a7565b5b6104068254610201565b610411828285610351565b5f60209050601f831160018114610442575f8415610430578287015190505b61043a85826103bf565b8655506104a1565b601f19841661045086610231565b5f5b8281101561047757848901518255600182019150602085019450602081019050610452565b868310156104945784890151610490601f8916826103a3565b8355505b6001600288020188555050505b505050505050565b616309806104b65f395ff3fe60806040526004361061022f575f3560e01c806387065deb1161012d578063c2d554ae116100aa578063e985e9c51161006e578063e985e9c5146108ad578063eb927a83146108e9578063f12b72ba14610925578063f8f5a54414610964578063fae5380c1461098c5761022f565b8063c2d554ae146107c7578063c87b56dd146107ef578063d19310d91461082b578063d32d579014610855578063dd142be0146108915761022f565b8063aff202b4116100f1578063aff202b4146106eb578063b142b4ec14610713578063b329bf5c1461073b578063b40b7eb014610763578063b88d4fde1461079f5761022f565b806387065deb1461060b5780638d7cf3e41461063557806395d89b411461065d578063a22cb46514610687578063a6cd5ff5146106af5761022f565b806330f37c7f116101bb5780634f6ccce71161017f5780634f6ccce714610503578063561892361461053f5780636352211e1461056957806370a08231146105a557806374b19a07146105e15761022f565b806330f37c7f146103dd5780633427ee941461041e57806334499fff1461045a5780633c64f04b1461049657806342842e0e146104db5761022f565b8063095ea7b311610202578063095ea7b3146102fd57806318160ddd146103255780631ba538cd1461034f57806323b872dd146103795780632f745c59146103a15761022f565b806301ffc9a714610233578063053992c51461026f57806306fdde0314610297578063081812fc146102c1575b5f5ffd5b34801561023e575f5ffd5b50610259600480360381019061025491906146dc565b6109b4565b6040516102669190614721565b60405180910390f35b34801561027a575f5ffd5b506102956004803603810190610290919061476d565b610a2d565b005b3480156102a2575f5ffd5b506102ab610c0c565b6040516102b8919061481b565b60405180910390f35b3480156102cc575f5ffd5b506102e760048036038101906102e2919061483b565b610c9b565b6040516102f491906148a5565b60405180910390f35b348015610308575f5ffd5b50610323600480360381019061031e91906148e8565b610cb6565b005b348015610330575f5ffd5b50610339610ccc565b6040516103469190614935565b60405180910390f35b34801561035a575f5ffd5b50610363610cd8565b60405161037091906148a5565b60405180910390f35b348015610384575f5ffd5b5061039f600480360381019061039a919061494e565b610cfd565b005b3480156103ac575f5ffd5b506103c760048036038101906103c291906148e8565b610dfc565b6040516103d49190614935565b60405180910390f35b3480156103e8575f5ffd5b5061040360048036038101906103fe919061483b565b610ea0565b60405161041596959493929190614a08565b60405180910390f35b348015610429575f5ffd5b50610444600480360381019061043f919061483b565b610f6a565b6040516104519190614721565b60405180910390f35b348015610465575f5ffd5b50610480600480360381019061047b919061483b565b610f87565b60405161048d9190614721565b60405180910390f35b3480156104a1575f5ffd5b506104bc60048036038101906104b79190614a98565b610fa4565b6040516104d29a99989796959493929190614ac3565b60405180910390f35b3480156104e6575f5ffd5b5061050160048036038101906104fc919061494e565b6110a5565b005b34801561050e575f5ffd5b506105296004803603810190610524919061483b565b6110c4565b6040516105369190614935565b60405180910390f35b34801561054a575f5ffd5b50610553611136565b6040516105609190614935565b60405180910390f35b348015610574575f5ffd5b5061058f600480360381019061058a919061483b565b61113f565b60405161059c91906148a5565b60405180910390f35b3480156105b0575f5ffd5b506105cb60048036038101906105c69190614b64565b611150565b6040516105d89190614935565b60405180910390f35b3480156105ec575f5ffd5b506105f5611206565b6040516106029190614935565b60405180910390f35b348015610616575f5ffd5b5061061f611212565b60405161062c9190614935565b60405180910390f35b348015610640575f5ffd5b5061065b6004803603810190610656919061483b565b611217565b005b348015610668575f5ffd5b50610671611303565b60405161067e919061481b565b60405180910390f35b348015610692575f5ffd5b506106ad60048036038101906106a89190614bb9565b611393565b005b3480156106ba575f5ffd5b506106d560048036038101906106d091906148e8565b6113a9565b6040516106e29190614bf7565b60405180910390f35b3480156106f6575f5ffd5b50610711600480360381019061070c9190614c71565b6113db565b005b34801561071e575f5ffd5b5061073960048036038101906107349190614a98565b61169b565b005b348015610746575f5ffd5b50610761600480360381019061075c9190614a98565b61183e565b005b34801561076e575f5ffd5b5061078960048036038101906107849190614d23565b611c90565b6040516107969190614935565b60405180910390f35b3480156107aa575f5ffd5b506107c560048036038101906107c09190614ebc565b611e5a565b005b3480156107d2575f5ffd5b506107ed60048036038101906107e89190614f3c565b611e7f565b005b3480156107fa575f5ffd5b506108156004803603810190610810919061483b565b6120eb565b604051610822919061481b565b60405180910390f35b348015610836575f5ffd5b5061083f612151565b60405161084c9190614935565b60405180910390f35b348015610860575f5ffd5b5061087b6004803603810190610876919061483b565b612156565b6040516108889190614935565b60405180910390f35b6108ab60048036038101906108a6919061483b565b6121e9565b005b3480156108b8575f5ffd5b506108d360048036038101906108ce9190614fad565b612614565b6040516108e09190614721565b60405180910390f35b3480156108f4575f5ffd5b5061090f600480360381019061090a919061483b565b6126a2565b60405161091c9190614feb565b60405180910390f35b348015610930575f5ffd5b5061094b6004803603810190610946919061476d565b61274f565b60405161095b9493929190615230565b60405180910390f35b34801561096f575f5ffd5b5061098a60048036038101906109859190614b64565b612a12565b005b348015610997575f5ffd5b506109b260048036038101906109ad919061528f565b612b5b565b005b5f7f780e9d63000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610a265750610a258261310f565b5b9050919050565b3373ffffffffffffffffffffffffffffffffffffffff16610a4d8361113f565b73ffffffffffffffffffffffffffffffffffffffff1614610aa3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9a9061534a565b60405180910390fd5b600d5f8381526020019081526020015f206002015f9054906101000a900460ff1615610afb576040517fff1e4dda00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600d5f8481526020019081526020015f2060040181905550600f5f8381526020019081526020015f205f9054906101000a900460ff16158015610b3e57505f81115b15610b97576001600f5f8481526020019081526020015f205f6101000a81548160ff021916908315150217905550601082908060018154018082558091505060019003905f5260205f20015f9091909190915055610bd0565b600f5f8381526020019081526020015f205f9054906101000a900460ff168015610bc057505f81145b15610bcf57610bce826131f0565b5b5b817fe23ea816dce6d7f5c0b85cbd597e7c3b97b2453791152c0b94e5e5c5f314d2f082604051610c009190614935565b60405180910390a25050565b60605f8054610c1a90615395565b80601f0160208091040260200160405190810160405280929190818152602001828054610c4690615395565b8015610c915780601f10610c6857610100808354040283529160200191610c91565b820191905f5260205f20905b815481529060010190602001808311610c7457829003601f168201915b5050505050905090565b5f610ca5826132fa565b50610caf82613380565b9050919050565b610cc88282610cc36133b9565b6133c0565b5050565b5f600880549050905090565b600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610d6d575f6040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610d6491906148a5565b60405180910390fd5b5f610d808383610d7b6133b9565b6133d2565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610df6578382826040517f64283d7b000000000000000000000000000000000000000000000000000000008152600401610ded939291906153c5565b60405180910390fd5b50505050565b5f610e0683611150565b8210610e4b5782826040517fa57d13dc000000000000000000000000000000000000000000000000000000008152600401610e429291906153fa565b60405180910390fd5b60065f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f2054905092915050565b600d602052805f5260405f205f91509050805f018054610ebf90615395565b80601f0160208091040260200160405190810160405280929190818152602001828054610eeb90615395565b8015610f365780601f10610f0d57610100808354040283529160200191610f36565b820191905f5260205f20905b815481529060010190602001808311610f1957829003601f168201915b505050505090806001015490806002015f9054906101000a900460ff16908060030154908060040154908060050154905086565b600f602052805f5260405f205f915054906101000a900460ff1681565b6011602052805f5260405f205f915054906101000a900460ff1681565b600e602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002015490806003015490806004018054610ffa90615395565b80601f016020809104026020016040519081016040528092919081815260200182805461102690615395565b80156110715780601f1061104857610100808354040283529160200191611071565b820191905f5260205f20905b81548152906001019060200180831161105457829003601f168201915b505050505090806005015490806006015490806007015f9054906101000a900460ff1690806008015490806009015490508a565b6110bf83838360405180602001604052805f815250611e5a565b505050565b5f6110cd610ccc565b8210611112575f826040517fa57d13dc0000000000000000000000000000000000000000000000000000000081526004016111099291906153fa565b60405180910390fd5b6008828154811061112657611125615421565b5b905f5260205f2001549050919050565b5f600b54905090565b5f611149826132fa565b9050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036111c1575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016111b891906148a5565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f601080549050905090565b60b481565b3373ffffffffffffffffffffffffffffffffffffffff166112378261113f565b73ffffffffffffffffffffffffffffffffffffffff161461128d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112849061534a565b60405180910390fd5b600f5f8281526020019081526020015f205f9054906101000a900460ff16156112ba576112b9816131f0565b5b5f600d5f8381526020019081526020015f2060040181905550807f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a250565b60606001805461131290615395565b80601f016020809104026020016040519081016040528092919081815260200182805461133e90615395565b80156113895780601f1061136057610100808354040283529160200191611389565b820191905f5260205f20905b81548152906001019060200180831161136c57829003601f168201915b5050505050905090565b6113a561139e6133b9565b83836134a0565b5050565b5f82826040516020016113bd9291906154b3565b60405160208183030381529060405280519060200120905092915050565b3373ffffffffffffffffffffffffffffffffffffffff166113fb8461113f565b73ffffffffffffffffffffffffffffffffffffffff1614611451576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114489061534a565b60405180910390fd5b600d5f8481526020019081526020015f206002015f9054906101000a900460ff16156114b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114a990615528565b60405180910390fd5b6003600d5f8581526020019081526020015f206003015410611509576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150090615590565b60405180910390fd5b600d5f8481526020019081526020015f206003015f81548092919061152d906155db565b9190505550827fa5d0a58799728745ca0e2b91a8e1b764e373f058529afb509f23b4b00a454fbe600d5f8681526020019081526020015f206003015460036115759190615622565b6040516115829190614935565b60405180910390a2600d5f8481526020019081526020015f206001015482826040516115af929190615683565b604051809103902003611696576001600d5f8581526020019081526020015f206002015f6101000a81548160ff021916908315150217905550600f5f8481526020019081526020015f205f9054906101000a900460ff161561165b57611614836131f0565b5f600d5f8581526020019081526020015f2060040181905550827f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a25b827f3138eb607d845be3efb1a7ea147da7816c8a05f683313c459e6bf953ea4f988e838360405161168d9291906156c7565b60405180910390a25b505050565b6116a3613609565b5f600e5f8381526020019081526020015f2090503373ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611747576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161173e90615733565b60405180910390fd5b5f81600801541161178d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117849061579b565b60405180910390fd5b60b481600801544261179f9190615622565b11156117e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117d790615803565b60405180910390fd5b6001816007015f6101000a81548160ff021916908315150217905550428160090181905550817f543093db8d78fd8619586d3a0be12a5736836393feede0888f262888c81ce4c360405160405180910390a25061183b61364f565b50565b611846613609565b5f600e5f8381526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036118ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118e19061586b565b60405180910390fd5b5f3373ffffffffffffffffffffffffffffffffffffffff16825f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161490505f3373ffffffffffffffffffffffffffffffffffffffff16611963846001015461113f565b73ffffffffffffffffffffffffffffffffffffffff161490505f5f905082156119d1575f8460090154146119cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119c3906158f9565b60405180910390fd5b600190505b8115611a80575f84600801541480156119f9575060b48460030154426119f79190615622565b115b15611a075760019050611a7f565b5f8460080154118015611a285750836007015f9054906101000a900460ff16155b8015611a43575060b4846008015442611a419190615622565b115b15611a515760019050611a7e565b5f8460090154118015611a73575060b4846009015442611a719190615622565b115b15611a7d57600190505b5b5b5b80611ac0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ab790615961565b60405180910390fd5b5f84600201541115611b9c575f845f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168560020154604051611b17906159a2565b5f6040518083038185875af1925050503d805f8114611b51576040519150601f19603f3d011682016040523d82523d5f602084013e611b56565b606091505b5050905080611b9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b9190615a00565b60405180910390fd5b505b5f60115f866001015481526020019081526020015f205f6101000a81548160ff021916908315150217905550847f2e936050b1807500251bb54605979b74ee4e0e31a0fcba9f12b51d99496c20fa60405160405180910390a2600e5f8681526020019081526020015f205f5f82015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f611c50919061461e565b600582015f9055600682015f9055600782015f6101000a81549060ff0219169055600882015f9055600982015f9055505050505050611c8d61364f565b50565b5f600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d17576040517f955c501b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b5f815480929190611d29906155db565b9190505590506040518060c0016040528086868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018481526020015f151581526020015f81526020015f815260200183815250600d5f8381526020019081526020015f205f820151815f019081611dc59190615bbe565b50602082015181600101556040820151816002015f6101000a81548160ff021916908315150217905550606082015181600301556080820151816004015560a08201518160050155905050611e1a3382613659565b807fa90e59f66e7533243b5959b6498caf4949957dbf8ccaa6b6534177c10041ea5433604051611e4a91906148a5565b60405180910390a2949350505050565b611e65848484610cfd565b611e79611e706133b9565b8585858561374c565b50505050565b611e87613609565b5f600e5f8681526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611f2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f229061586b565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16611f4f826001015461113f565b73ffffffffffffffffffffffffffffffffffffffff1614611fa5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f9c9061534a565b60405180910390fd5b60b4816003015442611fb79190615622565b1115611ff8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fef90615cd7565b60405180910390fd5b602084849050101561203f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161203690615d3f565b60405180910390fd5b5f84846020878790506120529190615622565b90809261206193929190615d65565b9061206c9190615da9565b90508484836004019182612081929190615e07565b50828260050181905550808260060181905550428260080181905550857f319414a72bfc3d93a989d08f1055fd74a1b953a652be46d0dff852ac157c12f2868686856040516120d39493929190615f00565b60405180910390a250506120e561364f565b50505050565b60606120f6826132fa565b505f6121006138f8565b90505f81511161211e5760405180602001604052805f815250612149565b806121288461390e565b604051602001612139929190615f78565b6040516020818303038152906040525b915050919050565b600381565b5f3373ffffffffffffffffffffffffffffffffffffffff166121778361113f565b73ffffffffffffffffffffffffffffffffffffffff16146121cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121c49061534a565b60405180910390fd5b600d5f8381526020019081526020015f20600501549050919050565b6121f1613609565b6121fa8161113f565b50600f5f8281526020019081526020015f205f9054906101000a900460ff16158061223857505f600d5f8381526020019081526020015f2060040154145b1561226f576040517fa7d67ebb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5f8281526020019081526020015f20600401543410156122bd576040517f356680b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5f8281526020019081526020015f206002015f9054906101000a900460ff1615612315576040517f6e40ff0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60115f8281526020019081526020015f205f9054906101000a900460ff161561236a576040517f74ed79ae00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61237533836113a9565b90505f73ffffffffffffffffffffffffffffffffffffffff16600e5f8381526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612417576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161240e90615fe5565b60405180910390fd5b600160115f8481526020019081526020015f205f6101000a81548160ff0219169083151502179055506040518061014001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018381526020013481526020014281526020015f67ffffffffffffffff81111561249457612493614d98565b5b6040519080825280601f01601f1916602001820160405280156124c65781602001600182028036833780820191505090505b5081526020015f5f1b81526020015f5f1b81526020015f151581526020015f81526020015f815250600e5f8381526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003015560808201518160040190816125769190615bbe565b5060a0820151816005015560c0820151816006015560e0820151816007015f6101000a81548160ff02191690831515021790555061010082015181600801556101208201518160090155905050813373ffffffffffffffffffffffffffffffffffffffff16827f2d18295f817f7e46b8d3401af48ee043761aba21f602005110a282939c3c4c7260405160405180910390a45061261161364f565b50565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60606126ad8261113f565b50600d5f8381526020019081526020015f205f0180546126cc90615395565b80601f01602080910402602001604051908101604052809291908181526020018280546126f890615395565b80156127435780601f1061271a57610100808354040283529160200191612743565b820191905f5260205f20905b81548152906001019060200180831161272657829003601f168201915b50505050509050919050565b6060806060805f60108054905090505f81878961276c9190616003565b1161278257868861277d9190616003565b612784565b815b90505f888211612794575f6127a1565b88826127a09190615622565b5b90508067ffffffffffffffff8111156127bd576127bc614d98565b5b6040519080825280602002602001820160405280156127eb5781602001602082028036833780820191505090505b5096508067ffffffffffffffff81111561280857612807614d98565b5b6040519080825280602002602001820160405280156128365781602001602082028036833780820191505090505b5095508067ffffffffffffffff81111561285357612852614d98565b5b6040519080825280602002602001820160405280156128815781602001602082028036833780820191505090505b5094508067ffffffffffffffff81111561289e5761289d614d98565b5b6040519080825280602002602001820160405280156128cc5781602001602082028036833780820191505090505b5093505f5f90505b81811015612a05575f6010828c6128eb9190616003565b815481106128fc576128fb615421565b5b905f5260205f20015490508089838151811061291b5761291a615421565b5b6020026020010181815250506129308161113f565b88838151811061294357612942615421565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600d5f8281526020019081526020015f20600401548783815181106129a5576129a4615421565b5b602002602001018181525050600d5f8281526020019081526020015f206002015f9054906101000a900460ff168683815181106129e5576129e4615421565b5b6020026020010190151590811515815250505080806001019150506128d4565b5050505092959194509250565b600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612a98576040517f7efb568f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600c5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f808ec13129987deb49ec337ab895a2cf7af16a4d0d55a51ddc054e2c7fb2515b60405160405180910390a35050565b612b63613609565b5f600e5f8681526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612c07576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bfe9061586b565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16612c2b826001015461113f565b73ffffffffffffffffffffffffffffffffffffffff1614612c81576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c789061534a565b60405180910390fd5b5f816009015411612cc7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cbe90616080565b60405180910390fd5b60b4816009015442612cd99190615622565b1115612d1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d11906160e8565b60405180910390fd5b80600501548484604051612d2f929190615683565b604051809103902014612d77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d6e90616150565b60405180910390fd5b806006015482604051602001612d8d919061616e565b6040516020818303038152906040528051906020012014612de3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dda906161d2565b60405180910390fd5b600d5f826001015481526020019081526020015f206002015f9054906101000a900460ff1615612e3f576040517f6e40ff0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8383600d5f846001015481526020019081526020015f205f019182612e65929190615e07565b5081600d5f836001015481526020019081526020015f20600501819055505f600d5f836001015481526020019081526020015f2060030181905550612ee033825f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836001015460405180602001604052805f8152506139d8565b5f3373ffffffffffffffffffffffffffffffffffffffff168260020154604051612f09906159a2565b5f6040518083038185875af1925050503d805f8114612f43576040519150601f19603f3d011682016040523d82523d5f602084013e612f48565b606091505b5050905080612f8c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f8390615a00565b60405180910390fd5b600f5f836001015481526020019081526020015f205f9054906101000a900460ff161561300f57612fc082600101546131f0565b5f600d5f846001015481526020019081526020015f206004018190555081600101547f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a25b5f60115f846001015481526020019081526020015f205f6101000a81548160ff021916908315150217905550600e5f8781526020019081526020015f205f5f82015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f613096919061461e565b600582015f9055600682015f9055600782015f6101000a81549060ff0219169055600882015f9055600982015f90555050857f062fb96142a4ea35fc5c48049c3a7d7a418829dea520220e03d76440bbe275c0846040516130f79190614935565b60405180910390a2505061310961364f565b50505050565b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806131d957507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806131e957506131e8826139fd565b5b9050919050565b600f5f8281526020019081526020015f205f9054906101000a900460ff16156132f7575f600f5f8381526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f90505b6010805490508110156132f557816010828154811061326257613261615421565b5b905f5260205f200154036132e857601060016010805490506132849190615622565b8154811061329557613294615421565b5b905f5260205f200154601082815481106132b2576132b1615421565b5b905f5260205f20018190555060108054806132d0576132cf6161f0565b5b600190038181905f5260205f20015f905590556132f5565b8080600101915050613240565b505b50565b5f5f61330583613a66565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361337757826040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161336e9190614935565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b6133cd8383836001613a9f565b505050565b5f5f6133df858585613c5e565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561344957505f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614155b1561349557600f5f8581526020019081526020015f205f9054906101000a900460ff161561347b5761347a846131f0565b5b5f600d5f8681526020019081526020015f20600401819055505b809150509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361351057816040517f5b08ba1800000000000000000000000000000000000000000000000000000000815260040161350791906148a5565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516135fc9190614721565b60405180910390a3505050565b6002600a5403613645576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600a81905550565b6001600a81905550565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036136c9575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016136c091906148a5565b60405180910390fd5b5f6136d583835f6133d2565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614613747575f6040517f73c6ac6e00000000000000000000000000000000000000000000000000000000815260040161373e91906148a5565b60405180910390fd5b505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b11156138f1578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b81526004016137aa949392919061621d565b6020604051808303815f875af19250505080156137e557506040513d601f19601f820116820180604052508101906137e2919061627b565b60015b613866573d805f8114613813576040519150601f19603f3d011682016040523d82523d5f602084013e613818565b606091505b505f81510361385e57836040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161385591906148a5565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146138ef57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016138e691906148a5565b60405180910390fd5b505b5050505050565b606060405180602001604052805f815250905090565b60605f600161391c84613d78565b0190505f8167ffffffffffffffff81111561393a57613939614d98565b5b6040519080825280601f01601f19166020018201604052801561396c5781602001600182028036833780820191505090505b5090505f82602001820190505b6001156139cd578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85816139c2576139c16162a6565b5b0494505f8503613979575b819350505050919050565b6139e3848484613ec9565b6139f76139ee6133b9565b8585858561374c565b50505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b8080613ad757505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15613c09575f613ae6846132fa565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015613b5057508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b8015613b635750613b618184612614565b155b15613ba557826040517fa9fbf51f000000000000000000000000000000000000000000000000000000008152600401613b9c91906148a5565b60405180910390fd5b8115613c0757838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b5f5f613c6b858585614031565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613cae57613ca98461423c565b613ced565b8473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614613cec57613ceb8185614280565b5b5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603613d2e57613d2984614357565b613d6d565b8473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614613d6c57613d6b8585614417565b5b5b809150509392505050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310613dd4577a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008381613dca57613dc96162a6565b5b0492506040810190505b6d04ee2d6d415b85acef81000000008310613e11576d04ee2d6d415b85acef81000000008381613e0757613e066162a6565b5b0492506020810190505b662386f26fc100008310613e4057662386f26fc100008381613e3657613e356162a6565b5b0492506010810190505b6305f5e1008310613e69576305f5e1008381613e5f57613e5e6162a6565b5b0492506008810190505b6127108310613e8e576127108381613e8457613e836162a6565b5b0492506004810190505b60648310613eb15760648381613ea757613ea66162a6565b5b0492506002810190505b600a8310613ec0576001810190505b80915050919050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603613f39575f6040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401613f3091906148a5565b60405180910390fd5b5f613f4583835f6133d2565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613fb757816040517f7e273289000000000000000000000000000000000000000000000000000000008152600401613fae9190614935565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461402b578382826040517f64283d7b000000000000000000000000000000000000000000000000000000008152600401614022939291906153c5565b60405180910390fd5b50505050565b5f5f61403c84613a66565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161461407d5761407c81848661449b565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614614108576140bc5f855f5f613a9f565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461418757600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b60088054905060095f8381526020019081526020015f2081905550600881908060018154018082558091505060019003905f5260205f20015f909190919091505550565b5f61428a83611150565b90505f60075f8481526020019081526020015f205490505f60065f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050828214614329575f815f8581526020019081526020015f2054905080825f8581526020019081526020015f20819055508260075f8381526020019081526020015f2081905550505b60075f8581526020019081526020015f205f9055805f8481526020019081526020015f205f90555050505050565b5f600160088054905061436a9190615622565b90505f60095f8481526020019081526020015f205490505f6008838154811061439657614395615421565b5b905f5260205f200154905080600883815481106143b6576143b5615421565b5b905f5260205f2001819055508160095f8381526020019081526020015f208190555060095f8581526020019081526020015f205f905560088054806143fe576143fd6161f0565b5b600190038181905f5260205f20015f9055905550505050565b5f600161442384611150565b61442d9190615622565b90508160065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f20819055508060075f8481526020019081526020015f2081905550505050565b6144a683838361455e565b614559575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361451a57806040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016145119190614935565b60405180910390fd5b81816040517f177e802f0000000000000000000000000000000000000000000000000000000081526004016145509291906153fa565b60405180910390fd5b505050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561461557508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806145d657506145d58484612614565b5b8061461457508273ffffffffffffffffffffffffffffffffffffffff166145fc83613380565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b50805461462a90615395565b5f825580601f1061463b5750614658565b601f0160209004905f5260205f2090810190614657919061465b565b5b50565b5b80821115614672575f815f90555060010161465c565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6146bb81614687565b81146146c5575f5ffd5b50565b5f813590506146d6816146b2565b92915050565b5f602082840312156146f1576146f061467f565b5b5f6146fe848285016146c8565b91505092915050565b5f8115159050919050565b61471b81614707565b82525050565b5f6020820190506147345f830184614712565b92915050565b5f819050919050565b61474c8161473a565b8114614756575f5ffd5b50565b5f8135905061476781614743565b92915050565b5f5f604083850312156147835761478261467f565b5b5f61479085828601614759565b92505060206147a185828601614759565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6147ed826147ab565b6147f781856147b5565b93506148078185602086016147c5565b614810816147d3565b840191505092915050565b5f6020820190508181035f83015261483381846147e3565b905092915050565b5f602082840312156148505761484f61467f565b5b5f61485d84828501614759565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61488f82614866565b9050919050565b61489f81614885565b82525050565b5f6020820190506148b85f830184614896565b92915050565b6148c781614885565b81146148d1575f5ffd5b50565b5f813590506148e2816148be565b92915050565b5f5f604083850312156148fe576148fd61467f565b5b5f61490b858286016148d4565b925050602061491c85828601614759565b9150509250929050565b61492f8161473a565b82525050565b5f6020820190506149485f830184614926565b92915050565b5f5f5f606084860312156149655761496461467f565b5b5f614972868287016148d4565b9350506020614983868287016148d4565b925050604061499486828701614759565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f6149c28261499e565b6149cc81856149a8565b93506149dc8185602086016147c5565b6149e5816147d3565b840191505092915050565b5f819050919050565b614a02816149f0565b82525050565b5f60c0820190508181035f830152614a2081896149b8565b9050614a2f60208301886149f9565b614a3c6040830187614712565b614a496060830186614926565b614a566080830185614926565b614a6360a0830184614926565b979650505050505050565b614a77816149f0565b8114614a81575f5ffd5b50565b5f81359050614a9281614a6e565b92915050565b5f60208284031215614aad57614aac61467f565b5b5f614aba84828501614a84565b91505092915050565b5f61014082019050614ad75f83018d614896565b614ae4602083018c614926565b614af1604083018b614926565b614afe606083018a614926565b8181036080830152614b1081896149b8565b9050614b1f60a08301886149f9565b614b2c60c08301876149f9565b614b3960e0830186614712565b614b47610100830185614926565b614b55610120830184614926565b9b9a5050505050505050505050565b5f60208284031215614b7957614b7861467f565b5b5f614b86848285016148d4565b91505092915050565b614b9881614707565b8114614ba2575f5ffd5b50565b5f81359050614bb381614b8f565b92915050565b5f5f60408385031215614bcf57614bce61467f565b5b5f614bdc858286016148d4565b9250506020614bed85828601614ba5565b9150509250929050565b5f602082019050614c0a5f8301846149f9565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112614c3157614c30614c10565b5b8235905067ffffffffffffffff811115614c4e57614c4d614c14565b5b602083019150836001820283011115614c6a57614c69614c18565b5b9250929050565b5f5f5f60408486031215614c8857614c8761467f565b5b5f614c9586828701614759565b935050602084013567ffffffffffffffff811115614cb657614cb5614683565b5b614cc286828701614c1c565b92509250509250925092565b5f5f83601f840112614ce357614ce2614c10565b5b8235905067ffffffffffffffff811115614d0057614cff614c14565b5b602083019150836001820283011115614d1c57614d1b614c18565b5b9250929050565b5f5f5f5f60608587031215614d3b57614d3a61467f565b5b5f85013567ffffffffffffffff811115614d5857614d57614683565b5b614d6487828801614cce565b94509450506020614d7787828801614a84565b9250506040614d8887828801614759565b91505092959194509250565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b614dce826147d3565b810181811067ffffffffffffffff82111715614ded57614dec614d98565b5b80604052505050565b5f614dff614676565b9050614e0b8282614dc5565b919050565b5f67ffffffffffffffff821115614e2a57614e29614d98565b5b614e33826147d3565b9050602081019050919050565b828183375f83830152505050565b5f614e60614e5b84614e10565b614df6565b905082815260208101848484011115614e7c57614e7b614d94565b5b614e87848285614e40565b509392505050565b5f82601f830112614ea357614ea2614c10565b5b8135614eb3848260208601614e4e565b91505092915050565b5f5f5f5f60808587031215614ed457614ed361467f565b5b5f614ee1878288016148d4565b9450506020614ef2878288016148d4565b9350506040614f0387828801614759565b925050606085013567ffffffffffffffff811115614f2457614f23614683565b5b614f3087828801614e8f565b91505092959194509250565b5f5f5f5f60608587031215614f5457614f5361467f565b5b5f614f6187828801614a84565b945050602085013567ffffffffffffffff811115614f8257614f81614683565b5b614f8e87828801614cce565b93509350506040614fa187828801614a84565b91505092959194509250565b5f5f60408385031215614fc357614fc261467f565b5b5f614fd0858286016148d4565b9250506020614fe1858286016148d4565b9150509250929050565b5f6020820190508181035f83015261500381846149b8565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61503d8161473a565b82525050565b5f61504e8383615034565b60208301905092915050565b5f602082019050919050565b5f6150708261500b565b61507a8185615015565b935061508583615025565b805f5b838110156150b557815161509c8882615043565b97506150a78361505a565b925050600181019050615088565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6150f481614885565b82525050565b5f61510583836150eb565b60208301905092915050565b5f602082019050919050565b5f615127826150c2565b61513181856150cc565b935061513c836150dc565b805f5b8381101561516c57815161515388826150fa565b975061515e83615111565b92505060018101905061513f565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6151ab81614707565b82525050565b5f6151bc83836151a2565b60208301905092915050565b5f602082019050919050565b5f6151de82615179565b6151e88185615183565b93506151f383615193565b805f5b8381101561522357815161520a88826151b1565b9750615215836151c8565b9250506001810190506151f6565b5085935050505092915050565b5f6080820190508181035f8301526152488187615066565b9050818103602083015261525c818661511d565b905081810360408301526152708185615066565b9050818103606083015261528481846151d4565b905095945050505050565b5f5f5f5f606085870312156152a7576152a661467f565b5b5f6152b487828801614a84565b945050602085013567ffffffffffffffff8111156152d5576152d4614683565b5b6152e187828801614cce565b935093505060406152f487828801614759565b91505092959194509250565b7f4e6f7420746f6b656e206f776e657200000000000000000000000000000000005f82015250565b5f615334600f836147b5565b915061533f82615300565b602082019050919050565b5f6020820190508181035f83015261536181615328565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806153ac57607f821691505b6020821081036153bf576153be615368565b5b50919050565b5f6060820190506153d85f830186614896565b6153e56020830185614926565b6153f26040830184614896565b949350505050565b5f60408201905061540d5f830185614896565b61541a6020830184614926565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f8160601b9050919050565b5f6154648261544e565b9050919050565b5f6154758261545a565b9050919050565b61548d61548882614885565b61546b565b82525050565b5f819050919050565b6154ad6154a88261473a565b615493565b82525050565b5f6154be828561547c565b6014820191506154ce828461549c565b6020820191508190509392505050565b7f436c756520616c726561647920736f6c766564000000000000000000000000005f82015250565b5f6155126013836147b5565b915061551d826154de565b602082019050919050565b5f6020820190508181035f83015261553f81615506565b9050919050565b7f4e6f20617474656d7074732072656d61696e696e6700000000000000000000005f82015250565b5f61557a6015836147b5565b915061558582615546565b602082019050919050565b5f6020820190508181035f8301526155a78161556e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6155e58261473a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615617576156166155ae565b5b600182019050919050565b5f61562c8261473a565b91506156378361473a565b925082820390508181111561564f5761564e6155ae565b5b92915050565b5f81905092915050565b5f61566a8385615655565b9350615677838584614e40565b82840190509392505050565b5f61568f82848661565f565b91508190509392505050565b5f6156a683856147b5565b93506156b3838584614e40565b6156bc836147d3565b840190509392505050565b5f6020820190508181035f8301526156e081848661569b565b90509392505050565b7f4e6f7420746865206275796572000000000000000000000000000000000000005f82015250565b5f61571d600d836147b5565b9150615728826156e9565b602082019050919050565b5f6020820190508181035f83015261574a81615711565b9050919050565b7f50726f6f66206e6f74207965742070726f7669646564000000000000000000005f82015250565b5f6157856016836147b5565b915061579082615751565b602082019050919050565b5f6020820190508181035f8301526157b281615779565b9050919050565b7f50726f6f6620766572696669636174696f6e20657870697265640000000000005f82015250565b5f6157ed601a836147b5565b91506157f8826157b9565b602082019050919050565b5f6020820190508181035f83015261581a816157e1565b9050919050565b7f5472616e7366657220646f6573206e6f742065786973740000000000000000005f82015250565b5f6158556017836147b5565b915061586082615821565b602082019050919050565b5f6020820190508181035f83015261588281615849565b9050919050565b7f43616e6e6f742063616e63656c2061667465722070726f6f66207665726966695f8201527f636174696f6e0000000000000000000000000000000000000000000000000000602082015250565b5f6158e36026836147b5565b91506158ee82615889565b604082019050919050565b5f6020820190508181035f830152615910816158d7565b9050919050565b7f4e6f7420617574686f72697a656420746f2063616e63656c00000000000000005f82015250565b5f61594b6018836147b5565b915061595682615917565b602082019050919050565b5f6020820190508181035f8301526159788161593f565b9050919050565b50565b5f61598d5f83615655565b91506159988261597f565b5f82019050919050565b5f6159ac82615982565b9150819050919050565b7f4661696c656420746f2073656e642045746865720000000000000000000000005f82015250565b5f6159ea6014836147b5565b91506159f5826159b6565b602082019050919050565b5f6020820190508181035f830152615a17816159de565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302615a7a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82615a3f565b615a848683615a3f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f615abf615aba615ab58461473a565b615a9c565b61473a565b9050919050565b5f819050919050565b615ad883615aa5565b615aec615ae482615ac6565b848454615a4b565b825550505050565b5f5f905090565b615b03615af4565b615b0e818484615acf565b505050565b5b81811015615b3157615b265f82615afb565b600181019050615b14565b5050565b601f821115615b7657615b4781615a1e565b615b5084615a30565b81016020851015615b5f578190505b615b73615b6b85615a30565b830182615b13565b50505b505050565b5f82821c905092915050565b5f615b965f1984600802615b7b565b1980831691505092915050565b5f615bae8383615b87565b9150826002028217905092915050565b615bc78261499e565b67ffffffffffffffff811115615be057615bdf614d98565b5b615bea8254615395565b615bf5828285615b35565b5f60209050601f831160018114615c26575f8415615c14578287015190505b615c1e8582615ba3565b865550615c85565b601f198416615c3486615a1e565b5f5b82811015615c5b57848901518255600182019150602085019450602081019050615c36565b86831015615c785784890151615c74601f891682615b87565b8355505b6001600288020188555050505b505050505050565b7f5472616e736665722065787069726564000000000000000000000000000000005f82015250565b5f615cc16010836147b5565b9150615ccc82615c8d565b602082019050919050565b5f6020820190508181035f830152615cee81615cb5565b9050919050565b7f50726f6f6620746f6f2073686f727400000000000000000000000000000000005f82015250565b5f615d29600f836147b5565b9150615d3482615cf5565b602082019050919050565b5f6020820190508181035f830152615d5681615d1d565b9050919050565b5f5ffd5b5f5ffd5b5f5f85851115615d7857615d77615d5d565b5b83861115615d8957615d88615d61565b5b6001850283019150848603905094509492505050565b5f82905092915050565b5f615db48383615d9f565b82615dbf81356149f0565b92506020821015615dff57615dfa7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83602003600802615a3f565b831692505b505092915050565b615e118383615d9f565b67ffffffffffffffff811115615e2a57615e29614d98565b5b615e348254615395565b615e3f828285615b35565b5f601f831160018114615e6c575f8415615e5a578287013590505b615e648582615ba3565b865550615ecb565b601f198416615e7a86615a1e565b5f5b82811015615ea157848901358255600182019150602085019450602081019050615e7c565b86831015615ebe5784890135615eba601f891682615b87565b8355505b6001600288020188555050505b50505050505050565b5f615edf83856149a8565b9350615eec838584614e40565b615ef5836147d3565b840190509392505050565b5f6060820190508181035f830152615f19818688615ed4565b9050615f2860208301856149f9565b615f3560408301846149f9565b95945050505050565b5f81905092915050565b5f615f52826147ab565b615f5c8185615f3e565b9350615f6c8185602086016147c5565b80840191505092915050565b5f615f838285615f48565b9150615f8f8284615f48565b91508190509392505050565b7f5472616e7366657220616c726561647920696e697469617465640000000000005f82015250565b5f615fcf601a836147b5565b9150615fda82615f9b565b602082019050919050565b5f6020820190508181035f830152615ffc81615fc3565b9050919050565b5f61600d8261473a565b91506160188361473a565b92508282019050808211156160305761602f6155ae565b5b92915050565b7f50726f6f66206e6f7420766572696669656400000000000000000000000000005f82015250565b5f61606a6012836147b5565b915061607582616036565b602082019050919050565b5f6020820190508181035f8301526160978161605e565b9050919050565b7f5472616e7366657220636f6d706c6574696f6e206578706972656400000000005f82015250565b5f6160d2601b836147b5565b91506160dd8261609e565b602082019050919050565b5f6020820190508181035f8301526160ff816160c6565b9050919050565b7f436f6e74656e742068617368206d69736d6174636800000000000000000000005f82015250565b5f61613a6015836147b5565b915061614582616106565b602082019050919050565b5f6020820190508181035f8301526161678161612e565b9050919050565b5f616179828461549c565b60208201915081905092915050565b7f522076616c75652068617368206d69736d6174636800000000000000000000005f82015250565b5f6161bc6015836147b5565b91506161c782616188565b602082019050919050565b5f6020820190508181035f8301526161e9816161b0565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f6080820190506162305f830187614896565b61623d6020830186614896565b61624a6040830185614926565b818103606083015261625c81846149b8565b905095945050505050565b5f81519050616275816146b2565b92915050565b5f602082840312156162905761628f61467f565b5b5f61629d84828501616267565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffdfea26469706673582212206366d95b5990645be027eb6ea478794a78f586ee24a137573f52648abb77fd4a64736f6c634300081c0033",
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

// MAXSOLVEATTEMPTS is a free data retrieval call binding the contract method 0xd19310d9.
//
// Solidity: function MAX_SOLVE_ATTEMPTS() view returns(uint256)
func (_Skavenge *SkavengeCaller) MAXSOLVEATTEMPTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "MAX_SOLVE_ATTEMPTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSOLVEATTEMPTS is a free data retrieval call binding the contract method 0xd19310d9.
//
// Solidity: function MAX_SOLVE_ATTEMPTS() view returns(uint256)
func (_Skavenge *SkavengeSession) MAXSOLVEATTEMPTS() (*big.Int, error) {
	return _Skavenge.Contract.MAXSOLVEATTEMPTS(&_Skavenge.CallOpts)
}

// MAXSOLVEATTEMPTS is a free data retrieval call binding the contract method 0xd19310d9.
//
// Solidity: function MAX_SOLVE_ATTEMPTS() view returns(uint256)
func (_Skavenge *SkavengeCallerSession) MAXSOLVEATTEMPTS() (*big.Int, error) {
	return _Skavenge.Contract.MAXSOLVEATTEMPTS(&_Skavenge.CallOpts)
}

// TRANSFERTIMEOUT is a free data retrieval call binding the contract method 0x87065deb.
//
// Solidity: function TRANSFER_TIMEOUT() view returns(uint256)
func (_Skavenge *SkavengeCaller) TRANSFERTIMEOUT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "TRANSFER_TIMEOUT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TRANSFERTIMEOUT is a free data retrieval call binding the contract method 0x87065deb.
//
// Solidity: function TRANSFER_TIMEOUT() view returns(uint256)
func (_Skavenge *SkavengeSession) TRANSFERTIMEOUT() (*big.Int, error) {
	return _Skavenge.Contract.TRANSFERTIMEOUT(&_Skavenge.CallOpts)
}

// TRANSFERTIMEOUT is a free data retrieval call binding the contract method 0x87065deb.
//
// Solidity: function TRANSFER_TIMEOUT() view returns(uint256)
func (_Skavenge *SkavengeCallerSession) TRANSFERTIMEOUT() (*big.Int, error) {
	return _Skavenge.Contract.TRANSFERTIMEOUT(&_Skavenge.CallOpts)
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
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 solveAttempts, uint256 salePrice, uint256 rValue)
func (_Skavenge *SkavengeCaller) Clues(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
	SolveAttempts     *big.Int
	SalePrice         *big.Int
	RValue            *big.Int
}, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "clues", arg0)

	outstruct := new(struct {
		EncryptedContents []byte
		SolutionHash      [32]byte
		IsSolved          bool
		SolveAttempts     *big.Int
		SalePrice         *big.Int
		RValue            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EncryptedContents = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.SolutionHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.IsSolved = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.SolveAttempts = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SalePrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RValue = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Clues is a free data retrieval call binding the contract method 0x30f37c7f.
//
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 solveAttempts, uint256 salePrice, uint256 rValue)
func (_Skavenge *SkavengeSession) Clues(arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
	SolveAttempts     *big.Int
	SalePrice         *big.Int
	RValue            *big.Int
}, error) {
	return _Skavenge.Contract.Clues(&_Skavenge.CallOpts, arg0)
}

// Clues is a free data retrieval call binding the contract method 0x30f37c7f.
//
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 solveAttempts, uint256 salePrice, uint256 rValue)
func (_Skavenge *SkavengeCallerSession) Clues(arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
	SolveAttempts     *big.Int
	SalePrice         *big.Int
	RValue            *big.Int
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

// MintClue is a paid mutator transaction binding the contract method 0xb40b7eb0.
//
// Solidity: function mintClue(bytes encryptedContents, bytes32 solutionHash, uint256 rValue) returns(uint256 tokenId)
func (_Skavenge *SkavengeTransactor) MintClue(opts *bind.TransactOpts, encryptedContents []byte, solutionHash [32]byte, rValue *big.Int) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "mintClue", encryptedContents, solutionHash, rValue)
}

// MintClue is a paid mutator transaction binding the contract method 0xb40b7eb0.
//
// Solidity: function mintClue(bytes encryptedContents, bytes32 solutionHash, uint256 rValue) returns(uint256 tokenId)
func (_Skavenge *SkavengeSession) MintClue(encryptedContents []byte, solutionHash [32]byte, rValue *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.MintClue(&_Skavenge.TransactOpts, encryptedContents, solutionHash, rValue)
}

// MintClue is a paid mutator transaction binding the contract method 0xb40b7eb0.
//
// Solidity: function mintClue(bytes encryptedContents, bytes32 solutionHash, uint256 rValue) returns(uint256 tokenId)
func (_Skavenge *SkavengeTransactorSession) MintClue(encryptedContents []byte, solutionHash [32]byte, rValue *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.MintClue(&_Skavenge.TransactOpts, encryptedContents, solutionHash, rValue)
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

// SetSalePrice is a paid mutator transaction binding the contract method 0x053992c5.
//
// Solidity: function setSalePrice(uint256 tokenId, uint256 price) returns()
func (_Skavenge *SkavengeTransactor) SetSalePrice(opts *bind.TransactOpts, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Skavenge.contract.Transact(opts, "setSalePrice", tokenId, price)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0x053992c5.
//
// Solidity: function setSalePrice(uint256 tokenId, uint256 price) returns()
func (_Skavenge *SkavengeSession) SetSalePrice(tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.SetSalePrice(&_Skavenge.TransactOpts, tokenId, price)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0x053992c5.
//
// Solidity: function setSalePrice(uint256 tokenId, uint256 price) returns()
func (_Skavenge *SkavengeTransactorSession) SetSalePrice(tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Skavenge.Contract.SetSalePrice(&_Skavenge.TransactOpts, tokenId, price)
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

// SkavengeClueAttemptedIterator is returned from FilterClueAttempted and is used to iterate over the raw logs and unpacked data for ClueAttempted events raised by the Skavenge contract.
type SkavengeClueAttemptedIterator struct {
	Event *SkavengeClueAttempted // Event containing the contract specifics and raw log

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
func (it *SkavengeClueAttemptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SkavengeClueAttempted)
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
		it.Event = new(SkavengeClueAttempted)
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
func (it *SkavengeClueAttemptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SkavengeClueAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SkavengeClueAttempted represents a ClueAttempted event raised by the Skavenge contract.
type SkavengeClueAttempted struct {
	TokenId           *big.Int
	RemainingAttempts *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClueAttempted is a free log retrieval operation binding the contract event 0xa5d0a58799728745ca0e2b91a8e1b764e373f058529afb509f23b4b00a454fbe.
//
// Solidity: event ClueAttempted(uint256 indexed tokenId, uint256 remainingAttempts)
func (_Skavenge *SkavengeFilterer) FilterClueAttempted(opts *bind.FilterOpts, tokenId []*big.Int) (*SkavengeClueAttemptedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "ClueAttempted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeClueAttemptedIterator{contract: _Skavenge.contract, event: "ClueAttempted", logs: logs, sub: sub}, nil
}

// WatchClueAttempted is a free log subscription operation binding the contract event 0xa5d0a58799728745ca0e2b91a8e1b764e373f058529afb509f23b4b00a454fbe.
//
// Solidity: event ClueAttempted(uint256 indexed tokenId, uint256 remainingAttempts)
func (_Skavenge *SkavengeFilterer) WatchClueAttempted(opts *bind.WatchOpts, sink chan<- *SkavengeClueAttempted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "ClueAttempted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SkavengeClueAttempted)
				if err := _Skavenge.contract.UnpackLog(event, "ClueAttempted", log); err != nil {
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

// ParseClueAttempted is a log parse operation binding the contract event 0xa5d0a58799728745ca0e2b91a8e1b764e373f058529afb509f23b4b00a454fbe.
//
// Solidity: event ClueAttempted(uint256 indexed tokenId, uint256 remainingAttempts)
func (_Skavenge *SkavengeFilterer) ParseClueAttempted(log types.Log) (*SkavengeClueAttempted, error) {
	event := new(SkavengeClueAttempted)
	if err := _Skavenge.contract.UnpackLog(event, "ClueAttempted", log); err != nil {
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
	TransferId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferCancelled is a free log retrieval operation binding the contract event 0x2e936050b1807500251bb54605979b74ee4e0e31a0fcba9f12b51d99496c20fa.
//
// Solidity: event TransferCancelled(bytes32 indexed transferId)
func (_Skavenge *SkavengeFilterer) FilterTransferCancelled(opts *bind.FilterOpts, transferId [][32]byte) (*SkavengeTransferCancelledIterator, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}

	logs, sub, err := _Skavenge.contract.FilterLogs(opts, "TransferCancelled", transferIdRule)
	if err != nil {
		return nil, err
	}
	return &SkavengeTransferCancelledIterator{contract: _Skavenge.contract, event: "TransferCancelled", logs: logs, sub: sub}, nil
}

// WatchTransferCancelled is a free log subscription operation binding the contract event 0x2e936050b1807500251bb54605979b74ee4e0e31a0fcba9f12b51d99496c20fa.
//
// Solidity: event TransferCancelled(bytes32 indexed transferId)
func (_Skavenge *SkavengeFilterer) WatchTransferCancelled(opts *bind.WatchOpts, sink chan<- *SkavengeTransferCancelled, transferId [][32]byte) (event.Subscription, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}

	logs, sub, err := _Skavenge.contract.WatchLogs(opts, "TransferCancelled", transferIdRule)
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

// ParseTransferCancelled is a log parse operation binding the contract event 0x2e936050b1807500251bb54605979b74ee4e0e31a0fcba9f12b51d99496c20fa.
//
// Solidity: event TransferCancelled(bytes32 indexed transferId)
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
