// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f9577e238fa9d545f759441d930618fa330f38eda57bc3a9e5acaa8ed0ae65d664736f6c634300060c0033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_IERC20 *IERC20Transactor) Allowance(opts *bind.TransactOpts, owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "allowance", owner, spender)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _IERC20.Contract.Allowance(&_IERC20.TransactOpts, owner, spender)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_IERC20 *IERC20TransactorSession) Allowance(owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _IERC20.Contract.Allowance(&_IERC20.TransactOpts, owner, spender)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_IERC20 *IERC20Transactor) BalanceOf(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "balanceOf", account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.TransactOpts, account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_IERC20 *IERC20TransactorSession) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.TransactOpts, account)
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_IERC20 *IERC20Transactor) TotalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "totalSupply")
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*types.Transaction, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.TransactOpts)
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_IERC20 *IERC20TransactorSession) TotalSupply() (*types.Transaction, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ILockProxyABI is the input ABI used to generate the binding from.
const ILockProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"assetHashMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"}],\"name\":\"bindAssetHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"targetProxyHash\",\"type\":\"bytes\"}],\"name\":\"bindProxyHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getBalanceFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerProxyContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"proxyHashMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"eccmpAddr\",\"type\":\"address\"}],\"name\":\"setManagerProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ILockProxyFuncSigs maps the 4-byte function signature to its string representation.
var ILockProxyFuncSigs = map[string]string{
	"4f7d9808": "assetHashMap(address,uint64)",
	"3348f63b": "bindAssetHash(address,uint64,bytes)",
	"379b98f6": "bindProxyHash(uint64,bytes)",
	"59c589a1": "getBalanceFor(address)",
	"84a6d055": "lock(address,uint64,bytes,uint256)",
	"d798f881": "managerProxyContract()",
	"9e5767aa": "proxyHashMap(uint64)",
	"af9980f0": "setManagerProxy(address)",
}

// ILockProxy is an auto generated Go binding around an Ethereum contract.
type ILockProxy struct {
	ILockProxyCaller     // Read-only binding to the contract
	ILockProxyTransactor // Write-only binding to the contract
	ILockProxyFilterer   // Log filterer for contract events
}

// ILockProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILockProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILockProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILockProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILockProxySession struct {
	Contract     *ILockProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILockProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILockProxyCallerSession struct {
	Contract *ILockProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ILockProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILockProxyTransactorSession struct {
	Contract     *ILockProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ILockProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILockProxyRaw struct {
	Contract *ILockProxy // Generic contract binding to access the raw methods on
}

// ILockProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILockProxyCallerRaw struct {
	Contract *ILockProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ILockProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILockProxyTransactorRaw struct {
	Contract *ILockProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILockProxy creates a new instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxy(address common.Address, backend bind.ContractBackend) (*ILockProxy, error) {
	contract, err := bindILockProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILockProxy{ILockProxyCaller: ILockProxyCaller{contract: contract}, ILockProxyTransactor: ILockProxyTransactor{contract: contract}, ILockProxyFilterer: ILockProxyFilterer{contract: contract}}, nil
}

// NewILockProxyCaller creates a new read-only instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxyCaller(address common.Address, caller bind.ContractCaller) (*ILockProxyCaller, error) {
	contract, err := bindILockProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILockProxyCaller{contract: contract}, nil
}

// NewILockProxyTransactor creates a new write-only instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ILockProxyTransactor, error) {
	contract, err := bindILockProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILockProxyTransactor{contract: contract}, nil
}

// NewILockProxyFilterer creates a new log filterer instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ILockProxyFilterer, error) {
	contract, err := bindILockProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILockProxyFilterer{contract: contract}, nil
}

// bindILockProxy binds a generic wrapper to an already deployed contract.
func bindILockProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILockProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILockProxy *ILockProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ILockProxy.Contract.ILockProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILockProxy *ILockProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILockProxy.Contract.ILockProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILockProxy *ILockProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILockProxy.Contract.ILockProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILockProxy *ILockProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ILockProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILockProxy *ILockProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILockProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILockProxy *ILockProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILockProxy.Contract.contract.Transact(opts, method, params...)
}

// AssetHashMap is a paid mutator transaction binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) returns(bytes)
func (_ILockProxy *ILockProxyTransactor) AssetHashMap(opts *bind.TransactOpts, arg0 common.Address, arg1 uint64) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "assetHashMap", arg0, arg1)
}

// AssetHashMap is a paid mutator transaction binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) returns(bytes)
func (_ILockProxy *ILockProxySession) AssetHashMap(arg0 common.Address, arg1 uint64) (*types.Transaction, error) {
	return _ILockProxy.Contract.AssetHashMap(&_ILockProxy.TransactOpts, arg0, arg1)
}

