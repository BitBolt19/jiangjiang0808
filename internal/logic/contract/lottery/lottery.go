// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lottery

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

// LotteryLotteryConfig is an auto generated low-level Go binding around an user-defined struct.
type LotteryLotteryConfig struct {
	StartTime      *big.Int
	EndTime        *big.Int
	WinnerRatio    *big.Int
	RequireEther   *big.Int
	RequiredNfts   []common.Address
	RequiredTokens []LotteryRequiredToken
	IsTrigger      bool
	IsActive       bool
}

// LotteryRequiredToken is an auto generated low-level Go binding around an user-defined struct.
type LotteryRequiredToken struct {
	Token  common.Address
	Amount *big.Int
}

// LotteryMetaData contains all meta data concerning the Lottery contract.
var LotteryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"UserParticipated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"UserWinner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"action\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_winnerRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requireEther\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_requiredNfts\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_requiredTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_requireTokensAmount\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"_isTrigger\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"addLotteryConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allUsers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getLotteryConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winnerRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requireEther\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"requiredNfts\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structLottery.RequiredToken[]\",\"name\":\"requiredTokens\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"isTrigger\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structLottery.LotteryConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLotteryConfigs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winnerRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requireEther\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"requiredNfts\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structLottery.RequiredToken[]\",\"name\":\"requiredTokens\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"isTrigger\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structLottery.LotteryConfig[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getUserParticipatedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getUserParticipatedList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getUserWinnerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiveAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCanLottery\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isUserWinner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lotteryConfigs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winnerRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requireEther\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isTrigger\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setEnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiveAddress\",\"type\":\"address\"}],\"name\":\"setReceiveAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerCurrentLottery\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"triggerLottery\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userParticipated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userParticipatedList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userWinner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userWinnerList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LotteryABI is the input ABI used to generate the binding from.
// Deprecated: Use LotteryMetaData.ABI instead.
var LotteryABI = LotteryMetaData.ABI

// Lottery is an auto generated Go binding around an Ethereum contract.
type Lottery struct {
	LotteryCaller     // Read-only binding to the contract
	LotteryTransactor // Write-only binding to the contract
	LotteryFilterer   // Log filterer for contract events
}

// LotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LotterySession struct {
	Contract     *Lottery          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LotteryCallerSession struct {
	Contract *LotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LotteryTransactorSession struct {
	Contract     *LotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryRaw struct {
	Contract *Lottery // Generic contract binding to access the raw methods on
}

// LotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryCallerRaw struct {
	Contract *LotteryCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LotteryTransactorRaw struct {
	Contract *LotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLottery creates a new instance of Lottery, bound to a specific deployed contract.
func NewLottery(address common.Address, backend bind.ContractBackend) (*Lottery, error) {
	contract, err := bindLottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lottery{LotteryCaller: LotteryCaller{contract: contract}, LotteryTransactor: LotteryTransactor{contract: contract}, LotteryFilterer: LotteryFilterer{contract: contract}}, nil
}

// NewLotteryCaller creates a new read-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryCaller(address common.Address, caller bind.ContractCaller) (*LotteryCaller, error) {
	contract, err := bindLottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryCaller{contract: contract}, nil
}

// NewLotteryTransactor creates a new write-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryTransactor, error) {
	contract, err := bindLottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryTransactor{contract: contract}, nil
}

// NewLotteryFilterer creates a new log filterer instance of Lottery, bound to a specific deployed contract.
func NewLotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryFilterer, error) {
	contract, err := bindLottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryFilterer{contract: contract}, nil
}

// bindLottery binds a generic wrapper to an already deployed contract.
func bindLottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LotteryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.LotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Lottery *LotteryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Lottery *LotterySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Lottery.Contract.DEFAULTADMINROLE(&_Lottery.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Lottery *LotteryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Lottery.Contract.DEFAULTADMINROLE(&_Lottery.CallOpts)
}

// AllUsers is a free data retrieval call binding the contract method 0xa2bdedf4.
//
// Solidity: function allUsers(uint256 ) view returns(address)
func (_Lottery *LotteryCaller) AllUsers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "allUsers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllUsers is a free data retrieval call binding the contract method 0xa2bdedf4.
//
// Solidity: function allUsers(uint256 ) view returns(address)
func (_Lottery *LotterySession) AllUsers(arg0 *big.Int) (common.Address, error) {
	return _Lottery.Contract.AllUsers(&_Lottery.CallOpts, arg0)
}

// AllUsers is a free data retrieval call binding the contract method 0xa2bdedf4.
//
// Solidity: function allUsers(uint256 ) view returns(address)
func (_Lottery *LotteryCallerSession) AllUsers(arg0 *big.Int) (common.Address, error) {
	return _Lottery.Contract.AllUsers(&_Lottery.CallOpts, arg0)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Lottery *LotteryCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Lottery *LotterySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Lottery.Contract.Eip712Domain(&_Lottery.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Lottery *LotteryCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Lottery.Contract.Eip712Domain(&_Lottery.CallOpts)
}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Lottery *LotteryCaller) Enable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "enable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Lottery *LotterySession) Enable() (bool, error) {
	return _Lottery.Contract.Enable(&_Lottery.CallOpts)
}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Lottery *LotteryCallerSession) Enable() (bool, error) {
	return _Lottery.Contract.Enable(&_Lottery.CallOpts)
}

