// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testContracts

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

// TestContractsMetaData contains all meta data concerning the TestContracts contract.
var TestContractsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"Start\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TestContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use TestContractsMetaData.ABI instead.
var TestContractsABI = TestContractsMetaData.ABI

// TestContracts is an auto generated Go binding around an Ethereum contract.
type TestContracts struct {
	TestContractsCaller     // Read-only binding to the contract
	TestContractsTransactor // Write-only binding to the contract
	TestContractsFilterer   // Log filterer for contract events
}

// TestContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestContractsSession struct {
	Contract     *TestContracts    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestContractsCallerSession struct {
	Contract *TestContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestContractsTransactorSession struct {
	Contract     *TestContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestContractsRaw struct {
	Contract *TestContracts // Generic contract binding to access the raw methods on
}

// TestContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestContractsCallerRaw struct {
	Contract *TestContractsCaller // Generic read-only contract binding to access the raw methods on
}

// TestContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestContractsTransactorRaw struct {
	Contract *TestContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestContracts creates a new instance of TestContracts, bound to a specific deployed contract.
func NewTestContracts(address common.Address, backend bind.ContractBackend) (*TestContracts, error) {
	contract, err := bindTestContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestContracts{TestContractsCaller: TestContractsCaller{contract: contract}, TestContractsTransactor: TestContractsTransactor{contract: contract}, TestContractsFilterer: TestContractsFilterer{contract: contract}}, nil
}

// NewTestContractsCaller creates a new read-only instance of TestContracts, bound to a specific deployed contract.
func NewTestContractsCaller(address common.Address, caller bind.ContractCaller) (*TestContractsCaller, error) {
	contract, err := bindTestContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestContractsCaller{contract: contract}, nil
}

// NewTestContractsTransactor creates a new write-only instance of TestContracts, bound to a specific deployed contract.
func NewTestContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*TestContractsTransactor, error) {
	contract, err := bindTestContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestContractsTransactor{contract: contract}, nil
}

// NewTestContractsFilterer creates a new log filterer instance of TestContracts, bound to a specific deployed contract.
func NewTestContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*TestContractsFilterer, error) {
	contract, err := bindTestContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestContractsFilterer{contract: contract}, nil
}

// bindTestContracts binds a generic wrapper to an already deployed contract.
func bindTestContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContracts *TestContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestContracts.Contract.TestContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContracts *TestContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContracts.Contract.TestContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContracts *TestContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContracts.Contract.TestContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContracts *TestContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestContracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContracts *TestContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContracts *TestContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContracts.Contract.contract.Transact(opts, method, params...)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_TestContracts *TestContractsCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestContracts.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_TestContracts *TestContractsSession) Retrieve() (*big.Int, error) {
	return _TestContracts.Contract.Retrieve(&_TestContracts.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_TestContracts *TestContractsCallerSession) Retrieve() (*big.Int, error) {
	return _TestContracts.Contract.Retrieve(&_TestContracts.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns(uint256)
func (_TestContracts *TestContractsTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _TestContracts.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns(uint256)
func (_TestContracts *TestContractsSession) Store(num *big.Int) (*types.Transaction, error) {
	return _TestContracts.Contract.Store(&_TestContracts.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns(uint256)
func (_TestContracts *TestContractsTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _TestContracts.Contract.Store(&_TestContracts.TransactOpts, num)
}

// TestContractsStartIterator is returned from FilterStart and is used to iterate over the raw logs and unpacked data for Start events raised by the TestContracts contract.
type TestContractsStartIterator struct {
	Event *TestContractsStart // Event containing the contract specifics and raw log

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
func (it *TestContractsStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestContractsStart)
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
		it.Event = new(TestContractsStart)
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
func (it *TestContractsStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestContractsStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestContractsStart represents a Start event raised by the TestContracts contract.
type TestContractsStart struct {
	Who common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStart is a free log retrieval operation binding the contract event 0x84d0447ca38875fa61115673259a210915bc1dd53a3c112d6f0790f15956a965.
//
// Solidity: event Start(address who)
func (_TestContracts *TestContractsFilterer) FilterStart(opts *bind.FilterOpts) (*TestContractsStartIterator, error) {

	logs, sub, err := _TestContracts.contract.FilterLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return &TestContractsStartIterator{contract: _TestContracts.contract, event: "Start", logs: logs, sub: sub}, nil
}

// WatchStart is a free log subscription operation binding the contract event 0x84d0447ca38875fa61115673259a210915bc1dd53a3c112d6f0790f15956a965.
//
// Solidity: event Start(address who)
func (_TestContracts *TestContractsFilterer) WatchStart(opts *bind.WatchOpts, sink chan<- *TestContractsStart) (event.Subscription, error) {

	logs, sub, err := _TestContracts.contract.WatchLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestContractsStart)
				if err := _TestContracts.contract.UnpackLog(event, "Start", log); err != nil {
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

// ParseStart is a log parse operation binding the contract event 0x84d0447ca38875fa61115673259a210915bc1dd53a3c112d6f0790f15956a965.
//
// Solidity: event Start(address who)
func (_TestContracts *TestContractsFilterer) ParseStart(log types.Log) (*TestContractsStart, error) {
	event := new(TestContractsStart)
	if err := _TestContracts.contract.UnpackLog(event, "Start", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
