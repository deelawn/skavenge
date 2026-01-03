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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialMinter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ClueNotForSale\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolvedClueCannotBeSold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolvedClueTransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferAlreadyInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedMinterUpdate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMinter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"AuthorizedMinterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"ClueMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"solution\",\"type\":\"string\"}],\"name\":\"ClueSolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rValueHash\",\"type\":\"bytes32\"}],\"name\":\"ProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"ProofVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SalePriceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SalePriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cancelledBy\",\"type\":\"address\"}],\"name\":\"TransferCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"name\":\"TransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TransferInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_TIMEOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_TIMEOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeTransferIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"solution\",\"type\":\"string\"}],\"name\":\"attemptSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorizedMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"cancelTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"clues\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"solutionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isSolved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cluesForSale\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"newEncryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"name\":\"completeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"generateTransferId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getClueContents\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getCluesForSale\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"solvedStatus\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalCluesForSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"initiatePurchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedContents\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"solutionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rValue\",\"type\":\"uint256\"}],\"name\":\"mintClue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"}],\"name\":\"provideProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"setSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initiatedAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newClueHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rValueHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"proofVerified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"proofProvidedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"updateAuthorizedMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051616d12380380616d1283398181016040528101906100319190610172565b6040518060400160405280600881526020017f536b6176656e67650000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f534b564700000000000000000000000000000000000000000000000000000000815250815f90816100ab91906103da565b5080600190816100bb91906103da565b5050506001600a819055506001600b8190555080600c5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506104a9565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61014182610118565b9050919050565b61015181610137565b811461015b575f5ffd5b50565b5f8151905061016c81610148565b92915050565b5f6020828403121561018757610186610114565b5b5f6101948482850161015e565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061021857607f821691505b60208210810361022b5761022a6101d4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261028d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610252565b6102978683610252565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102db6102d66102d1846102af565b6102b8565b6102af565b9050919050565b5f819050919050565b6102f4836102c1565b610308610300826102e2565b84845461025e565b825550505050565b5f5f905090565b61031f610310565b61032a8184846102eb565b505050565b5b8181101561034d576103425f82610317565b600181019050610330565b5050565b601f8211156103925761036381610231565b61036c84610243565b8101602085101561037b578190505b61038f61038785610243565b83018261032f565b50505b505050565b5f82821c905092915050565b5f6103b25f1984600802610397565b1980831691505092915050565b5f6103ca83836103a3565b9150826002028217905092915050565b6103e38261019d565b67ffffffffffffffff8111156103fc576103fb6101a7565b5b6104068254610201565b610411828285610351565b5f60209050601f831160018114610442575f8415610430578287015190505b61043a85826103bf565b8655506104a1565b601f19841661045086610231565b5f5b8281101561047757848901518255600182019150602085019450602081019050610452565b868310156104945784890151610490601f8916826103a3565b8355505b6001600288020188555050505b505050505050565b61685c806104b65f395ff3fe60806040526004361061023a575f3560e01c806379096ee81161012d578063c2d554ae116100aa578063e985e9c51161006e578063e985e9c5146108f4578063eb927a8314610930578063f12b72ba1461096c578063f8f5a544146109ab578063fae5380c146109d35761023a565b8063c2d554ae1461080e578063c87b56dd14610836578063d32d579014610872578063dd142be0146108ae578063de38eb3a146108ca5761023a565b8063aff202b4116100f1578063aff202b414610732578063b142b4ec1461075a578063b329bf5c14610782578063b40b7eb0146107aa578063b88d4fde146107e65761023a565b806379096ee8146106405780638d7cf3e41461067c57806395d89b41146106a4578063a22cb465146106ce578063a6cd5ff5146106f65761023a565b806334499fff116101bb578063561892361161017f578063561892361461054c5780636352211e146105765780636c39cc34146105b257806370a08231146105da57806374b19a07146106165761023a565b806334499fff1461043d5780633c64f04b1461047957806342842e0e146104be5780634f6ccce7146104e6578063543ad1df146105225761023a565b80631ba538cd116102025780631ba538cd1461033257806323b872dd1461035c5780632f745c591461038457806330f37c7f146103c05780633427ee94146104015761023a565b806301ffc9a71461023e57806306fdde031461027a578063081812fc146102a4578063095ea7b3146102e057806318160ddd14610308575b5f5ffd5b348015610249575f5ffd5b50610264600480360381019061025f9190614b88565b6109fb565b6040516102719190614bcd565b60405180910390f35b348015610285575f5ffd5b5061028e610a74565b60405161029b9190614c56565b60405180910390f35b3480156102af575f5ffd5b506102ca60048036038101906102c59190614ca9565b610b03565b6040516102d79190614d13565b60405180910390f35b3480156102eb575f5ffd5b5061030660048036038101906103019190614d56565b610b1e565b005b348015610313575f5ffd5b5061031c610b34565b6040516103299190614da3565b60405180910390f35b34801561033d575f5ffd5b50610346610b40565b6040516103539190614d13565b60405180910390f35b348015610367575f5ffd5b50610382600480360381019061037d9190614dbc565b610b65565b005b34801561038f575f5ffd5b506103aa60048036038101906103a59190614d56565b610c64565b6040516103b79190614da3565b60405180910390f35b3480156103cb575f5ffd5b506103e660048036038101906103e19190614ca9565b610d08565b6040516103f896959493929190614e76565b60405180910390f35b34801561040c575f5ffd5b5061042760048036038101906104229190614ca9565b610dd2565b6040516104349190614bcd565b60405180910390f35b348015610448575f5ffd5b50610463600480360381019061045e9190614ca9565b610def565b6040516104709190614bcd565b60405180910390f35b348015610484575f5ffd5b5061049f600480360381019061049a9190614f06565b610e0c565b6040516104b59a99989796959493929190614f31565b60405180910390f35b3480156104c9575f5ffd5b506104e460048036038101906104df9190614dbc565b610f0d565b005b3480156104f1575f5ffd5b5061050c60048036038101906105079190614ca9565b610f2c565b6040516105199190614da3565b60405180910390f35b34801561052d575f5ffd5b50610536610f9e565b6040516105439190614da3565b60405180910390f35b348015610557575f5ffd5b50610560610fa2565b60405161056d9190614da3565b60405180910390f35b348015610581575f5ffd5b5061059c60048036038101906105979190614ca9565b610fab565b6040516105a99190614d13565b60405180910390f35b3480156105bd575f5ffd5b506105d860048036038101906105d39190614fd2565b610fbc565b005b3480156105e5575f5ffd5b5061060060048036038101906105fb9190615022565b611210565b60405161060d9190614da3565b60405180910390f35b348015610621575f5ffd5b5061062a6112c6565b6040516106379190614da3565b60405180910390f35b34801561064b575f5ffd5b5061066660048036038101906106619190614ca9565b6112d2565b604051610673919061504d565b60405180910390f35b348015610687575f5ffd5b506106a2600480360381019061069d9190614ca9565b6112e7565b005b3480156106af575f5ffd5b506106b86113d3565b6040516106c59190614c56565b60405180910390f35b3480156106d9575f5ffd5b506106f460048036038101906106ef9190615090565b611463565b005b348015610701575f5ffd5b5061071c60048036038101906107179190614d56565b611479565b604051610729919061504d565b60405180910390f35b34801561073d575f5ffd5b506107586004803603810190610753919061512f565b6114ab565b005b348015610765575f5ffd5b50610780600480360381019061077b9190614f06565b611693565b005b34801561078d575f5ffd5b506107a860048036038101906107a39190614f06565b61184d565b005b3480156107b5575f5ffd5b506107d060048036038101906107cb91906151e1565b611cf4565b6040516107dd9190614da3565b60405180910390f35b3480156107f1575f5ffd5b5061080c6004803603810190610807919061537a565b611ebe565b005b348015610819575f5ffd5b50610834600480360381019061082f91906153fa565b611ee3565b005b348015610841575f5ffd5b5061085c60048036038101906108579190614ca9565b61229f565b6040516108699190614c56565b60405180910390f35b34801561087d575f5ffd5b5061089860048036038101906108939190614ca9565b612305565b6040516108a59190614da3565b60405180910390f35b6108c860048036038101906108c39190614ca9565b612322565b005b3480156108d5575f5ffd5b506108de612763565b6040516108eb9190614da3565b60405180910390f35b3480156108ff575f5ffd5b5061091a6004803603810190610915919061546b565b61276a565b6040516109279190614bcd565b60405180910390f35b34801561093b575f5ffd5b5061095660048036038101906109519190614ca9565b6127f8565b60405161096391906154a9565b60405180910390f35b348015610977575f5ffd5b50610992600480360381019061098d91906154c9565b6128a5565b6040516109a2949392919061572c565b60405180910390f35b3480156109b6575f5ffd5b506109d160048036038101906109cc9190615022565b612b68565b005b3480156109de575f5ffd5b506109f960048036038101906109f4919061578b565b612cb1565b005b5f7f780e9d63000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610a6d5750610a6c82613348565b5b9050919050565b60605f8054610a8290615829565b80601f0160208091040260200160405190810160405280929190818152602001828054610aae90615829565b8015610af95780601f10610ad057610100808354040283529160200191610af9565b820191905f5260205f20905b815481529060010190602001808311610adc57829003601f168201915b5050505050905090565b5f610b0d82613429565b50610b17826134af565b9050919050565b610b308282610b2b6134e8565b6134ef565b5050565b5f600880549050905090565b600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610bd5575f6040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610bcc9190614d13565b60405180910390fd5b5f610be88383610be36134e8565b613501565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610c5e578382826040517f64283d7b000000000000000000000000000000000000000000000000000000008152600401610c5593929190615859565b60405180910390fd5b50505050565b5f610c6e83611210565b8210610cb35782826040517fa57d13dc000000000000000000000000000000000000000000000000000000008152600401610caa92919061588e565b60405180910390fd5b60065f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f2054905092915050565b600d602052805f5260405f205f91509050805f018054610d2790615829565b80601f0160208091040260200160405190810160405280929190818152602001828054610d5390615829565b8015610d9e5780601f10610d7557610100808354040283529160200191610d9e565b820191905f5260205f20905b815481529060010190602001808311610d8157829003601f168201915b505050505090806001015490806002015f9054906101000a900460ff16908060030154908060040154908060050154905086565b600f602052805f5260405f205f915054906101000a900460ff1681565b6011602052805f5260405f205f915054906101000a900460ff1681565b600e602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002015490806003015490806004018054610e6290615829565b80601f0160208091040260200160405190810160405280929190818152602001828054610e8e90615829565b8015610ed95780601f10610eb057610100808354040283529160200191610ed9565b820191905f5260205f20905b815481529060010190602001808311610ebc57829003601f168201915b505050505090806005015490806006015490806007015f9054906101000a900460ff1690806008015490806009015490508a565b610f2783838360405180602001604052805f815250611ebe565b505050565b5f610f35610b34565b8210610f7a575f826040517fa57d13dc000000000000000000000000000000000000000000000000000000008152600401610f7192919061588e565b60405180910390fd5b60088281548110610f8e57610f8d6158b5565b5b905f5260205f2001549050919050565b5f81565b5f600b54905090565b5f610fb582613429565b9050919050565b3373ffffffffffffffffffffffffffffffffffffffff16610fdc84610fab565b73ffffffffffffffffffffffffffffffffffffffff1614611032576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110299061592c565b60405180910390fd5b600d5f8481526020019081526020015f206002015f9054906101000a900460ff161561108a576040517fff1e4dda00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8211156110fe575f81101580156110a55750620151808111155b6110e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110db90615994565b60405180910390fd5b80600d5f8581526020019081526020015f20600501819055505b81600d5f8581526020019081526020015f2060030181905550600f5f8481526020019081526020015f205f9054906101000a900460ff1615801561114157505f82115b1561119a576001600f5f8581526020019081526020015f205f6101000a81548160ff021916908315150217905550601083908060018154018082558091505060019003905f5260205f20015f90919091909150556111d3565b600f5f8481526020019081526020015f205f9054906101000a900460ff1680156111c357505f82145b156111d2576111d1836135cf565b5b5b827fe23ea816dce6d7f5c0b85cbd597e7c3b97b2453791152c0b94e5e5c5f314d2f0836040516112039190614da3565b60405180910390a2505050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611281575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016112789190614d13565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f601080549050905090565b6012602052805f5260405f205f915090505481565b3373ffffffffffffffffffffffffffffffffffffffff1661130782610fab565b73ffffffffffffffffffffffffffffffffffffffff161461135d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113549061592c565b60405180910390fd5b600f5f8281526020019081526020015f205f9054906101000a900460ff161561138a57611389816135cf565b5b5f600d5f8381526020019081526020015f2060030181905550807f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a250565b6060600180546113e290615829565b80601f016020809104026020016040519081016040528092919081815260200182805461140e90615829565b80156114595780601f1061143057610100808354040283529160200191611459565b820191905f5260205f20905b81548152906001019060200180831161143c57829003601f168201915b5050505050905090565b61147561146e6134e8565b838361394c565b5050565b5f828260405160200161148d929190615a17565b60405160208183030381529060405280519060200120905092915050565b3373ffffffffffffffffffffffffffffffffffffffff166114cb84610fab565b73ffffffffffffffffffffffffffffffffffffffff1614611521576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115189061592c565b60405180910390fd5b600d5f8481526020019081526020015f206002015f9054906101000a900460ff1615611582576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161157990615a8c565b60405180910390fd5b600d5f8481526020019081526020015f206001015482826040516115a7929190615ad8565b60405180910390200361168e576001600d5f8581526020019081526020015f206002015f6101000a81548160ff021916908315150217905550600f5f8481526020019081526020015f205f9054906101000a900460ff16156116535761160c836135cf565b5f600d5f8581526020019081526020015f2060030181905550827f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a25b827f3138eb607d845be3efb1a7ea147da7816c8a05f683313c459e6bf953ea4f988e8383604051611685929190615b1c565b60405180910390a25b505050565b61169b613ab5565b5f600e5f8381526020019081526020015f2090503373ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461173f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161173690615b88565b60405180910390fd5b5f816008015411611785576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161177c90615bf0565b60405180910390fd5b600d5f826001015481526020019081526020015f20600501548160080154426117ae9190615c3b565b11156117ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117e690615cb8565b60405180910390fd5b6001816007015f6101000a81548160ff021916908315150217905550428160090181905550817f543093db8d78fd8619586d3a0be12a5736836393feede0888f262888c81ce4c360405160405180910390a25061184a613afb565b50565b611855613ab5565b5f600e5f8381526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036118f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118f090615d20565b60405180910390fd5b5f3373ffffffffffffffffffffffffffffffffffffffff16825f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161490505f3373ffffffffffffffffffffffffffffffffffffffff166119728460010154610fab565b73ffffffffffffffffffffffffffffffffffffffff161490505f5f90508215611a51575f84600801541480156119ce5750600d5f856001015481526020019081526020015f20600501548460030154426119cc9190615c3b565b115b156119dc5760019050611a50565b5f84600801541180156119fd5750836007015f9054906101000a900460ff16155b15611a0b5760019050611a4f565b5f8460090154118015611a445750600d5f856001015481526020019081526020015f2060050154846009015442611a429190615c3b565b115b15611a4e57600190505b5b5b5b8115611ab5575f8460080154118015611a785750836007015f9054906101000a900460ff16155b8015611aaa5750600d5f856001015481526020019081526020015f2060050154846008015442611aa89190615c3b565b115b15611ab457600190505b5b80611af5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611aec90615d88565b60405180910390fd5b5f84600201541115611bd1575f845f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168560020154604051611b4c90615dc9565b5f6040518083038185875af1925050503d805f8114611b86576040519150601f19603f3d011682016040523d82523d5f602084013e611b8b565b606091505b5050905080611bcf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bc690615e27565b60405180910390fd5b505b5f60115f866001015481526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f856001015481526020019081526020015f205f90553373ffffffffffffffffffffffffffffffffffffffff16857f1ed784ea0b4551753ccb1bbf1711421d8a07aff605d39bb9d770c25943aea48560405160405180910390a3600e5f8681526020019081526020015f205f5f82015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f611cb49190614aca565b600582015f9055600682015f9055600782015f6101000a81549060ff0219169055600882015f9055600982015f9055505050505050611cf1613afb565b50565b5f600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d7b576040517f955c501b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b5f815480929190611d8d90615e45565b9190505590506040518060c0016040528086868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018481526020015f151581526020015f81526020018381526020015f815250600d5f8381526020019081526020015f205f820151815f019081611e29919061602c565b50602082015181600101556040820151816002015f6101000a81548160ff021916908315150217905550606082015181600301556080820151816004015560a08201518160050155905050611e7e3382613b05565b807fa90e59f66e7533243b5959b6498caf4949957dbf8ccaa6b6534177c10041ea5433604051611eae9190614d13565b60405180910390a2949350505050565b611ec9848484610b65565b611edd611ed46134e8565b85858585613bf8565b50505050565b611eeb613ab5565b5f600e5f8681526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611f8f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f8690615d20565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16611fb38260010154610fab565b73ffffffffffffffffffffffffffffffffffffffff1614612009576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120009061592c565b60405180910390fd5b600d5f826001015481526020019081526020015f20600501548160030154426120329190615c3b565b1115612073576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161206a90616145565b60405180910390fd5b60248484905010156120ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120b1906161ad565b60405180910390fd5b5f848460038181106120cf576120ce6158b5565b5b9050013560f81c60f81b60f81c60ff166008868660028181106120f5576120f46158b5565b5b9050013560f81c60f81b60f81c60ff1663ffffffff16901b601087876001818110612123576121226158b5565b5b9050013560f81c60f81b60f81c60ff1663ffffffff16901b601888885f8181106121505761214f6158b5565b5b9050013560f81c60f81b60f81c60ff1663ffffffff16901b1717179050602081600461217c91906161da565b61218691906161da565b63ffffffff168585905010156121d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121c89061625b565b60405180910390fd5b5f858560208460046121e391906161da565b6121ed9190616279565b63ffffffff169084600461220191906161da565b63ffffffff1692612214939291906162b8565b9061221f91906162fc565b9050858584600401918261223492919061635a565b50838360050181905550808360060181905550428360080181905550867f319414a72bfc3d93a989d08f1055fd74a1b953a652be46d0dff852ac157c12f2878787856040516122869493929190616453565b60405180910390a2505050612299613afb565b50505050565b60606122aa82613429565b505f6122b4613da4565b90505f8151116122d25760405180602001604052805f8152506122fd565b806122dc84613dba565b6040516020016122ed9291906164cb565b6040516020818303038152906040525b915050919050565b5f600d5f8381526020019081526020015f20600401549050919050565b61232a613ab5565b61233381610fab565b50600f5f8281526020019081526020015f205f9054906101000a900460ff16158061237157505f600d5f8381526020019081526020015f2060030154145b156123a8576040517fa7d67ebb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5f8281526020019081526020015f20600301543410156123f6576040517f356680b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5f8281526020019081526020015f206002015f9054906101000a900460ff161561244e576040517f6e40ff0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60115f8281526020019081526020015f205f9054906101000a900460ff16156124a3576040517f74ed79ae00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6124ae3383611479565b90505f73ffffffffffffffffffffffffffffffffffffffff16600e5f8381526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612550576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161254790616538565b60405180910390fd5b600160115f8481526020019081526020015f205f6101000a81548160ff0219169083151502179055508060125f8481526020019081526020015f20819055506040518061014001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018381526020013481526020014281526020015f67ffffffffffffffff8111156125e3576125e2615256565b5b6040519080825280601f01601f1916602001820160405280156126155781602001600182028036833780820191505090505b5081526020015f5f1b81526020015f5f1b81526020015f151581526020015f81526020015f815250600e5f8381526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003015560808201518160040190816126c5919061602c565b5060a0820151816005015560c0820151816006015560e0820151816007015f6101000a81548160ff02191690831515021790555061010082015181600801556101208201518160090155905050813373ffffffffffffffffffffffffffffffffffffffff16827f2d18295f817f7e46b8d3401af48ee043761aba21f602005110a282939c3c4c7260405160405180910390a450612760613afb565b50565b6201518081565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b606061280382610fab565b50600d5f8381526020019081526020015f205f01805461282290615829565b80601f016020809104026020016040519081016040528092919081815260200182805461284e90615829565b80156128995780601f1061287057610100808354040283529160200191612899565b820191905f5260205f20905b81548152906001019060200180831161287c57829003601f168201915b50505050509050919050565b6060806060805f60108054905090505f8187896128c29190616556565b116128d85786886128d39190616556565b6128da565b815b90505f8882116128ea575f6128f7565b88826128f69190615c3b565b5b90508067ffffffffffffffff81111561291357612912615256565b5b6040519080825280602002602001820160405280156129415781602001602082028036833780820191505090505b5096508067ffffffffffffffff81111561295e5761295d615256565b5b60405190808252806020026020018201604052801561298c5781602001602082028036833780820191505090505b5095508067ffffffffffffffff8111156129a9576129a8615256565b5b6040519080825280602002602001820160405280156129d75781602001602082028036833780820191505090505b5094508067ffffffffffffffff8111156129f4576129f3615256565b5b604051908082528060200260200182016040528015612a225781602001602082028036833780820191505090505b5093505f5f90505b81811015612b5b575f6010828c612a419190616556565b81548110612a5257612a516158b5565b5b905f5260205f200154905080898381518110612a7157612a706158b5565b5b602002602001018181525050612a8681610fab565b888381518110612a9957612a986158b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600d5f8281526020019081526020015f2060030154878381518110612afb57612afa6158b5565b5b602002602001018181525050600d5f8281526020019081526020015f206002015f9054906101000a900460ff16868381518110612b3b57612b3a6158b5565b5b602002602001019015159081151581525050508080600101915050612a2a565b5050505092959194509250565b600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612bee576040517f7efb568f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f600c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600c5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f808ec13129987deb49ec337ab895a2cf7af16a4d0d55a51ddc054e2c7fb2515b60405160405180910390a35050565b612cb9613ab5565b5f600e5f8681526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612d5d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d5490615d20565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16612d818260010154610fab565b73ffffffffffffffffffffffffffffffffffffffff1614612dd7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dce9061592c565b60405180910390fd5b5f816009015411612e1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e14906165d3565b60405180910390fd5b600d5f826001015481526020019081526020015f2060050154816009015442612e469190615c3b565b1115612e87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e7e9061663b565b60405180910390fd5b80600501548484604051612e9c929190615ad8565b604051809103902014612ee4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612edb906166a3565b60405180910390fd5b806006015482604051602001612efa91906166c1565b6040516020818303038152906040528051906020012014612f50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f4790616725565b60405180910390fd5b600d5f826001015481526020019081526020015f206002015f9054906101000a900460ff1615612fac576040517f6e40ff0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8383600d5f846001015481526020019081526020015f205f019182612fd292919061635a565b5081600d5f836001015481526020019081526020015f20600401819055505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f826002015490505f83600101549050600f5f8281526020019081526020015f205f9054906101000a900460ff1615613173575f600f5f8381526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f90505b60108054905081101561312b578160108281548110613098576130976158b5565b5b905f5260205f2001540361311e57601060016010805490506130ba9190615c3b565b815481106130cb576130ca6158b5565b5b905f5260205f200154601082815481106130e8576130e76158b5565b5b905f5260205f200181905550601080548061310657613105616743565b5b600190038181905f5260205f20015f9055905561312b565b8080600101915050613076565b505f600d5f8381526020019081526020015f2060030181905550807f06dad65ef75f5ad325b5f8a967c83ae3a120f2d6f9bd9928a7a6c71d6f26898960405160405180910390a25b5f60115f8381526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f8281526020019081526020015f205f9055600e5f8981526020019081526020015f205f5f82015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f61320a9190614aca565b600582015f9055600682015f9055600782015f6101000a81549060ff0219169055600882015f9055600982015f9055505061325533848360405180602001604052805f815250613e84565b5f3373ffffffffffffffffffffffffffffffffffffffff168360405161327a90615dc9565b5f6040518083038185875af1925050503d805f81146132b4576040519150601f19603f3d011682016040523d82523d5f602084013e6132b9565b606091505b50509050806132fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132f490615e27565b60405180910390fd5b887f062fb96142a4ea35fc5c48049c3a7d7a418829dea520220e03d76440bbe275c08760405161332d9190614da3565b60405180910390a25050505050613342613afb565b50505050565b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061341257507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80613422575061342182613ea9565b5b9050919050565b5f5f61343483613f12565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036134a657826040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161349d9190614da3565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b6134fc8383836001613f4b565b505050565b5f5f61350e85858561410a565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561357857505f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614155b156135c457600f5f8581526020019081526020015f205f9054906101000a900460ff16156135aa576135a9846135cf565b5b5f600d5f8681526020019081526020015f20600301819055505b809150509392505050565b600f5f8281526020019081526020015f205f9054906101000a900460ff1615613949575f60125f8381526020019081526020015f205490505f5f1b8114613864575f600e5f8381526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613862575f81600201541115613755575f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1682600201546040516136d090615dc9565b5f6040518083038185875af1925050503d805f811461370a576040519150601f19603f3d011682016040523d82523d5f602084013e61370f565b606091505b5050905080613753576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161374a90615e27565b60405180910390fd5b505b5f60115f8581526020019081526020015f205f6101000a81548160ff02191690831515021790555060125f8481526020019081526020015f205f90553373ffffffffffffffffffffffffffffffffffffffff16827f1ed784ea0b4551753ccb1bbf1711421d8a07aff605d39bb9d770c25943aea48560405160405180910390a3600e5f8381526020019081526020015f205f5f82015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f6138309190614aca565b600582015f9055600682015f9055600782015f6101000a81549060ff0219169055600882015f9055600982015f905550505b505b5f600f5f8481526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f90505b6010805490508110156139465782601082815481106138b3576138b26158b5565b5b905f5260205f2001540361393957601060016010805490506138d59190615c3b565b815481106138e6576138e56158b5565b5b905f5260205f20015460108281548110613903576139026158b5565b5b905f5260205f200181905550601080548061392157613920616743565b5b600190038181905f5260205f20015f90559055613946565b8080600101915050613891565b50505b50565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036139bc57816040517f5b08ba180000000000000000000000000000000000000000000000000000000081526004016139b39190614d13565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051613aa89190614bcd565b60405180910390a3505050565b6002600a5403613af1576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600a81905550565b6001600a81905550565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603613b75575f6040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401613b6c9190614d13565b60405180910390fd5b5f613b8183835f613501565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614613bf3575f6040517f73c6ac6e000000000000000000000000000000000000000000000000000000008152600401613bea9190614d13565b60405180910390fd5b505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115613d9d578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b8152600401613c569493929190616770565b6020604051808303815f875af1925050508015613c9157506040513d601f19601f82011682018060405250810190613c8e91906167ce565b60015b613d12573d805f8114613cbf576040519150601f19603f3d011682016040523d82523d5f602084013e613cc4565b606091505b505f815103613d0a57836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401613d019190614d13565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614613d9b57836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401613d929190614d13565b60405180910390fd5b505b5050505050565b606060405180602001604052805f815250905090565b60605f6001613dc884614224565b0190505f8167ffffffffffffffff811115613de657613de5615256565b5b6040519080825280601f01601f191660200182016040528015613e185781602001600182028036833780820191505090505b5090505f82602001820190505b600115613e79578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8581613e6e57613e6d6167f9565b5b0494505f8503613e25575b819350505050919050565b613e8f848484614375565b613ea3613e9a6134e8565b85858585613bf8565b50505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b8080613f8357505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b156140b5575f613f9284613429565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015613ffc57508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561400f575061400d818461276a565b155b1561405157826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016140489190614d13565b60405180910390fd5b81156140b357838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b5f5f6141178585856144dd565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361415a57614155846146e8565b614199565b8473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461419857614197818561472c565b5b5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16036141da576141d584614803565b614219565b8473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146142185761421785856148c3565b5b5b809150509392505050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310614280577a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008381614276576142756167f9565b5b0492506040810190505b6d04ee2d6d415b85acef810000000083106142bd576d04ee2d6d415b85acef810000000083816142b3576142b26167f9565b5b0492506020810190505b662386f26fc1000083106142ec57662386f26fc1000083816142e2576142e16167f9565b5b0492506010810190505b6305f5e1008310614315576305f5e100838161430b5761430a6167f9565b5b0492506008810190505b612710831061433a5761271083816143305761432f6167f9565b5b0492506004810190505b6064831061435d5760648381614353576143526167f9565b5b0492506002810190505b600a831061436c576001810190505b80915050919050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036143e5575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016143dc9190614d13565b60405180910390fd5b5f6143f183835f613501565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361446357816040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161445a9190614da3565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146144d7578382826040517f64283d7b0000000000000000000000000000000000000000000000000000000081526004016144ce93929190615859565b60405180910390fd5b50505050565b5f5f6144e884613f12565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161461452957614528818486614947565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146145b4576145685f855f5f613f4b565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461463357600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b60088054905060095f8381526020019081526020015f2081905550600881908060018154018082558091505060019003905f5260205f20015f909190919091505550565b5f61473683611210565b90505f60075f8481526020019081526020015f205490505f60065f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090508282146147d5575f815f8581526020019081526020015f2054905080825f8581526020019081526020015f20819055508260075f8381526020019081526020015f2081905550505b60075f8581526020019081526020015f205f9055805f8481526020019081526020015f205f90555050505050565b5f60016008805490506148169190615c3b565b90505f60095f8481526020019081526020015f205490505f60088381548110614842576148416158b5565b5b905f5260205f20015490508060088381548110614862576148616158b5565b5b905f5260205f2001819055508160095f8381526020019081526020015f208190555060095f8581526020019081526020015f205f905560088054806148aa576148a9616743565b5b600190038181905f5260205f20015f9055905550505050565b5f60016148cf84611210565b6148d99190615c3b565b90508160065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f20819055508060075f8481526020019081526020015f2081905550505050565b614952838383614a0a565b614a05575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036149c657806040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016149bd9190614da3565b60405180910390fd5b81816040517f177e802f0000000000000000000000000000000000000000000000000000000081526004016149fc92919061588e565b60405180910390fd5b505050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015614ac157508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480614a825750614a81848461276a565b5b80614ac057508273ffffffffffffffffffffffffffffffffffffffff16614aa8836134af565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b508054614ad690615829565b5f825580601f10614ae75750614b04565b601f0160209004905f5260205f2090810190614b039190614b07565b5b50565b5b80821115614b1e575f815f905550600101614b08565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b614b6781614b33565b8114614b71575f5ffd5b50565b5f81359050614b8281614b5e565b92915050565b5f60208284031215614b9d57614b9c614b2b565b5b5f614baa84828501614b74565b91505092915050565b5f8115159050919050565b614bc781614bb3565b82525050565b5f602082019050614be05f830184614bbe565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f614c2882614be6565b614c328185614bf0565b9350614c42818560208601614c00565b614c4b81614c0e565b840191505092915050565b5f6020820190508181035f830152614c6e8184614c1e565b905092915050565b5f819050919050565b614c8881614c76565b8114614c92575f5ffd5b50565b5f81359050614ca381614c7f565b92915050565b5f60208284031215614cbe57614cbd614b2b565b5b5f614ccb84828501614c95565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f614cfd82614cd4565b9050919050565b614d0d81614cf3565b82525050565b5f602082019050614d265f830184614d04565b92915050565b614d3581614cf3565b8114614d3f575f5ffd5b50565b5f81359050614d5081614d2c565b92915050565b5f5f60408385031215614d6c57614d6b614b2b565b5b5f614d7985828601614d42565b9250506020614d8a85828601614c95565b9150509250929050565b614d9d81614c76565b82525050565b5f602082019050614db65f830184614d94565b92915050565b5f5f5f60608486031215614dd357614dd2614b2b565b5b5f614de086828701614d42565b9350506020614df186828701614d42565b9250506040614e0286828701614c95565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f614e3082614e0c565b614e3a8185614e16565b9350614e4a818560208601614c00565b614e5381614c0e565b840191505092915050565b5f819050919050565b614e7081614e5e565b82525050565b5f60c0820190508181035f830152614e8e8189614e26565b9050614e9d6020830188614e67565b614eaa6040830187614bbe565b614eb76060830186614d94565b614ec46080830185614d94565b614ed160a0830184614d94565b979650505050505050565b614ee581614e5e565b8114614eef575f5ffd5b50565b5f81359050614f0081614edc565b92915050565b5f60208284031215614f1b57614f1a614b2b565b5b5f614f2884828501614ef2565b91505092915050565b5f61014082019050614f455f83018d614d04565b614f52602083018c614d94565b614f5f604083018b614d94565b614f6c606083018a614d94565b8181036080830152614f7e8189614e26565b9050614f8d60a0830188614e67565b614f9a60c0830187614e67565b614fa760e0830186614bbe565b614fb5610100830185614d94565b614fc3610120830184614d94565b9b9a5050505050505050505050565b5f5f5f60608486031215614fe957614fe8614b2b565b5b5f614ff686828701614c95565b935050602061500786828701614c95565b925050604061501886828701614c95565b9150509250925092565b5f6020828403121561503757615036614b2b565b5b5f61504484828501614d42565b91505092915050565b5f6020820190506150605f830184614e67565b92915050565b61506f81614bb3565b8114615079575f5ffd5b50565b5f8135905061508a81615066565b92915050565b5f5f604083850312156150a6576150a5614b2b565b5b5f6150b385828601614d42565b92505060206150c48582860161507c565b9150509250929050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126150ef576150ee6150ce565b5b8235905067ffffffffffffffff81111561510c5761510b6150d2565b5b602083019150836001820283011115615128576151276150d6565b5b9250929050565b5f5f5f6040848603121561514657615145614b2b565b5b5f61515386828701614c95565b935050602084013567ffffffffffffffff81111561517457615173614b2f565b5b615180868287016150da565b92509250509250925092565b5f5f83601f8401126151a1576151a06150ce565b5b8235905067ffffffffffffffff8111156151be576151bd6150d2565b5b6020830191508360018202830111156151da576151d96150d6565b5b9250929050565b5f5f5f5f606085870312156151f9576151f8614b2b565b5b5f85013567ffffffffffffffff81111561521657615215614b2f565b5b6152228782880161518c565b9450945050602061523587828801614ef2565b925050604061524687828801614c95565b91505092959194509250565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61528c82614c0e565b810181811067ffffffffffffffff821117156152ab576152aa615256565b5b80604052505050565b5f6152bd614b22565b90506152c98282615283565b919050565b5f67ffffffffffffffff8211156152e8576152e7615256565b5b6152f182614c0e565b9050602081019050919050565b828183375f83830152505050565b5f61531e615319846152ce565b6152b4565b90508281526020810184848401111561533a57615339615252565b5b6153458482856152fe565b509392505050565b5f82601f830112615361576153606150ce565b5b813561537184826020860161530c565b91505092915050565b5f5f5f5f6080858703121561539257615391614b2b565b5b5f61539f87828801614d42565b94505060206153b087828801614d42565b93505060406153c187828801614c95565b925050606085013567ffffffffffffffff8111156153e2576153e1614b2f565b5b6153ee8782880161534d565b91505092959194509250565b5f5f5f5f6060858703121561541257615411614b2b565b5b5f61541f87828801614ef2565b945050602085013567ffffffffffffffff8111156154405761543f614b2f565b5b61544c8782880161518c565b9350935050604061545f87828801614ef2565b91505092959194509250565b5f5f6040838503121561548157615480614b2b565b5b5f61548e85828601614d42565b925050602061549f85828601614d42565b9150509250929050565b5f6020820190508181035f8301526154c18184614e26565b905092915050565b5f5f604083850312156154df576154de614b2b565b5b5f6154ec85828601614c95565b92505060206154fd85828601614c95565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61553981614c76565b82525050565b5f61554a8383615530565b60208301905092915050565b5f602082019050919050565b5f61556c82615507565b6155768185615511565b935061558183615521565b805f5b838110156155b1578151615598888261553f565b97506155a383615556565b925050600181019050615584565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6155f081614cf3565b82525050565b5f61560183836155e7565b60208301905092915050565b5f602082019050919050565b5f615623826155be565b61562d81856155c8565b9350615638836155d8565b805f5b8381101561566857815161564f88826155f6565b975061565a8361560d565b92505060018101905061563b565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6156a781614bb3565b82525050565b5f6156b8838361569e565b60208301905092915050565b5f602082019050919050565b5f6156da82615675565b6156e4818561567f565b93506156ef8361568f565b805f5b8381101561571f57815161570688826156ad565b9750615711836156c4565b9250506001810190506156f2565b5085935050505092915050565b5f6080820190508181035f8301526157448187615562565b905081810360208301526157588186615619565b9050818103604083015261576c8185615562565b9050818103606083015261578081846156d0565b905095945050505050565b5f5f5f5f606085870312156157a3576157a2614b2b565b5b5f6157b087828801614ef2565b945050602085013567ffffffffffffffff8111156157d1576157d0614b2f565b5b6157dd8782880161518c565b935093505060406157f087828801614c95565b91505092959194509250565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061584057607f821691505b602082108103615853576158526157fc565b5b50919050565b5f60608201905061586c5f830186614d04565b6158796020830185614d94565b6158866040830184614d04565b949350505050565b5f6040820190506158a15f830185614d04565b6158ae6020830184614d94565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e6f7420746f6b656e206f776e657200000000000000000000000000000000005f82015250565b5f615916600f83614bf0565b9150615921826158e2565b602082019050919050565b5f6020820190508181035f8301526159438161590a565b9050919050565b7f496e76616c69642074696d656f757400000000000000000000000000000000005f82015250565b5f61597e600f83614bf0565b91506159898261594a565b602082019050919050565b5f6020820190508181035f8301526159ab81615972565b9050919050565b5f8160601b9050919050565b5f6159c8826159b2565b9050919050565b5f6159d9826159be565b9050919050565b6159f16159ec82614cf3565b6159cf565b82525050565b5f819050919050565b615a11615a0c82614c76565b6159f7565b82525050565b5f615a2282856159e0565b601482019150615a328284615a00565b6020820191508190509392505050565b7f436c756520616c726561647920736f6c766564000000000000000000000000005f82015250565b5f615a76601383614bf0565b9150615a8182615a42565b602082019050919050565b5f6020820190508181035f830152615aa381615a6a565b9050919050565b5f81905092915050565b5f615abf8385615aaa565b9350615acc8385846152fe565b82840190509392505050565b5f615ae4828486615ab4565b91508190509392505050565b5f615afb8385614bf0565b9350615b088385846152fe565b615b1183614c0e565b840190509392505050565b5f6020820190508181035f830152615b35818486615af0565b90509392505050565b7f4e6f7420746865206275796572000000000000000000000000000000000000005f82015250565b5f615b72600d83614bf0565b9150615b7d82615b3e565b602082019050919050565b5f6020820190508181035f830152615b9f81615b66565b9050919050565b7f50726f6f66206e6f74207965742070726f7669646564000000000000000000005f82015250565b5f615bda601683614bf0565b9150615be582615ba6565b602082019050919050565b5f6020820190508181035f830152615c0781615bce565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f615c4582614c76565b9150615c5083614c76565b9250828203905081811115615c6857615c67615c0e565b5b92915050565b7f50726f6f6620766572696669636174696f6e20657870697265640000000000005f82015250565b5f615ca2601a83614bf0565b9150615cad82615c6e565b602082019050919050565b5f6020820190508181035f830152615ccf81615c96565b9050919050565b7f5472616e7366657220646f6573206e6f742065786973740000000000000000005f82015250565b5f615d0a601783614bf0565b9150615d1582615cd6565b602082019050919050565b5f6020820190508181035f830152615d3781615cfe565b9050919050565b7f4e6f7420617574686f72697a656420746f2063616e63656c00000000000000005f82015250565b5f615d72601883614bf0565b9150615d7d82615d3e565b602082019050919050565b5f6020820190508181035f830152615d9f81615d66565b9050919050565b50565b5f615db45f83615aaa565b9150615dbf82615da6565b5f82019050919050565b5f615dd382615da9565b9150819050919050565b7f4661696c656420746f2073656e642045746865720000000000000000000000005f82015250565b5f615e11601483614bf0565b9150615e1c82615ddd565b602082019050919050565b5f6020820190508181035f830152615e3e81615e05565b9050919050565b5f615e4f82614c76565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615e8157615e80615c0e565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302615ee87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82615ead565b615ef28683615ead565b95508019841693508086168417925050509392505050565b5f819050919050565b5f615f2d615f28615f2384614c76565b615f0a565b614c76565b9050919050565b5f819050919050565b615f4683615f13565b615f5a615f5282615f34565b848454615eb9565b825550505050565b5f5f905090565b615f71615f62565b615f7c818484615f3d565b505050565b5b81811015615f9f57615f945f82615f69565b600181019050615f82565b5050565b601f821115615fe457615fb581615e8c565b615fbe84615e9e565b81016020851015615fcd578190505b615fe1615fd985615e9e565b830182615f81565b50505b505050565b5f82821c905092915050565b5f6160045f1984600802615fe9565b1980831691505092915050565b5f61601c8383615ff5565b9150826002028217905092915050565b61603582614e0c565b67ffffffffffffffff81111561604e5761604d615256565b5b6160588254615829565b616063828285615fa3565b5f60209050601f831160018114616094575f8415616082578287015190505b61608c8582616011565b8655506160f3565b601f1984166160a286615e8c565b5f5b828110156160c9578489015182556001820191506020850194506020810190506160a4565b868310156160e657848901516160e2601f891682615ff5565b8355505b6001600288020188555050505b505050505050565b7f5472616e736665722065787069726564000000000000000000000000000000005f82015250565b5f61612f601083614bf0565b915061613a826160fb565b602082019050919050565b5f6020820190508181035f83015261615c81616123565b9050919050565b7f50726f6f6620746f6f2073686f727400000000000000000000000000000000005f82015250565b5f616197600f83614bf0565b91506161a282616163565b602082019050919050565b5f6020820190508181035f8301526161c48161618b565b9050919050565b5f63ffffffff82169050919050565b5f6161e4826161cb565b91506161ef836161cb565b9250828201905063ffffffff81111561620b5761620a615c0e565b5b92915050565b7f496e76616c69642070726f6f66207374727563747572650000000000000000005f82015250565b5f616245601783614bf0565b915061625082616211565b602082019050919050565b5f6020820190508181035f83015261627281616239565b9050919050565b5f616283826161cb565b915061628e836161cb565b9250828203905063ffffffff8111156162aa576162a9615c0e565b5b92915050565b5f5ffd5b5f5ffd5b5f5f858511156162cb576162ca6162b0565b5b838611156162dc576162db6162b4565b5b6001850283019150848603905094509492505050565b5f82905092915050565b5f61630783836162f2565b826163128135614e5e565b925060208210156163525761634d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83602003600802615ead565b831692505b505092915050565b61636483836162f2565b67ffffffffffffffff81111561637d5761637c615256565b5b6163878254615829565b616392828285615fa3565b5f601f8311600181146163bf575f84156163ad578287013590505b6163b78582616011565b86555061641e565b601f1984166163cd86615e8c565b5f5b828110156163f4578489013582556001820191506020850194506020810190506163cf565b86831015616411578489013561640d601f891682615ff5565b8355505b6001600288020188555050505b50505050505050565b5f6164328385614e16565b935061643f8385846152fe565b61644883614c0e565b840190509392505050565b5f6060820190508181035f83015261646c818688616427565b905061647b6020830185614e67565b6164886040830184614e67565b95945050505050565b5f81905092915050565b5f6164a582614be6565b6164af8185616491565b93506164bf818560208601614c00565b80840191505092915050565b5f6164d6828561649b565b91506164e2828461649b565b91508190509392505050565b7f5472616e7366657220616c726561647920696e697469617465640000000000005f82015250565b5f616522601a83614bf0565b915061652d826164ee565b602082019050919050565b5f6020820190508181035f83015261654f81616516565b9050919050565b5f61656082614c76565b915061656b83614c76565b925082820190508082111561658357616582615c0e565b5b92915050565b7f50726f6f66206e6f7420766572696669656400000000000000000000000000005f82015250565b5f6165bd601283614bf0565b91506165c882616589565b602082019050919050565b5f6020820190508181035f8301526165ea816165b1565b9050919050565b7f5472616e7366657220636f6d706c6574696f6e206578706972656400000000005f82015250565b5f616625601b83614bf0565b9150616630826165f1565b602082019050919050565b5f6020820190508181035f83015261665281616619565b9050919050565b7f436f6e74656e742068617368206d69736d6174636800000000000000000000005f82015250565b5f61668d601583614bf0565b915061669882616659565b602082019050919050565b5f6020820190508181035f8301526166ba81616681565b9050919050565b5f6166cc8284615a00565b60208201915081905092915050565b7f522076616c75652068617368206d69736d6174636800000000000000000000005f82015250565b5f61670f601583614bf0565b915061671a826166db565b602082019050919050565b5f6020820190508181035f83015261673c81616703565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f6080820190506167835f830187614d04565b6167906020830186614d04565b61679d6040830185614d94565b81810360608301526167af8184614e26565b905095945050505050565b5f815190506167c881614b5e565b92915050565b5f602082840312156167e3576167e2614b2b565b5b5f6167f0848285016167ba565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffdfea2646970667358221220af07b23850337d69c9676d70d3d464a8d73fa5b3ffb332f231eb4b689f1b581664736f6c634300081c0033",
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
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 salePrice, uint256 rValue, uint256 timeout)
func (_Skavenge *SkavengeCaller) Clues(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
	SalePrice         *big.Int
	RValue            *big.Int
	Timeout           *big.Int
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

	return *outstruct, err

}

// Clues is a free data retrieval call binding the contract method 0x30f37c7f.
//
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 salePrice, uint256 rValue, uint256 timeout)
func (_Skavenge *SkavengeSession) Clues(arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
	SalePrice         *big.Int
	RValue            *big.Int
	Timeout           *big.Int
}, error) {
	return _Skavenge.Contract.Clues(&_Skavenge.CallOpts, arg0)
}

// Clues is a free data retrieval call binding the contract method 0x30f37c7f.
//
// Solidity: function clues(uint256 ) view returns(bytes encryptedContents, bytes32 solutionHash, bool isSolved, uint256 salePrice, uint256 rValue, uint256 timeout)
func (_Skavenge *SkavengeCallerSession) Clues(arg0 *big.Int) (struct {
	EncryptedContents []byte
	SolutionHash      [32]byte
	IsSolved          bool
	SalePrice         *big.Int
	RValue            *big.Int
	Timeout           *big.Int
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