// GetCurrentRound is a free data retrieval call binding the contract method 0xa32bf597.
//
// Solidity: function getCurrentRound() view returns(uint256)
func (_Lottery *LotteryCaller) GetCurrentRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getCurrentRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRound is a free data retrieval call binding the contract method 0xa32bf597.
//
// Solidity: function getCurrentRound() view returns(uint256)
func (_Lottery *LotterySession) GetCurrentRound() (*big.Int, error) {
	return _Lottery.Contract.GetCurrentRound(&_Lottery.CallOpts)
}

// GetCurrentRound is a free data retrieval call binding the contract method 0xa32bf597.
//
// Solidity: function getCurrentRound() view returns(uint256)
func (_Lottery *LotteryCallerSession) GetCurrentRound() (*big.Int, error) {
	return _Lottery.Contract.GetCurrentRound(&_Lottery.CallOpts)
}

// GetLotteryConfig is a free data retrieval call binding the contract method 0x19457a1c.
//
// Solidity: function getLotteryConfig(uint256 _round) view returns((uint256,uint256,uint256,uint256,address[],(address,uint256)[],bool,bool))
func (_Lottery *LotteryCaller) GetLotteryConfig(opts *bind.CallOpts, _round *big.Int) (LotteryLotteryConfig, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getLotteryConfig", _round)

	if err != nil {
		return *new(LotteryLotteryConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(LotteryLotteryConfig)).(*LotteryLotteryConfig)

	return out0, err

}

// GetLotteryConfig is a free data retrieval call binding the contract method 0x19457a1c.
//
// Solidity: function getLotteryConfig(uint256 _round) view returns((uint256,uint256,uint256,uint256,address[],(address,uint256)[],bool,bool))
func (_Lottery *LotterySession) GetLotteryConfig(_round *big.Int) (LotteryLotteryConfig, error) {
	return _Lottery.Contract.GetLotteryConfig(&_Lottery.CallOpts, _round)
}

// GetLotteryConfig is a free data retrieval call binding the contract method 0x19457a1c.
//
// Solidity: function getLotteryConfig(uint256 _round) view returns((uint256,uint256,uint256,uint256,address[],(address,uint256)[],bool,bool))
func (_Lottery *LotteryCallerSession) GetLotteryConfig(_round *big.Int) (LotteryLotteryConfig, error) {
	return _Lottery.Contract.GetLotteryConfig(&_Lottery.CallOpts, _round)
}

// GetLotteryConfigs is a free data retrieval call binding the contract method 0x2f6e8946.
//
// Solidity: function getLotteryConfigs() view returns((uint256,uint256,uint256,uint256,address[],(address,uint256)[],bool,bool)[])
func (_Lottery *LotteryCaller) GetLotteryConfigs(opts *bind.CallOpts) ([]LotteryLotteryConfig, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getLotteryConfigs")

	if err != nil {
		return *new([]LotteryLotteryConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]LotteryLotteryConfig)).(*[]LotteryLotteryConfig)

	return out0, err

}

// GetLotteryConfigs is a free data retrieval call binding the contract method 0x2f6e8946.
//
// Solidity: function getLotteryConfigs() view returns((uint256,uint256,uint256,uint256,address[],(address,uint256)[],bool,bool)[])
func (_Lottery *LotterySession) GetLotteryConfigs() ([]LotteryLotteryConfig, error) {
	return _Lottery.Contract.GetLotteryConfigs(&_Lottery.CallOpts)
}

