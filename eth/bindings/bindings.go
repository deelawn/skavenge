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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialMinter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ClueNotForSale\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolvedClueCannotBeSold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolvedClueTransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferAlreadyInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedMinterUpdate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMinter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"AuthorizedMinterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"ClueMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"solution\",\"type\":\"string\"}],\"name\":\"ClueSolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rValueHash\",\"type\":\"bytes32\"}],\"name\":\"ProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"ProofVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SalePriceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SalePriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cancelledBy\",\"type\":\"address\"}],\"name\":\"TransferCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"name\":\"TransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TransferInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TRANSFER_TIMEOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeTransferIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"solution\",\"type\":\"string\"}],\"name\":\"attemptSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorizedMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"cancelTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"clues\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"solutionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isSolved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cluesForSale\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"newEncryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"name\":\"completeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"generateTransferId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getClueContents\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getCluesForSale\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"solvedStatus\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalCluesForSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"initiatePurchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"solutionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"name\":\"mintClue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"}],\"name\":\"provideProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initiatedAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rValueHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"proofVerified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"proofProvidedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"updateAuthorizedMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051616ba8380380616ba883398181016040528101906100319190610172565b6040518060400160405280600881526020017f536b6176656e67650000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f534b564700000000000000000000000000000000000000000000000000000000815250815f90816100ab91906103da565b5080600190816100bb91906103da565b5050506001600a819055506001600b8190555080600c5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506104a9565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61014182610118565b9050919050565b61015181610137565b811461015b575f5ffd5b50565b5f8151905061016c81610148565b92915050565b5f6020828403121561018757610186610114565b5b5f6101948482850161015e565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061021857607f821691505b60208210810361022b5761022a6101d4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261028d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610252565b6102978683610252565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102db6102d66102d1846102af565b6102b8565b6102af565b9050919050565b5f819050919050565b6102f4836102c1565b610308610300826102e2565b84845461025e565b825550505050565b5f5f905090565b61031f610310565b61032a8184846102eb565b505050565b5b8181101561034d576103425f82610317565b600181019050610330565b5050565b601f8211156103925761036381610231565b61036c84610243565b8101602085101561037b578190505b61038f61038785610243565b83018261032f565b50505b505050565b5f82821c905092915050565b5f6103b25f1984600802610397565b1980831691505092915050565b5f6103ca83836103a3565b9150826002028217905092915050565b6103e38261019d565b67ffffffffffffffff8111156103fc576103fb6101a7565b5b6104068254610201565b610411828285610351565b5f60209050601f831160018114610442575f8415610430578287015190505b61043a85826103bf565b8655506104a1565b601f19841661045086610231565b5f5b8281101561047757848901518255600182019150602085019450602081019050610452565b868310156104945784890151610490601f8916826103a3565b8355505b6001600288020188555050505b505050505050565b6166f2806104b65f395ff3fe60806040526004361061022f575f3560e01c806379096ee81161012d578063b88d4fde116100aa578063e985e9c51161006e578063e985e9c5146108be578063eb927a83146108fa578063f12b72ba14610936578063f8f5a54414610975578063fae5380c1461099d5761022f565b8063b88d4fde146107da578063c2d554ae14610802578063c87b56dd1461082a578063d32d579014610866578063dd142be0146108a25761022f565b8063a6cd5ff5116100f1578063a6cd5ff5146106ea578063aff202b414610726578063b142b4ec1461074e578063b329bf5c14610776578063b40b7eb01461079e5761022f565b806379096ee81461060a57806387065deb146106465780638d7cf3e41461067057806395d89b4114610698578063a22cb465146106c25761022f565b806330f37c7f116101bb5780634f6ccce71161017f5780634f6ccce714610502578063561892361461053e5780636352211e1461056857806370a08231146105a457806374b19a07146105e05761022f565b806330f37c7f146103dd5780633427ee941461041d57806334499fff146104595780633c64f04b1461049557806342842e0e146104da5761022f565b8063095ea7b311610202578063095ea7b3146102fd57806318160ddd146103255780631ba538cd1461034f57806323b872dd146103795780632f745c59146103a15761022f565b806301ffc9a714610233578063053992c51461026f57806306fdde0314610297578063081812fc146102c1575b5f5ffd5b34801561023e575f5ffd5b5061025960048036038101906102549190614a56565b6109c5565b6040516102669190614a9b565b60405180910390f35b34801561027a575f5ffd5b5061029560048036038101906102909190614ae7565b610a3e565b005b3480156102a2575f5ffd5b506102ab610c1d565b6040516102b89190614b95565b60405180910390f35b3480156102cc575f5ffd5b506102e760048036038101906102e29190614bb5565b610cac565b6040516102f49190614c1f565b60405180910390f35b348015610308575f5ffd5b50610323600480360381019061031e9190614c62565b610cc7565b005b348015610330575f5ffd5b50610339610cdd565b6040516103469190614caf565b60405180910390f35b34801561035a575f5ffd5b50610363610ce9565b6040516103709190614c1f565b60405180910390f35b348015610384575f5ffd5b5061039f600480360381019061039a9190614cc8565b610d0e565b005b3480156103ac575f5ffd5b506103c760048036038101906103c29190614c62565b610e0d565b6040516103d49190614caf565b60405180910390f35b3480156103e8575f5ffd5b5061040360048036038101906103fe9190614bb5565b610eb1565b604051610414959493929190614d82565b60405180910390f35b348015610428575f5ffd5b50610443600480360381019061043e9190614bb5565b610f75565b6040516104509190614a9b565b60405180910390f35b348015610464575f5ffd5b5061047f600480360381019061047a9190614bb5565b610f92565b60405161048c9190614a9b565b60405180910390f35b3480156104a0575f5ffd5b506104bb60048036038101906104b69190614e04565b610faf565b6040516104d19a99989796959493929190614e2f565b60405180910390f35b3480156104e5575f5ffd5b5061050060048036038101906104fb9190614cc8565b6110b0565b005b34801561050d575f5ffd5b5061052860048036038101906105239190614bb5565b6110cf565b6040516105359190614caf565b60405180910390f35b348015610549575f5ffd5b50610552611141565b60405161055f9190614caf565b60405180910390f35b348015610573575f5ffd5b5061058e60048036038101906105899190614bb5565b61114a565b60405161059b9190614c1f565b60405180910390f35b3480156105af575f5ffd5b506105ca60048036038101906105c59190614ed0565b61115b565b6040516105d79190614caf565b60405180910390f35b3480156105eb575f5ffd5b506105f4611211565b6040516106019190614caf565b60405180910390f35b348015610615575f5ffd5b50610630600480360381019061062b9190614bb5565b61121d565b60405161063d9190614efb565b60405180910390f35b348015610651575f5ffd5b5061065a611232565b6040516106679190614caf565b60405180910390f35b34801561067b575f5ffd5b5061069660048036038101906106919190614bb5565b611237565b005b3480156106a3575f5ffd5b506106ac611323565b6040516106b99190614b95565b60405180910390f35b3480156106cd575f5ffd5b506106e860048036038101906106e39190614f3e565b6113b3565b005b3480156106f5575f5ffd5b50610710600480360381019061070b9190614c62565b6113c9565b60405161071d9190614efb565b60405180910390f35b348015610731575f5ffd5b5061074c60048036038101906107479190614fdd565b6113fb565b005b348015610759575f5ffd5b50610774600480360381019061076f9190614e04565b6115e3565b005b348015610781575f5ffd5b5061079c60048036038101906107979190614e04565b611786565b005b3480156107a9575f5ffd5b506107c460048036038101906107bf919061508f565b611c07565b6040516107d19190614caf565b60405180910390f35b3480156107e5575f5ffd5b5061080060048036038101906107fb9190615228565b611dc1565b005b34801561080d575f5ffd5b50610828600480360381019061082391906152a8565b611de6565b005b348015610835575f5ffd5b50610850600480360381019061084b9190614bb5565b61218b565b60405161085d9190614b95565b60405180910390f35b348015610871575f5ffd5b5061088c60048036038101906108879190614bb5565b6121f1565b6040516108999190614caf565b60405180910390f35b6108bc60048036038101906108b79190614bb5565b61220e565b005b3480156108c9575f5ffd5b506108e460048036038101906108df9190615319565b61264f565b6040516108f19190614a9b565b60405180910390f35b348015610905575f5ffd5b50610920600480360381019061091b9190614bb5565b6126dd565b60405161092d9190615357565b60405180910390f35b348015610941575f5ffd5b5061095c60048036038101906109579190614ae7565b61278a565b60405161096c949392919061559c565b60405180910390f35b348015610980575f5ffd5b5061099b60048036038101906109969190614ed0565b612a4d565b005b3480156109a8575f5ffd5b506109c360048036038101906109be91906155fb565b612b96565b005b5f7f780e9d63000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610a375750610a3682613216565b5b9050919050565b3373ffffffffffffffffffffffffffffffffffffffff16610a5e8361114a565b73ffffffffffffffffffffffffffffffffffffffff1614610ab4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aab906156b6565b60405180910390fd5b600d5f8381526020019081526020015f206002015f9054906101000a900460ff1615610b0c576040517fff1e4dda00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600d5f8481526020019081526020015f2060030181905550600f5f8381526020019081526020015f205f9054906101000a900460ff16158015610b4f57505f81115b15610ba8576001600f5f8481526020019081526020015f205f6101000a81548160ff021916908315150217905550601082908060018154018082558091505060019003905f5260205f20015f9091909190915055610be1565b600f5f8381526020019081526020015f205f9054906101000a900460ff168015610bd157505f81145b15610be057610bdf826132f7565b5b5b817fe23ea816dce6d7f5c0b85cbd597e7c3b97b2453791152c0b94e5e5c5f314d2f082604051610c119190614caf565b60405180910390a25050565b60605f8054610c2b90615701565b80601f0160208091040260200160405190810160405280929190818152602001828054610c5790615701565b8015610ca25780601f10610c7957610100808354040283529160200191610ca2565b820191905f5260205f20905b815481529060010190602001808311610c8557829003601f168201915b5050505050905090565b5f610cb682613674565b50610cc0826136fa565b9050919050565b610cd98282610cd4613733565b61373a565b5050565b5f600880549050905090565b600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610d7e575f6040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610d759190614c1f565b60405180910390fd5b5f610d918383610d8c613733565b61374c565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610e07578382826040517f64283d7b000000000000000000000000000000000000000000000000000000008152600401610dfe93929190615731565b60405180910390fd5b50505050565b5f610e178361115b565b8210610e5c5782826040517fa57d13dc000000000000000000000000000000000000000000000000000000008152600401610e53929190615766565b60405180910390fd5b60065f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f2054905092915050565b600d602052805f5260405f205f91509050805f018054610ed090615701565b80601f0160208091040260200160405190810160405280929190818152602001828054610efc90615701565b8015610f475780601f10610f1e57610100808354040283529160200191610f47565b820191905f5260205f20905b815481529060010190602001808311610f2a57829003601f168201915b505050505090806001015490806002015f9054906101000a900460ff16908060030154908060040154905085565b600f602052805f5260405f205f915054906101000a900460ff1681565b6011602052805f5260405f205f915054906101000a900460ff1681565b600e602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101549080600201549080600301549080600401805461100590615701565b80601f016020809104026020016040519081016040528092919081815260200182805461103190615701565b801561107c5780601f106110535761010080835404028352916020019161107c565b820191905f5260205f20905b81548152906001019060200180831161105f57829003601f168201915b505050505090806005015490806006015490806007015f9054906101000a900460ff1690806008015490806009015490508a565b6110ca83838360405180602001604052805f815250611dc1565b505050565b5f6110d8610cdd565b821061111d575f826040517fa57d13dc000000000000000000000000000000000000000000000000000000008152600401611114929190615766565b60405180910390fd5b600882815481106111315761113061578d565b5b905f5260205f2001549050919050565b5f600b54905090565b5f61115482613674565b9050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036111cc575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016111c39190614c1f565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f601080549050905090565b6012602052805f5260405f205f915090505481565b60b481565b3373ffffffffffffffffffffffffffffffffffffffff166112578261114a565b73ffffffffffffffffffffffffffffffffffffffff16146112ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112a4906156b6565b60405180910390fd5b600f5f8281526020019081526020015f205f9054906101000a900460ff16156112da576112d9816132f7565b5b5f600d5f8381526020019081526020015f2060030181905550807f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a250565b60606001805461133290615701565b80601f016020809104026020016040519081016040528092919081815260200182805461135e90615701565b80156113a95780601f10611380576101008083540402835291602001916113a9565b820191905f5260205f20905b81548152906001019060200180831161138c57829003601f168201915b5050505050905090565b6113c56113be613733565b838361381a565b5050565b5f82826040516020016113dd92919061581f565b60405160208183030381529060405280519060200120905092915050565b3373ffffffffffffffffffffffffffffffffffffffff1661141b8461114a565b73ffffffffffffffffffffffffffffffffffffffff1614611471576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611468906156b6565b60405180910390fd5b600d5f8481526020019081526020015f206002015f9054906101000a900460ff16156114d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114c990615894565b60405180910390fd5b600d5f8481526020019081526020015f206001015482826040516114f79291906158e0565b6040518091039020036115de576001600d5f8581526020019081526020015f206002015f6101000a81548160ff021916908315150217905550600f5f8481526020019081526020015f205f9054906101000a900460ff16156115a35761155c836132f7565b5f600d5f8581526020019081526020015f2060030181905550827f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a25b827f3138eb607d845be3efb1a7ea147da7816c8a05f683313c459e6bf953ea4f988e83836040516115d5929190615924565b60405180910390a25b505050565b6115eb613983565b5f600e5f8381526020019081526020015f2090503373ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461168f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161168690615990565b60405180910390fd5b5f8160080154116116d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116cc906159f8565b60405180910390fd5b60b48160080154426116e79190615a43565b1115611728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161171f90615ac0565b60405180910390fd5b6001816007015f6101000a81548160ff021916908315150217905550428160090181905550817f543093db8d78fd8619586d3a0be12a5736836393feede0888f262888c81ce4c360405160405180910390a2506117836139c9565b50565b61178e613983565b5f600e5f8381526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611832576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161182990615b28565b60405180910390fd5b5f3373ffffffffffffffffffffffffffffffffffffffff16825f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161490505f3373ffffffffffffffffffffffffffffffffffffffff166118ab846001015461114a565b73ffffffffffffffffffffffffffffffffffffffff161490505f5f90508215611919575f846009015414611914576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161190b90615bb6565b60405180910390fd5b600190505b81156119c8575f8460080154148015611941575060b484600301544261193f9190615a43565b115b1561194f57600190506119c7565b5f84600801541180156119705750836007015f9054906101000a900460ff16155b801561198b575060b48460080154426119899190615a43565b115b1561199957600190506119c6565b5f84600901541180156119bb575060b48460090154426119b99190615a43565b115b156119c557600190505b5b5b5b80611a08576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119ff90615c1e565b60405180910390fd5b5f84600201541115611ae4575f845f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168560020154604051611a5f90615c5f565b5f6040518083038185875af1925050503d805f8114611a99576040519150601f19603f3d011682016040523d82523d5f602084013e611a9e565b606091505b5050905080611ae2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ad990615cbd565b60405180910390fd5b505b5f60115f866001015481526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f856001015481526020019081526020015f205f90553373ffffffffffffffffffffffffffffffffffffffff16857f1ed784ea0b4551753ccb1bbf1711421d8a07aff605d39bb9d770c25943aea48560405160405180910390a3600e5f8681526020019081526020015f205f5f82015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f611bc79190614998565b600582015f9055600682015f9055600782015f6101000a81549060ff0219169055600882015f9055600982015f9055505050505050611c046139c9565b50565b5f600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611c8e576040517f955c501b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b5f815480929190611ca090615cdb565b9190505590506040518060a0016040528086868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018481526020015f151581526020015f815260200183815250600d5f8381526020019081526020015f205f820151815f019081611d369190615ec2565b50602082015181600101556040820151816002015f6101000a81548160ff0219169083151502179055506060820151816003015560808201518160040155905050611d8133826139d3565b807fa90e59f66e7533243b5959b6498caf4949957dbf8ccaa6b6534177c10041ea5433604051611db19190614c1f565b60405180910390a2949350505050565b611dcc848484610d0e565b611de0611dd7613733565b85858585613ac6565b50505050565b611dee613983565b5f600e5f8681526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611e92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e8990615b28565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16611eb6826001015461114a565b73ffffffffffffffffffffffffffffffffffffffff1614611f0c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f03906156b6565b60405180910390fd5b60b4816003015442611f1e9190615a43565b1115611f5f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f5690615fdb565b60405180910390fd5b6024848490501015611fa6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f9d90616043565b60405180910390fd5b5f84846003818110611fbb57611fba61578d565b5b9050013560f81c60f81b60f81c60ff16600886866002818110611fe157611fe061578d565b5b9050013560f81c60f81b60f81c60ff1663ffffffff16901b60108787600181811061200f5761200e61578d565b5b9050013560f81c60f81b60f81c60ff1663ffffffff16901b601888885f81811061203c5761203b61578d565b5b9050013560f81c60f81b60f81c60ff1663ffffffff16901b171717905060208160046120689190616070565b6120729190616070565b63ffffffff168585905010156120bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120b4906160f1565b60405180910390fd5b5f858560208460046120cf9190616070565b6120d9919061610f565b63ffffffff16908460046120ed9190616070565b63ffffffff16926121009392919061614e565b9061210b9190616192565b905085858460040191826121209291906161f0565b50838360050181905550808360060181905550428360080181905550867f319414a72bfc3d93a989d08f1055fd74a1b953a652be46d0dff852ac157c12f28787878560405161217294939291906162e9565b60405180910390a25050506121856139c9565b50505050565b606061219682613674565b505f6121a0613c72565b90505f8151116121be5760405180602001604052805f8152506121e9565b806121c884613c88565b6040516020016121d9929190616361565b6040516020818303038152906040525b915050919050565b5f600d5f8381526020019081526020015f20600401549050919050565b612216613983565b61221f8161114a565b50600f5f8281526020019081526020015f205f9054906101000a900460ff16158061225d57505f600d5f8381526020019081526020015f2060030154145b15612294576040517fa7d67ebb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5f8281526020019081526020015f20600301543410156122e2576040517f356680b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5f8281526020019081526020015f206002015f9054906101000a900460ff161561233a576040517f6e40ff0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60115f8281526020019081526020015f205f9054906101000a900460ff161561238f576040517f74ed79ae00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61239a33836113c9565b90505f73ffffffffffffffffffffffffffffffffffffffff16600e5f8381526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461243c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612433906163ce565b60405180910390fd5b600160115f8481526020019081526020015f205f6101000a81548160ff0219169083151502179055508060125f8481526020019081526020015f20819055506040518061014001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018381526020013481526020014281526020015f67ffffffffffffffff8111156124cf576124ce615104565b5b6040519080825280601f01601f1916602001820160405280156125015781602001600182028036833780820191505090505b5081526020015f5f1b81526020015f5f1b81526020015f151581526020015f81526020015f815250600e5f8381526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003015560808201518160040190816125b19190615ec2565b5060a0820151816005015560c0820151816006015560e0820151816007015f6101000a81548160ff02191690831515021790555061010082015181600801556101208201518160090155905050813373ffffffffffffffffffffffffffffffffffffffff16827f2d18295f817f7e46b8d3401af48ee043761aba21f602005110a282939c3c4c7260405160405180910390a45061264c6139c9565b50565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60606126e88261114a565b50600d5f8381526020019081526020015f205f01805461270790615701565b80601f016020809104026020016040519081016040528092919081815260200182805461273390615701565b801561277e5780601f106127555761010080835404028352916020019161277e565b820191905f5260205f20905b81548152906001019060200180831161276157829003601f168201915b50505050509050919050565b6060806060805f60108054905090505f8187896127a791906163ec565b116127bd5786886127b891906163ec565b6127bf565b815b90505f8882116127cf575f6127dc565b88826127db9190615a43565b5b90508067ffffffffffffffff8111156127f8576127f7615104565b5b6040519080825280602002602001820160405280156128265781602001602082028036833780820191505090505b5096508067ffffffffffffffff81111561284357612842615104565b5b6040519080825280602002602001820160405280156128715781602001602082028036833780820191505090505b5095508067ffffffffffffffff81111561288e5761288d615104565b5b6040519080825280602002602001820160405280156128bc5781602001602082028036833780820191505090505b5094508067ffffffffffffffff8111156128d9576128d8615104565b5b6040519080825280602002602001820160405280156129075781602001602082028036833780820191505090505b5093505f5f90505b81811015612a40575f6010828c61292691906163ec565b815481106129375761293661578d565b5b905f5260205f2001549050808983815181106129565761295561578d565b5b60200260200101818152505061296b8161114a565b88838151811061297e5761297d61578d565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600d5f8281526020019081526020015f20600301548783815181106129e0576129df61578d565b5b602002602001018181525050600d5f8281526020019081526020015f206002015f9054906101000a900460ff16868381518110612a2057612a1f61578d565b5b60200260200101901515908115158152505050808060010191505061290f565b5050505092959194509250565b600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612ad3576040517f7efb568f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600c5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f808ec13129987deb49ec337ab895a2cf7af16a4d0d55a51ddc054e2c7fb2515b60405160405180910390a35050565b612b9e613983565b5f600e5f8681526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612c42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c3990615b28565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16612c66826001015461114a565b73ffffffffffffffffffffffffffffffffffffffff1614612cbc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cb3906156b6565b60405180910390fd5b5f816009015411612d02576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cf990616469565b60405180910390fd5b60b4816009015442612d149190615a43565b1115612d55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d4c906164d1565b60405180910390fd5b80600501548484604051612d6a9291906158e0565b604051809103902014612db2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612da990616539565b60405180910390fd5b806006015482604051602001612dc89190616557565b6040516020818303038152906040528051906020012014612e1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e15906165bb565b60405180910390fd5b600d5f826001015481526020019081526020015f206002015f9054906101000a900460ff1615612e7a576040517f6e40ff0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8383600d5f846001015481526020019081526020015f205f019182612ea09291906161f0565b5081600d5f836001015481526020019081526020015f20600401819055505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f826002015490505f83600101549050600f5f8281526020019081526020015f205f9054906101000a900460ff1615613041575f600f5f8381526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f90505b601080549050811015612ff9578160108281548110612f6657612f6561578d565b5b905f5260205f20015403612fec5760106001601080549050612f889190615a43565b81548110612f9957612f9861578d565b5b905f5260205f20015460108281548110612fb657612fb561578d565b5b905f5260205f2001819055506010805480612fd457612fd36165d9565b5b600190038181905f5260205f20015f90559055612ff9565b8080600101915050612f44565b505f600d5f8381526020019081526020015f2060030181905550807f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a25b5f60115f8381526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f8281526020019081526020015f205f9055600e5f8981526020019081526020015f205f5f82015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f6130d89190614998565b600582015f9055600682015f9055600782015f6101000a81549060ff0219169055600882015f9055600982015f9055505061312333848360405180602001604052805f815250613d52565b5f3373ffffffffffffffffffffffffffffffffffffffff168360405161314890615c5f565b5f6040518083038185875af1925050503d805f8114613182576040519150601f19603f3d011682016040523d82523d5f602084013e613187565b606091505b50509050806131cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131c290615cbd565b60405180910390fd5b887f062fb96142a4ea35fc5c48049c3a7d7a418829dea520220e03d76440bbe275c0876040516131fb9190614caf565b60405180910390a250505050506132106139c9565b50505050565b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806132e057507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806132f057506132ef82613d77565b5b9050919050565b600f5f8281526020019081526020015f205f9054906101000a900460ff1615613671575f60125f8381526020019081526020015f205490505f5f1b811461358c575f600e5f8381526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461358a575f8160020154111561347d575f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1682600201546040516133f890615c5f565b5f6040518083038185875af1925050503d805f8114613432576040519150601f19603f3d011682016040523d82523d5f602084013e613437565b606091505b505090508061347b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161347290615cbd565b60405180910390fd5b505b5f60115f8581526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f8481526020019081526020015f205f90553373ffffffffffffffffffffffffffffffffffffffff16827f1ed784ea0b4551753ccb1bbf1711421d8a07aff605d39bb9d770c25943aea48560405160405180910390a3600e5f8381526020019081526020015f205f5f82015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f6135589190614998565b600582015f9055600682015f9055600782015f6101000a81549060ff0219169055600882015f9055600982015f905550505b505b5f600f5f8481526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f90505b60108054905081101561366e5782601082815481106135db576135da61578d565b5b905f5260205f2001540361366157601060016010805490506135fd9190615a43565b8154811061360e5761360d61578d565b5b905f5260205f2001546010828154811061362b5761362a61578d565b5b905f5260205f2001819055506010805480613649576136486165d9565b5b600190038181905f5260205f20015f9055905561366e565b80806001019150506135b9565b50505b50565b5f5f61367f83613de0565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036136f157826040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016136e89190614caf565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b6137478383836001613e19565b505050565b5f5f613759858585613fd8565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141580156137c357505f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614155b1561380f57600f5f8581526020019081526020015f205f9054906101000a900460ff16156137f5576137f4846132f7565b5b5f600d5f8681526020019081526020015f20600301819055505b809150509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361388a57816040517f5b08ba180000000000000000000000000000000000000000000000000000000081526004016138819190614c1f565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516139769190614a9b565b60405180910390a3505050565b6002600a54036139bf576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600a81905550565b6001600a81905550565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603613a43575f6040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401613a3a9190614c1f565b60405180910390fd5b5f613a4f83835f61374c565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614613ac1575f6040517f73c6ac6e000000000000000000000000000000000000000000000000000000008152600401613ab89190614c1f565b60405180910390fd5b505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115613c6b578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b8152600401613b249493929190616606565b6020604051808303815f875af1925050508015613b5f57506040513d601f19601f82011682018060405250810190613b5c9190616664565b60015b613be0573d805f8114613b8d576040519150601f19603f3d011682016040523d82523d5f602084013e613b92565b606091505b505f815103613bd857836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401613bcf9190614c1f565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614613c6957836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401613c609190614c1f565b60405180910390fd5b505b5050505050565b606060405180602001604052805f815250905090565b60605f6001613c96846140f2565b0190505f8167ffffffffffffffff811115613cb457613cb3615104565b5b6040519080825280601f01601f191660200182016040528015613ce65781602001600182028036833780820191505090505b5090505f82602001820190505b600115613d47578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8581613d3c57613d3b61668f565b5b0494505f8503613cf3575b819350505050919050565b613d5d848484614243565b613d71613d68613733565b85858585613ac6565b50505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b8080613e5157505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15613f83575f613e6084613674565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015613eca57508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b8015613edd5750613edb818461264f565b155b15613f1f57826040517fa9fbf51f000000000000000000000000000000000000000000000000000000008152600401613f169190614c1f565b60405180910390fd5b8115613f8157838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b5f5f613fe58585856143ab565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361402857614023846145b6565b614067565b8473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146140665761406581856145fa565b5b5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16036140a8576140a3846146d1565b6140e7565b8473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146140e6576140e58585614791565b5b5b809150509392505050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061414e577a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083816141445761414361668f565b5b0492506040810190505b6d04ee2d6d415b85acef8100000000831061418b576d04ee2d6d415b85acef810000000083816141815761418061668f565b5b0492506020810190505b662386f26fc1000083106141ba57662386f26fc1000083816141b0576141af61668f565b5b0492506010810190505b6305f5e10083106141e3576305f5e10083816141d9576141d861668f565b5b0492506008810190505b61271083106142085761271083816141fe576141fd61668f565b5b0492506004810190505b6064831061422b57606483816142215761422061668f565b5b0492506002810190505b600a831061423a576001810190505b80915050919050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036142b3575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016142aa9190614c1f565b60405180910390fd5b5f6142bf83835f61374c565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361433157816040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016143289190614caf565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146143a5578382826040517f64283d7b00000000000000000000000000000000000000000000000000000000815260040161439c93929190615731565b60405180910390fd5b50505050565b5f5f6143b684613de0565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146143f7576143f6818486614815565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614614482576144365f855f5f613e19565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461450157600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b60088054905060095f8381526020019081526020015f2081905550600881908060018154018082558091505060019003905f5260205f20015f909190919091505550565b5f6146048361115b565b90505f60075f8481526020019081526020015f205490505f60065f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090508282146146a3575f815f8581526020019081526020015f2054905080825f8581526020019081526020015f20819055508260075f8381526020019081526020015f2081905550505b60075f8581526020019081526020015f205f9055805f8481526020019081526020015f205f90555050505050565b5f60016008805490506146e49190615a43565b90505f60095f8481526020019081526020015f205490505f600883815481106147105761470f61578d565b5b905f5260205f200154905080600883815481106147305761472f61578d565b5b905f5260205f2001819055508160095f8381526020019081526020015f208190555060095f8581526020019081526020015f205f90556008805480614778576147776165d9565b5b600190038181905f5260205f20015f9055905550505050565b5f600161479d8461115b565b6147a79190615a43565b90508160065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f20819055508060075f8481526020019081526020015f2081905550505050565b6148208383836148d8565b6148d3575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361489457806040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161488b9190614caf565b60405180910390fd5b81816040517f177e802f0000000000000000000000000000000000000000000000000000000081526004016148ca929190615766565b60405180910390fd5b505050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561498f57508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480614950575061494f848461264f565b5b8061498e57508273ffffffffffffffffffffffffffffffffffffffff16614976836136fa565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5080546149a490615701565b5f825580601f106149b557506149d2565b601f0160209004905f5260205f20908101906149d191906149d5565b5b50565b5b808211156149ec575f815f9055506001016149d6565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b614a3581614a01565b8114614a3f575f5ffd5b50565b5f81359050614a5081614a2c565b92915050565b5f60208284031215614a6b57614a6a6149f9565b5b5f614a7884828501614a42565b91505092915050565b5f8115159050919050565b614a9581614a81565b82525050565b5f602082019050614aae5f830184614a8c565b92915050565b5f819050919050565b614ac681614ab4565b8114614ad0575f5ffd5b50565b5f81359050614ae181614abd565b92915050565b5f5f60408385031215614afd57614afc6149f9565b5b5f614b0a85828601614ad3565b9250506020614b1b85828601614ad3565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f614b6782614b25565b614b718185614b2f565b9350614b81818560208601614b3f565b614b8a81614b4d565b840191505092915050565b5f6020820190508181035f830152614bad8184614b5d565b905092915050565b5f60208284031215614bca57614bc96149f9565b5b5f614bd784828501614ad3565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f614c0982614be0565b9050919050565b614c1981614bff565b82525050565b5f602082019050614c325f830184614c10565b92915050565b614c4181614bff565b8114614c4b575f5ffd5b50565b5f81359050614c5c81614c38565b92915050565b5f5f60408385031215614c7857614c776149f9565b5b5f614c8585828601614c4e565b9250506020614c9685828601614ad3565b9150509250929050565b614ca981614ab4565b82525050565b5f602082019050614cc25f830184614ca0565b92915050565b5f5f5f60608486031215614cdf57614cde6149f9565b5b5f614cec86828701614c4e565b9350506020614cfd86828701614c4e565b9250506040614d0e86828701614ad3565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f614d3c82614d18565b614d468185614d22565b9350614d56818560208601614b3f565b614d5f81614b4d565b840191505092915050565b5f819050919050565b614d7c81614d6a565b82525050565b5f60a0820190508181035f830152614d9a8188614d32565b9050614da96020830187614d73565b614db66040830186614a8c565b614dc36060830185614ca0565b614dd06080830184614ca0565b9695505050505050565b614de381614d6a565b8114614ded575f5ffd5b50565b5f81359050614dfe81614dda565b92915050565b5f60208284031215614e1957614e186149f9565b5b5f614e2684828501614df0565b91505092915050565b5f61014082019050614e435f83018d614c10565b614e50602083018c614ca0565b614e5d604083018b614ca0565b614e6a606083018a614ca0565b8181036080830152614e7c8189614d32565b9050614e8b60a0830188614d73565b614e9860c0830187614d73565b614ea560e0830186614a8c565b614eb3610100830185614ca0565b614ec1610120830184614ca0565b9b9a5050505050505050505050565b5f60208284031215614ee557614ee46149f9565b5b5f614ef284828501614c4e565b91505092915050565b5f602082019050614f0e5f830184614d73565b92915050565b614f1d81614a81565b8114614f27575f5ffd5b50565b5f81359050614f3881614f14565b92915050565b5f5f60408385031215614f5457614f536149f9565b5b5f614f6185828601614c4e565b9250506020614f7285828601614f2a565b9150509250929050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112614f9d57614f9c614f7c565b5b8235905067ffffffffffffffff811115614fba57614fb9614f80565b5b602083019150836001820283011115614fd657614fd5614f84565b5b9250929050565b5f5f5f60408486031215614ff457614ff36149f9565b5b5f61500186828701614ad3565b935050602084013567ffffffffffffffff811115615022576150216149fd565b5b61502e86828701614f88565b92509250509250925092565b5f5f83601f84011261504f5761504e614f7c565b5b8235905067ffffffffffffffff81111561506c5761506b614f80565b5b60208301915083600182028301111561508857615087614f84565b5b9250929050565b5f5f5f5f606085870312156150a7576150a66149f9565b5b5f85013567ffffffffffffffff8111156150c4576150c36149fd565b5b6150d08782880161503a565b945094505060206150e387828801614df0565b92505060406150f487828801614ad3565b91505092959194509250565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61513a82614b4d565b810181811067ffffffffffffffff8211171561515957615158615104565b5b80604052505050565b5f61516b6149f0565b90506151778282615131565b919050565b5f67ffffffffffffffff82111561519657615195615104565b5b61519f82614b4d565b9050602081019050919050565b828183375f83830152505050565b5f6151cc6151c78461517c565b615162565b9050828152602081018484840111156151e8576151e7615100565b5b6151f38482856151ac565b509392505050565b5f82601f83011261520f5761520e614f7c565b5b813561521f8482602086016151ba565b91505092915050565b5f5f5f5f608085870312156152405761523f6149f9565b5b5f61524d87828801614c4e565b945050602061525e87828801614c4e565b935050604061526f87828801614ad3565b925050606085013567ffffffffffffffff8111156152905761528f6149fd565b5b61529c878288016151fb565b91505092959194509250565b5f5f5f5f606085870312156152c0576152bf6149f9565b5b5f6152cd87828801614df0565b945050602085013567ffffffffffffffff8111156152ee576152ed6149fd565b5b6152fa8782880161503a565b9350935050604061530d87828801614df0565b91505092959194509250565b5f5f6040838503121561532f5761532e6149f9565b5b5f61533c85828601614c4e565b925050602061534d85828601614c4e565b9150509250929050565b5f6020820190508181035f83015261536f8184614d32565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6153a981614ab4565b82525050565b5f6153ba83836153a0565b60208301905092915050565b5f602082019050919050565b5f6153dc82615377565b6153e68185615381565b93506153f183615391565b805f5b8381101561542157815161540888826153af565b9750615413836153c6565b9250506001810190506153f4565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61546081614bff565b82525050565b5f6154718383615457565b60208301905092915050565b5f602082019050919050565b5f6154938261542e565b61549d8185615438565b93506154a883615448565b805f5b838110156154d85781516154bf8882615466565b97506154ca8361547d565b9250506001810190506154ab565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61551781614a81565b82525050565b5f615528838361550e565b60208301905092915050565b5f602082019050919050565b5f61554a826154e5565b61555481856154ef565b935061555f836154ff565b805f5b8381101561558f578151615576888261551d565b975061558183615534565b925050600181019050615562565b5085935050505092915050565b5f6080820190508181035f8301526155b481876153d2565b905081810360208301526155c88186615489565b905081810360408301526155dc81856153d2565b905081810360608301526155f08184615540565b905095945050505050565b5f5f5f5f60608587031215615613576156126149f9565b5b5f61562087828801614df0565b945050602085013567ffffffffffffffff811115615641576156406149fd565b5b61564d8782880161503a565b9350935050604061566087828801614ad3565b91505092959194509250565b7f4e6f7420746f6b656e206f776e657200000000000000000000000000000000005f82015250565b5f6156a0600f83614b2f565b91506156ab8261566c565b602082019050919050565b5f6020820190508181035f8301526156cd81615694565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061571857607f821691505b60208210810361572b5761572a6156d4565b5b50919050565b5f6060820190506157445f830186614c10565b6157516020830185614ca0565b61575e6040830184614c10565b949350505050565b5f6040820190506157795f830185614c10565b6157866020830184614ca0565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f8160601b9050919050565b5f6157d0826157ba565b9050919050565b5f6157e1826157c6565b9050919050565b6157f96157f482614bff565b6157d7565b82525050565b5f819050919050565b61581961581482614ab4565b6157ff565b82525050565b5f61582a82856157e8565b60148201915061583a8284615808565b6020820191508190509392505050565b7f436c756520616c726561647920736f6c766564000000000000000000000000005f82015250565b5f61587e601383614b2f565b91506158898261584a565b602082019050919050565b5f6020820190508181035f8301526158ab81615872565b9050919050565b5f81905092915050565b5f6158c783856158b2565b93506158d48385846151ac565b82840190509392505050565b5f6158ec8284866158bc565b91508190509392505050565b5f6159038385614b2f565b93506159108385846151ac565b61591983614b4d565b840190509392505050565b5f6020820190508181035f83015261593d8184866158f8565b90509392505050565b7f4e6f7420746865206275796572000000000000000000000000000000000000005f82015250565b5f61597a600d83614b2f565b915061598582615946565b602082019050919050565b5f6020820190508181035f8301526159a78161596e565b9050919050565b7f50726f6f66206e6f74207965742070726f7669646564000000000000000000005f82015250565b5f6159e2601683614b2f565b91506159ed826159ae565b602082019050919050565b5f6020820190508181035f830152615a0f816159d6565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f615a4d82614ab4565b9150615a5883614ab4565b9250828203905081811115615a7057615a6f615a16565b5b92915050565b7f50726f6f6620766572696669636174696f6e20657870697265640000000000005f82015250565b5f615aaa601a83614b2f565b9150615ab582615a76565b602082019050919050565b5f6020820190508181035f830152615ad781615a9e565b9050919050565b7f5472616e7366657220646f6573206e6f742065786973740000000000000000005f82015250565b5f615b12601783614b2f565b9150615b1d82615ade565b602082019050919050565b5f6020820190508181035f830152615b3f81615b06565b9050919050565b7f43616e6e6f742063616e63656c2061667465722070726f6f66207665726966695f8201527f636174696f6e0000000000000000000000000000000000000000000000000000602082015250565b5f615ba0602683614b2f565b9150615bab82615b46565b604082019050919050565b5f6020820190508181035f830152615bcd81615b94565b9050919050565b7f4e6f7420617574686f72697a656420746f2063616e63656c00000000000000005f82015250565b5f615c08601883614b2f565b9150615c1382615bd4565b602082019050919050565b5f6020820190508181035f830152615c3581615bfc565b9050919050565b50565b5f615c4a5f836158b2565b9150615c5582615c3c565b5f82019050919050565b5f615c6982615c3f565b9150819050919050565b7f4661696c656420746f2073656e642045746865720000000000000000000000005f82015250565b5f615ca7601483614b2f565b9150615cb282615c73565b602082019050919050565b5f6020820190508181035f830152615cd481615c9b565b9050919050565b5f615ce582614ab4565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615d1757615d16615a16565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302615d7e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82615d43565b615d888683615d43565b95508019841693508086168417925050509392505050565b5f819050919050565b5f615dc3615dbe615db984614ab4565b615da0565b614ab4565b9050919050565b5f819050919050565b615ddc83615da9565b615df0615de882615dca565b848454615d4f565b825550505050565b5f5f905090565b615e07615df8565b615e12818484615dd3565b505050565b5b81811015615e3557615e2a5f82615dff565b600181019050615e18565b5050565b601f821115615e7a57615e4b81615d22565b615e5484615d34565b81016020851015615e63578190505b615e77615e6f85615d34565b830182615e17565b50505b505050565b5f82821c905092915050565b5f615e9a5f1984600802615e7f565b1980831691505092915050565b5f615eb28383615e8b565b9150826002028217905092915050565b615ecb82614d18565b67ffffffffffffffff811115615ee457615ee3615104565b5b615eee8254615701565b615ef9828285615e39565b5f60209050601f831160018114615f2a575f8415615f18578287015190505b615f228582615ea7565b865550615f89565b601f198416615f3886615d22565b5f5b82811015615f5f57848901518255600182019150602085019450602081019050615f3a565b86831015615f7c5784890151615f78601f891682615e8b565b8355505b6001600288020188555050505b505050505050565b7f5472616e736665722065787069726564000000000000000000000000000000005f82015250565b5f615fc5601083614b2f565b9150615fd082615f91565b602082019050919050565b5f6020820190508181035f830152615ff281615fb9565b9050919050565b7f50726f6f6620746f6f2073686f727400000000000000000000000000000000005f82015250565b5f61602d600f83614b2f565b915061603882615ff9565b602082019050919050565b5f6020820190508181035f83015261605a81616021565b9050919050565b5f63ffffffff82169050919050565b5f61607a82616061565b915061608583616061565b9250828201905063ffffffff8111156160a1576160a0615a16565b5b92915050565b7f496e76616c69642070726f6f66207374727563747572650000000000000000005f82015250565b5f6160db601783614b2f565b91506160e6826160a7565b602082019050919050565b5f6020820190508181035f830152616108816160cf565b9050919050565b5f61611982616061565b915061612483616061565b9250828203905063ffffffff8111156161405761613f615a16565b5b92915050565b5f5ffd5b5f5ffd5b5f5f8585111561616157616160616146565b5b838611156161725761617161614a565b5b6001850283019150848603905094509492505050565b5f82905092915050565b5f61619d8383616188565b826161a88135614d6a565b925060208210156161e8576161e37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83602003600802615d43565b831692505b505092915050565b6161fa8383616188565b67ffffffffffffffff81111561621357616212615104565b5b61621d8254615701565b616228828285615e39565b5f601f831160018114616255575f8415616243578287013590505b61624d8582615ea7565b8655506162b4565b601f19841661626386615d22565b5f5b8281101561628a57848901358255600182019150602085019450602081019050616265565b868310156162a757848901356162a3601f891682615e8b565b8355505b6001600288020188555050505b50505050505050565b5f6162c88385614d22565b93506162d58385846151ac565b6162de83614b4d565b840190509392505050565b5f6060820190508181035f8301526163028186886162bd565b90506163116020830185614d73565b61631e6040830184614d73565b95945050505050565b5f81905092915050565b5f61633b82614b25565b6163458185616327565b9350616355818560208601614b3f565b80840191505092915050565b5f61636c8285616331565b91506163788284616331565b91508190509392505050565b7f5472616e7366657220616c726561647920696e697469617465640000000000005f82015250565b5f6163b8601a83614b2f565b91506163c382616384565b602082019050919050565b5f6020820190508181035f8301526163e5816163ac565b9050919050565b5f6163f682614ab4565b915061640183614ab4565b925082820190508082111561641957616418615a16565b5b92915050565b7f50726f6f66206e6f7420766572696669656400000000000000000000000000005f82015250565b5f616453601283614b2f565b915061645e8261641f565b602082019050919050565b5f6020820190508181035f83015261648081616447565b9050919050565b7f5472616e7366657220636f6d706c6574696f6e206578706972656400000000005f82015250565b5f6164bb601b83614b2f565b91506164c682616487565b602082019050919050565b5f6020820190508181035f8301526164e8816164af565b9050919050565b7f436f6e74656e742068617368206d69736d6174636800000000000000000000005f82015250565b5f616523601583614b2f565b915061652e826164ef565b602082019050919050565b5f6020820190508181035f83015261655081616517565b9050919050565b5f6165628284615808565b60208201915081905092915050565b7f522076616c75652068617368206d69736d6174636800000000000000000000005f82015250565b5f6165a5601583614b2f565b91506165b082616571565b602082019050919050565b5f6020820190508181035f8301526165d281616599565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f6080820190506166195f830187614c10565b6166266020830186614c10565b6166336040830185614ca0565b81810360608301526166458184614d32565b905095945050505050565b5f8151905061665e81614a2c565b92915050565b5f60208284031215616679576166786149f9565b5b5f61668684828501616650565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffdfea26469706673582212204b50cfd8d681656cf1a9fc06ac4425179122aa8a2fafeb3b85963be020fa723364736f6c634300081c0033",
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
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 salePrice, uint256 rValue)
func (_Skavenge *SkavengeCaller) Clues(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
	SalePrice         *big.Int
	RValue            *big.Int
}, error) {
	var out []interface{}
	err := _Skavenge.contract.Call(opts, &out, "clues", arg0)

	outstruct := new(struct {
		EncryptedContents []byte
		SolutionHash      [32]byte
		IsSolved          bool
		SalePrice         *big.Int
		RValue            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EncryptedContents = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.SolutionHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.IsSolved = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.SalePrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Clues is a free data retrieval call binding the contract method 0x30f37c7f.
//
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 salePrice, uint256 rValue)
func (_Skavenge *SkavengeSession) Clues(arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
	SalePrice         *big.Int
	RValue            *big.Int
}, error) {
	return _Skavenge.Contract.Clues(&_Skavenge.CallOpts, arg0)
}

// Clues is a free data retrieval call binding the contract method 0x30f37c7f.
//
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 salePrice, uint256 rValue)
func (_Skavenge *SkavengeCallerSession) Clues(arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
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
