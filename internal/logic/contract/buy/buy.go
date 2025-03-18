// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package buy

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

// BuyBatchPrice is an auto generated low-level Go binding around an user-defined struct.
type BuyBatchPrice struct {
	Amount   *big.Int
	EthPrice *big.Int
	UsdPrice *big.Int
	NezPrice *big.Int
}

// BuyMetaData contains all meta data concerning the Buy contract.
var BuyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BuyBlindBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BuyNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OpenBlindBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferBlindBox\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ethPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_usdPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nezPrice\",\"type\":\"uint256\"}],\"name\":\"addBatchPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batchPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nezPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blindBoxEnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"buyBlindBox\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyEnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"buyNft\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"buyNfts\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearNftPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getBatchPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nezPrice\",\"type\":\"uint256\"}],\"internalType\":\"structBuy.BatchPrice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getNftPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"getUserBlindBoxes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"healthcheck\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_receiveAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nez\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftPrices\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ethPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nezPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"openBlindBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_buyEnable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_blindBoxEnable\",\"type\":\"bool\"}],\"name\":\"setBuyAndBlindBoxEnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setEnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ethPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_usdPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nezPrice\",\"type\":\"uint256\"}],\"name\":\"setNftPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiveAddress\",\"type\":\"address\"}],\"name\":\"setReceiveAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"setTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferBlindBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBlindBoxes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BuyABI is the input ABI used to generate the binding from.
// Deprecated: Use BuyMetaData.ABI instead.
var BuyABI = BuyMetaData.ABI

// Buy is an auto generated Go binding around an Ethereum contract.
type Buy struct {
	BuyCaller     // Read-only binding to the contract
	BuyTransactor // Write-only binding to the contract
	BuyFilterer   // Log filterer for contract events
}

// BuyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuySession struct {
	Contract     *Buy              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuyCallerSession struct {
	Contract *BuyCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BuyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuyTransactorSession struct {
	Contract     *BuyTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuyRaw struct {
	Contract *Buy // Generic contract binding to access the raw methods on
}

// BuyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuyCallerRaw struct {
	Contract *BuyCaller // Generic read-only contract binding to access the raw methods on
}

// BuyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuyTransactorRaw struct {
	Contract *BuyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuy creates a new instance of Buy, bound to a specific deployed contract.
func NewBuy(address common.Address, backend bind.ContractBackend) (*Buy, error) {
	contract, err := bindBuy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Buy{BuyCaller: BuyCaller{contract: contract}, BuyTransactor: BuyTransactor{contract: contract}, BuyFilterer: BuyFilterer{contract: contract}}, nil
}

// NewBuyCaller creates a new read-only instance of Buy, bound to a specific deployed contract.
func NewBuyCaller(address common.Address, caller bind.ContractCaller) (*BuyCaller, error) {
	contract, err := bindBuy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuyCaller{contract: contract}, nil
}

// NewBuyTransactor creates a new write-only instance of Buy, bound to a specific deployed contract.
func NewBuyTransactor(address common.Address, transactor bind.ContractTransactor) (*BuyTransactor, error) {
	contract, err := bindBuy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuyTransactor{contract: contract}, nil
}

// NewBuyFilterer creates a new log filterer instance of Buy, bound to a specific deployed contract.
func NewBuyFilterer(address common.Address, filterer bind.ContractFilterer) (*BuyFilterer, error) {
	contract, err := bindBuy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuyFilterer{contract: contract}, nil
}

// bindBuy binds a generic wrapper to an already deployed contract.
func bindBuy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BuyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buy *BuyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Buy.Contract.BuyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buy *BuyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.Contract.BuyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buy *BuyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buy.Contract.BuyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buy *BuyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Buy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buy *BuyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buy *BuyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buy.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Buy *BuyCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Buy *BuySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Buy.Contract.DEFAULTADMINROLE(&_Buy.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Buy *BuyCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Buy.Contract.DEFAULTADMINROLE(&_Buy.CallOpts)
}

// BatchPrices is a free data retrieval call binding the contract method 0x9aa0dd82.
//
// Solidity: function batchPrices(uint256 ) view returns(uint256 amount, uint256 ethPrice, uint256 usdPrice, uint256 nezPrice)
func (_Buy *BuyCaller) BatchPrices(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount   *big.Int
	EthPrice *big.Int
	UsdPrice *big.Int
	NezPrice *big.Int
}, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "batchPrices", arg0)

	outstruct := new(struct {
		Amount   *big.Int
		EthPrice *big.Int
		UsdPrice *big.Int
		NezPrice *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EthPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UsdPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NezPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BatchPrices is a free data retrieval call binding the contract method 0x9aa0dd82.
//
// Solidity: function batchPrices(uint256 ) view returns(uint256 amount, uint256 ethPrice, uint256 usdPrice, uint256 nezPrice)
func (_Buy *BuySession) BatchPrices(arg0 *big.Int) (struct {
	Amount   *big.Int
	EthPrice *big.Int
	UsdPrice *big.Int
	NezPrice *big.Int
}, error) {
	return _Buy.Contract.BatchPrices(&_Buy.CallOpts, arg0)
}

// BatchPrices is a free data retrieval call binding the contract method 0x9aa0dd82.
//
// Solidity: function batchPrices(uint256 ) view returns(uint256 amount, uint256 ethPrice, uint256 usdPrice, uint256 nezPrice)
func (_Buy *BuyCallerSession) BatchPrices(arg0 *big.Int) (struct {
	Amount   *big.Int
	EthPrice *big.Int
	UsdPrice *big.Int
	NezPrice *big.Int
}, error) {
	return _Buy.Contract.BatchPrices(&_Buy.CallOpts, arg0)
}

// BlindBoxEnable is a free data retrieval call binding the contract method 0x9808fa5c.
//
// Solidity: function blindBoxEnable() view returns(bool)
func (_Buy *BuyCaller) BlindBoxEnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "blindBoxEnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlindBoxEnable is a free data retrieval call binding the contract method 0x9808fa5c.
//
// Solidity: function blindBoxEnable() view returns(bool)
func (_Buy *BuySession) BlindBoxEnable() (bool, error) {
	return _Buy.Contract.BlindBoxEnable(&_Buy.CallOpts)
}

// BlindBoxEnable is a free data retrieval call binding the contract method 0x9808fa5c.
//
// Solidity: function blindBoxEnable() view returns(bool)
func (_Buy *BuyCallerSession) BlindBoxEnable() (bool, error) {
	return _Buy.Contract.BlindBoxEnable(&_Buy.CallOpts)
}

// BuyEnable is a free data retrieval call binding the contract method 0x3d389faf.
//
// Solidity: function buyEnable() view returns(bool)
func (_Buy *BuyCaller) BuyEnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "buyEnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BuyEnable is a free data retrieval call binding the contract method 0x3d389faf.
//
// Solidity: function buyEnable() view returns(bool)
func (_Buy *BuySession) BuyEnable() (bool, error) {
	return _Buy.Contract.BuyEnable(&_Buy.CallOpts)
}

// BuyEnable is a free data retrieval call binding the contract method 0x3d389faf.
//
// Solidity: function buyEnable() view returns(bool)
func (_Buy *BuyCallerSession) BuyEnable() (bool, error) {
	return _Buy.Contract.BuyEnable(&_Buy.CallOpts)
}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Buy *BuyCaller) Enable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "enable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Buy *BuySession) Enable() (bool, error) {
	return _Buy.Contract.Enable(&_Buy.CallOpts)
}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Buy *BuyCallerSession) Enable() (bool, error) {
	return _Buy.Contract.Enable(&_Buy.CallOpts)
}

