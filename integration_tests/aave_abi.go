// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package integration_tests

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
)

// AaveV2MetaData contains all meta data concerning the AaveV2 contract.
var AaveV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"mockAccrueFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"mockClaimAndUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"mockEnterPosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[9]\",\"name\":\"route\",\"type\":\"address[9]\"},{\"indexed\":false,\"internalType\":\"uint256[3][4]\",\"name\":\"swapParams\",\"type\":\"uint256[3][4]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minAssetsOut\",\"type\":\"uint256\"}],\"name\":\"mockRebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minAssetsOut\",\"type\":\"uint256\"}],\"name\":\"mockReinvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"mockSetDepositLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newFeesDistributor\",\"type\":\"bytes32\"}],\"name\":\"mockSetFeesDistributor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"mockSetLiquidityLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"shutdown\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"exitPosition\",\"type\":\"bool\"}],\"name\":\"mockSetShutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"position\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"trust\",\"type\":\"bool\"}],\"name\":\"mockSetTrust\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mockSweep\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"mockTransferFees\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accrueFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAndUnstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enterPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[9]\",\"name\":\"route\",\"type\":\"address[9]\"},{\"internalType\":\"uint256[3][4]\",\"name\":\"swapParams\",\"type\":\"uint256[3][4]\"},{\"internalType\":\"uint256\",\"name\":\"minAssetsOut\",\"type\":\"uint256\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAssetsOut\",\"type\":\"uint256\"}],\"name\":\"reinvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setDepositLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newFeesDistributor\",\"type\":\"bytes32\"}],\"name\":\"setFeesDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setLiquidityLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"shutdown\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exitPosition\",\"type\":\"bool\"}],\"name\":\"setShutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"position\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"trust\",\"type\":\"bool\"}],\"name\":\"setTrust\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveV2MetaData.ABI instead.
var AaveV2ABI = AaveV2MetaData.ABI

// AaveV2 is an auto generated Go binding around an Ethereum contract.
type AaveV2 struct {
	AaveV2Caller     // Read-only binding to the contract
	AaveV2Transactor // Write-only binding to the contract
	AaveV2Filterer   // Log filterer for contract events
}

// AaveV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveV2Session struct {
	Contract     *AaveV2           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveV2CallerSession struct {
	Contract *AaveV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AaveV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveV2TransactorSession struct {
	Contract     *AaveV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveV2Raw struct {
	Contract *AaveV2 // Generic contract binding to access the raw methods on
}

// AaveV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveV2CallerRaw struct {
	Contract *AaveV2Caller // Generic read-only contract binding to access the raw methods on
}

// AaveV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveV2TransactorRaw struct {
	Contract *AaveV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveV2 creates a new instance of AaveV2, bound to a specific deployed contract.
func NewAaveV2(address common.Address, backend bind.ContractBackend) (*AaveV2, error) {
	contract, err := bindAaveV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveV2{AaveV2Caller: AaveV2Caller{contract: contract}, AaveV2Transactor: AaveV2Transactor{contract: contract}, AaveV2Filterer: AaveV2Filterer{contract: contract}}, nil
}

// NewAaveV2Caller creates a new read-only instance of AaveV2, bound to a specific deployed contract.
func NewAaveV2Caller(address common.Address, caller bind.ContractCaller) (*AaveV2Caller, error) {
	contract, err := bindAaveV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV2Caller{contract: contract}, nil
}

// NewAaveV2Transactor creates a new write-only instance of AaveV2, bound to a specific deployed contract.
func NewAaveV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveV2Transactor, error) {
	contract, err := bindAaveV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV2Transactor{contract: contract}, nil
}

// NewAaveV2Filterer creates a new log filterer instance of AaveV2, bound to a specific deployed contract.
func NewAaveV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveV2Filterer, error) {
	contract, err := bindAaveV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveV2Filterer{contract: contract}, nil
}

// bindAaveV2 binds a generic wrapper to an already deployed contract.
func bindAaveV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV2 *AaveV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV2.Contract.AaveV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV2 *AaveV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2.Contract.AaveV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV2 *AaveV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV2.Contract.AaveV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV2 *AaveV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV2 *AaveV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV2 *AaveV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV2.Contract.contract.Transact(opts, method, params...)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_AaveV2 *AaveV2Caller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AaveV2.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_AaveV2 *AaveV2Session) IsShutdown() (bool, error) {
	return _AaveV2.Contract.IsShutdown(&_AaveV2.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_AaveV2 *AaveV2CallerSession) IsShutdown() (bool, error) {
	return _AaveV2.Contract.IsShutdown(&_AaveV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV2 *AaveV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV2 *AaveV2Session) Owner() (common.Address, error) {
	return _AaveV2.Contract.Owner(&_AaveV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV2 *AaveV2CallerSession) Owner() (common.Address, error) {
	return _AaveV2.Contract.Owner(&_AaveV2.CallOpts)
}

// AccrueFees is a paid mutator transaction binding the contract method 0x37a4e834.
//
// Solidity: function accrueFees() returns()
func (_AaveV2 *AaveV2Transactor) AccrueFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "accrueFees")
}

// AccrueFees is a paid mutator transaction binding the contract method 0x37a4e834.
//
// Solidity: function accrueFees() returns()
func (_AaveV2 *AaveV2Session) AccrueFees() (*types.Transaction, error) {
	return _AaveV2.Contract.AccrueFees(&_AaveV2.TransactOpts)
}

// AccrueFees is a paid mutator transaction binding the contract method 0x37a4e834.
//
// Solidity: function accrueFees() returns()
func (_AaveV2 *AaveV2TransactorSession) AccrueFees() (*types.Transaction, error) {
	return _AaveV2.Contract.AccrueFees(&_AaveV2.TransactOpts)
}

// ClaimAndUnstake is a paid mutator transaction binding the contract method 0xad004e20.
//
// Solidity: function claimAndUnstake() returns(uint256 claimed)
func (_AaveV2 *AaveV2Transactor) ClaimAndUnstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "claimAndUnstake")
}

