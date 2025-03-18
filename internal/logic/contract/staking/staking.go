// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allNfts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allStakedUsers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nez\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_allNfts\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_reStakeNfts\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_stakeNfts\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_unstakeNfts\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDailyNezReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nez\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryAllStakedUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"queryUserReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reStakeNfts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"}],\"name\":\"setAllNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setEnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDailyNezReward\",\"type\":\"uint256\"}],\"name\":\"setMaxDailyNezReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nez\",\"type\":\"address\"}],\"name\":\"setNez\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"}],\"name\":\"setReStakeNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"}],\"name\":\"setStakeNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"}],\"name\":\"setUnstakeNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeNfts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unstakeNfts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"userIsStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userStakeNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"unstaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Staking.Contract.DEFAULTADMINROLE(&_Staking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Staking.Contract.DEFAULTADMINROLE(&_Staking.CallOpts)
}

// AllNfts is a free data retrieval call binding the contract method 0xc529de65.
//
// Solidity: function allNfts(uint256 ) view returns(address)
func (_Staking *StakingCaller) AllNfts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "allNfts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllNfts is a free data retrieval call binding the contract method 0xc529de65.
//
// Solidity: function allNfts(uint256 ) view returns(address)
func (_Staking *StakingSession) AllNfts(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.AllNfts(&_Staking.CallOpts, arg0)
}

// AllNfts is a free data retrieval call binding the contract method 0xc529de65.
//
// Solidity: function allNfts(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) AllNfts(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.AllNfts(&_Staking.CallOpts, arg0)
}

// AllStakedUsers is a free data retrieval call binding the contract method 0x1ac918f4.
//
// Solidity: function allStakedUsers(uint256 ) view returns(address)
func (_Staking *StakingCaller) AllStakedUsers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "allStakedUsers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllStakedUsers is a free data retrieval call binding the contract method 0x1ac918f4.
//
// Solidity: function allStakedUsers(uint256 ) view returns(address)
func (_Staking *StakingSession) AllStakedUsers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.AllStakedUsers(&_Staking.CallOpts, arg0)
}

// AllStakedUsers is a free data retrieval call binding the contract method 0x1ac918f4.
//
// Solidity: function allStakedUsers(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) AllStakedUsers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.AllStakedUsers(&_Staking.CallOpts, arg0)
}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Staking *StakingCaller) Enable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "enable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Staking *StakingSession) Enable() (bool, error) {
	return _Staking.Contract.Enable(&_Staking.CallOpts)
}