// AssetHashMap is a paid mutator transaction binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) returns(bytes)
func (_ILockProxy *ILockProxyTransactorSession) AssetHashMap(arg0 common.Address, arg1 uint64) (*types.Transaction, error) {
	return _ILockProxy.Contract.AssetHashMap(&_ILockProxy.TransactOpts, arg0, arg1)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_ILockProxy *ILockProxyTransactor) BindAssetHash(opts *bind.TransactOpts, fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "bindAssetHash", fromAssetHash, toChainId, toAssetHash)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_ILockProxy *ILockProxySession) BindAssetHash(fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindAssetHash(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAssetHash)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_ILockProxy *ILockProxyTransactorSession) BindAssetHash(fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindAssetHash(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAssetHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_ILockProxy *ILockProxyTransactor) BindProxyHash(opts *bind.TransactOpts, toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "bindProxyHash", toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_ILockProxy *ILockProxySession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindProxyHash(&_ILockProxy.TransactOpts, toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_ILockProxy *ILockProxyTransactorSession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindProxyHash(&_ILockProxy.TransactOpts, toChainId, targetProxyHash)
}

// GetBalanceFor is a paid mutator transaction binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address ) returns(uint256)
func (_ILockProxy *ILockProxyTransactor) GetBalanceFor(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "getBalanceFor", arg0)
}

// GetBalanceFor is a paid mutator transaction binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address ) returns(uint256)
func (_ILockProxy *ILockProxySession) GetBalanceFor(arg0 common.Address) (*types.Transaction, error) {
	return _ILockProxy.Contract.GetBalanceFor(&_ILockProxy.TransactOpts, arg0)
}

// GetBalanceFor is a paid mutator transaction binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address ) returns(uint256)
func (_ILockProxy *ILockProxyTransactorSession) GetBalanceFor(arg0 common.Address) (*types.Transaction, error) {
	return _ILockProxy.Contract.GetBalanceFor(&_ILockProxy.TransactOpts, arg0)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) returns(bool)
func (_ILockProxy *ILockProxyTransactor) Lock(opts *bind.TransactOpts, fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "lock", fromAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) returns(bool)
func (_ILockProxy *ILockProxySession) Lock(fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _ILockProxy.Contract.Lock(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) returns(bool)
func (_ILockProxy *ILockProxyTransactorSession) Lock(fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _ILockProxy.Contract.Lock(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAddress, amount)
}

// ManagerProxyContract is a paid mutator transaction binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() returns(address)
func (_ILockProxy *ILockProxyTransactor) ManagerProxyContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "managerProxyContract")
}

// ManagerProxyContract is a paid mutator transaction binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() returns(address)
func (_ILockProxy *ILockProxySession) ManagerProxyContract() (*types.Transaction, error) {
	return _ILockProxy.Contract.ManagerProxyContract(&_ILockProxy.TransactOpts)
}

// ManagerProxyContract is a paid mutator transaction binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() returns(address)
func (_ILockProxy *ILockProxyTransactorSession) ManagerProxyContract() (*types.Transaction, error) {
	return _ILockProxy.Contract.ManagerProxyContract(&_ILockProxy.TransactOpts)
}

// ProxyHashMap is a paid mutator transaction binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) returns(bytes)
func (_ILockProxy *ILockProxyTransactor) ProxyHashMap(opts *bind.TransactOpts, arg0 uint64) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "proxyHashMap", arg0)
}

// ProxyHashMap is a paid mutator transaction binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) returns(bytes)
func (_ILockProxy *ILockProxySession) ProxyHashMap(arg0 uint64) (*types.Transaction, error) {
	return _ILockProxy.Contract.ProxyHashMap(&_ILockProxy.TransactOpts, arg0)
}

// ProxyHashMap is a paid mutator transaction binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) returns(bytes)
func (_ILockProxy *ILockProxyTransactorSession) ProxyHashMap(arg0 uint64) (*types.Transaction, error) {
	return _ILockProxy.Contract.ProxyHashMap(&_ILockProxy.TransactOpts, arg0)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address eccmpAddr) returns()
func (_ILockProxy *ILockProxyTransactor) SetManagerProxy(opts *bind.TransactOpts, eccmpAddr common.Address) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "setManagerProxy", eccmpAddr)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address eccmpAddr) returns()
func (_ILockProxy *ILockProxySession) SetManagerProxy(eccmpAddr common.Address) (*types.Transaction, error) {
	return _ILockProxy.Contract.SetManagerProxy(&_ILockProxy.TransactOpts, eccmpAddr)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address eccmpAddr) returns()
