// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fastlaneOnlineSolver

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

// Swap is an auto generated low-level Go binding around an user-defined struct.
type Swap struct {
	DexType  uint8
	PoolAddr common.Address
	TokenIn  common.Address
	TokenOut common.Address
}

// FastlaneOnlineSolverMetaData contains all meta data concerning the FastlaneOnlineSolver contract.
var FastlaneOnlineSolverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"atlas\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"solverOpFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"solverOpData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"atlasSolverCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumDexType\",\"name\":\"dexType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"internalType\":\"structSwap[]\",\"name\":\"swapPath\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FastlaneOnlineSolverABI is the input ABI used to generate the binding from.
// Deprecated: Use FastlaneOnlineSolverMetaData.ABI instead.
var FastlaneOnlineSolverABI = FastlaneOnlineSolverMetaData.ABI

// FastlaneOnlineSolver is an auto generated Go binding around an Ethereum contract.
type FastlaneOnlineSolver struct {
	FastlaneOnlineSolverCaller     // Read-only binding to the contract
	FastlaneOnlineSolverTransactor // Write-only binding to the contract
	FastlaneOnlineSolverFilterer   // Log filterer for contract events
}

// FastlaneOnlineSolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type FastlaneOnlineSolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastlaneOnlineSolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FastlaneOnlineSolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastlaneOnlineSolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FastlaneOnlineSolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastlaneOnlineSolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FastlaneOnlineSolverSession struct {
	Contract     *FastlaneOnlineSolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FastlaneOnlineSolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FastlaneOnlineSolverCallerSession struct {
	Contract *FastlaneOnlineSolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// FastlaneOnlineSolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FastlaneOnlineSolverTransactorSession struct {
	Contract     *FastlaneOnlineSolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FastlaneOnlineSolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type FastlaneOnlineSolverRaw struct {
	Contract *FastlaneOnlineSolver // Generic contract binding to access the raw methods on
}

// FastlaneOnlineSolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FastlaneOnlineSolverCallerRaw struct {
	Contract *FastlaneOnlineSolverCaller // Generic read-only contract binding to access the raw methods on
}

// FastlaneOnlineSolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FastlaneOnlineSolverTransactorRaw struct {
	Contract *FastlaneOnlineSolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFastlaneOnlineSolver creates a new instance of FastlaneOnlineSolver, bound to a specific deployed contract.
func NewFastlaneOnlineSolver(address common.Address, backend bind.ContractBackend) (*FastlaneOnlineSolver, error) {
	contract, err := bindFastlaneOnlineSolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FastlaneOnlineSolver{FastlaneOnlineSolverCaller: FastlaneOnlineSolverCaller{contract: contract}, FastlaneOnlineSolverTransactor: FastlaneOnlineSolverTransactor{contract: contract}, FastlaneOnlineSolverFilterer: FastlaneOnlineSolverFilterer{contract: contract}}, nil
}

// NewFastlaneOnlineSolverCaller creates a new read-only instance of FastlaneOnlineSolver, bound to a specific deployed contract.
func NewFastlaneOnlineSolverCaller(address common.Address, caller bind.ContractCaller) (*FastlaneOnlineSolverCaller, error) {
	contract, err := bindFastlaneOnlineSolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FastlaneOnlineSolverCaller{contract: contract}, nil
}

// NewFastlaneOnlineSolverTransactor creates a new write-only instance of FastlaneOnlineSolver, bound to a specific deployed contract.
func NewFastlaneOnlineSolverTransactor(address common.Address, transactor bind.ContractTransactor) (*FastlaneOnlineSolverTransactor, error) {
	contract, err := bindFastlaneOnlineSolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FastlaneOnlineSolverTransactor{contract: contract}, nil
}

// NewFastlaneOnlineSolverFilterer creates a new log filterer instance of FastlaneOnlineSolver, bound to a specific deployed contract.
func NewFastlaneOnlineSolverFilterer(address common.Address, filterer bind.ContractFilterer) (*FastlaneOnlineSolverFilterer, error) {
	contract, err := bindFastlaneOnlineSolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FastlaneOnlineSolverFilterer{contract: contract}, nil
}