// GetLotteryConfigs is a free data retrieval call binding the contract method 0x2f6e8946.
//
// Solidity: function getLotteryConfigs() view returns((uint256,uint256,uint256,uint256,address[],(address,uint256)[],bool,bool)[])
func (_Lottery *LotteryCallerSession) GetLotteryConfigs() ([]LotteryLotteryConfig, error) {
	return _Lottery.Contract.GetLotteryConfigs(&_Lottery.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Lottery *LotteryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Lottery *LotterySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Lottery.Contract.GetRoleAdmin(&_Lottery.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Lottery *LotteryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Lottery.Contract.GetRoleAdmin(&_Lottery.CallOpts, role)
}

// GetUserParticipatedCount is a free data retrieval call binding the contract method 0x397f1066.
//
// Solidity: function getUserParticipatedCount(uint256 _round) view returns(uint256)
func (_Lottery *LotteryCaller) GetUserParticipatedCount(opts *bind.CallOpts, _round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getUserParticipatedCount", _round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserParticipatedCount is a free data retrieval call binding the contract method 0x397f1066.
//
// Solidity: function getUserParticipatedCount(uint256 _round) view returns(uint256)
func (_Lottery *LotterySession) GetUserParticipatedCount(_round *big.Int) (*big.Int, error) {
	return _Lottery.Contract.GetUserParticipatedCount(&_Lottery.CallOpts, _round)
}

// GetUserParticipatedCount is a free data retrieval call binding the contract method 0x397f1066.
//
// Solidity: function getUserParticipatedCount(uint256 _round) view returns(uint256)
func (_Lottery *LotteryCallerSession) GetUserParticipatedCount(_round *big.Int) (*big.Int, error) {
	return _Lottery.Contract.GetUserParticipatedCount(&_Lottery.CallOpts, _round)
}

// GetUserParticipatedList is a free data retrieval call binding the contract method 0x9d97ea1b.
//
// Solidity: function getUserParticipatedList(uint256 _round) view returns(address[])
func (_Lottery *LotteryCaller) GetUserParticipatedList(opts *bind.CallOpts, _round *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getUserParticipatedList", _round)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUserParticipatedList is a free data retrieval call binding the contract method 0x9d97ea1b.
//
// Solidity: function getUserParticipatedList(uint256 _round) view returns(address[])
func (_Lottery *LotterySession) GetUserParticipatedList(_round *big.Int) ([]common.Address, error) {
	return _Lottery.Contract.GetUserParticipatedList(&_Lottery.CallOpts, _round)
}

// GetUserParticipatedList is a free data retrieval call binding the contract method 0x9d97ea1b.
//
// Solidity: function getUserParticipatedList(uint256 _round) view returns(address[])
func (_Lottery *LotteryCallerSession) GetUserParticipatedList(_round *big.Int) ([]common.Address, error) {
	return _Lottery.Contract.GetUserParticipatedList(&_Lottery.CallOpts, _round)
}

// GetUserWinnerCount is a free data retrieval call binding the contract method 0x272ecd2b.
//
// Solidity: function getUserWinnerCount(uint256 _round) view returns(uint256)
func (_Lottery *LotteryCaller) GetUserWinnerCount(opts *bind.CallOpts, _round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getUserWinnerCount", _round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserWinnerCount is a free data retrieval call binding the contract method 0x272ecd2b.
//
// Solidity: function getUserWinnerCount(uint256 _round) view returns(uint256)
func (_Lottery *LotterySession) GetUserWinnerCount(_round *big.Int) (*big.Int, error) {
	return _Lottery.Contract.GetUserWinnerCount(&_Lottery.CallOpts, _round)
}

// GetUserWinnerCount is a free data retrieval call binding the contract method 0x272ecd2b.
//
// Solidity: function getUserWinnerCount(uint256 _round) view returns(uint256)
func (_Lottery *LotteryCallerSession) GetUserWinnerCount(_round *big.Int) (*big.Int, error) {
	return _Lottery.Contract.GetUserWinnerCount(&_Lottery.CallOpts, _round)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Lottery *LotteryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Lottery *LotterySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Lottery.Contract.HasRole(&_Lottery.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Lottery *LotteryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Lottery.Contract.HasRole(&_Lottery.CallOpts, role, account)
}

// IsCanLottery is a free data retrieval call binding the contract method 0x36beb58f.
//
// Solidity: function isCanLottery() view returns(bool)
func (_Lottery *LotteryCaller) IsCanLottery(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "isCanLottery")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCanLottery is a free data retrieval call binding the contract method 0x36beb58f.
//
// Solidity: function isCanLottery() view returns(bool)
func (_Lottery *LotterySession) IsCanLottery() (bool, error) {
	return _Lottery.Contract.IsCanLottery(&_Lottery.CallOpts)
}

// IsCanLottery is a free data retrieval call binding the contract method 0x36beb58f.
//
// Solidity: function isCanLottery() view returns(bool)
func (_Lottery *LotteryCallerSession) IsCanLottery() (bool, error) {
	return _Lottery.Contract.IsCanLottery(&_Lottery.CallOpts)
}

// IsUserWinner is a free data retrieval call binding the contract method 0x45a07a79.
//
// Solidity: function isUserWinner(uint256 _round, address _user) view returns(bool)
func (_Lottery *LotteryCaller) IsUserWinner(opts *bind.CallOpts, _round *big.Int, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "isUserWinner", _round, _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserWinner is a free data retrieval call binding the contract method 0x45a07a79.
//
// Solidity: function isUserWinner(uint256 _round, address _user) view returns(bool)
func (_Lottery *LotterySession) IsUserWinner(_round *big.Int, _user common.Address) (bool, error) {
	return _Lottery.Contract.IsUserWinner(&_Lottery.CallOpts, _round, _user)
}

// IsUserWinner is a free data retrieval call binding the contract method 0x45a07a79.
//
// Solidity: function isUserWinner(uint256 _round, address _user) view returns(bool)
func (_Lottery *LotteryCallerSession) IsUserWinner(_round *big.Int, _user common.Address) (bool, error) {
	return _Lottery.Contract.IsUserWinner(&_Lottery.CallOpts, _round, _user)
}

// LotteryConfigs is a free data retrieval call binding the contract method 0x46ac96b7.
//
// Solidity: function lotteryConfigs(uint256 ) view returns(uint256 startTime, uint256 endTime, uint256 winnerRatio, uint256 requireEther, bool isTrigger, bool isActive)
func (_Lottery *LotteryCaller) LotteryConfigs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartTime    *big.Int
	EndTime      *big.Int
	WinnerRatio  *big.Int
	RequireEther *big.Int
	IsTrigger    bool
	IsActive     bool
}, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "lotteryConfigs", arg0)

	outstruct := new(struct {
		StartTime    *big.Int
		EndTime      *big.Int
		WinnerRatio  *big.Int
		RequireEther *big.Int
		IsTrigger    bool
		IsActive     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WinnerRatio = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RequireEther = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsTrigger = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.IsActive = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// LotteryConfigs is a free data retrieval call binding the contract method 0x46ac96b7.
//
// Solidity: function lotteryConfigs(uint256 ) view returns(uint256 startTime, uint256 endTime, uint256 winnerRatio, uint256 requireEther, bool isTrigger, bool isActive)
func (_Lottery *LotterySession) LotteryConfigs(arg0 *big.Int) (struct {
	StartTime    *big.Int
	EndTime      *big.Int
	WinnerRatio  *big.Int
	RequireEther *big.Int
	IsTrigger    bool
	IsActive     bool
}, error) {
	return _Lottery.Contract.LotteryConfigs(&_Lottery.CallOpts, arg0)
}

// LotteryConfigs is a free data retrieval call binding the contract method 0x46ac96b7.
//
// Solidity: function lotteryConfigs(uint256 ) view returns(uint256 startTime, uint256 endTime, uint256 winnerRatio, uint256 requireEther, bool isTrigger, bool isActive)
func (_Lottery *LotteryCallerSession) LotteryConfigs(arg0 *big.Int) (struct {
	StartTime    *big.Int
	EndTime      *big.Int
	WinnerRatio  *big.Int
	RequireEther *big.Int
	IsTrigger    bool
	IsActive     bool
}, error) {
	return _Lottery.Contract.LotteryConfigs(&_Lottery.CallOpts, arg0)
}

// ReceiveAddress is a free data retrieval call binding the contract method 0xfffe42e9.
//
// Solidity: function receiveAddress() view returns(address)
func (_Lottery *LotteryCaller) ReceiveAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "receiveAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceiveAddress is a free data retrieval call binding the contract method 0xfffe42e9.
//
// Solidity: function receiveAddress() view returns(address)
func (_Lottery *LotterySession) ReceiveAddress() (common.Address, error) {
	return _Lottery.Contract.ReceiveAddress(&_Lottery.CallOpts)
}

// ReceiveAddress is a free data retrieval call binding the contract method 0xfffe42e9.
//
// Solidity: function receiveAddress() view returns(address)
func (_Lottery *LotteryCallerSession) ReceiveAddress() (common.Address, error) {
	return _Lottery.Contract.ReceiveAddress(&_Lottery.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lottery *LotteryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lottery *LotterySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Lottery.Contract.SupportsInterface(&_Lottery.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lottery *LotteryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Lottery.Contract.SupportsInterface(&_Lottery.CallOpts, interfaceId)
}

// UserExists is a free data retrieval call binding the contract method 0x0e666e49.
//
// Solidity: function userExists(address ) view returns(bool)
func (_Lottery *LotteryCaller) UserExists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "userExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserExists is a free data retrieval call binding the contract method 0x0e666e49.
//
// Solidity: function userExists(address ) view returns(bool)
func (_Lottery *LotterySession) UserExists(arg0 common.Address) (bool, error) {
	return _Lottery.Contract.UserExists(&_Lottery.CallOpts, arg0)
}

// UserExists is a free data retrieval call binding the contract method 0x0e666e49.
//
// Solidity: function userExists(address ) view returns(bool)
func (_Lottery *LotteryCallerSession) UserExists(arg0 common.Address) (bool, error) {
	return _Lottery.Contract.UserExists(&_Lottery.CallOpts, arg0)
}

// UserParticipated is a free data retrieval call binding the contract method 0x3ad25ebc.
//
// Solidity: function userParticipated(uint256 , address ) view returns(bool)
func (_Lottery *LotteryCaller) UserParticipated(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "userParticipated", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserParticipated is a free data retrieval call binding the contract method 0x3ad25ebc.
//
// Solidity: function userParticipated(uint256 , address ) view returns(bool)
func (_Lottery *LotterySession) UserParticipated(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Lottery.Contract.UserParticipated(&_Lottery.CallOpts, arg0, arg1)
}

// UserParticipated is a free data retrieval call binding the contract method 0x3ad25ebc.
//
// Solidity: function userParticipated(uint256 , address ) view returns(bool)
func (_Lottery *LotteryCallerSession) UserParticipated(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Lottery.Contract.UserParticipated(&_Lottery.CallOpts, arg0, arg1)
}

// UserParticipatedList is a free data retrieval call binding the contract method 0x58c7cda0.
//
// Solidity: function userParticipatedList(uint256 , uint256 ) view returns(address)
func (_Lottery *LotteryCaller) UserParticipatedList(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "userParticipatedList", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserParticipatedList is a free data retrieval call binding the contract method 0x58c7cda0.
//
// Solidity: function userParticipatedList(uint256 , uint256 ) view returns(address)
func (_Lottery *LotterySession) UserParticipatedList(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Lottery.Contract.UserParticipatedList(&_Lottery.CallOpts, arg0, arg1)
}

// UserParticipatedList is a free data retrieval call binding the contract method 0x58c7cda0.
//
// Solidity: function userParticipatedList(uint256 , uint256 ) view returns(address)
func (_Lottery *LotteryCallerSession) UserParticipatedList(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Lottery.Contract.UserParticipatedList(&_Lottery.CallOpts, arg0, arg1)
}

// UserWinner is a free data retrieval call binding the contract method 0x7258fb90.
//
// Solidity: function userWinner(uint256 , address ) view returns(bool)
func (_Lottery *LotteryCaller) UserWinner(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "userWinner", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserWinner is a free data retrieval call binding the contract method 0x7258fb90.
//
// Solidity: function userWinner(uint256 , address ) view returns(bool)
func (_Lottery *LotterySession) UserWinner(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Lottery.Contract.UserWinner(&_Lottery.CallOpts, arg0, arg1)
}

// UserWinner is a free data retrieval call binding the contract method 0x7258fb90.
//
// Solidity: function userWinner(uint256 , address ) view returns(bool)
func (_Lottery *LotteryCallerSession) UserWinner(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Lottery.Contract.UserWinner(&_Lottery.CallOpts, arg0, arg1)
}

// UserWinnerList is a free data retrieval call binding the contract method 0xf4868005.
//
// Solidity: function userWinnerList(uint256 , uint256 ) view returns(address)
func (_Lottery *LotteryCaller) UserWinnerList(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "userWinnerList", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserWinnerList is a free data retrieval call binding the contract method 0xf4868005.
//
// Solidity: function userWinnerList(uint256 , uint256 ) view returns(address)
func (_Lottery *LotterySession) UserWinnerList(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Lottery.Contract.UserWinnerList(&_Lottery.CallOpts, arg0, arg1)
}

// UserWinnerList is a free data retrieval call binding the contract method 0xf4868005.
//
// Solidity: function userWinnerList(uint256 , uint256 ) view returns(address)
func (_Lottery *LotteryCallerSession) UserWinnerList(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Lottery.Contract.UserWinnerList(&_Lottery.CallOpts, arg0, arg1)
}

// Action is a paid mutator transaction binding the contract method 0x0a7a1c4d.
//
// Solidity: function action() payable returns()
func (_Lottery *LotteryTransactor) Action(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "action")
}

// Action is a paid mutator transaction binding the contract method 0x0a7a1c4d.
//
// Solidity: function action() payable returns()
func (_Lottery *LotterySession) Action() (*types.Transaction, error) {
	return _Lottery.Contract.Action(&_Lottery.TransactOpts)
}

// Action is a paid mutator transaction binding the contract method 0x0a7a1c4d.
//
// Solidity: function action() payable returns()
func (_Lottery *LotteryTransactorSession) Action() (*types.Transaction, error) {
	return _Lottery.Contract.Action(&_Lottery.TransactOpts)
}

// AddLotteryConfig is a paid mutator transaction binding the contract method 0xc0766a56.
//
// Solidity: function addLotteryConfig(uint256 _startTime, uint256 _endTime, uint256 _winnerRatio, uint256 _requireEther, address[] _requiredNfts, address[] _requiredTokens, uint256[] _requireTokensAmount, bool _isTrigger, bool _isActive) returns()
func (_Lottery *LotteryTransactor) AddLotteryConfig(opts *bind.TransactOpts, _startTime *big.Int, _endTime *big.Int, _winnerRatio *big.Int, _requireEther *big.Int, _requiredNfts []common.Address, _requiredTokens []common.Address, _requireTokensAmount []*big.Int, _isTrigger bool, _isActive bool) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "addLotteryConfig", _startTime, _endTime, _winnerRatio, _requireEther, _requiredNfts, _requiredTokens, _requireTokensAmount, _isTrigger, _isActive)
}

// AddLotteryConfig is a paid mutator transaction binding the contract method 0xc0766a56.
//
// Solidity: function addLotteryConfig(uint256 _startTime, uint256 _endTime, uint256 _winnerRatio, uint256 _requireEther, address[] _requiredNfts, address[] _requiredTokens, uint256[] _requireTokensAmount, bool _isTrigger, bool _isActive) returns()
func (_Lottery *LotterySession) AddLotteryConfig(_startTime *big.Int, _endTime *big.Int, _winnerRatio *big.Int, _requireEther *big.Int, _requiredNfts []common.Address, _requiredTokens []common.Address, _requireTokensAmount []*big.Int, _isTrigger bool, _isActive bool) (*types.Transaction, error) {
	return _Lottery.Contract.AddLotteryConfig(&_Lottery.TransactOpts, _startTime, _endTime, _winnerRatio, _requireEther, _requiredNfts, _requiredTokens, _requireTokensAmount, _isTrigger, _isActive)
}

// AddLotteryConfig is a paid mutator transaction binding the contract method 0xc0766a56.
//
// Solidity: function addLotteryConfig(uint256 _startTime, uint256 _endTime, uint256 _winnerRatio, uint256 _requireEther, address[] _requiredNfts, address[] _requiredTokens, uint256[] _requireTokensAmount, bool _isTrigger, bool _isActive) returns()
func (_Lottery *LotteryTransactorSession) AddLotteryConfig(_startTime *big.Int, _endTime *big.Int, _winnerRatio *big.Int, _requireEther *big.Int, _requiredNfts []common.Address, _requiredTokens []common.Address, _requireTokensAmount []*big.Int, _isTrigger bool, _isActive bool) (*types.Transaction, error) {
	return _Lottery.Contract.AddLotteryConfig(&_Lottery.TransactOpts, _startTime, _endTime, _winnerRatio, _requireEther, _requiredNfts, _requiredTokens, _requireTokensAmount, _isTrigger, _isActive)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Lottery *LotteryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Lottery *LotterySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.GrantRole(&_Lottery.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Lottery *LotteryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.GrantRole(&_Lottery.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _receiveAddress) returns()
func (_Lottery *LotteryTransactor) Initialize(opts *bind.TransactOpts, _receiveAddress common.Address) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "initialize", _receiveAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _receiveAddress) returns()
func (_Lottery *LotterySession) Initialize(_receiveAddress common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.Initialize(&_Lottery.TransactOpts, _receiveAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _receiveAddress) returns()
func (_Lottery *LotteryTransactorSession) Initialize(_receiveAddress common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.Initialize(&_Lottery.TransactOpts, _receiveAddress)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Lottery *LotteryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Lottery *LotterySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.RenounceRole(&_Lottery.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Lottery *LotteryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.RenounceRole(&_Lottery.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Lottery *LotteryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Lottery *LotterySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.RevokeRole(&_Lottery.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Lottery *LotteryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.RevokeRole(&_Lottery.TransactOpts, role, account)
}

// SetEnable is a paid mutator transaction binding the contract method 0x7726bed3.
//
// Solidity: function setEnable(bool _enable) returns()
func (_Lottery *LotteryTransactor) SetEnable(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "setEnable", _enable)
}

// SetEnable is a paid mutator transaction binding the contract method 0x7726bed3.
//
// Solidity: function setEnable(bool _enable) returns()
func (_Lottery *LotterySession) SetEnable(_enable bool) (*types.Transaction, error) {
	return _Lottery.Contract.SetEnable(&_Lottery.TransactOpts, _enable)
}

// SetEnable is a paid mutator transaction binding the contract method 0x7726bed3.
//
// Solidity: function setEnable(bool _enable) returns()
func (_Lottery *LotteryTransactorSession) SetEnable(_enable bool) (*types.Transaction, error) {
	return _Lottery.Contract.SetEnable(&_Lottery.TransactOpts, _enable)
}

// SetReceiveAddress is a paid mutator transaction binding the contract method 0x5ec4b7a8.
//
// Solidity: function setReceiveAddress(address _receiveAddress) returns()
func (_Lottery *LotteryTransactor) SetReceiveAddress(opts *bind.TransactOpts, _receiveAddress common.Address) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "setReceiveAddress", _receiveAddress)
}

// SetReceiveAddress is a paid mutator transaction binding the contract method 0x5ec4b7a8.
//
// Solidity: function setReceiveAddress(address _receiveAddress) returns()
func (_Lottery *LotterySession) SetReceiveAddress(_receiveAddress common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.SetReceiveAddress(&_Lottery.TransactOpts, _receiveAddress)
}

// SetReceiveAddress is a paid mutator transaction binding the contract method 0x5ec4b7a8.
//
// Solidity: function setReceiveAddress(address _receiveAddress) returns()
func (_Lottery *LotteryTransactorSession) SetReceiveAddress(_receiveAddress common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.SetReceiveAddress(&_Lottery.TransactOpts, _receiveAddress)
}

// TriggerCurrentLottery is a paid mutator transaction binding the contract method 0xe78f84d8.
//
// Solidity: function triggerCurrentLottery() payable returns()
func (_Lottery *LotteryTransactor) TriggerCurrentLottery(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "triggerCurrentLottery")
}

// TriggerCurrentLottery is a paid mutator transaction binding the contract method 0xe78f84d8.
//
// Solidity: function triggerCurrentLottery() payable returns()
func (_Lottery *LotterySession) TriggerCurrentLottery() (*types.Transaction, error) {
	return _Lottery.Contract.TriggerCurrentLottery(&_Lottery.TransactOpts)
}

// TriggerCurrentLottery is a paid mutator transaction binding the contract method 0xe78f84d8.
//
// Solidity: function triggerCurrentLottery() payable returns()
func (_Lottery *LotteryTransactorSession) TriggerCurrentLottery() (*types.Transaction, error) {
	return _Lottery.Contract.TriggerCurrentLottery(&_Lottery.TransactOpts)
}

// TriggerLottery is a paid mutator transaction binding the contract method 0xadcf23b4.
//
// Solidity: function triggerLottery(uint256 _round) payable returns()
func (_Lottery *LotteryTransactor) TriggerLottery(opts *bind.TransactOpts, _round *big.Int) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "triggerLottery", _round)
}

// TriggerLottery is a paid mutator transaction binding the contract method 0xadcf23b4.
//
// Solidity: function triggerLottery(uint256 _round) payable returns()
func (_Lottery *LotterySession) TriggerLottery(_round *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.TriggerLottery(&_Lottery.TransactOpts, _round)
}

// TriggerLottery is a paid mutator transaction binding the contract method 0xadcf23b4.
//
// Solidity: function triggerLottery(uint256 _round) payable returns()
func (_Lottery *LotteryTransactorSession) TriggerLottery(_round *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.TriggerLottery(&_Lottery.TransactOpts, _round)
}

// LotteryEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Lottery contract.
type LotteryEIP712DomainChangedIterator struct {
	Event *LotteryEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *LotteryEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryEIP712DomainChanged)
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
		it.Event = new(LotteryEIP712DomainChanged)
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
func (it *LotteryEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryEIP712DomainChanged represents a EIP712DomainChanged event raised by the Lottery contract.
type LotteryEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Lottery *LotteryFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*LotteryEIP712DomainChangedIterator, error) {

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &LotteryEIP712DomainChangedIterator{contract: _Lottery.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Lottery *LotteryFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *LotteryEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryEIP712DomainChanged)
				if err := _Lottery.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Lottery *LotteryFilterer) ParseEIP712DomainChanged(log types.Log) (*LotteryEIP712DomainChanged, error) {
	event := new(LotteryEIP712DomainChanged)
	if err := _Lottery.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Lottery contract.
type LotteryInitializedIterator struct {
	Event *LotteryInitialized // Event containing the contract specifics and raw log

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
func (it *LotteryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryInitialized)
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
		it.Event = new(LotteryInitialized)
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
func (it *LotteryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryInitialized represents a Initialized event raised by the Lottery contract.
type LotteryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lottery *LotteryFilterer) FilterInitialized(opts *bind.FilterOpts) (*LotteryInitializedIterator, error) {

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LotteryInitializedIterator{contract: _Lottery.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lottery *LotteryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LotteryInitialized) (event.Subscription, error) {

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryInitialized)
				if err := _Lottery.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lottery *LotteryFilterer) ParseInitialized(log types.Log) (*LotteryInitialized, error) {
	event := new(LotteryInitialized)
	if err := _Lottery.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Lottery contract.
type LotteryRoleAdminChangedIterator struct {
	Event *LotteryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LotteryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryRoleAdminChanged)
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
		it.Event = new(LotteryRoleAdminChanged)
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
func (it *LotteryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryRoleAdminChanged represents a RoleAdminChanged event raised by the Lottery contract.
type LotteryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Lottery *LotteryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LotteryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LotteryRoleAdminChangedIterator{contract: _Lottery.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Lottery *LotteryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LotteryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryRoleAdminChanged)
				if err := _Lottery.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Lottery *LotteryFilterer) ParseRoleAdminChanged(log types.Log) (*LotteryRoleAdminChanged, error) {
	event := new(LotteryRoleAdminChanged)
	if err := _Lottery.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Lottery contract.
type LotteryRoleGrantedIterator struct {
	Event *LotteryRoleGranted // Event containing the contract specifics and raw log

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
func (it *LotteryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryRoleGranted)
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
		it.Event = new(LotteryRoleGranted)
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
func (it *LotteryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryRoleGranted represents a RoleGranted event raised by the Lottery contract.
type LotteryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lottery *LotteryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LotteryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LotteryRoleGrantedIterator{contract: _Lottery.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lottery *LotteryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LotteryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryRoleGranted)
				if err := _Lottery.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lottery *LotteryFilterer) ParseRoleGranted(log types.Log) (*LotteryRoleGranted, error) {
	event := new(LotteryRoleGranted)
	if err := _Lottery.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Lottery contract.
type LotteryRoleRevokedIterator struct {
	Event *LotteryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LotteryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryRoleRevoked)
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
		it.Event = new(LotteryRoleRevoked)
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
func (it *LotteryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryRoleRevoked represents a RoleRevoked event raised by the Lottery contract.
type LotteryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lottery *LotteryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LotteryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LotteryRoleRevokedIterator{contract: _Lottery.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lottery *LotteryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LotteryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryRoleRevoked)
				if err := _Lottery.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lottery *LotteryFilterer) ParseRoleRevoked(log types.Log) (*LotteryRoleRevoked, error) {
	event := new(LotteryRoleRevoked)
	if err := _Lottery.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryUserParticipatedIterator is returned from FilterUserParticipated and is used to iterate over the raw logs and unpacked data for UserParticipated events raised by the Lottery contract.
type LotteryUserParticipatedIterator struct {
	Event *LotteryUserParticipated // Event containing the contract specifics and raw log

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
func (it *LotteryUserParticipatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryUserParticipated)
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
		it.Event = new(LotteryUserParticipated)
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
func (it *LotteryUserParticipatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryUserParticipatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryUserParticipated represents a UserParticipated event raised by the Lottery contract.
type LotteryUserParticipated struct {
	User  common.Address
	Round *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserParticipated is a free log retrieval operation binding the contract event 0xd00162658a963ed1bfc4ff82d4c5ddf3d2b28835dea3f61c483ceda1f59e1971.
//
// Solidity: event UserParticipated(address indexed user, uint256 round)
func (_Lottery *LotteryFilterer) FilterUserParticipated(opts *bind.FilterOpts, user []common.Address) (*LotteryUserParticipatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "UserParticipated", userRule)
	if err != nil {
		return nil, err
	}
	return &LotteryUserParticipatedIterator{contract: _Lottery.contract, event: "UserParticipated", logs: logs, sub: sub}, nil
}

// WatchUserParticipated is a free log subscription operation binding the contract event 0xd00162658a963ed1bfc4ff82d4c5ddf3d2b28835dea3f61c483ceda1f59e1971.
//
// Solidity: event UserParticipated(address indexed user, uint256 round)
func (_Lottery *LotteryFilterer) WatchUserParticipated(opts *bind.WatchOpts, sink chan<- *LotteryUserParticipated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "UserParticipated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryUserParticipated)
				if err := _Lottery.contract.UnpackLog(event, "UserParticipated", log); err != nil {
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

// ParseUserParticipated is a log parse operation binding the contract event 0xd00162658a963ed1bfc4ff82d4c5ddf3d2b28835dea3f61c483ceda1f59e1971.
//
// Solidity: event UserParticipated(address indexed user, uint256 round)
func (_Lottery *LotteryFilterer) ParseUserParticipated(log types.Log) (*LotteryUserParticipated, error) {
	event := new(LotteryUserParticipated)
	if err := _Lottery.contract.UnpackLog(event, "UserParticipated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryUserWinnerIterator is returned from FilterUserWinner and is used to iterate over the raw logs and unpacked data for UserWinner events raised by the Lottery contract.
type LotteryUserWinnerIterator struct {
	Event *LotteryUserWinner // Event containing the contract specifics and raw log

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
func (it *LotteryUserWinnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryUserWinner)
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
		it.Event = new(LotteryUserWinner)
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
func (it *LotteryUserWinnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryUserWinnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryUserWinner represents a UserWinner event raised by the Lottery contract.
type LotteryUserWinner struct {
	User  common.Address
	Round *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserWinner is a free log retrieval operation binding the contract event 0xe57442c958f1a6ae94215b7eee52d07f8404de028288a16fc181c24a4836f127.
//
// Solidity: event UserWinner(address indexed user, uint256 round)
func (_Lottery *LotteryFilterer) FilterUserWinner(opts *bind.FilterOpts, user []common.Address) (*LotteryUserWinnerIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "UserWinner", userRule)
	if err != nil {
		return nil, err
	}
	return &LotteryUserWinnerIterator{contract: _Lottery.contract, event: "UserWinner", logs: logs, sub: sub}, nil
}

// WatchUserWinner is a free log subscription operation binding the contract event 0xe57442c958f1a6ae94215b7eee52d07f8404de028288a16fc181c24a4836f127.
//
// Solidity: event UserWinner(address indexed user, uint256 round)
func (_Lottery *LotteryFilterer) WatchUserWinner(opts *bind.WatchOpts, sink chan<- *LotteryUserWinner, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "UserWinner", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryUserWinner)
				if err := _Lottery.contract.UnpackLog(event, "UserWinner", log); err != nil {
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

// ParseUserWinner is a log parse operation binding the contract event 0xe57442c958f1a6ae94215b7eee52d07f8404de028288a16fc181c24a4836f127.
//
// Solidity: event UserWinner(address indexed user, uint256 round)
func (_Lottery *LotteryFilterer) ParseUserWinner(log types.Log) (*LotteryUserWinner, error) {
	event := new(LotteryUserWinner)
	if err := _Lottery.contract.UnpackLog(event, "UserWinner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