// GetBatchIndex is a free data retrieval call binding the contract method 0x44ec3c07.
//
// Solidity: function getBatchIndex(uint256 _id) view returns(uint256)
func (_Buy *BuyCaller) GetBatchIndex(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "getBatchIndex", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchIndex is a free data retrieval call binding the contract method 0x44ec3c07.
//
// Solidity: function getBatchIndex(uint256 _id) view returns(uint256)
func (_Buy *BuySession) GetBatchIndex(_id *big.Int) (*big.Int, error) {
	return _Buy.Contract.GetBatchIndex(&_Buy.CallOpts, _id)
}

// GetBatchIndex is a free data retrieval call binding the contract method 0x44ec3c07.
//
// Solidity: function getBatchIndex(uint256 _id) view returns(uint256)
func (_Buy *BuyCallerSession) GetBatchIndex(_id *big.Int) (*big.Int, error) {
	return _Buy.Contract.GetBatchIndex(&_Buy.CallOpts, _id)
}

// GetBatchPrice is a free data retrieval call binding the contract method 0x2bd92dd3.
//
// Solidity: function getBatchPrice(uint256 _id) view returns((uint256,uint256,uint256,uint256))
func (_Buy *BuyCaller) GetBatchPrice(opts *bind.CallOpts, _id *big.Int) (BuyBatchPrice, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "getBatchPrice", _id)

	if err != nil {
		return *new(BuyBatchPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(BuyBatchPrice)).(*BuyBatchPrice)

	return out0, err

}

// GetBatchPrice is a free data retrieval call binding the contract method 0x2bd92dd3.
//
// Solidity: function getBatchPrice(uint256 _id) view returns((uint256,uint256,uint256,uint256))
func (_Buy *BuySession) GetBatchPrice(_id *big.Int) (BuyBatchPrice, error) {
	return _Buy.Contract.GetBatchPrice(&_Buy.CallOpts, _id)
}

// GetBatchPrice is a free data retrieval call binding the contract method 0x2bd92dd3.
//
// Solidity: function getBatchPrice(uint256 _id) view returns((uint256,uint256,uint256,uint256))
func (_Buy *BuyCallerSession) GetBatchPrice(_id *big.Int) (BuyBatchPrice, error) {
	return _Buy.Contract.GetBatchPrice(&_Buy.CallOpts, _id)
}

// GetNftPrice is a free data retrieval call binding the contract method 0x2cf7d7e6.
//
// Solidity: function getNftPrice(address _token) view returns(uint256, uint256, uint256)
func (_Buy *BuyCaller) GetNftPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "getNftPrice", _token)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetNftPrice is a free data retrieval call binding the contract method 0x2cf7d7e6.
//
// Solidity: function getNftPrice(address _token) view returns(uint256, uint256, uint256)
func (_Buy *BuySession) GetNftPrice(_token common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Buy.Contract.GetNftPrice(&_Buy.CallOpts, _token)
}