// ClaimAndUnstake is a paid mutator transaction binding the contract method 0xad004e20.
//
// Solidity: function claimAndUnstake() returns(uint256 claimed)
func (_AaveV2 *AaveV2Session) ClaimAndUnstake() (*types.Transaction, error) {
	return _AaveV2.Contract.ClaimAndUnstake(&_AaveV2.TransactOpts)
}

// ClaimAndUnstake is a paid mutator transaction binding the contract method 0xad004e20.
//
// Solidity: function claimAndUnstake() returns(uint256 claimed)
func (_AaveV2 *AaveV2TransactorSession) ClaimAndUnstake() (*types.Transaction, error) {
	return _AaveV2.Contract.ClaimAndUnstake(&_AaveV2.TransactOpts)
}

// EnterPosition is a paid mutator transaction binding the contract method 0x3dc6eabf.
//
// Solidity: function enterPosition() returns()
func (_AaveV2 *AaveV2Transactor) EnterPosition(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "enterPosition")
}

// EnterPosition is a paid mutator transaction binding the contract method 0x3dc6eabf.
//
// Solidity: function enterPosition() returns()
func (_AaveV2 *AaveV2Session) EnterPosition() (*types.Transaction, error) {
	return _AaveV2.Contract.EnterPosition(&_AaveV2.TransactOpts)
}

// EnterPosition is a paid mutator transaction binding the contract method 0x3dc6eabf.
//
// Solidity: function enterPosition() returns()
func (_AaveV2 *AaveV2TransactorSession) EnterPosition() (*types.Transaction, error) {
	return _AaveV2.Contract.EnterPosition(&_AaveV2.TransactOpts)
}

// Rebalance is a paid mutator transaction binding the contract method 0x15f4c611.
//
// Solidity: function rebalance(address[9] route, uint256[3][4] swapParams, uint256 minAssetsOut) returns()
func (_AaveV2 *AaveV2Transactor) Rebalance(opts *bind.TransactOpts, route [9]common.Address, swapParams [4][3]*big.Int, minAssetsOut *big.Int) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "rebalance", route, swapParams, minAssetsOut)
}

// Rebalance is a paid mutator transaction binding the contract method 0x15f4c611.
//
// Solidity: function rebalance(address[9] route, uint256[3][4] swapParams, uint256 minAssetsOut) returns()
func (_AaveV2 *AaveV2Session) Rebalance(route [9]common.Address, swapParams [4][3]*big.Int, minAssetsOut *big.Int) (*types.Transaction, error) {
	return _AaveV2.Contract.Rebalance(&_AaveV2.TransactOpts, route, swapParams, minAssetsOut)
}

// Rebalance is a paid mutator transaction binding the contract method 0x15f4c611.
//
// Solidity: function rebalance(address[9] route, uint256[3][4] swapParams, uint256 minAssetsOut) returns()
func (_AaveV2 *AaveV2TransactorSession) Rebalance(route [9]common.Address, swapParams [4][3]*big.Int, minAssetsOut *big.Int) (*types.Transaction, error) {
	return _AaveV2.Contract.Rebalance(&_AaveV2.TransactOpts, route, swapParams, minAssetsOut)
}

// Reinvest is a paid mutator transaction binding the contract method 0x83b4918b.
//
// Solidity: function reinvest(uint256 minAssetsOut) returns()
func (_AaveV2 *AaveV2Transactor) Reinvest(opts *bind.TransactOpts, minAssetsOut *big.Int) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "reinvest", minAssetsOut)
}

// Reinvest is a paid mutator transaction binding the contract method 0x83b4918b.
//
// Solidity: function reinvest(uint256 minAssetsOut) returns()
func (_AaveV2 *AaveV2Session) Reinvest(minAssetsOut *big.Int) (*types.Transaction, error) {
	return _AaveV2.Contract.Reinvest(&_AaveV2.TransactOpts, minAssetsOut)
}

// Reinvest is a paid mutator transaction binding the contract method 0x83b4918b.
//
// Solidity: function reinvest(uint256 minAssetsOut) returns()
func (_AaveV2 *AaveV2TransactorSession) Reinvest(minAssetsOut *big.Int) (*types.Transaction, error) {
	return _AaveV2.Contract.Reinvest(&_AaveV2.TransactOpts, minAssetsOut)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV2 *AaveV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV2 *AaveV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _AaveV2.Contract.RenounceOwnership(&_AaveV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV2 *AaveV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveV2.Contract.RenounceOwnership(&_AaveV2.TransactOpts)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_AaveV2 *AaveV2Transactor) SetDepositLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "setDepositLimit", limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_AaveV2 *AaveV2Session) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _AaveV2.Contract.SetDepositLimit(&_AaveV2.TransactOpts, limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_AaveV2 *AaveV2TransactorSession) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _AaveV2.Contract.SetDepositLimit(&_AaveV2.TransactOpts, limit)
}

// SetFeesDistributor is a paid mutator transaction binding the contract method 0x6e85f183.
//
// Solidity: function setFeesDistributor(bytes32 newFeesDistributor) returns()
func (_AaveV2 *AaveV2Transactor) SetFeesDistributor(opts *bind.TransactOpts, newFeesDistributor [32]byte) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "setFeesDistributor", newFeesDistributor)
}

// SetFeesDistributor is a paid mutator transaction binding the contract method 0x6e85f183.
//
// Solidity: function setFeesDistributor(bytes32 newFeesDistributor) returns()
func (_AaveV2 *AaveV2Session) SetFeesDistributor(newFeesDistributor [32]byte) (*types.Transaction, error) {
	return _AaveV2.Contract.SetFeesDistributor(&_AaveV2.TransactOpts, newFeesDistributor)
}