// bindFastlaneOnlineSolver binds a generic wrapper to an already deployed contract.
func bindFastlaneOnlineSolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FastlaneOnlineSolverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastlaneOnlineSolver *FastlaneOnlineSolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastlaneOnlineSolver.Contract.FastlaneOnlineSolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastlaneOnlineSolver *FastlaneOnlineSolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.Contract.FastlaneOnlineSolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastlaneOnlineSolver *FastlaneOnlineSolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.Contract.FastlaneOnlineSolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastlaneOnlineSolver *FastlaneOnlineSolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastlaneOnlineSolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastlaneOnlineSolver *FastlaneOnlineSolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastlaneOnlineSolver *FastlaneOnlineSolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.Contract.contract.Transact(opts, method, params...)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_FastlaneOnlineSolver *FastlaneOnlineSolverCaller) WETHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastlaneOnlineSolver.contract.Call(opts, &out, "WETH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_FastlaneOnlineSolver *FastlaneOnlineSolverSession) WETHADDRESS() (common.Address, error) {
	return _FastlaneOnlineSolver.Contract.WETHADDRESS(&_FastlaneOnlineSolver.CallOpts)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_FastlaneOnlineSolver *FastlaneOnlineSolverCallerSession) WETHADDRESS() (common.Address, error) {
	return _FastlaneOnlineSolver.Contract.WETHADDRESS(&_FastlaneOnlineSolver.CallOpts)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x024181a6.
//
// Solidity: function atlasSolverCall(address solverOpFrom, address executionEnvironment, address bidToken, uint256 bidAmount, bytes solverOpData, bytes ) payable returns()
func (_FastlaneOnlineSolver *FastlaneOnlineSolverTransactor) AtlasSolverCall(opts *bind.TransactOpts, solverOpFrom common.Address, executionEnvironment common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, arg5 []byte) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.contract.Transact(opts, "atlasSolverCall", solverOpFrom, executionEnvironment, bidToken, bidAmount, solverOpData, arg5)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x024181a6.
//
// Solidity: function atlasSolverCall(address solverOpFrom, address executionEnvironment, address bidToken, uint256 bidAmount, bytes solverOpData, bytes ) payable returns()
func (_FastlaneOnlineSolver *FastlaneOnlineSolverSession) AtlasSolverCall(solverOpFrom common.Address, executionEnvironment common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, arg5 []byte) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.Contract.AtlasSolverCall(&_FastlaneOnlineSolver.TransactOpts, solverOpFrom, executionEnvironment, bidToken, bidAmount, solverOpData, arg5)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x024181a6.
//
// Solidity: function atlasSolverCall(address solverOpFrom, address executionEnvironment, address bidToken, uint256 bidAmount, bytes solverOpData, bytes ) payable returns()
func (_FastlaneOnlineSolver *FastlaneOnlineSolverTransactorSession) AtlasSolverCall(solverOpFrom common.Address, executionEnvironment common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, arg5 []byte) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.Contract.AtlasSolverCall(&_FastlaneOnlineSolver.TransactOpts, solverOpFrom, executionEnvironment, bidToken, bidAmount, solverOpData, arg5)
}

// Execute is a paid mutator transaction binding the contract method 0xcf09e3dd.
//
// Solidity: function execute((uint8,address,address,address)[] swapPath, address tokenUserSells, address tokenUserBuys, address bidToken, uint256 bidAmount) payable returns()
func (_FastlaneOnlineSolver *FastlaneOnlineSolverTransactor) Execute(opts *bind.TransactOpts, swapPath []Swap, tokenUserSells common.Address, tokenUserBuys common.Address, bidToken common.Address, bidAmount *big.Int) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.contract.Transact(opts, "execute", swapPath, tokenUserSells, tokenUserBuys, bidToken, bidAmount)
}

// Execute is a paid mutator transaction binding the contract method 0xcf09e3dd.
//
// Solidity: function execute((uint8,address,address,address)[] swapPath, address tokenUserSells, address tokenUserBuys, address bidToken, uint256 bidAmount) payable returns()
func (_FastlaneOnlineSolver *FastlaneOnlineSolverSession) Execute(swapPath []Swap, tokenUserSells common.Address, tokenUserBuys common.Address, bidToken common.Address, bidAmount *big.Int) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.Contract.Execute(&_FastlaneOnlineSolver.TransactOpts, swapPath, tokenUserSells, tokenUserBuys, bidToken, bidAmount)
}

// Execute is a paid mutator transaction binding the contract method 0xcf09e3dd.
//
// Solidity: function execute((uint8,address,address,address)[] swapPath, address tokenUserSells, address tokenUserBuys, address bidToken, uint256 bidAmount) payable returns()
func (_FastlaneOnlineSolver *FastlaneOnlineSolverTransactorSession) Execute(swapPath []Swap, tokenUserSells common.Address, tokenUserBuys common.Address, bidToken common.Address, bidAmount *big.Int) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.Contract.Execute(&_FastlaneOnlineSolver.TransactOpts, swapPath, tokenUserSells, tokenUserBuys, bidToken, bidAmount)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_FastlaneOnlineSolver *FastlaneOnlineSolverTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_FastlaneOnlineSolver *FastlaneOnlineSolverSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.Contract.UniswapV3SwapCallback(&_FastlaneOnlineSolver.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_FastlaneOnlineSolver *FastlaneOnlineSolverTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _FastlaneOnlineSolver.Contract.UniswapV3SwapCallback(&_FastlaneOnlineSolver.TransactOpts, amount0Delta, amount1Delta, _data)
}