// GetNftPrice is a free data retrieval call binding the contract method 0x2cf7d7e6.
//
// Solidity: function getNftPrice(address _token) view returns(uint256, uint256, uint256)
func (_Buy *BuyCallerSession) GetNftPrice(_token common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Buy.Contract.GetNftPrice(&_Buy.CallOpts, _token)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Buy *BuyCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Buy *BuySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Buy.Contract.GetRoleAdmin(&_Buy.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Buy *BuyCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Buy.Contract.GetRoleAdmin(&_Buy.CallOpts, role)
}

// GetUserBlindBoxes is a free data retrieval call binding the contract method 0xf9f6689a.
//
// Solidity: function getUserBlindBoxes(address _member) view returns(uint256)
func (_Buy *BuyCaller) GetUserBlindBoxes(opts *bind.CallOpts, _member common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "getUserBlindBoxes", _member)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserBlindBoxes is a free data retrieval call binding the contract method 0xf9f6689a.
//
// Solidity: function getUserBlindBoxes(address _member) view returns(uint256)
func (_Buy *BuySession) GetUserBlindBoxes(_member common.Address) (*big.Int, error) {
	return _Buy.Contract.GetUserBlindBoxes(&_Buy.CallOpts, _member)
}

// GetUserBlindBoxes is a free data retrieval call binding the contract method 0xf9f6689a.
//
// Solidity: function getUserBlindBoxes(address _member) view returns(uint256)
func (_Buy *BuyCallerSession) GetUserBlindBoxes(_member common.Address) (*big.Int, error) {
	return _Buy.Contract.GetUserBlindBoxes(&_Buy.CallOpts, _member)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Buy *BuyCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Buy *BuySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Buy.Contract.HasRole(&_Buy.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Buy *BuyCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Buy.Contract.HasRole(&_Buy.CallOpts, role, account)
}

// Healthcheck is a free data retrieval call binding the contract method 0x4f42fc02.
//
// Solidity: function healthcheck() view returns()
func (_Buy *BuyCaller) Healthcheck(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "healthcheck")

	if err != nil {
		return err
	}

	return err

}

// Healthcheck is a free data retrieval call binding the contract method 0x4f42fc02.
//
// Solidity: function healthcheck() view returns()
func (_Buy *BuySession) Healthcheck() error {
	return _Buy.Contract.Healthcheck(&_Buy.CallOpts)
}

// Healthcheck is a free data retrieval call binding the contract method 0x4f42fc02.
//
// Solidity: function healthcheck() view returns()
func (_Buy *BuyCallerSession) Healthcheck() error {
	return _Buy.Contract.Healthcheck(&_Buy.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_Buy *BuyCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_Buy *BuySession) Id() (*big.Int, error) {
	return _Buy.Contract.Id(&_Buy.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_Buy *BuyCallerSession) Id() (*big.Int, error) {
	return _Buy.Contract.Id(&_Buy.CallOpts)
}

// Nez is a free data retrieval call binding the contract method 0xd5bf4d4b.
//
// Solidity: function nez() view returns(address)
func (_Buy *BuyCaller) Nez(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "nez")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nez is a free data retrieval call binding the contract method 0xd5bf4d4b.
//
// Solidity: function nez() view returns(address)
func (_Buy *BuySession) Nez() (common.Address, error) {
	return _Buy.Contract.Nez(&_Buy.CallOpts)
}

// Nez is a free data retrieval call binding the contract method 0xd5bf4d4b.
//
// Solidity: function nez() view returns(address)
func (_Buy *BuyCallerSession) Nez() (common.Address, error) {
	return _Buy.Contract.Nez(&_Buy.CallOpts)
}

// NftPrices is a free data retrieval call binding the contract method 0xd9f10a2b.
//
// Solidity: function nftPrices(uint256 ) view returns(address token, uint256 ethPrice, uint256 usdPrice, uint256 nezPrice)
func (_Buy *BuyCaller) NftPrices(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token    common.Address
	EthPrice *big.Int
	UsdPrice *big.Int
	NezPrice *big.Int
}, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "nftPrices", arg0)

	outstruct := new(struct {
		Token    common.Address
		EthPrice *big.Int
		UsdPrice *big.Int
		NezPrice *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EthPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UsdPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NezPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NftPrices is a free data retrieval call binding the contract method 0xd9f10a2b.
//
// Solidity: function nftPrices(uint256 ) view returns(address token, uint256 ethPrice, uint256 usdPrice, uint256 nezPrice)
func (_Buy *BuySession) NftPrices(arg0 *big.Int) (struct {
	Token    common.Address
	EthPrice *big.Int
	UsdPrice *big.Int
	NezPrice *big.Int
}, error) {
	return _Buy.Contract.NftPrices(&_Buy.CallOpts, arg0)
}

// NftPrices is a free data retrieval call binding the contract method 0xd9f10a2b.
//
// Solidity: function nftPrices(uint256 ) view returns(address token, uint256 ethPrice, uint256 usdPrice, uint256 nezPrice)
func (_Buy *BuyCallerSession) NftPrices(arg0 *big.Int) (struct {
	Token    common.Address
	EthPrice *big.Int
	UsdPrice *big.Int
	NezPrice *big.Int
}, error) {
	return _Buy.Contract.NftPrices(&_Buy.CallOpts, arg0)
}

// ReceiveAddress is a free data retrieval call binding the contract method 0xfffe42e9.
//
// Solidity: function receiveAddress() view returns(address)
func (_Buy *BuyCaller) ReceiveAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "receiveAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceiveAddress is a free data retrieval call binding the contract method 0xfffe42e9.
//
// Solidity: function receiveAddress() view returns(address)
func (_Buy *BuySession) ReceiveAddress() (common.Address, error) {
	return _Buy.Contract.ReceiveAddress(&_Buy.CallOpts)
}

// ReceiveAddress is a free data retrieval call binding the contract method 0xfffe42e9.
//
// Solidity: function receiveAddress() view returns(address)
func (_Buy *BuyCallerSession) ReceiveAddress() (common.Address, error) {
	return _Buy.Contract.ReceiveAddress(&_Buy.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Buy *BuyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Buy *BuySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Buy.Contract.SupportsInterface(&_Buy.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Buy *BuyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Buy.Contract.SupportsInterface(&_Buy.CallOpts, interfaceId)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Buy *BuyCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Buy *BuySession) Usdt() (common.Address, error) {
	return _Buy.Contract.Usdt(&_Buy.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Buy *BuyCallerSession) Usdt() (common.Address, error) {
	return _Buy.Contract.Usdt(&_Buy.CallOpts)
}

// UserBlindBoxes is a free data retrieval call binding the contract method 0xca6bd0fb.
//
// Solidity: function userBlindBoxes(address ) view returns(uint256)
func (_Buy *BuyCaller) UserBlindBoxes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "userBlindBoxes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBlindBoxes is a free data retrieval call binding the contract method 0xca6bd0fb.
//
// Solidity: function userBlindBoxes(address ) view returns(uint256)
func (_Buy *BuySession) UserBlindBoxes(arg0 common.Address) (*big.Int, error) {
	return _Buy.Contract.UserBlindBoxes(&_Buy.CallOpts, arg0)
}

// UserBlindBoxes is a free data retrieval call binding the contract method 0xca6bd0fb.
//
// Solidity: function userBlindBoxes(address ) view returns(uint256)
func (_Buy *BuyCallerSession) UserBlindBoxes(arg0 common.Address) (*big.Int, error) {
	return _Buy.Contract.UserBlindBoxes(&_Buy.CallOpts, arg0)
}

// AddBatchPrice is a paid mutator transaction binding the contract method 0xc400a2a7.
//
// Solidity: function addBatchPrice(uint256 _amount, uint256 _ethPrice, uint256 _usdPrice, uint256 _nezPrice) returns()
func (_Buy *BuyTransactor) AddBatchPrice(opts *bind.TransactOpts, _amount *big.Int, _ethPrice *big.Int, _usdPrice *big.Int, _nezPrice *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "addBatchPrice", _amount, _ethPrice, _usdPrice, _nezPrice)
}

// AddBatchPrice is a paid mutator transaction binding the contract method 0xc400a2a7.
//
// Solidity: function addBatchPrice(uint256 _amount, uint256 _ethPrice, uint256 _usdPrice, uint256 _nezPrice) returns()
func (_Buy *BuySession) AddBatchPrice(_amount *big.Int, _ethPrice *big.Int, _usdPrice *big.Int, _nezPrice *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.AddBatchPrice(&_Buy.TransactOpts, _amount, _ethPrice, _usdPrice, _nezPrice)
}

// AddBatchPrice is a paid mutator transaction binding the contract method 0xc400a2a7.
//
// Solidity: function addBatchPrice(uint256 _amount, uint256 _ethPrice, uint256 _usdPrice, uint256 _nezPrice) returns()
func (_Buy *BuyTransactorSession) AddBatchPrice(_amount *big.Int, _ethPrice *big.Int, _usdPrice *big.Int, _nezPrice *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.AddBatchPrice(&_Buy.TransactOpts, _amount, _ethPrice, _usdPrice, _nezPrice)
}

// BuyBlindBox is a paid mutator transaction binding the contract method 0x62977ae7.
//
// Solidity: function buyBlindBox(uint256 _amount) payable returns()
func (_Buy *BuyTransactor) BuyBlindBox(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "buyBlindBox", _amount)
}

// BuyBlindBox is a paid mutator transaction binding the contract method 0x62977ae7.
//
// Solidity: function buyBlindBox(uint256 _amount) payable returns()
func (_Buy *BuySession) BuyBlindBox(_amount *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.BuyBlindBox(&_Buy.TransactOpts, _amount)
}

// BuyBlindBox is a paid mutator transaction binding the contract method 0x62977ae7.
//
// Solidity: function buyBlindBox(uint256 _amount) payable returns()
func (_Buy *BuyTransactorSession) BuyBlindBox(_amount *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.BuyBlindBox(&_Buy.TransactOpts, _amount)
}

// BuyNft is a paid mutator transaction binding the contract method 0x4ddf4f8f.
//
// Solidity: function buyNft(address _token, uint256 _amount) payable returns()
func (_Buy *BuyTransactor) BuyNft(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "buyNft", _token, _amount)
}

// BuyNft is a paid mutator transaction binding the contract method 0x4ddf4f8f.
//
// Solidity: function buyNft(address _token, uint256 _amount) payable returns()
func (_Buy *BuySession) BuyNft(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.BuyNft(&_Buy.TransactOpts, _token, _amount)
}

// BuyNft is a paid mutator transaction binding the contract method 0x4ddf4f8f.
//
// Solidity: function buyNft(address _token, uint256 _amount) payable returns()
func (_Buy *BuyTransactorSession) BuyNft(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.BuyNft(&_Buy.TransactOpts, _token, _amount)
}

// BuyNfts is a paid mutator transaction binding the contract method 0xda079669.
//
// Solidity: function buyNfts(address[] _tokens, uint256[] _amounts) payable returns()
func (_Buy *BuyTransactor) BuyNfts(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "buyNfts", _tokens, _amounts)
}

// BuyNfts is a paid mutator transaction binding the contract method 0xda079669.
//
// Solidity: function buyNfts(address[] _tokens, uint256[] _amounts) payable returns()
func (_Buy *BuySession) BuyNfts(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Buy.Contract.BuyNfts(&_Buy.TransactOpts, _tokens, _amounts)
}

// BuyNfts is a paid mutator transaction binding the contract method 0xda079669.
//
// Solidity: function buyNfts(address[] _tokens, uint256[] _amounts) payable returns()
func (_Buy *BuyTransactorSession) BuyNfts(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Buy.Contract.BuyNfts(&_Buy.TransactOpts, _tokens, _amounts)
}

// ClearNftPrice is a paid mutator transaction binding the contract method 0x3e241973.
//
// Solidity: function clearNftPrice() returns()
func (_Buy *BuyTransactor) ClearNftPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "clearNftPrice")
}

// ClearNftPrice is a paid mutator transaction binding the contract method 0x3e241973.
//
// Solidity: function clearNftPrice() returns()
func (_Buy *BuySession) ClearNftPrice() (*types.Transaction, error) {
	return _Buy.Contract.ClearNftPrice(&_Buy.TransactOpts)
}

// ClearNftPrice is a paid mutator transaction binding the contract method 0x3e241973.
//
// Solidity: function clearNftPrice() returns()
func (_Buy *BuyTransactorSession) ClearNftPrice() (*types.Transaction, error) {
	return _Buy.Contract.ClearNftPrice(&_Buy.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Buy *BuyTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Buy *BuySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Buy.Contract.GrantRole(&_Buy.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Buy *BuyTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Buy.Contract.GrantRole(&_Buy.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _receiveAddress) returns()
func (_Buy *BuyTransactor) Initialize(opts *bind.TransactOpts, _tokens []common.Address, _receiveAddress common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "initialize", _tokens, _receiveAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _receiveAddress) returns()
func (_Buy *BuySession) Initialize(_tokens []common.Address, _receiveAddress common.Address) (*types.Transaction, error) {
	return _Buy.Contract.Initialize(&_Buy.TransactOpts, _tokens, _receiveAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _receiveAddress) returns()
func (_Buy *BuyTransactorSession) Initialize(_tokens []common.Address, _receiveAddress common.Address) (*types.Transaction, error) {
	return _Buy.Contract.Initialize(&_Buy.TransactOpts, _tokens, _receiveAddress)
}

// OpenBlindBox is a paid mutator transaction binding the contract method 0xd2df81ae.
//
// Solidity: function openBlindBox(uint256 _amount) returns()
func (_Buy *BuyTransactor) OpenBlindBox(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "openBlindBox", _amount)
}

// OpenBlindBox is a paid mutator transaction binding the contract method 0xd2df81ae.
//
// Solidity: function openBlindBox(uint256 _amount) returns()
func (_Buy *BuySession) OpenBlindBox(_amount *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.OpenBlindBox(&_Buy.TransactOpts, _amount)
}

// OpenBlindBox is a paid mutator transaction binding the contract method 0xd2df81ae.
//
// Solidity: function openBlindBox(uint256 _amount) returns()
func (_Buy *BuyTransactorSession) OpenBlindBox(_amount *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.OpenBlindBox(&_Buy.TransactOpts, _amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Buy *BuyTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Buy *BuySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Buy.Contract.RenounceRole(&_Buy.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Buy *BuyTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Buy.Contract.RenounceRole(&_Buy.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Buy *BuyTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Buy *BuySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Buy.Contract.RevokeRole(&_Buy.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Buy *BuyTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Buy.Contract.RevokeRole(&_Buy.TransactOpts, role, account)
}

// SetBuyAndBlindBoxEnable is a paid mutator transaction binding the contract method 0x7851f7b6.
//
// Solidity: function setBuyAndBlindBoxEnable(bool _buyEnable, bool _blindBoxEnable) returns()
func (_Buy *BuyTransactor) SetBuyAndBlindBoxEnable(opts *bind.TransactOpts, _buyEnable bool, _blindBoxEnable bool) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "setBuyAndBlindBoxEnable", _buyEnable, _blindBoxEnable)
}

// SetBuyAndBlindBoxEnable is a paid mutator transaction binding the contract method 0x7851f7b6.
//
// Solidity: function setBuyAndBlindBoxEnable(bool _buyEnable, bool _blindBoxEnable) returns()
func (_Buy *BuySession) SetBuyAndBlindBoxEnable(_buyEnable bool, _blindBoxEnable bool) (*types.Transaction, error) {
	return _Buy.Contract.SetBuyAndBlindBoxEnable(&_Buy.TransactOpts, _buyEnable, _blindBoxEnable)
}

// SetBuyAndBlindBoxEnable is a paid mutator transaction binding the contract method 0x7851f7b6.
//
// Solidity: function setBuyAndBlindBoxEnable(bool _buyEnable, bool _blindBoxEnable) returns()
func (_Buy *BuyTransactorSession) SetBuyAndBlindBoxEnable(_buyEnable bool, _blindBoxEnable bool) (*types.Transaction, error) {
	return _Buy.Contract.SetBuyAndBlindBoxEnable(&_Buy.TransactOpts, _buyEnable, _blindBoxEnable)
}

// SetEnable is a paid mutator transaction binding the contract method 0x7726bed3.
//
// Solidity: function setEnable(bool _enable) returns()
func (_Buy *BuyTransactor) SetEnable(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "setEnable", _enable)
}

// SetEnable is a paid mutator transaction binding the contract method 0x7726bed3.
//
// Solidity: function setEnable(bool _enable) returns()
func (_Buy *BuySession) SetEnable(_enable bool) (*types.Transaction, error) {
	return _Buy.Contract.SetEnable(&_Buy.TransactOpts, _enable)
}

// SetEnable is a paid mutator transaction binding the contract method 0x7726bed3.
//
// Solidity: function setEnable(bool _enable) returns()
func (_Buy *BuyTransactorSession) SetEnable(_enable bool) (*types.Transaction, error) {
	return _Buy.Contract.SetEnable(&_Buy.TransactOpts, _enable)
}

// SetNftPrice is a paid mutator transaction binding the contract method 0x0f3c4fe4.
//
// Solidity: function setNftPrice(address _token, uint256 _ethPrice, uint256 _usdPrice, uint256 _nezPrice) returns()
func (_Buy *BuyTransactor) SetNftPrice(opts *bind.TransactOpts, _token common.Address, _ethPrice *big.Int, _usdPrice *big.Int, _nezPrice *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "setNftPrice", _token, _ethPrice, _usdPrice, _nezPrice)
}

// SetNftPrice is a paid mutator transaction binding the contract method 0x0f3c4fe4.
//
// Solidity: function setNftPrice(address _token, uint256 _ethPrice, uint256 _usdPrice, uint256 _nezPrice) returns()
func (_Buy *BuySession) SetNftPrice(_token common.Address, _ethPrice *big.Int, _usdPrice *big.Int, _nezPrice *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.SetNftPrice(&_Buy.TransactOpts, _token, _ethPrice, _usdPrice, _nezPrice)
}

// SetNftPrice is a paid mutator transaction binding the contract method 0x0f3c4fe4.
//
// Solidity: function setNftPrice(address _token, uint256 _ethPrice, uint256 _usdPrice, uint256 _nezPrice) returns()
func (_Buy *BuyTransactorSession) SetNftPrice(_token common.Address, _ethPrice *big.Int, _usdPrice *big.Int, _nezPrice *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.SetNftPrice(&_Buy.TransactOpts, _token, _ethPrice, _usdPrice, _nezPrice)
}

// SetReceiveAddress is a paid mutator transaction binding the contract method 0x5ec4b7a8.
//
// Solidity: function setReceiveAddress(address _receiveAddress) returns()
func (_Buy *BuyTransactor) SetReceiveAddress(opts *bind.TransactOpts, _receiveAddress common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "setReceiveAddress", _receiveAddress)
}

// SetReceiveAddress is a paid mutator transaction binding the contract method 0x5ec4b7a8.
//
// Solidity: function setReceiveAddress(address _receiveAddress) returns()
func (_Buy *BuySession) SetReceiveAddress(_receiveAddress common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetReceiveAddress(&_Buy.TransactOpts, _receiveAddress)
}

// SetReceiveAddress is a paid mutator transaction binding the contract method 0x5ec4b7a8.
//
// Solidity: function setReceiveAddress(address _receiveAddress) returns()
func (_Buy *BuyTransactorSession) SetReceiveAddress(_receiveAddress common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetReceiveAddress(&_Buy.TransactOpts, _receiveAddress)
}

// SetTokens is a paid mutator transaction binding the contract method 0x625adaf2.
//
// Solidity: function setTokens(address[] _tokens) returns()
func (_Buy *BuyTransactor) SetTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "setTokens", _tokens)
}

// SetTokens is a paid mutator transaction binding the contract method 0x625adaf2.
//
// Solidity: function setTokens(address[] _tokens) returns()
func (_Buy *BuySession) SetTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetTokens(&_Buy.TransactOpts, _tokens)
}