// SetFeesDistributor is a paid mutator transaction binding the contract method 0x6e85f183.
//
// Solidity: function setFeesDistributor(bytes32 newFeesDistributor) returns()
func (_AaveV2 *AaveV2TransactorSession) SetFeesDistributor(newFeesDistributor [32]byte) (*types.Transaction, error) {
	return _AaveV2.Contract.SetFeesDistributor(&_AaveV2.TransactOpts, newFeesDistributor)
}

// SetLiquidityLimit is a paid mutator transaction binding the contract method 0xdf05a52a.
//
// Solidity: function setLiquidityLimit(uint256 limit) returns()
func (_AaveV2 *AaveV2Transactor) SetLiquidityLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "setLiquidityLimit", limit)
}

// SetLiquidityLimit is a paid mutator transaction binding the contract method 0xdf05a52a.
//
// Solidity: function setLiquidityLimit(uint256 limit) returns()
func (_AaveV2 *AaveV2Session) SetLiquidityLimit(limit *big.Int) (*types.Transaction, error) {
	return _AaveV2.Contract.SetLiquidityLimit(&_AaveV2.TransactOpts, limit)
}

// SetLiquidityLimit is a paid mutator transaction binding the contract method 0xdf05a52a.
//
// Solidity: function setLiquidityLimit(uint256 limit) returns()
func (_AaveV2 *AaveV2TransactorSession) SetLiquidityLimit(limit *big.Int) (*types.Transaction, error) {
	return _AaveV2.Contract.SetLiquidityLimit(&_AaveV2.TransactOpts, limit)
}

// SetShutdown is a paid mutator transaction binding the contract method 0xbb27280b.
//
// Solidity: function setShutdown(bool shutdown, bool exitPosition) returns()
func (_AaveV2 *AaveV2Transactor) SetShutdown(opts *bind.TransactOpts, shutdown bool, exitPosition bool) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "setShutdown", shutdown, exitPosition)
}

// SetShutdown is a paid mutator transaction binding the contract method 0xbb27280b.
//
// Solidity: function setShutdown(bool shutdown, bool exitPosition) returns()
func (_AaveV2 *AaveV2Session) SetShutdown(shutdown bool, exitPosition bool) (*types.Transaction, error) {
	return _AaveV2.Contract.SetShutdown(&_AaveV2.TransactOpts, shutdown, exitPosition)
}

// SetShutdown is a paid mutator transaction binding the contract method 0xbb27280b.
//
// Solidity: function setShutdown(bool shutdown, bool exitPosition) returns()
func (_AaveV2 *AaveV2TransactorSession) SetShutdown(shutdown bool, exitPosition bool) (*types.Transaction, error) {
	return _AaveV2.Contract.SetShutdown(&_AaveV2.TransactOpts, shutdown, exitPosition)
}

// SetTrust is a paid mutator transaction binding the contract method 0xcab59238.
//
// Solidity: function setTrust(address position, bool trust) returns()
func (_AaveV2 *AaveV2Transactor) SetTrust(opts *bind.TransactOpts, position common.Address, trust bool) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "setTrust", position, trust)
}

// SetTrust is a paid mutator transaction binding the contract method 0xcab59238.
//
// Solidity: function setTrust(address position, bool trust) returns()
func (_AaveV2 *AaveV2Session) SetTrust(position common.Address, trust bool) (*types.Transaction, error) {
	return _AaveV2.Contract.SetTrust(&_AaveV2.TransactOpts, position, trust)
}

// SetTrust is a paid mutator transaction binding the contract method 0xcab59238.
//
// Solidity: function setTrust(address position, bool trust) returns()
func (_AaveV2 *AaveV2TransactorSession) SetTrust(position common.Address, trust bool) (*types.Transaction, error) {
	return _AaveV2.Contract.SetTrust(&_AaveV2.TransactOpts, position, trust)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address token, address to) returns()
func (_AaveV2 *AaveV2Transactor) Sweep(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "sweep", token, to)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address token, address to) returns()
func (_AaveV2 *AaveV2Session) Sweep(token common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveV2.Contract.Sweep(&_AaveV2.TransactOpts, token, to)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address token, address to) returns()
func (_AaveV2 *AaveV2TransactorSession) Sweep(token common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveV2.Contract.Sweep(&_AaveV2.TransactOpts, token, to)
}

// TransferFees is a paid mutator transaction binding the contract method 0xc2fbe7bc.
//
// Solidity: function transferFees() returns()
func (_AaveV2 *AaveV2Transactor) TransferFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "transferFees")
}

// TransferFees is a paid mutator transaction binding the contract method 0xc2fbe7bc.
//
// Solidity: function transferFees() returns()
func (_AaveV2 *AaveV2Session) TransferFees() (*types.Transaction, error) {
	return _AaveV2.Contract.TransferFees(&_AaveV2.TransactOpts)
}