func (_ILockProxy *ILockProxyTransactorSession) SetManagerProxy(eccmpAddr common.Address) (*types.Transaction, error) {
	return _ILockProxy.Contract.SetManagerProxy(&_ILockProxy.TransactOpts, eccmpAddr)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Ownable *OwnableTransactor) Owner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "owner")
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Ownable *OwnableSession) Owner() (*types.Transaction, error) {
	return _Ownable.Contract.Owner(&_Ownable.TransactOpts)
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Ownable *OwnableTransactorSession) Owner() (*types.Transaction, error) {
	return _Ownable.Contract.Owner(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = map[string]string{
	"5c975abb": "paused()",
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_Pausable *PausableTransactor) Paused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "paused")
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_Pausable *PausableSession) Paused() (*types.Transaction, error) {
	return _Pausable.Contract.Paused(&_Pausable.TransactOpts)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_Pausable *PausableTransactorSession) Paused() (*types.Transaction, error) {
	return _Pausable.Contract.Paused(&_Pausable.TransactOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

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
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
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
		it.Event = new(PausablePaused)
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
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
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
		it.Event = new(PausableUnpaused)
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
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PolyWrapperABI is the input ABI used to generate the binding from.
const PolyWrapperABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"net\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"PolyWrapperLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"txHash\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"efee\",\"type\":\"uint256\"}],\"name\":\"PolyWrapperSpeedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"extractFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockProxy\",\"outputs\":[{\"internalType\":\"contractILockProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lockProxy\",\"type\":\"address\"}],\"name\":\"setLockProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"txHash\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"speedUp\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PolyWrapperFuncSigs maps the 4-byte function signature to its string representation.
var PolyWrapperFuncSigs = map[string]string{
	"9a8a0592": "chainId()",
	"1745399d": "extractFee(address)",
	"c415b95c": "feeCollector()",
	"60de1a9b": "lock(address,uint64,bytes,uint256,uint256,uint256)",
	"9d4dc021": "lockProxy()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"a42dce80": "setFeeCollector(address)",
	"6f2b6ee6": "setLockProxy(address)",
	"d3ed7c76": "speedUp(address,bytes,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// PolyWrapperBin is the compiled bytecode used for deploying new contracts.
var PolyWrapperBin = "0x60806040523480156200001157600080fd5b5060405162001ac738038062001ac7833981810160405260408110156200003757600080fd5b50805160209091015160006200004c620000e7565b600080546001600160a01b0319166001600160a01b03831690811782556040519293509160008051602062001aa7833981519152908290a3506000805460ff60a01b191690556001805580620000d2576040805162461bcd60e51b8152602060048201526006602482015265085b1959d85b60d21b604482015290519081900360640190fd5b620000dd82620000eb565b60025550620001e9565b3390565b620000f5620000e7565b6000546001600160a01b0390811691161462000158576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b0381166200019f5760405162461bcd60e51b815260040180806020018281038252602681526020018062001a816026913960400191505060405180910390fd5b600080546040516001600160a01b038085169392169160008051602062001aa783398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b61188880620001f96000396000f3fe6080604052600436106100dd5760003560e01c80638da5cb5b1161007f578063a42dce8011610059578063a42dce80146102ef578063c415b95c14610322578063d3ed7c7614610337578063f2fde38b146103ef576100dd565b80638da5cb5b146102825780639a8a0592146102b35780639d4dc021146102da576100dd565b806360de1a9b116100bb57806360de1a9b146101555780636f2b6ee614610225578063715018a6146102585780638456cb591461026d576100dd565b80631745399d146100e25780633f4ba83a146101175780635c975abb1461012c575b600080fd5b3480156100ee57600080fd5b506101156004803603602081101561010557600080fd5b50356001600160a01b0316610422565b005b34801561012357600080fd5b50610115610544565b34801561013857600080fd5b506101416105a6565b604080519115158252519081900360200190f35b610115600480360360c081101561016b57600080fd5b6001600160a01b038235169167ffffffffffffffff602082013516918101906060810160408201356401000000008111156101a557600080fd5b8201836020820111156101b757600080fd5b803590602001918460018302840111640100000000831117156101d957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602081013590604001356105b6565b34801561023157600080fd5b506101156004803603602081101561024857600080fd5b50356001600160a01b0316610815565b34801561026457600080fd5b50610115610953565b34801561027957600080fd5b506101156109f5565b34801561028e57600080fd5b50610297610a55565b604080516001600160a01b039092168252519081900360200190f35b3480156102bf57600080fd5b506102c8610a64565b60408051918252519081900360200190f35b3480156102e657600080fd5b50610297610a6a565b3480156102fb57600080fd5b506101156004803603602081101561031257600080fd5b50356001600160a01b0316610a79565b34801561032e57600080fd5b50610297610b3e565b6101156004803603606081101561034d57600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561037857600080fd5b82018360208201111561038a57600080fd5b803590602001918460018302840111640100000000831117156103ac57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610b4d915050565b3480156103fb57600080fd5b506101156004803603602081101561041257600080fd5b50356001600160a01b0316610cae565b6003546001600160a01b03163314610471576040805162461bcd60e51b815260206004820152600d60248201526c10b332b2a1b7b63632b1ba37b960991b604482015290519081900360640190fd5b6001600160a01b0381166104b15760405133904780156108fc02916000818181858888f193505050501580156104ab573d6000803e3d6000fd5b50610541565b600354604080516370a0823160e01b81523060048201529051610541926001600160a01b0390811692908516916370a0823191602480820192602092909190829003018186803b15801561050457600080fd5b505afa158015610518573d6000803e3d6000fd5b505050506040513d602081101561052e57600080fd5b50516001600160a01b0384169190610da6565b50565b61054c610dfd565b6000546001600160a01b0390811691161461059c576040805162461bcd60e51b815260206004820181905260248201526000805160206117d3833981519152604482015290519081900360640190fd5b6105a4610e01565b565b600054600160a01b900460ff1690565b6002600154141561060e576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002600155600054600160a01b900460ff1615610665576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6002548567ffffffffffffffff161415801561068a575067ffffffffffffffff851615155b6106c8576040805162461bcd60e51b815260206004820152600a602482015269085d1bd0da185a5b925960b21b604482015290519081900360640190fd5b818311610713576040805162461bcd60e51b8152602060048201526014602482015273616d6f756e74206c657373207468616e2066656560601b604482015290519081900360640190fd5b61071d8684610ea9565b61073286868661072d8787610f1e565b610f67565b336001600160a01b0387167f2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e878761076a8888610f1e565b8787604051808667ffffffffffffffff16815260200180602001858152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b838110156107cb5781810151838201526020016107b3565b50505050905090810190601f1680156107f85780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a350506001805550505050565b61081d610dfd565b6000546001600160a01b0390811691161461086d576040805162461bcd60e51b815260206004820181905260248201526000805160206117d3833981519152604482015290519081900360640190fd5b6001600160a01b03811661088057600080fd5b600480546001600160a01b0319166001600160a01b03838116919091178083556040805163d798f88160e01b81529051600094929093169263d798f88192828101926020929190829003018186803b1580156108db57600080fd5b505afa1580156108ef573d6000803e3d6000fd5b505050506040513d602081101561090557600080fd5b50516001600160a01b03161415610541576040805162461bcd60e51b815260206004820152600d60248201526c6e6f74206c6f636b70726f787960981b604482015290519081900360640190fd5b61095b610dfd565b6000546001600160a01b039081169116146109ab576040805162461bcd60e51b815260206004820181905260248201526000805160206117d3833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6109fd610dfd565b6000546001600160a01b03908116911614610a4d576040805162461bcd60e51b815260206004820181905260248201526000805160206117d3833981519152604482015290519081900360640190fd5b6105a4611264565b6000546001600160a01b031690565b60025481565b6004546001600160a01b031681565b610a81610dfd565b6000546001600160a01b03908116911614610ad1576040805162461bcd60e51b815260206004820181905260248201526000805160206117d3833981519152604482015290519081900360640190fd5b6001600160a01b038116610b1c576040805162461bcd60e51b815260206004820152600d60248201526c656d747079206164647265737360981b604482015290519081900360640190fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6003546001600160a01b031681565b60026001541415610ba5576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002600155600054600160a01b900460ff1615610bfc576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610c068382610ea9565b336001600160a01b0316826040518082805190602001908083835b60208310610c405780518252601f199092019160209182019101610c21565b51815160209384036101000a60001901801990921691161790526040805192909401829003822088835293519395506001600160a01b038a1694507ff6579aef3e0d086d986c5d6972659f8a0d8602ef7945b054be1b88e088773ef69391829003019150a450506001805550565b610cb6610dfd565b6000546001600160a01b03908116911614610d06576040805162461bcd60e51b815260206004820181905260248201526000805160206117d3833981519152604482015290519081900360640190fd5b6001600160a01b038116610d4b5760405162461bcd60e51b81526004018080602001828103825260268152602001806117876026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610df89084906112f2565b505050565b3390565b600054600160a01b900460ff16610e56576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610e8c610dfd565b604080516001600160a01b039092168252519081900360200190a1565b6001600160a01b038216610f0557803414610f00576040805162461bcd60e51b815260206004820152601260248201527134b739bab33334b1b4b2b73a1032ba3432b960711b604482015290519081900360640190fd5b610f1a565b610f1a6001600160a01b0383163330846113a3565b5050565b6000610f6083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506113fd565b9392505050565b6001600160a01b0384166110d257600460009054906101000a90046001600160a01b03166001600160a01b03166384a6d05582868686866040518663ffffffff1660e01b815260040180856001600160a01b031681526020018467ffffffffffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561100f578181015183820152602001610ff7565b50505050905090810190601f16801561103c5780820380516001836020036101000a031916815260200191505b50955050505050506020604051808303818588803b15801561105d57600080fd5b505af1158015611071573d6000803e3d6000fd5b50505050506040513d602081101561108857600080fd5b50516110cd576040805162461bcd60e51b815260206004820152600f60248201526e1b1bd8dac8195d1a195c8819985a5b608a1b604482015290519081900360640190fd5b61125e565b6004546110ed906001600160a01b0386811691166000611494565b600454611107906001600160a01b03868116911683611494565b600460009054906101000a90046001600160a01b03166001600160a01b03166384a6d055858585856040518563ffffffff1660e01b815260040180856001600160a01b031681526020018467ffffffffffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156111a0578181015183820152602001611188565b50505050905090810190601f1680156111cd5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156111ef57600080fd5b505af1158015611203573d6000803e3d6000fd5b505050506040513d602081101561121957600080fd5b505161125e576040805162461bcd60e51b815260206004820152600f60248201526e1b1bd8dac8195c98cc8c0819985a5b608a1b604482015290519081900360640190fd5b50505050565b600054600160a01b900460ff16156112b6576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610e8c610dfd565b6060611347826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166115a79092919063ffffffff16565b805190915015610df85780806020019051602081101561136657600080fd5b5051610df85760405162461bcd60e51b815260040180806020018281038252602a8152602001806117f3602a913960400191505060405180910390fd5b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b17905261125e9085906112f2565b6000818484111561148c5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611451578181015183820152602001611439565b50505050905090810190601f16801561147e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b80158061151a575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b1580156114ec57600080fd5b505afa158015611500573d6000803e3d6000fd5b505050506040513d602081101561151657600080fd5b5051155b6115555760405162461bcd60e51b815260040180806020018281038252603681526020018061181d6036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052610df89084906112f2565b60606115b684846000856115be565b949350505050565b6060824710156115ff5760405162461bcd60e51b81526004018080602001828103825260268152602001806117ad6026913960400191505060405180910390fd5b6116088561171a565b611659576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b602083106116985780518252601f199092019160209182019101611679565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146116fa576040519150601f19603f3d011682016040523d82523d6000602084013e6116ff565b606091505b509150915061170f828286611720565b979650505050505050565b3b151590565b6060831561172f575081610f60565b82511561173f5782518084602001fd5b60405162461bcd60e51b815260206004820181815284516024840152845185939192839260440191908501908083836000831561145157818101518382015260200161143956fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565645361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a264697066735822122020315d5da31df18d3d37fd22a89e01a631002f3b4013ba57125ccdc57a40e9d764736f6c634300060c00334f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573738be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"

// DeployPolyWrapper deploys a new Ethereum contract, binding an instance of PolyWrapper to it.
func DeployPolyWrapper(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _chainId *big.Int) (common.Address, *types.Transaction, *PolyWrapper, error) {
	parsed, err := abi.JSON(strings.NewReader(PolyWrapperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PolyWrapperBin), backend, _owner, _chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PolyWrapper{PolyWrapperCaller: PolyWrapperCaller{contract: contract}, PolyWrapperTransactor: PolyWrapperTransactor{contract: contract}, PolyWrapperFilterer: PolyWrapperFilterer{contract: contract}}, nil
}

// PolyWrapper is an auto generated Go binding around an Ethereum contract.
type PolyWrapper struct {
	PolyWrapperCaller     // Read-only binding to the contract
	PolyWrapperTransactor // Write-only binding to the contract
	PolyWrapperFilterer   // Log filterer for contract events
}

// PolyWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolyWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolyWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolyWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolyWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolyWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolyWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolyWrapperSession struct {
	Contract     *PolyWrapper      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolyWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolyWrapperCallerSession struct {
	Contract *PolyWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PolyWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolyWrapperTransactorSession struct {
	Contract     *PolyWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PolyWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolyWrapperRaw struct {
	Contract *PolyWrapper // Generic contract binding to access the raw methods on
}

// PolyWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolyWrapperCallerRaw struct {
	Contract *PolyWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// PolyWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolyWrapperTransactorRaw struct {
	Contract *PolyWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolyWrapper creates a new instance of PolyWrapper, bound to a specific deployed contract.
func NewPolyWrapper(address common.Address, backend bind.ContractBackend) (*PolyWrapper, error) {
	contract, err := bindPolyWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolyWrapper{PolyWrapperCaller: PolyWrapperCaller{contract: contract}, PolyWrapperTransactor: PolyWrapperTransactor{contract: contract}, PolyWrapperFilterer: PolyWrapperFilterer{contract: contract}}, nil
}

// NewPolyWrapperCaller creates a new read-only instance of PolyWrapper, bound to a specific deployed contract.
func NewPolyWrapperCaller(address common.Address, caller bind.ContractCaller) (*PolyWrapperCaller, error) {
	contract, err := bindPolyWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperCaller{contract: contract}, nil
}

// NewPolyWrapperTransactor creates a new write-only instance of PolyWrapper, bound to a specific deployed contract.
func NewPolyWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*PolyWrapperTransactor, error) {
	contract, err := bindPolyWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperTransactor{contract: contract}, nil
}

// NewPolyWrapperFilterer creates a new log filterer instance of PolyWrapper, bound to a specific deployed contract.
func NewPolyWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*PolyWrapperFilterer, error) {
	contract, err := bindPolyWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperFilterer{contract: contract}, nil
}

// bindPolyWrapper binds a generic wrapper to an already deployed contract.
func bindPolyWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolyWrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolyWrapper *PolyWrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PolyWrapper.Contract.PolyWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolyWrapper *PolyWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapper.Contract.PolyWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolyWrapper *PolyWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolyWrapper.Contract.PolyWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolyWrapper *PolyWrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PolyWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolyWrapper *PolyWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolyWrapper *PolyWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolyWrapper.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a paid mutator transaction binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() returns(uint256)
func (_PolyWrapper *PolyWrapperTransactor) ChainId(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "chainId")
}

// ChainId is a paid mutator transaction binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() returns(uint256)
func (_PolyWrapper *PolyWrapperSession) ChainId() (*types.Transaction, error) {
	return _PolyWrapper.Contract.ChainId(&_PolyWrapper.TransactOpts)
}

// ChainId is a paid mutator transaction binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() returns(uint256)
func (_PolyWrapper *PolyWrapperTransactorSession) ChainId() (*types.Transaction, error) {
	return _PolyWrapper.Contract.ChainId(&_PolyWrapper.TransactOpts)
}

// ExtractFee is a paid mutator transaction binding the contract method 0x1745399d.
//
// Solidity: function extractFee(address token) returns()
func (_PolyWrapper *PolyWrapperTransactor) ExtractFee(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "extractFee", token)
}

// ExtractFee is a paid mutator transaction binding the contract method 0x1745399d.
//
// Solidity: function extractFee(address token) returns()
func (_PolyWrapper *PolyWrapperSession) ExtractFee(token common.Address) (*types.Transaction, error) {
	return _PolyWrapper.Contract.ExtractFee(&_PolyWrapper.TransactOpts, token)
}

// ExtractFee is a paid mutator transaction binding the contract method 0x1745399d.
//
// Solidity: function extractFee(address token) returns()
func (_PolyWrapper *PolyWrapperTransactorSession) ExtractFee(token common.Address) (*types.Transaction, error) {
	return _PolyWrapper.Contract.ExtractFee(&_PolyWrapper.TransactOpts, token)
}

// FeeCollector is a paid mutator transaction binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() returns(address)
func (_PolyWrapper *PolyWrapperTransactor) FeeCollector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "feeCollector")
}