// SetTokens is a paid mutator transaction binding the contract method 0x625adaf2.
//
// Solidity: function setTokens(address[] _tokens) returns()
func (_Buy *BuyTransactorSession) SetTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetTokens(&_Buy.TransactOpts, _tokens)
}

// TransferBlindBox is a paid mutator transaction binding the contract method 0x674b6cd9.
//
// Solidity: function transferBlindBox(address _to, uint256 _amount) returns()
func (_Buy *BuyTransactor) TransferBlindBox(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "transferBlindBox", _to, _amount)
}

// TransferBlindBox is a paid mutator transaction binding the contract method 0x674b6cd9.
//
// Solidity: function transferBlindBox(address _to, uint256 _amount) returns()
func (_Buy *BuySession) TransferBlindBox(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.TransferBlindBox(&_Buy.TransactOpts, _to, _amount)
}

// TransferBlindBox is a paid mutator transaction binding the contract method 0x674b6cd9.
//
// Solidity: function transferBlindBox(address _to, uint256 _amount) returns()
func (_Buy *BuyTransactorSession) TransferBlindBox(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.TransferBlindBox(&_Buy.TransactOpts, _to, _amount)
}

// BuyBuyBlindBoxIterator is returned from FilterBuyBlindBox and is used to iterate over the raw logs and unpacked data for BuyBlindBox events raised by the Buy contract.
type BuyBuyBlindBoxIterator struct {
	Event *BuyBuyBlindBox // Event containing the contract specifics and raw log

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
func (it *BuyBuyBlindBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBuyBlindBox)
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
		it.Event = new(BuyBuyBlindBox)
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
func (it *BuyBuyBlindBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBuyBlindBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBuyBlindBox represents a BuyBlindBox event raised by the Buy contract.
type BuyBuyBlindBox struct {
	Member common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuyBlindBox is a free log retrieval operation binding the contract event 0x996fb5df80ed4cb9750c208a37f2da0935db8267fdf2dbd81e2fba06e57a55d9.
//
// Solidity: event BuyBlindBox(address indexed member, uint256 amount)
func (_Buy *BuyFilterer) FilterBuyBlindBox(opts *bind.FilterOpts, member []common.Address) (*BuyBuyBlindBoxIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "BuyBlindBox", memberRule)
	if err != nil {
		return nil, err
	}
	return &BuyBuyBlindBoxIterator{contract: _Buy.contract, event: "BuyBlindBox", logs: logs, sub: sub}, nil
}

// WatchBuyBlindBox is a free log subscription operation binding the contract event 0x996fb5df80ed4cb9750c208a37f2da0935db8267fdf2dbd81e2fba06e57a55d9.
//
// Solidity: event BuyBlindBox(address indexed member, uint256 amount)
func (_Buy *BuyFilterer) WatchBuyBlindBox(opts *bind.WatchOpts, sink chan<- *BuyBuyBlindBox, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "BuyBlindBox", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBuyBlindBox)
				if err := _Buy.contract.UnpackLog(event, "BuyBlindBox", log); err != nil {
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

// ParseBuyBlindBox is a log parse operation binding the contract event 0x996fb5df80ed4cb9750c208a37f2da0935db8267fdf2dbd81e2fba06e57a55d9.
//
// Solidity: event BuyBlindBox(address indexed member, uint256 amount)
func (_Buy *BuyFilterer) ParseBuyBlindBox(log types.Log) (*BuyBuyBlindBox, error) {
	event := new(BuyBuyBlindBox)
	if err := _Buy.contract.UnpackLog(event, "BuyBlindBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyBuyNftIterator is returned from FilterBuyNft and is used to iterate over the raw logs and unpacked data for BuyNft events raised by the Buy contract.
type BuyBuyNftIterator struct {
	Event *BuyBuyNft // Event containing the contract specifics and raw log

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
func (it *BuyBuyNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyBuyNft)
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
		it.Event = new(BuyBuyNft)
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
func (it *BuyBuyNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyBuyNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyBuyNft represents a BuyNft event raised by the Buy contract.
type BuyBuyNft struct {
	Member common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuyNft is a free log retrieval operation binding the contract event 0x812b46001b22802a4c4e896a8160a696158e1059db4b9a20cf580be71a9d58fb.
//
// Solidity: event BuyNft(address indexed member, address token, uint256 amount)
func (_Buy *BuyFilterer) FilterBuyNft(opts *bind.FilterOpts, member []common.Address) (*BuyBuyNftIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "BuyNft", memberRule)
	if err != nil {
		return nil, err
	}
	return &BuyBuyNftIterator{contract: _Buy.contract, event: "BuyNft", logs: logs, sub: sub}, nil
}

// WatchBuyNft is a free log subscription operation binding the contract event 0x812b46001b22802a4c4e896a8160a696158e1059db4b9a20cf580be71a9d58fb.
//
// Solidity: event BuyNft(address indexed member, address token, uint256 amount)
func (_Buy *BuyFilterer) WatchBuyNft(opts *bind.WatchOpts, sink chan<- *BuyBuyNft, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "BuyNft", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyBuyNft)
				if err := _Buy.contract.UnpackLog(event, "BuyNft", log); err != nil {
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

// ParseBuyNft is a log parse operation binding the contract event 0x812b46001b22802a4c4e896a8160a696158e1059db4b9a20cf580be71a9d58fb.
//
// Solidity: event BuyNft(address indexed member, address token, uint256 amount)
func (_Buy *BuyFilterer) ParseBuyNft(log types.Log) (*BuyBuyNft, error) {
	event := new(BuyBuyNft)
	if err := _Buy.contract.UnpackLog(event, "BuyNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Buy contract.
type BuyInitializedIterator struct {
	Event *BuyInitialized // Event containing the contract specifics and raw log

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
func (it *BuyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyInitialized)
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
		it.Event = new(BuyInitialized)
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
func (it *BuyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyInitialized represents a Initialized event raised by the Buy contract.
type BuyInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Buy *BuyFilterer) FilterInitialized(opts *bind.FilterOpts) (*BuyInitializedIterator, error) {

	logs, sub, err := _Buy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BuyInitializedIterator{contract: _Buy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Buy *BuyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BuyInitialized) (event.Subscription, error) {

	logs, sub, err := _Buy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyInitialized)
				if err := _Buy.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Buy *BuyFilterer) ParseInitialized(log types.Log) (*BuyInitialized, error) {
	event := new(BuyInitialized)
	if err := _Buy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyOpenBlindBoxIterator is returned from FilterOpenBlindBox and is used to iterate over the raw logs and unpacked data for OpenBlindBox events raised by the Buy contract.
type BuyOpenBlindBoxIterator struct {
	Event *BuyOpenBlindBox // Event containing the contract specifics and raw log

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
func (it *BuyOpenBlindBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyOpenBlindBox)
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
		it.Event = new(BuyOpenBlindBox)
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
func (it *BuyOpenBlindBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyOpenBlindBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyOpenBlindBox represents a OpenBlindBox event raised by the Buy contract.
type BuyOpenBlindBox struct {
	Member common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOpenBlindBox is a free log retrieval operation binding the contract event 0xe71889d549f147c5826377e104cbc41f8344c596fa65b8af5de4320911a161e2.
//
// Solidity: event OpenBlindBox(address indexed member, address token, uint256 amount)
func (_Buy *BuyFilterer) FilterOpenBlindBox(opts *bind.FilterOpts, member []common.Address) (*BuyOpenBlindBoxIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "OpenBlindBox", memberRule)
	if err != nil {
		return nil, err
	}
	return &BuyOpenBlindBoxIterator{contract: _Buy.contract, event: "OpenBlindBox", logs: logs, sub: sub}, nil
}

// WatchOpenBlindBox is a free log subscription operation binding the contract event 0xe71889d549f147c5826377e104cbc41f8344c596fa65b8af5de4320911a161e2.
//
// Solidity: event OpenBlindBox(address indexed member, address token, uint256 amount)
func (_Buy *BuyFilterer) WatchOpenBlindBox(opts *bind.WatchOpts, sink chan<- *BuyOpenBlindBox, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "OpenBlindBox", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyOpenBlindBox)
				if err := _Buy.contract.UnpackLog(event, "OpenBlindBox", log); err != nil {
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

// ParseOpenBlindBox is a log parse operation binding the contract event 0xe71889d549f147c5826377e104cbc41f8344c596fa65b8af5de4320911a161e2.
//
// Solidity: event OpenBlindBox(address indexed member, address token, uint256 amount)
func (_Buy *BuyFilterer) ParseOpenBlindBox(log types.Log) (*BuyOpenBlindBox, error) {
	event := new(BuyOpenBlindBox)
	if err := _Buy.contract.UnpackLog(event, "OpenBlindBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Buy contract.
type BuyRoleAdminChangedIterator struct {
	Event *BuyRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BuyRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyRoleAdminChanged)
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
		it.Event = new(BuyRoleAdminChanged)
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
func (it *BuyRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyRoleAdminChanged represents a RoleAdminChanged event raised by the Buy contract.
type BuyRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Buy *BuyFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BuyRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Buy.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BuyRoleAdminChangedIterator{contract: _Buy.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Buy *BuyFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BuyRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Buy.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyRoleAdminChanged)
				if err := _Buy.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Buy *BuyFilterer) ParseRoleAdminChanged(log types.Log) (*BuyRoleAdminChanged, error) {
	event := new(BuyRoleAdminChanged)
	if err := _Buy.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Buy contract.
type BuyRoleGrantedIterator struct {
	Event *BuyRoleGranted // Event containing the contract specifics and raw log

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
func (it *BuyRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyRoleGranted)
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
		it.Event = new(BuyRoleGranted)
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
func (it *BuyRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyRoleGranted represents a RoleGranted event raised by the Buy contract.
type BuyRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Buy *BuyFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BuyRoleGrantedIterator, error) {

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

	logs, sub, err := _Buy.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BuyRoleGrantedIterator{contract: _Buy.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Buy *BuyFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BuyRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Buy.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyRoleGranted)
				if err := _Buy.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Buy *BuyFilterer) ParseRoleGranted(log types.Log) (*BuyRoleGranted, error) {
	event := new(BuyRoleGranted)
	if err := _Buy.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Buy contract.
type BuyRoleRevokedIterator struct {
	Event *BuyRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BuyRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyRoleRevoked)
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
		it.Event = new(BuyRoleRevoked)
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
func (it *BuyRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyRoleRevoked represents a RoleRevoked event raised by the Buy contract.
type BuyRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Buy *BuyFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BuyRoleRevokedIterator, error) {

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

	logs, sub, err := _Buy.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BuyRoleRevokedIterator{contract: _Buy.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Buy *BuyFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BuyRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Buy.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyRoleRevoked)
				if err := _Buy.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Buy *BuyFilterer) ParseRoleRevoked(log types.Log) (*BuyRoleRevoked, error) {
	event := new(BuyRoleRevoked)
	if err := _Buy.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyTransferBlindBoxIterator is returned from FilterTransferBlindBox and is used to iterate over the raw logs and unpacked data for TransferBlindBox events raised by the Buy contract.
type BuyTransferBlindBoxIterator struct {
	Event *BuyTransferBlindBox // Event containing the contract specifics and raw log

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
func (it *BuyTransferBlindBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyTransferBlindBox)
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
		it.Event = new(BuyTransferBlindBox)
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
func (it *BuyTransferBlindBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyTransferBlindBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyTransferBlindBox represents a TransferBlindBox event raised by the Buy contract.
type BuyTransferBlindBox struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferBlindBox is a free log retrieval operation binding the contract event 0xb2c52e347b57f86e9af8a6df6d503f9bf32b25cc357cb421054c43dc041fe2ed.
//
// Solidity: event TransferBlindBox(address indexed from, address indexed to, uint256 amount)
func (_Buy *BuyFilterer) FilterTransferBlindBox(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BuyTransferBlindBoxIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "TransferBlindBox", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BuyTransferBlindBoxIterator{contract: _Buy.contract, event: "TransferBlindBox", logs: logs, sub: sub}, nil
}

// WatchTransferBlindBox is a free log subscription operation binding the contract event 0xb2c52e347b57f86e9af8a6df6d503f9bf32b25cc357cb421054c43dc041fe2ed.
//
// Solidity: event TransferBlindBox(address indexed from, address indexed to, uint256 amount)
func (_Buy *BuyFilterer) WatchTransferBlindBox(opts *bind.WatchOpts, sink chan<- *BuyTransferBlindBox, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "TransferBlindBox", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyTransferBlindBox)
				if err := _Buy.contract.UnpackLog(event, "TransferBlindBox", log); err != nil {
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

// ParseTransferBlindBox is a log parse operation binding the contract event 0xb2c52e347b57f86e9af8a6df6d503f9bf32b25cc357cb421054c43dc041fe2ed.
//
// Solidity: event TransferBlindBox(address indexed from, address indexed to, uint256 amount)
func (_Buy *BuyFilterer) ParseTransferBlindBox(log types.Log) (*BuyTransferBlindBox, error) {
	event := new(BuyTransferBlindBox)
	if err := _Buy.contract.UnpackLog(event, "TransferBlindBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