// TransferFees is a paid mutator transaction binding the contract method 0xc2fbe7bc.
//
// Solidity: function transferFees() returns()
func (_AaveV2 *AaveV2TransactorSession) TransferFees() (*types.Transaction, error) {
	return _AaveV2.Contract.TransferFees(&_AaveV2.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV2 *AaveV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AaveV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV2 *AaveV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveV2.Contract.TransferOwnership(&_AaveV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV2 *AaveV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveV2.Contract.TransferOwnership(&_AaveV2.TransactOpts, newOwner)
}

// AaveV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AaveV2 contract.
type AaveV2OwnershipTransferredIterator struct {
	Event *AaveV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AaveV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2OwnershipTransferred)
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
		it.Event = new(AaveV2OwnershipTransferred)
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
func (it *AaveV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2OwnershipTransferred represents a OwnershipTransferred event raised by the AaveV2 contract.
type AaveV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV2 *AaveV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AaveV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2OwnershipTransferredIterator{contract: _AaveV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV2 *AaveV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AaveV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2OwnershipTransferred)
				if err := _AaveV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV2 *AaveV2Filterer) ParseOwnershipTransferred(log types.Log) (*AaveV2OwnershipTransferred, error) {
	event := new(AaveV2OwnershipTransferred)
	if err := _AaveV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockAccrueFeesIterator is returned from FilterMockAccrueFees and is used to iterate over the raw logs and unpacked data for MockAccrueFees events raised by the AaveV2 contract.
type AaveV2MockAccrueFeesIterator struct {
	Event *AaveV2MockAccrueFees // Event containing the contract specifics and raw log

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
func (it *AaveV2MockAccrueFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockAccrueFees)
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
		it.Event = new(AaveV2MockAccrueFees)
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
func (it *AaveV2MockAccrueFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockAccrueFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockAccrueFees represents a MockAccrueFees event raised by the AaveV2 contract.
type AaveV2MockAccrueFees struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMockAccrueFees is a free log retrieval operation binding the contract event 0xeea290d1596cec66f2811f6ec6d4b3e366f8f9b263e15ea87fb624ab69924042.
//
// Solidity: event mockAccrueFees()
func (_AaveV2 *AaveV2Filterer) FilterMockAccrueFees(opts *bind.FilterOpts) (*AaveV2MockAccrueFeesIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockAccrueFees")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockAccrueFeesIterator{contract: _AaveV2.contract, event: "mockAccrueFees", logs: logs, sub: sub}, nil
}

// WatchMockAccrueFees is a free log subscription operation binding the contract event 0xeea290d1596cec66f2811f6ec6d4b3e366f8f9b263e15ea87fb624ab69924042.
//
// Solidity: event mockAccrueFees()
func (_AaveV2 *AaveV2Filterer) WatchMockAccrueFees(opts *bind.WatchOpts, sink chan<- *AaveV2MockAccrueFees) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockAccrueFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockAccrueFees)
				if err := _AaveV2.contract.UnpackLog(event, "mockAccrueFees", log); err != nil {
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

// ParseMockAccrueFees is a log parse operation binding the contract event 0xeea290d1596cec66f2811f6ec6d4b3e366f8f9b263e15ea87fb624ab69924042.
//
// Solidity: event mockAccrueFees()
func (_AaveV2 *AaveV2Filterer) ParseMockAccrueFees(log types.Log) (*AaveV2MockAccrueFees, error) {
	event := new(AaveV2MockAccrueFees)
	if err := _AaveV2.contract.UnpackLog(event, "mockAccrueFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockClaimAndUnstakeIterator is returned from FilterMockClaimAndUnstake and is used to iterate over the raw logs and unpacked data for MockClaimAndUnstake events raised by the AaveV2 contract.
type AaveV2MockClaimAndUnstakeIterator struct {
	Event *AaveV2MockClaimAndUnstake // Event containing the contract specifics and raw log

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
func (it *AaveV2MockClaimAndUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockClaimAndUnstake)
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
		it.Event = new(AaveV2MockClaimAndUnstake)
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
func (it *AaveV2MockClaimAndUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockClaimAndUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockClaimAndUnstake represents a MockClaimAndUnstake event raised by the AaveV2 contract.
type AaveV2MockClaimAndUnstake struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMockClaimAndUnstake is a free log retrieval operation binding the contract event 0xbfb260c8d19d48ed37a6d7379cdb9622833ea0d1f8389d3b16b0a9dc5cce17b8.
//
// Solidity: event mockClaimAndUnstake()
func (_AaveV2 *AaveV2Filterer) FilterMockClaimAndUnstake(opts *bind.FilterOpts) (*AaveV2MockClaimAndUnstakeIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockClaimAndUnstake")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockClaimAndUnstakeIterator{contract: _AaveV2.contract, event: "mockClaimAndUnstake", logs: logs, sub: sub}, nil
}

// WatchMockClaimAndUnstake is a free log subscription operation binding the contract event 0xbfb260c8d19d48ed37a6d7379cdb9622833ea0d1f8389d3b16b0a9dc5cce17b8.
//
// Solidity: event mockClaimAndUnstake()
func (_AaveV2 *AaveV2Filterer) WatchMockClaimAndUnstake(opts *bind.WatchOpts, sink chan<- *AaveV2MockClaimAndUnstake) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockClaimAndUnstake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockClaimAndUnstake)
				if err := _AaveV2.contract.UnpackLog(event, "mockClaimAndUnstake", log); err != nil {
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

// ParseMockClaimAndUnstake is a log parse operation binding the contract event 0xbfb260c8d19d48ed37a6d7379cdb9622833ea0d1f8389d3b16b0a9dc5cce17b8.
//
// Solidity: event mockClaimAndUnstake()
func (_AaveV2 *AaveV2Filterer) ParseMockClaimAndUnstake(log types.Log) (*AaveV2MockClaimAndUnstake, error) {
	event := new(AaveV2MockClaimAndUnstake)
	if err := _AaveV2.contract.UnpackLog(event, "mockClaimAndUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockEnterPositionIterator is returned from FilterMockEnterPosition and is used to iterate over the raw logs and unpacked data for MockEnterPosition events raised by the AaveV2 contract.
type AaveV2MockEnterPositionIterator struct {
	Event *AaveV2MockEnterPosition // Event containing the contract specifics and raw log

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
func (it *AaveV2MockEnterPositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockEnterPosition)
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
		it.Event = new(AaveV2MockEnterPosition)
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
func (it *AaveV2MockEnterPositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockEnterPositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockEnterPosition represents a MockEnterPosition event raised by the AaveV2 contract.
type AaveV2MockEnterPosition struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMockEnterPosition is a free log retrieval operation binding the contract event 0xac91e57d60d4977f906c8747139df337f4553e15652b9306669094fe4f77ac07.
//
// Solidity: event mockEnterPosition()
func (_AaveV2 *AaveV2Filterer) FilterMockEnterPosition(opts *bind.FilterOpts) (*AaveV2MockEnterPositionIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockEnterPosition")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockEnterPositionIterator{contract: _AaveV2.contract, event: "mockEnterPosition", logs: logs, sub: sub}, nil
}

// WatchMockEnterPosition is a free log subscription operation binding the contract event 0xac91e57d60d4977f906c8747139df337f4553e15652b9306669094fe4f77ac07.
//
// Solidity: event mockEnterPosition()
func (_AaveV2 *AaveV2Filterer) WatchMockEnterPosition(opts *bind.WatchOpts, sink chan<- *AaveV2MockEnterPosition) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockEnterPosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockEnterPosition)
				if err := _AaveV2.contract.UnpackLog(event, "mockEnterPosition", log); err != nil {
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

// ParseMockEnterPosition is a log parse operation binding the contract event 0xac91e57d60d4977f906c8747139df337f4553e15652b9306669094fe4f77ac07.
//
// Solidity: event mockEnterPosition()
func (_AaveV2 *AaveV2Filterer) ParseMockEnterPosition(log types.Log) (*AaveV2MockEnterPosition, error) {
	event := new(AaveV2MockEnterPosition)
	if err := _AaveV2.contract.UnpackLog(event, "mockEnterPosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockRebalanceIterator is returned from FilterMockRebalance and is used to iterate over the raw logs and unpacked data for MockRebalance events raised by the AaveV2 contract.
type AaveV2MockRebalanceIterator struct {
	Event *AaveV2MockRebalance // Event containing the contract specifics and raw log

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
func (it *AaveV2MockRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockRebalance)
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
		it.Event = new(AaveV2MockRebalance)
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
func (it *AaveV2MockRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockRebalance represents a MockRebalance event raised by the AaveV2 contract.
type AaveV2MockRebalance struct {
	Route        [9]common.Address
	SwapParams   [4][3]*big.Int
	MinAssetsOut *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMockRebalance is a free log retrieval operation binding the contract event 0x30f4d32b8116c263624c7aee2a3725a23891b165da52311c9e389d9d44ee8b94.
//
// Solidity: event mockRebalance(address[9] route, uint256[3][4] swapParams, uint256 minAssetsOut)
func (_AaveV2 *AaveV2Filterer) FilterMockRebalance(opts *bind.FilterOpts) (*AaveV2MockRebalanceIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockRebalance")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockRebalanceIterator{contract: _AaveV2.contract, event: "mockRebalance", logs: logs, sub: sub}, nil
}

// WatchMockRebalance is a free log subscription operation binding the contract event 0x30f4d32b8116c263624c7aee2a3725a23891b165da52311c9e389d9d44ee8b94.
//
// Solidity: event mockRebalance(address[9] route, uint256[3][4] swapParams, uint256 minAssetsOut)
func (_AaveV2 *AaveV2Filterer) WatchMockRebalance(opts *bind.WatchOpts, sink chan<- *AaveV2MockRebalance) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockRebalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockRebalance)
				if err := _AaveV2.contract.UnpackLog(event, "mockRebalance", log); err != nil {
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

// ParseMockRebalance is a log parse operation binding the contract event 0x30f4d32b8116c263624c7aee2a3725a23891b165da52311c9e389d9d44ee8b94.
//
// Solidity: event mockRebalance(address[9] route, uint256[3][4] swapParams, uint256 minAssetsOut)
func (_AaveV2 *AaveV2Filterer) ParseMockRebalance(log types.Log) (*AaveV2MockRebalance, error) {
	event := new(AaveV2MockRebalance)
	if err := _AaveV2.contract.UnpackLog(event, "mockRebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockReinvestIterator is returned from FilterMockReinvest and is used to iterate over the raw logs and unpacked data for MockReinvest events raised by the AaveV2 contract.
type AaveV2MockReinvestIterator struct {
	Event *AaveV2MockReinvest // Event containing the contract specifics and raw log

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
func (it *AaveV2MockReinvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockReinvest)
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
		it.Event = new(AaveV2MockReinvest)
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
func (it *AaveV2MockReinvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockReinvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockReinvest represents a MockReinvest event raised by the AaveV2 contract.
type AaveV2MockReinvest struct {
	MinAssetsOut *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMockReinvest is a free log retrieval operation binding the contract event 0x30c7b12b693373870dedbf5ae1c074038300870753880d4383acbd89d0727ebb.
//
// Solidity: event mockReinvest(uint256 minAssetsOut)
func (_AaveV2 *AaveV2Filterer) FilterMockReinvest(opts *bind.FilterOpts) (*AaveV2MockReinvestIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockReinvest")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockReinvestIterator{contract: _AaveV2.contract, event: "mockReinvest", logs: logs, sub: sub}, nil
}

// WatchMockReinvest is a free log subscription operation binding the contract event 0x30c7b12b693373870dedbf5ae1c074038300870753880d4383acbd89d0727ebb.
//
// Solidity: event mockReinvest(uint256 minAssetsOut)
func (_AaveV2 *AaveV2Filterer) WatchMockReinvest(opts *bind.WatchOpts, sink chan<- *AaveV2MockReinvest) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockReinvest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockReinvest)
				if err := _AaveV2.contract.UnpackLog(event, "mockReinvest", log); err != nil {
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

// ParseMockReinvest is a log parse operation binding the contract event 0x30c7b12b693373870dedbf5ae1c074038300870753880d4383acbd89d0727ebb.
//
// Solidity: event mockReinvest(uint256 minAssetsOut)
func (_AaveV2 *AaveV2Filterer) ParseMockReinvest(log types.Log) (*AaveV2MockReinvest, error) {
	event := new(AaveV2MockReinvest)
	if err := _AaveV2.contract.UnpackLog(event, "mockReinvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockSetDepositLimitIterator is returned from FilterMockSetDepositLimit and is used to iterate over the raw logs and unpacked data for MockSetDepositLimit events raised by the AaveV2 contract.
type AaveV2MockSetDepositLimitIterator struct {
	Event *AaveV2MockSetDepositLimit // Event containing the contract specifics and raw log

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
func (it *AaveV2MockSetDepositLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockSetDepositLimit)
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
		it.Event = new(AaveV2MockSetDepositLimit)
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
func (it *AaveV2MockSetDepositLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockSetDepositLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockSetDepositLimit represents a MockSetDepositLimit event raised by the AaveV2 contract.
type AaveV2MockSetDepositLimit struct {
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMockSetDepositLimit is a free log retrieval operation binding the contract event 0x45bfc2de9a8e67c77e8bb4f799dba14b1f33b20360a5376ad130ad5a21acf881.
//
// Solidity: event mockSetDepositLimit(uint256 limit)
func (_AaveV2 *AaveV2Filterer) FilterMockSetDepositLimit(opts *bind.FilterOpts) (*AaveV2MockSetDepositLimitIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockSetDepositLimit")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockSetDepositLimitIterator{contract: _AaveV2.contract, event: "mockSetDepositLimit", logs: logs, sub: sub}, nil
}

// WatchMockSetDepositLimit is a free log subscription operation binding the contract event 0x45bfc2de9a8e67c77e8bb4f799dba14b1f33b20360a5376ad130ad5a21acf881.
//
// Solidity: event mockSetDepositLimit(uint256 limit)
func (_AaveV2 *AaveV2Filterer) WatchMockSetDepositLimit(opts *bind.WatchOpts, sink chan<- *AaveV2MockSetDepositLimit) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockSetDepositLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockSetDepositLimit)
				if err := _AaveV2.contract.UnpackLog(event, "mockSetDepositLimit", log); err != nil {
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

// ParseMockSetDepositLimit is a log parse operation binding the contract event 0x45bfc2de9a8e67c77e8bb4f799dba14b1f33b20360a5376ad130ad5a21acf881.
//
// Solidity: event mockSetDepositLimit(uint256 limit)
func (_AaveV2 *AaveV2Filterer) ParseMockSetDepositLimit(log types.Log) (*AaveV2MockSetDepositLimit, error) {
	event := new(AaveV2MockSetDepositLimit)
	if err := _AaveV2.contract.UnpackLog(event, "mockSetDepositLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockSetFeesDistributorIterator is returned from FilterMockSetFeesDistributor and is used to iterate over the raw logs and unpacked data for MockSetFeesDistributor events raised by the AaveV2 contract.
type AaveV2MockSetFeesDistributorIterator struct {
	Event *AaveV2MockSetFeesDistributor // Event containing the contract specifics and raw log

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
func (it *AaveV2MockSetFeesDistributorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockSetFeesDistributor)
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
		it.Event = new(AaveV2MockSetFeesDistributor)
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
func (it *AaveV2MockSetFeesDistributorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockSetFeesDistributorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockSetFeesDistributor represents a MockSetFeesDistributor event raised by the AaveV2 contract.
type AaveV2MockSetFeesDistributor struct {
	NewFeesDistributor [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMockSetFeesDistributor is a free log retrieval operation binding the contract event 0x011f7dbcc31d2199a0d6fbb2dac3a05053f7af1c127dcccb02f7138b059b4b25.
//
// Solidity: event mockSetFeesDistributor(bytes32 newFeesDistributor)
func (_AaveV2 *AaveV2Filterer) FilterMockSetFeesDistributor(opts *bind.FilterOpts) (*AaveV2MockSetFeesDistributorIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockSetFeesDistributor")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockSetFeesDistributorIterator{contract: _AaveV2.contract, event: "mockSetFeesDistributor", logs: logs, sub: sub}, nil
}

// WatchMockSetFeesDistributor is a free log subscription operation binding the contract event 0x011f7dbcc31d2199a0d6fbb2dac3a05053f7af1c127dcccb02f7138b059b4b25.
//
// Solidity: event mockSetFeesDistributor(bytes32 newFeesDistributor)
func (_AaveV2 *AaveV2Filterer) WatchMockSetFeesDistributor(opts *bind.WatchOpts, sink chan<- *AaveV2MockSetFeesDistributor) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockSetFeesDistributor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockSetFeesDistributor)
				if err := _AaveV2.contract.UnpackLog(event, "mockSetFeesDistributor", log); err != nil {
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

// ParseMockSetFeesDistributor is a log parse operation binding the contract event 0x011f7dbcc31d2199a0d6fbb2dac3a05053f7af1c127dcccb02f7138b059b4b25.
//
// Solidity: event mockSetFeesDistributor(bytes32 newFeesDistributor)
func (_AaveV2 *AaveV2Filterer) ParseMockSetFeesDistributor(log types.Log) (*AaveV2MockSetFeesDistributor, error) {
	event := new(AaveV2MockSetFeesDistributor)
	if err := _AaveV2.contract.UnpackLog(event, "mockSetFeesDistributor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockSetLiquidityLimitIterator is returned from FilterMockSetLiquidityLimit and is used to iterate over the raw logs and unpacked data for MockSetLiquidityLimit events raised by the AaveV2 contract.
type AaveV2MockSetLiquidityLimitIterator struct {
	Event *AaveV2MockSetLiquidityLimit // Event containing the contract specifics and raw log

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
func (it *AaveV2MockSetLiquidityLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockSetLiquidityLimit)
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
		it.Event = new(AaveV2MockSetLiquidityLimit)
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
func (it *AaveV2MockSetLiquidityLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockSetLiquidityLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockSetLiquidityLimit represents a MockSetLiquidityLimit event raised by the AaveV2 contract.
type AaveV2MockSetLiquidityLimit struct {
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMockSetLiquidityLimit is a free log retrieval operation binding the contract event 0xb992a37793e6e84172de1507059850072e4a3af03ce22c8247c6b9428030aa3e.
//
// Solidity: event mockSetLiquidityLimit(uint256 limit)
func (_AaveV2 *AaveV2Filterer) FilterMockSetLiquidityLimit(opts *bind.FilterOpts) (*AaveV2MockSetLiquidityLimitIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockSetLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockSetLiquidityLimitIterator{contract: _AaveV2.contract, event: "mockSetLiquidityLimit", logs: logs, sub: sub}, nil
}

// WatchMockSetLiquidityLimit is a free log subscription operation binding the contract event 0xb992a37793e6e84172de1507059850072e4a3af03ce22c8247c6b9428030aa3e.
//
// Solidity: event mockSetLiquidityLimit(uint256 limit)
func (_AaveV2 *AaveV2Filterer) WatchMockSetLiquidityLimit(opts *bind.WatchOpts, sink chan<- *AaveV2MockSetLiquidityLimit) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockSetLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockSetLiquidityLimit)
				if err := _AaveV2.contract.UnpackLog(event, "mockSetLiquidityLimit", log); err != nil {
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

// ParseMockSetLiquidityLimit is a log parse operation binding the contract event 0xb992a37793e6e84172de1507059850072e4a3af03ce22c8247c6b9428030aa3e.
//
// Solidity: event mockSetLiquidityLimit(uint256 limit)
func (_AaveV2 *AaveV2Filterer) ParseMockSetLiquidityLimit(log types.Log) (*AaveV2MockSetLiquidityLimit, error) {
	event := new(AaveV2MockSetLiquidityLimit)
	if err := _AaveV2.contract.UnpackLog(event, "mockSetLiquidityLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockSetShutdownIterator is returned from FilterMockSetShutdown and is used to iterate over the raw logs and unpacked data for MockSetShutdown events raised by the AaveV2 contract.
type AaveV2MockSetShutdownIterator struct {
	Event *AaveV2MockSetShutdown // Event containing the contract specifics and raw log

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
func (it *AaveV2MockSetShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockSetShutdown)
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
		it.Event = new(AaveV2MockSetShutdown)
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
func (it *AaveV2MockSetShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockSetShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockSetShutdown represents a MockSetShutdown event raised by the AaveV2 contract.
type AaveV2MockSetShutdown struct {
	Shutdown     bool
	ExitPosition bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMockSetShutdown is a free log retrieval operation binding the contract event 0x5afb9c112875465fbf21f4b0923a311393e37940d92ea70da88f4f2532cf8e05.
//
// Solidity: event mockSetShutdown(bool shutdown, bool exitPosition)
func (_AaveV2 *AaveV2Filterer) FilterMockSetShutdown(opts *bind.FilterOpts) (*AaveV2MockSetShutdownIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockSetShutdown")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockSetShutdownIterator{contract: _AaveV2.contract, event: "mockSetShutdown", logs: logs, sub: sub}, nil
}

// WatchMockSetShutdown is a free log subscription operation binding the contract event 0x5afb9c112875465fbf21f4b0923a311393e37940d92ea70da88f4f2532cf8e05.
//
// Solidity: event mockSetShutdown(bool shutdown, bool exitPosition)
func (_AaveV2 *AaveV2Filterer) WatchMockSetShutdown(opts *bind.WatchOpts, sink chan<- *AaveV2MockSetShutdown) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockSetShutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockSetShutdown)
				if err := _AaveV2.contract.UnpackLog(event, "mockSetShutdown", log); err != nil {
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

// ParseMockSetShutdown is a log parse operation binding the contract event 0x5afb9c112875465fbf21f4b0923a311393e37940d92ea70da88f4f2532cf8e05.
//
// Solidity: event mockSetShutdown(bool shutdown, bool exitPosition)
func (_AaveV2 *AaveV2Filterer) ParseMockSetShutdown(log types.Log) (*AaveV2MockSetShutdown, error) {
	event := new(AaveV2MockSetShutdown)
	if err := _AaveV2.contract.UnpackLog(event, "mockSetShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockSetTrustIterator is returned from FilterMockSetTrust and is used to iterate over the raw logs and unpacked data for MockSetTrust events raised by the AaveV2 contract.
type AaveV2MockSetTrustIterator struct {
	Event *AaveV2MockSetTrust // Event containing the contract specifics and raw log

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
func (it *AaveV2MockSetTrustIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockSetTrust)
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
		it.Event = new(AaveV2MockSetTrust)
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
func (it *AaveV2MockSetTrustIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockSetTrustIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockSetTrust represents a MockSetTrust event raised by the AaveV2 contract.
type AaveV2MockSetTrust struct {
	Position common.Address
	Trust    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMockSetTrust is a free log retrieval operation binding the contract event 0x0f0c347f3545e0dc49b875ada970dcf21943a65ef1c3a7ba91aa13c35b1aa4a9.
//
// Solidity: event mockSetTrust(address position, bool trust)
func (_AaveV2 *AaveV2Filterer) FilterMockSetTrust(opts *bind.FilterOpts) (*AaveV2MockSetTrustIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockSetTrust")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockSetTrustIterator{contract: _AaveV2.contract, event: "mockSetTrust", logs: logs, sub: sub}, nil
}

// WatchMockSetTrust is a free log subscription operation binding the contract event 0x0f0c347f3545e0dc49b875ada970dcf21943a65ef1c3a7ba91aa13c35b1aa4a9.
//
// Solidity: event mockSetTrust(address position, bool trust)
func (_AaveV2 *AaveV2Filterer) WatchMockSetTrust(opts *bind.WatchOpts, sink chan<- *AaveV2MockSetTrust) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockSetTrust")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockSetTrust)
				if err := _AaveV2.contract.UnpackLog(event, "mockSetTrust", log); err != nil {
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

// ParseMockSetTrust is a log parse operation binding the contract event 0x0f0c347f3545e0dc49b875ada970dcf21943a65ef1c3a7ba91aa13c35b1aa4a9.
//
// Solidity: event mockSetTrust(address position, bool trust)
func (_AaveV2 *AaveV2Filterer) ParseMockSetTrust(log types.Log) (*AaveV2MockSetTrust, error) {
	event := new(AaveV2MockSetTrust)
	if err := _AaveV2.contract.UnpackLog(event, "mockSetTrust", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockSweepIterator is returned from FilterMockSweep and is used to iterate over the raw logs and unpacked data for MockSweep events raised by the AaveV2 contract.
type AaveV2MockSweepIterator struct {
	Event *AaveV2MockSweep // Event containing the contract specifics and raw log

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
func (it *AaveV2MockSweepIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockSweep)
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
		it.Event = new(AaveV2MockSweep)
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
func (it *AaveV2MockSweepIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockSweepIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockSweep represents a MockSweep event raised by the AaveV2 contract.
type AaveV2MockSweep struct {
	Token common.Address
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMockSweep is a free log retrieval operation binding the contract event 0x01708d8c4133bc9d37e3ab9d6571e7bb8dac5304e30008172855af779a282f9d.
//
// Solidity: event mockSweep(address token, address to)
func (_AaveV2 *AaveV2Filterer) FilterMockSweep(opts *bind.FilterOpts) (*AaveV2MockSweepIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockSweep")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockSweepIterator{contract: _AaveV2.contract, event: "mockSweep", logs: logs, sub: sub}, nil
}

// WatchMockSweep is a free log subscription operation binding the contract event 0x01708d8c4133bc9d37e3ab9d6571e7bb8dac5304e30008172855af779a282f9d.
//
// Solidity: event mockSweep(address token, address to)
func (_AaveV2 *AaveV2Filterer) WatchMockSweep(opts *bind.WatchOpts, sink chan<- *AaveV2MockSweep) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockSweep")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockSweep)
				if err := _AaveV2.contract.UnpackLog(event, "mockSweep", log); err != nil {
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

// ParseMockSweep is a log parse operation binding the contract event 0x01708d8c4133bc9d37e3ab9d6571e7bb8dac5304e30008172855af779a282f9d.
//
// Solidity: event mockSweep(address token, address to)
func (_AaveV2 *AaveV2Filterer) ParseMockSweep(log types.Log) (*AaveV2MockSweep, error) {
	event := new(AaveV2MockSweep)
	if err := _AaveV2.contract.UnpackLog(event, "mockSweep", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV2MockTransferFeesIterator is returned from FilterMockTransferFees and is used to iterate over the raw logs and unpacked data for MockTransferFees events raised by the AaveV2 contract.
type AaveV2MockTransferFeesIterator struct {
	Event *AaveV2MockTransferFees // Event containing the contract specifics and raw log

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
func (it *AaveV2MockTransferFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2MockTransferFees)
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
		it.Event = new(AaveV2MockTransferFees)
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
func (it *AaveV2MockTransferFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2MockTransferFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2MockTransferFees represents a MockTransferFees event raised by the AaveV2 contract.
type AaveV2MockTransferFees struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMockTransferFees is a free log retrieval operation binding the contract event 0x9e1dd58aa34e20b52e0befa7246182b86d06fd56ea0a521d4d3dc8d6241cc43a.
//
// Solidity: event mockTransferFees()
func (_AaveV2 *AaveV2Filterer) FilterMockTransferFees(opts *bind.FilterOpts) (*AaveV2MockTransferFeesIterator, error) {

	logs, sub, err := _AaveV2.contract.FilterLogs(opts, "mockTransferFees")
	if err != nil {
		return nil, err
	}
	return &AaveV2MockTransferFeesIterator{contract: _AaveV2.contract, event: "mockTransferFees", logs: logs, sub: sub}, nil
}

// WatchMockTransferFees is a free log subscription operation binding the contract event 0x9e1dd58aa34e20b52e0befa7246182b86d06fd56ea0a521d4d3dc8d6241cc43a.
//
// Solidity: event mockTransferFees()
func (_AaveV2 *AaveV2Filterer) WatchMockTransferFees(opts *bind.WatchOpts, sink chan<- *AaveV2MockTransferFees) (event.Subscription, error) {

	logs, sub, err := _AaveV2.contract.WatchLogs(opts, "mockTransferFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2MockTransferFees)
				if err := _AaveV2.contract.UnpackLog(event, "mockTransferFees", log); err != nil {
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

// ParseMockTransferFees is a log parse operation binding the contract event 0x9e1dd58aa34e20b52e0befa7246182b86d06fd56ea0a521d4d3dc8d6241cc43a.
//
// Solidity: event mockTransferFees()
func (_AaveV2 *AaveV2Filterer) ParseMockTransferFees(log types.Log) (*AaveV2MockTransferFees, error) {
	event := new(AaveV2MockTransferFees)
	if err := _AaveV2.contract.UnpackLog(event, "mockTransferFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