// FeeCollector is a paid mutator transaction binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() returns(address)
func (_PolyWrapper *PolyWrapperSession) FeeCollector() (*types.Transaction, error) {
	return _PolyWrapper.Contract.FeeCollector(&_PolyWrapper.TransactOpts)
}

// FeeCollector is a paid mutator transaction binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() returns(address)
func (_PolyWrapper *PolyWrapperTransactorSession) FeeCollector() (*types.Transaction, error) {
	return _PolyWrapper.Contract.FeeCollector(&_PolyWrapper.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) returns()
func (_PolyWrapper *PolyWrapperTransactor) Lock(opts *bind.TransactOpts, fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "lock", fromAsset, toChainId, toAddress, amount, fee, id)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) returns()
func (_PolyWrapper *PolyWrapperSession) Lock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _PolyWrapper.Contract.Lock(&_PolyWrapper.TransactOpts, fromAsset, toChainId, toAddress, amount, fee, id)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) returns()
func (_PolyWrapper *PolyWrapperTransactorSession) Lock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _PolyWrapper.Contract.Lock(&_PolyWrapper.TransactOpts, fromAsset, toChainId, toAddress, amount, fee, id)
}

// LockProxy is a paid mutator transaction binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() returns(address)
func (_PolyWrapper *PolyWrapperTransactor) LockProxy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "lockProxy")
}