// Enable is a free data retrieval call binding the contract method 0xa3907d71.
//
// Solidity: function enable() view returns(bool)
func (_Staking *StakingCallerSession) Enable() (bool, error) {
	return _Staking.Contract.Enable(&_Staking.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Staking.Contract.GetRoleAdmin(&_Staking.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Staking.Contract.GetRoleAdmin(&_Staking.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Staking.Contract.HasRole(&_Staking.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Staking.Contract.HasRole(&_Staking.CallOpts, role, account)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address ) view returns(bool)
func (_Staking *StakingCaller) IsStaked(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isStaked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address ) view returns(bool)
func (_Staking *StakingSession) IsStaked(arg0 common.Address) (bool, error) {
	return _Staking.Contract.IsStaked(&_Staking.CallOpts, arg0)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address ) view returns(bool)
func (_Staking *StakingCallerSession) IsStaked(arg0 common.Address) (bool, error) {
	return _Staking.Contract.IsStaked(&_Staking.CallOpts, arg0)
}

// MaxDailyNezReward is a free data retrieval call binding the contract method 0xf1da30f3.
//
// Solidity: function maxDailyNezReward() view returns(uint256)
func (_Staking *StakingCaller) MaxDailyNezReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "maxDailyNezReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDailyNezReward is a free data retrieval call binding the contract method 0xf1da30f3.
//
// Solidity: function maxDailyNezReward() view returns(uint256)
func (_Staking *StakingSession) MaxDailyNezReward() (*big.Int, error) {
	return _Staking.Contract.MaxDailyNezReward(&_Staking.CallOpts)
}

// MaxDailyNezReward is a free data retrieval call binding the contract method 0xf1da30f3.
//
// Solidity: function maxDailyNezReward() view returns(uint256)
func (_Staking *StakingCallerSession) MaxDailyNezReward() (*big.Int, error) {
	return _Staking.Contract.MaxDailyNezReward(&_Staking.CallOpts)
}

// Nez is a free data retrieval call binding the contract method 0xd5bf4d4b.
//
// Solidity: function nez() view returns(address)
func (_Staking *StakingCaller) Nez(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "nez")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nez is a free data retrieval call binding the contract method 0xd5bf4d4b.
//
// Solidity: function nez() view returns(address)
func (_Staking *StakingSession) Nez() (common.Address, error) {
	return _Staking.Contract.Nez(&_Staking.CallOpts)
}

// Nez is a free data retrieval call binding the contract method 0xd5bf4d4b.
//
// Solidity: function nez() view returns(address)
func (_Staking *StakingCallerSession) Nez() (common.Address, error) {
	return _Staking.Contract.Nez(&_Staking.CallOpts)
}

// QueryAllStakedUsers is a free data retrieval call binding the contract method 0x6f21290e.
//
// Solidity: function queryAllStakedUsers() view returns(address[])
func (_Staking *StakingCaller) QueryAllStakedUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "queryAllStakedUsers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// QueryAllStakedUsers is a free data retrieval call binding the contract method 0x6f21290e.
//
// Solidity: function queryAllStakedUsers() view returns(address[])
func (_Staking *StakingSession) QueryAllStakedUsers() ([]common.Address, error) {
	return _Staking.Contract.QueryAllStakedUsers(&_Staking.CallOpts)
}

// QueryAllStakedUsers is a free data retrieval call binding the contract method 0x6f21290e.
//
// Solidity: function queryAllStakedUsers() view returns(address[])
func (_Staking *StakingCallerSession) QueryAllStakedUsers() ([]common.Address, error) {
	return _Staking.Contract.QueryAllStakedUsers(&_Staking.CallOpts)
}

// QueryUserReward is a free data retrieval call binding the contract method 0x182ee8b3.
//
// Solidity: function queryUserReward(address _user) view returns(uint256)
func (_Staking *StakingCaller) QueryUserReward(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "queryUserReward", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryUserReward is a free data retrieval call binding the contract method 0x182ee8b3.
//
// Solidity: function queryUserReward(address _user) view returns(uint256)
func (_Staking *StakingSession) QueryUserReward(_user common.Address) (*big.Int, error) {
	return _Staking.Contract.QueryUserReward(&_Staking.CallOpts, _user)
}

// QueryUserReward is a free data retrieval call binding the contract method 0x182ee8b3.
//
// Solidity: function queryUserReward(address _user) view returns(uint256)
func (_Staking *StakingCallerSession) QueryUserReward(_user common.Address) (*big.Int, error) {
	return _Staking.Contract.QueryUserReward(&_Staking.CallOpts, _user)
}

// ReStakeNfts is a free data retrieval call binding the contract method 0x644cbb88.
//
// Solidity: function reStakeNfts(address ) view returns(bool)
func (_Staking *StakingCaller) ReStakeNfts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "reStakeNfts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReStakeNfts is a free data retrieval call binding the contract method 0x644cbb88.
//
// Solidity: function reStakeNfts(address ) view returns(bool)
func (_Staking *StakingSession) ReStakeNfts(arg0 common.Address) (bool, error) {
	return _Staking.Contract.ReStakeNfts(&_Staking.CallOpts, arg0)
}

// ReStakeNfts is a free data retrieval call binding the contract method 0x644cbb88.
//
// Solidity: function reStakeNfts(address ) view returns(bool)
func (_Staking *StakingCallerSession) ReStakeNfts(arg0 common.Address) (bool, error) {
	return _Staking.Contract.ReStakeNfts(&_Staking.CallOpts, arg0)
}

// StakeNfts is a free data retrieval call binding the contract method 0x906e31b4.
//
// Solidity: function stakeNfts(address ) view returns(bool)
func (_Staking *StakingCaller) StakeNfts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakeNfts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StakeNfts is a free data retrieval call binding the contract method 0x906e31b4.
//
// Solidity: function stakeNfts(address ) view returns(bool)
func (_Staking *StakingSession) StakeNfts(arg0 common.Address) (bool, error) {
	return _Staking.Contract.StakeNfts(&_Staking.CallOpts, arg0)
}

// StakeNfts is a free data retrieval call binding the contract method 0x906e31b4.
//
// Solidity: function stakeNfts(address ) view returns(bool)
func (_Staking *StakingCallerSession) StakeNfts(arg0 common.Address) (bool, error) {
	return _Staking.Contract.StakeNfts(&_Staking.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Staking.Contract.SupportsInterface(&_Staking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Staking.Contract.SupportsInterface(&_Staking.CallOpts, interfaceId)
}

// UnstakeNfts is a free data retrieval call binding the contract method 0x033c59f7.
//
// Solidity: function unstakeNfts(address ) view returns(bool)
func (_Staking *StakingCaller) UnstakeNfts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "unstakeNfts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UnstakeNfts is a free data retrieval call binding the contract method 0x033c59f7.
//
// Solidity: function unstakeNfts(address ) view returns(bool)
func (_Staking *StakingSession) UnstakeNfts(arg0 common.Address) (bool, error) {
	return _Staking.Contract.UnstakeNfts(&_Staking.CallOpts, arg0)
}

// UnstakeNfts is a free data retrieval call binding the contract method 0x033c59f7.
//
// Solidity: function unstakeNfts(address ) view returns(bool)
func (_Staking *StakingCallerSession) UnstakeNfts(arg0 common.Address) (bool, error) {
	return _Staking.Contract.UnstakeNfts(&_Staking.CallOpts, arg0)
}

// UserIsStaked is a free data retrieval call binding the contract method 0xf2c16a28.
//
// Solidity: function userIsStaked(address _user) view returns(bool)
func (_Staking *StakingCaller) UserIsStaked(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userIsStaked", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserIsStaked is a free data retrieval call binding the contract method 0xf2c16a28.
//
// Solidity: function userIsStaked(address _user) view returns(bool)
func (_Staking *StakingSession) UserIsStaked(_user common.Address) (bool, error) {
	return _Staking.Contract.UserIsStaked(&_Staking.CallOpts, _user)
}

// UserIsStaked is a free data retrieval call binding the contract method 0xf2c16a28.
//
// Solidity: function userIsStaked(address _user) view returns(bool)
func (_Staking *StakingCallerSession) UserIsStaked(_user common.Address) (bool, error) {
	return _Staking.Contract.UserIsStaked(&_Staking.CallOpts, _user)
}

// UserStakeNfts is a free data retrieval call binding the contract method 0x36ff56b7.
//
// Solidity: function userStakeNfts(address , address ) view returns(uint256)
func (_Staking *StakingCaller) UserStakeNfts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userStakeNfts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserStakeNfts is a free data retrieval call binding the contract method 0x36ff56b7.
//
// Solidity: function userStakeNfts(address , address ) view returns(uint256)
func (_Staking *StakingSession) UserStakeNfts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Staking.Contract.UserStakeNfts(&_Staking.CallOpts, arg0, arg1)
}

// UserStakeNfts is a free data retrieval call binding the contract method 0x36ff56b7.
//
// Solidity: function userStakeNfts(address , address ) view returns(uint256)
func (_Staking *StakingCallerSession) UserStakeNfts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Staking.Contract.UserStakeNfts(&_Staking.CallOpts, arg0, arg1)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 stakeTime, uint256 claimTime, uint256 endTime, bool unstaked)
func (_Staking *StakingCaller) UserStakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	StakeTime *big.Int
	ClaimTime *big.Int
	EndTime   *big.Int
	Unstaked  bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userStakes", arg0)

	outstruct := new(struct {
		StakeTime *big.Int
		ClaimTime *big.Int
		EndTime   *big.Int
		Unstaked  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ClaimTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Unstaked = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 stakeTime, uint256 claimTime, uint256 endTime, bool unstaked)
func (_Staking *StakingSession) UserStakes(arg0 common.Address) (struct {
	StakeTime *big.Int
	ClaimTime *big.Int
	EndTime   *big.Int
	Unstaked  bool
}, error) {
	return _Staking.Contract.UserStakes(&_Staking.CallOpts, arg0)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 stakeTime, uint256 claimTime, uint256 endTime, bool unstaked)
func (_Staking *StakingCallerSession) UserStakes(arg0 common.Address) (struct {
	StakeTime *big.Int
	ClaimTime *big.Int
	EndTime   *big.Int
	Unstaked  bool
}, error) {
	return _Staking.Contract.UserStakes(&_Staking.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Staking *StakingTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Staking *StakingSession) Claim() (*types.Transaction, error) {
	return _Staking.Contract.Claim(&_Staking.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Staking *StakingTransactorSession) Claim() (*types.Transaction, error) {
	return _Staking.Contract.Claim(&_Staking.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.GrantRole(&_Staking.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.GrantRole(&_Staking.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xa204526b.
//
// Solidity: function initialize(address _nez, address[] _allNfts, address[] _reStakeNfts, address[] _stakeNfts, address[] _unstakeNfts) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, _nez common.Address, _allNfts []common.Address, _reStakeNfts []common.Address, _stakeNfts []common.Address, _unstakeNfts []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", _nez, _allNfts, _reStakeNfts, _stakeNfts, _unstakeNfts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa204526b.
//
// Solidity: function initialize(address _nez, address[] _allNfts, address[] _reStakeNfts, address[] _stakeNfts, address[] _unstakeNfts) returns()
func (_Staking *StakingSession) Initialize(_nez common.Address, _allNfts []common.Address, _reStakeNfts []common.Address, _stakeNfts []common.Address, _unstakeNfts []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _nez, _allNfts, _reStakeNfts, _stakeNfts, _unstakeNfts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa204526b.
//
// Solidity: function initialize(address _nez, address[] _allNfts, address[] _reStakeNfts, address[] _stakeNfts, address[] _unstakeNfts) returns()
func (_Staking *StakingTransactorSession) Initialize(_nez common.Address, _allNfts []common.Address, _reStakeNfts []common.Address, _stakeNfts []common.Address, _unstakeNfts []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _nez, _allNfts, _reStakeNfts, _stakeNfts, _unstakeNfts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Staking *StakingTransactor) OnERC721Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "onERC721Received", _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Staking *StakingSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Staking.Contract.OnERC721Received(&_Staking.TransactOpts, _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Staking *StakingTransactorSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Staking.Contract.OnERC721Received(&_Staking.TransactOpts, _operator, _from, _tokenId, _data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Staking *StakingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Staking *StakingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RenounceRole(&_Staking.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Staking *StakingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RenounceRole(&_Staking.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RevokeRole(&_Staking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RevokeRole(&_Staking.TransactOpts, role, account)
}

// SetAllNfts is a paid mutator transaction binding the contract method 0xb007c619.
//
// Solidity: function setAllNfts(address[] _nfts) returns()
func (_Staking *StakingTransactor) SetAllNfts(opts *bind.TransactOpts, _nfts []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setAllNfts", _nfts)
}

// SetAllNfts is a paid mutator transaction binding the contract method 0xb007c619.
//
// Solidity: function setAllNfts(address[] _nfts) returns()
func (_Staking *StakingSession) SetAllNfts(_nfts []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetAllNfts(&_Staking.TransactOpts, _nfts)
}

// SetAllNfts is a paid mutator transaction binding the contract method 0xb007c619.
//
// Solidity: function setAllNfts(address[] _nfts) returns()
func (_Staking *StakingTransactorSession) SetAllNfts(_nfts []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetAllNfts(&_Staking.TransactOpts, _nfts)
}

// SetEnable is a paid mutator transaction binding the contract method 0x7726bed3.
//
// Solidity: function setEnable(bool _enable) returns()
func (_Staking *StakingTransactor) SetEnable(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setEnable", _enable)
}

// SetEnable is a paid mutator transaction binding the contract method 0x7726bed3.
//
// Solidity: function setEnable(bool _enable) returns()
func (_Staking *StakingSession) SetEnable(_enable bool) (*types.Transaction, error) {
	return _Staking.Contract.SetEnable(&_Staking.TransactOpts, _enable)
}

// SetEnable is a paid mutator transaction binding the contract method 0x7726bed3.
//
// Solidity: function setEnable(bool _enable) returns()
func (_Staking *StakingTransactorSession) SetEnable(_enable bool) (*types.Transaction, error) {
	return _Staking.Contract.SetEnable(&_Staking.TransactOpts, _enable)
}

// SetMaxDailyNezReward is a paid mutator transaction binding the contract method 0xa9d02fd0.
//
// Solidity: function setMaxDailyNezReward(uint256 _maxDailyNezReward) returns()
func (_Staking *StakingTransactor) SetMaxDailyNezReward(opts *bind.TransactOpts, _maxDailyNezReward *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setMaxDailyNezReward", _maxDailyNezReward)
}

// SetMaxDailyNezReward is a paid mutator transaction binding the contract method 0xa9d02fd0.
//
// Solidity: function setMaxDailyNezReward(uint256 _maxDailyNezReward) returns()
func (_Staking *StakingSession) SetMaxDailyNezReward(_maxDailyNezReward *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxDailyNezReward(&_Staking.TransactOpts, _maxDailyNezReward)
}

// SetMaxDailyNezReward is a paid mutator transaction binding the contract method 0xa9d02fd0.
//
// Solidity: function setMaxDailyNezReward(uint256 _maxDailyNezReward) returns()
func (_Staking *StakingTransactorSession) SetMaxDailyNezReward(_maxDailyNezReward *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxDailyNezReward(&_Staking.TransactOpts, _maxDailyNezReward)
}

// SetNez is a paid mutator transaction binding the contract method 0x54903328.
//
// Solidity: function setNez(address _nez) returns()
func (_Staking *StakingTransactor) SetNez(opts *bind.TransactOpts, _nez common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setNez", _nez)
}

// SetNez is a paid mutator transaction binding the contract method 0x54903328.
//
// Solidity: function setNez(address _nez) returns()
func (_Staking *StakingSession) SetNez(_nez common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetNez(&_Staking.TransactOpts, _nez)
}

// SetNez is a paid mutator transaction binding the contract method 0x54903328.
//
// Solidity: function setNez(address _nez) returns()
func (_Staking *StakingTransactorSession) SetNez(_nez common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetNez(&_Staking.TransactOpts, _nez)
}

// SetReStakeNfts is a paid mutator transaction binding the contract method 0x20102889.
//
// Solidity: function setReStakeNfts(address[] _nfts) returns()
func (_Staking *StakingTransactor) SetReStakeNfts(opts *bind.TransactOpts, _nfts []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setReStakeNfts", _nfts)
}

// SetReStakeNfts is a paid mutator transaction binding the contract method 0x20102889.
//
// Solidity: function setReStakeNfts(address[] _nfts) returns()
func (_Staking *StakingSession) SetReStakeNfts(_nfts []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetReStakeNfts(&_Staking.TransactOpts, _nfts)
}

// SetReStakeNfts is a paid mutator transaction binding the contract method 0x20102889.
//
// Solidity: function setReStakeNfts(address[] _nfts) returns()
func (_Staking *StakingTransactorSession) SetReStakeNfts(_nfts []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetReStakeNfts(&_Staking.TransactOpts, _nfts)
}

// SetStakeNfts is a paid mutator transaction binding the contract method 0xf344be64.
//
// Solidity: function setStakeNfts(address[] _nfts) returns()
func (_Staking *StakingTransactor) SetStakeNfts(opts *bind.TransactOpts, _nfts []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setStakeNfts", _nfts)
}

// SetStakeNfts is a paid mutator transaction binding the contract method 0xf344be64.
//
// Solidity: function setStakeNfts(address[] _nfts) returns()
func (_Staking *StakingSession) SetStakeNfts(_nfts []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetStakeNfts(&_Staking.TransactOpts, _nfts)
}

// SetStakeNfts is a paid mutator transaction binding the contract method 0xf344be64.
//
// Solidity: function setStakeNfts(address[] _nfts) returns()
func (_Staking *StakingTransactorSession) SetStakeNfts(_nfts []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetStakeNfts(&_Staking.TransactOpts, _nfts)
}

// SetUnstakeNfts is a paid mutator transaction binding the contract method 0x61dfbf06.
//
// Solidity: function setUnstakeNfts(address[] _nfts) returns()
func (_Staking *StakingTransactor) SetUnstakeNfts(opts *bind.TransactOpts, _nfts []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setUnstakeNfts", _nfts)
}

// SetUnstakeNfts is a paid mutator transaction binding the contract method 0x61dfbf06.
//
// Solidity: function setUnstakeNfts(address[] _nfts) returns()
func (_Staking *StakingSession) SetUnstakeNfts(_nfts []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetUnstakeNfts(&_Staking.TransactOpts, _nfts)
}

// SetUnstakeNfts is a paid mutator transaction binding the contract method 0x61dfbf06.
//
// Solidity: function setUnstakeNfts(address[] _nfts) returns()
func (_Staking *StakingTransactorSession) SetUnstakeNfts(_nfts []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetUnstakeNfts(&_Staking.TransactOpts, _nfts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Staking *StakingSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Staking *StakingTransactorSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_Staking *StakingTransactor) Test(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "test")
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_Staking *StakingSession) Test() (*types.Transaction, error) {
	return _Staking.Contract.Test(&_Staking.TransactOpts)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_Staking *StakingTransactorSession) Test() (*types.Transaction, error) {
	return _Staking.Contract.Test(&_Staking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingTransactor) Unstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unstake")
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingSession) Unstake() (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingTransactorSession) Unstake() (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts)
}

// StakingClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Staking contract.
type StakingClaimedIterator struct {
	Event *StakingClaimed // Event containing the contract specifics and raw log

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
func (it *StakingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingClaimed)
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
		it.Event = new(StakingClaimed)
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
func (it *StakingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingClaimed represents a Claimed event raised by the Staking contract.
type StakingClaimed struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed user, uint256 reward)
func (_Staking *StakingFilterer) FilterClaimed(opts *bind.FilterOpts, user []common.Address) (*StakingClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingClaimedIterator{contract: _Staking.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed user, uint256 reward)
func (_Staking *StakingFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *StakingClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingClaimed)
				if err := _Staking.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed user, uint256 reward)
func (_Staking *StakingFilterer) ParseClaimed(log types.Log) (*StakingClaimed, error) {
	event := new(StakingClaimed)
	if err := _Staking.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Staking contract.
type StakingInitializedIterator struct {
	Event *StakingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInitialized)
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
		it.Event = new(StakingInitialized)
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
func (it *StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInitialized represents a Initialized event raised by the Staking contract.
type StakingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Staking *StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingInitializedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingInitializedIterator{contract: _Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Staking *StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInitialized)
				if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Staking *StakingFilterer) ParseInitialized(log types.Log) (*StakingInitialized, error) {
	event := new(StakingInitialized)
	if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Staking contract.
type StakingRoleAdminChangedIterator struct {
	Event *StakingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleAdminChanged)
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
		it.Event = new(StakingRoleAdminChanged)
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
func (it *StakingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleAdminChanged represents a RoleAdminChanged event raised by the Staking contract.
type StakingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Staking *StakingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleAdminChangedIterator{contract: _Staking.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Staking *StakingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleAdminChanged)
				if err := _Staking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleAdminChanged(log types.Log) (*StakingRoleAdminChanged, error) {
	event := new(StakingRoleAdminChanged)
	if err := _Staking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Staking contract.
type StakingRoleGrantedIterator struct {
	Event *StakingRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleGranted)
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
		it.Event = new(StakingRoleGranted)
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
func (it *StakingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleGranted represents a RoleGranted event raised by the Staking contract.
type StakingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingRoleGrantedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleGrantedIterator{contract: _Staking.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleGranted)
				if err := _Staking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleGranted(log types.Log) (*StakingRoleGranted, error) {
	event := new(StakingRoleGranted)
	if err := _Staking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Staking contract.
type StakingRoleRevokedIterator struct {
	Event *StakingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleRevoked)
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
		it.Event = new(StakingRoleRevoked)
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
func (it *StakingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleRevoked represents a RoleRevoked event raised by the Staking contract.
type StakingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingRoleRevokedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleRevokedIterator{contract: _Staking.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleRevoked)
				if err := _Staking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleRevoked(log types.Log) (*StakingRoleRevoked, error) {
	event := new(StakingRoleRevoked)
	if err := _Staking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Staking contract.
type StakingStakedIterator struct {
	Event *StakingStaked // Event containing the contract specifics and raw log

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
func (it *StakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStaked)
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
		it.Event = new(StakingStaked)
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
func (it *StakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStaked represents a Staked event raised by the Staking contract.
type StakingStaked struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x77338642d9284a44296d29a273e04b8ab6b15c7d2439094cd460b7e4f0b33074.
//
// Solidity: event Staked(address indexed user)
func (_Staking *StakingFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*StakingStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakedIterator{contract: _Staking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x77338642d9284a44296d29a273e04b8ab6b15c7d2439094cd460b7e4f0b33074.
//
// Solidity: event Staked(address indexed user)
func (_Staking *StakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStaked)
				if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x77338642d9284a44296d29a273e04b8ab6b15c7d2439094cd460b7e4f0b33074.
//
// Solidity: event Staked(address indexed user)
func (_Staking *StakingFilterer) ParseStaked(log types.Log) (*StakingStaked, error) {
	event := new(StakingStaked)
	if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Staking contract.
type StakingUnstakedIterator struct {
	Event *StakingUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnstaked)
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
		it.Event = new(StakingUnstaked)
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
func (it *StakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnstaked represents a Unstaked event raised by the Staking contract.
type StakingUnstaked struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x908e667f6c2b13b8062954eb100253ea804c21222b190449e40d967a3ac0ff13.
//
// Solidity: event Unstaked(address indexed user)
func (_Staking *StakingFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*StakingUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnstakedIterator{contract: _Staking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x908e667f6c2b13b8062954eb100253ea804c21222b190449e40d967a3ac0ff13.
//
// Solidity: event Unstaked(address indexed user)
func (_Staking *StakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakingUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnstaked)
				if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x908e667f6c2b13b8062954eb100253ea804c21222b190449e40d967a3ac0ff13.
//
// Solidity: event Unstaked(address indexed user)
func (_Staking *StakingFilterer) ParseUnstaked(log types.Log) (*StakingUnstaked, error) {
	event := new(StakingUnstaked)
	if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