// LockProxy is a paid mutator transaction binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() returns(address)
func (_PolyWrapper *PolyWrapperSession) LockProxy() (*types.Transaction, error) {
	return _PolyWrapper.Contract.LockProxy(&_PolyWrapper.TransactOpts)
}

// LockProxy is a paid mutator transaction binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() returns(address)
func (_PolyWrapper *PolyWrapperTransactorSession) LockProxy() (*types.Transaction, error) {
	return _PolyWrapper.Contract.LockProxy(&_PolyWrapper.TransactOpts)
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_PolyWrapper *PolyWrapperTransactor) Owner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "owner")
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_PolyWrapper *PolyWrapperSession) Owner() (*types.Transaction, error) {
	return _PolyWrapper.Contract.Owner(&_PolyWrapper.TransactOpts)
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_PolyWrapper *PolyWrapperTransactorSession) Owner() (*types.Transaction, error) {
	return _PolyWrapper.Contract.Owner(&_PolyWrapper.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolyWrapper *PolyWrapperTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolyWrapper *PolyWrapperSession) Pause() (*types.Transaction, error) {
	return _PolyWrapper.Contract.Pause(&_PolyWrapper.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolyWrapper *PolyWrapperTransactorSession) Pause() (*types.Transaction, error) {
	return _PolyWrapper.Contract.Pause(&_PolyWrapper.TransactOpts)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_PolyWrapper *PolyWrapperTransactor) Paused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "paused")
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_PolyWrapper *PolyWrapperSession) Paused() (*types.Transaction, error) {
	return _PolyWrapper.Contract.Paused(&_PolyWrapper.TransactOpts)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_PolyWrapper *PolyWrapperTransactorSession) Paused() (*types.Transaction, error) {
	return _PolyWrapper.Contract.Paused(&_PolyWrapper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolyWrapper *PolyWrapperTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolyWrapper *PolyWrapperSession) RenounceOwnership() (*types.Transaction, error) {
	return _PolyWrapper.Contract.RenounceOwnership(&_PolyWrapper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolyWrapper *PolyWrapperTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PolyWrapper.Contract.RenounceOwnership(&_PolyWrapper.TransactOpts)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address collector) returns()
func (_PolyWrapper *PolyWrapperTransactor) SetFeeCollector(opts *bind.TransactOpts, collector common.Address) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "setFeeCollector", collector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address collector) returns()
func (_PolyWrapper *PolyWrapperSession) SetFeeCollector(collector common.Address) (*types.Transaction, error) {
	return _PolyWrapper.Contract.SetFeeCollector(&_PolyWrapper.TransactOpts, collector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address collector) returns()
func (_PolyWrapper *PolyWrapperTransactorSession) SetFeeCollector(collector common.Address) (*types.Transaction, error) {
	return _PolyWrapper.Contract.SetFeeCollector(&_PolyWrapper.TransactOpts, collector)
}

// SetLockProxy is a paid mutator transaction binding the contract method 0x6f2b6ee6.
//
// Solidity: function setLockProxy(address _lockProxy) returns()
func (_PolyWrapper *PolyWrapperTransactor) SetLockProxy(opts *bind.TransactOpts, _lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "setLockProxy", _lockProxy)
}

// SetLockProxy is a paid mutator transaction binding the contract method 0x6f2b6ee6.
//
// Solidity: function setLockProxy(address _lockProxy) returns()
func (_PolyWrapper *PolyWrapperSession) SetLockProxy(_lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapper.Contract.SetLockProxy(&_PolyWrapper.TransactOpts, _lockProxy)
}

// SetLockProxy is a paid mutator transaction binding the contract method 0x6f2b6ee6.
//
// Solidity: function setLockProxy(address _lockProxy) returns()
func (_PolyWrapper *PolyWrapperTransactorSession) SetLockProxy(_lockProxy common.Address) (*types.Transaction, error) {
	return _PolyWrapper.Contract.SetLockProxy(&_PolyWrapper.TransactOpts, _lockProxy)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) returns()
func (_PolyWrapper *PolyWrapperTransactor) SpeedUp(opts *bind.TransactOpts, fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "speedUp", fromAsset, txHash, fee)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) returns()
func (_PolyWrapper *PolyWrapperSession) SpeedUp(fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _PolyWrapper.Contract.SpeedUp(&_PolyWrapper.TransactOpts, fromAsset, txHash, fee)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) returns()
func (_PolyWrapper *PolyWrapperTransactorSession) SpeedUp(fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _PolyWrapper.Contract.SpeedUp(&_PolyWrapper.TransactOpts, fromAsset, txHash, fee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolyWrapper *PolyWrapperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolyWrapper *PolyWrapperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolyWrapper.Contract.TransferOwnership(&_PolyWrapper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolyWrapper *PolyWrapperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolyWrapper.Contract.TransferOwnership(&_PolyWrapper.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolyWrapper *PolyWrapperTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolyWrapper.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolyWrapper *PolyWrapperSession) Unpause() (*types.Transaction, error) {
	return _PolyWrapper.Contract.Unpause(&_PolyWrapper.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolyWrapper *PolyWrapperTransactorSession) Unpause() (*types.Transaction, error) {
	return _PolyWrapper.Contract.Unpause(&_PolyWrapper.TransactOpts)
}

// PolyWrapperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PolyWrapper contract.
type PolyWrapperOwnershipTransferredIterator struct {
	Event *PolyWrapperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolyWrapperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolyWrapperOwnershipTransferred)
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
		it.Event = new(PolyWrapperOwnershipTransferred)
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
func (it *PolyWrapperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolyWrapperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolyWrapperOwnershipTransferred represents a OwnershipTransferred event raised by the PolyWrapper contract.
type PolyWrapperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PolyWrapper *PolyWrapperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolyWrapperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolyWrapper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperOwnershipTransferredIterator{contract: _PolyWrapper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PolyWrapper *PolyWrapperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolyWrapperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolyWrapper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolyWrapperOwnershipTransferred)
				if err := _PolyWrapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PolyWrapper *PolyWrapperFilterer) ParseOwnershipTransferred(log types.Log) (*PolyWrapperOwnershipTransferred, error) {
	event := new(PolyWrapperOwnershipTransferred)
	if err := _PolyWrapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PolyWrapperPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PolyWrapper contract.
type PolyWrapperPausedIterator struct {
	Event *PolyWrapperPaused // Event containing the contract specifics and raw log

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
func (it *PolyWrapperPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolyWrapperPaused)
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
		it.Event = new(PolyWrapperPaused)
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
func (it *PolyWrapperPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolyWrapperPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolyWrapperPaused represents a Paused event raised by the PolyWrapper contract.
type PolyWrapperPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolyWrapper *PolyWrapperFilterer) FilterPaused(opts *bind.FilterOpts) (*PolyWrapperPausedIterator, error) {

	logs, sub, err := _PolyWrapper.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PolyWrapperPausedIterator{contract: _PolyWrapper.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolyWrapper *PolyWrapperFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PolyWrapperPaused) (event.Subscription, error) {

	logs, sub, err := _PolyWrapper.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolyWrapperPaused)
				if err := _PolyWrapper.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolyWrapper *PolyWrapperFilterer) ParsePaused(log types.Log) (*PolyWrapperPaused, error) {
	event := new(PolyWrapperPaused)
	if err := _PolyWrapper.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PolyWrapperPolyWrapperLockIterator is returned from FilterPolyWrapperLock and is used to iterate over the raw logs and unpacked data for PolyWrapperLock events raised by the PolyWrapper contract.
type PolyWrapperPolyWrapperLockIterator struct {
	Event *PolyWrapperPolyWrapperLock // Event containing the contract specifics and raw log

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
func (it *PolyWrapperPolyWrapperLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolyWrapperPolyWrapperLock)
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
		it.Event = new(PolyWrapperPolyWrapperLock)
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
func (it *PolyWrapperPolyWrapperLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolyWrapperPolyWrapperLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolyWrapperPolyWrapperLock represents a PolyWrapperLock event raised by the PolyWrapper contract.
type PolyWrapperPolyWrapperLock struct {
	FromAsset common.Address
	Sender    common.Address
	ToChainId uint64
	ToAddress []byte
	Net       *big.Int
	Fee       *big.Int
	Id        *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPolyWrapperLock is a free log retrieval operation binding the contract event 0x2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e.
//
// Solidity: event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint256 net, uint256 fee, uint256 id)
func (_PolyWrapper *PolyWrapperFilterer) FilterPolyWrapperLock(opts *bind.FilterOpts, fromAsset []common.Address, sender []common.Address) (*PolyWrapperPolyWrapperLockIterator, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolyWrapper.contract.FilterLogs(opts, "PolyWrapperLock", fromAssetRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperPolyWrapperLockIterator{contract: _PolyWrapper.contract, event: "PolyWrapperLock", logs: logs, sub: sub}, nil
}

// WatchPolyWrapperLock is a free log subscription operation binding the contract event 0x2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e.
//
// Solidity: event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint256 net, uint256 fee, uint256 id)
func (_PolyWrapper *PolyWrapperFilterer) WatchPolyWrapperLock(opts *bind.WatchOpts, sink chan<- *PolyWrapperPolyWrapperLock, fromAsset []common.Address, sender []common.Address) (event.Subscription, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolyWrapper.contract.WatchLogs(opts, "PolyWrapperLock", fromAssetRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolyWrapperPolyWrapperLock)
				if err := _PolyWrapper.contract.UnpackLog(event, "PolyWrapperLock", log); err != nil {
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

// ParsePolyWrapperLock is a log parse operation binding the contract event 0x2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e.
//
// Solidity: event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint256 net, uint256 fee, uint256 id)
func (_PolyWrapper *PolyWrapperFilterer) ParsePolyWrapperLock(log types.Log) (*PolyWrapperPolyWrapperLock, error) {
	event := new(PolyWrapperPolyWrapperLock)
	if err := _PolyWrapper.contract.UnpackLog(event, "PolyWrapperLock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PolyWrapperPolyWrapperSpeedUpIterator is returned from FilterPolyWrapperSpeedUp and is used to iterate over the raw logs and unpacked data for PolyWrapperSpeedUp events raised by the PolyWrapper contract.
type PolyWrapperPolyWrapperSpeedUpIterator struct {
	Event *PolyWrapperPolyWrapperSpeedUp // Event containing the contract specifics and raw log

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
func (it *PolyWrapperPolyWrapperSpeedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolyWrapperPolyWrapperSpeedUp)
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
		it.Event = new(PolyWrapperPolyWrapperSpeedUp)
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
func (it *PolyWrapperPolyWrapperSpeedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolyWrapperPolyWrapperSpeedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolyWrapperPolyWrapperSpeedUp represents a PolyWrapperSpeedUp event raised by the PolyWrapper contract.
type PolyWrapperPolyWrapperSpeedUp struct {
	FromAsset common.Address
	TxHash    common.Hash
	Sender    common.Address
	Efee      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPolyWrapperSpeedUp is a free log retrieval operation binding the contract event 0xf6579aef3e0d086d986c5d6972659f8a0d8602ef7945b054be1b88e088773ef6.
//
// Solidity: event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint256 efee)
func (_PolyWrapper *PolyWrapperFilterer) FilterPolyWrapperSpeedUp(opts *bind.FilterOpts, fromAsset []common.Address, txHash [][]byte, sender []common.Address) (*PolyWrapperPolyWrapperSpeedUpIterator, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolyWrapper.contract.FilterLogs(opts, "PolyWrapperSpeedUp", fromAssetRule, txHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolyWrapperPolyWrapperSpeedUpIterator{contract: _PolyWrapper.contract, event: "PolyWrapperSpeedUp", logs: logs, sub: sub}, nil
}

// WatchPolyWrapperSpeedUp is a free log subscription operation binding the contract event 0xf6579aef3e0d086d986c5d6972659f8a0d8602ef7945b054be1b88e088773ef6.
//
// Solidity: event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint256 efee)
func (_PolyWrapper *PolyWrapperFilterer) WatchPolyWrapperSpeedUp(opts *bind.WatchOpts, sink chan<- *PolyWrapperPolyWrapperSpeedUp, fromAsset []common.Address, txHash [][]byte, sender []common.Address) (event.Subscription, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolyWrapper.contract.WatchLogs(opts, "PolyWrapperSpeedUp", fromAssetRule, txHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolyWrapperPolyWrapperSpeedUp)
				if err := _PolyWrapper.contract.UnpackLog(event, "PolyWrapperSpeedUp", log); err != nil {
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

// ParsePolyWrapperSpeedUp is a log parse operation binding the contract event 0xf6579aef3e0d086d986c5d6972659f8a0d8602ef7945b054be1b88e088773ef6.
//
// Solidity: event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint256 efee)
func (_PolyWrapper *PolyWrapperFilterer) ParsePolyWrapperSpeedUp(log types.Log) (*PolyWrapperPolyWrapperSpeedUp, error) {
	event := new(PolyWrapperPolyWrapperSpeedUp)
	if err := _PolyWrapper.contract.UnpackLog(event, "PolyWrapperSpeedUp", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PolyWrapperUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PolyWrapper contract.
type PolyWrapperUnpausedIterator struct {
	Event *PolyWrapperUnpaused // Event containing the contract specifics and raw log

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
func (it *PolyWrapperUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolyWrapperUnpaused)
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
		it.Event = new(PolyWrapperUnpaused)
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
func (it *PolyWrapperUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolyWrapperUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolyWrapperUnpaused represents a Unpaused event raised by the PolyWrapper contract.
type PolyWrapperUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolyWrapper *PolyWrapperFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PolyWrapperUnpausedIterator, error) {

	logs, sub, err := _PolyWrapper.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PolyWrapperUnpausedIterator{contract: _PolyWrapper.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolyWrapper *PolyWrapperFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PolyWrapperUnpaused) (event.Subscription, error) {

	logs, sub, err := _PolyWrapper.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolyWrapperUnpaused)
				if err := _PolyWrapper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolyWrapper *PolyWrapperFilterer) ParseUnpaused(log types.Log) (*PolyWrapperUnpaused, error) {
	event := new(PolyWrapperUnpaused)
	if err := _PolyWrapper.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[]"

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122097bcb93b0c78fdb1e91887501df90a83953936c9a96569922498bad8115ae83d64736f6c634300060c0033"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122030f033b4b702e7b46074bc0078ed329556d628c3f9a750ca05a6faf7079b0d6364736f6c634300060c0033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

