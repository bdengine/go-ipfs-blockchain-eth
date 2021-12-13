// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipfs

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

// IPFSPeer is an auto generated low-level Go binding around an user-defined struct.
type IPFSPeer struct {
	PeerId      string
	AddressList []string
	Valid       bool
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"Success\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"SetOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"SetPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"SetToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"addFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"AddressList\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structIPFS.Peer\",\"name\":\"_peer\",\"type\":\"tuple\"}],\"name\":\"addPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addrPeerMap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"PeerId\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"fileMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getPeerList\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"AddressList\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structIPFS.Peer[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"head\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"peerHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"peerList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peerNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"pidAddrMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"removeFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"removePeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"addressList\",\"type\":\"string[]\"}],\"name\":\"updateAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// AddrPeerMap is a free data retrieval call binding the contract method 0x9d6c39a2.
//
// Solidity: function addrPeerMap(address ) view returns(string PeerId, bool valid)
func (_Contracts *ContractsCaller) AddrPeerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	PeerId string
	Valid  bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "addrPeerMap", arg0)

	outstruct := new(struct {
		PeerId string
		Valid  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PeerId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// AddrPeerMap is a free data retrieval call binding the contract method 0x9d6c39a2.
//
// Solidity: function addrPeerMap(address ) view returns(string PeerId, bool valid)
func (_Contracts *ContractsSession) AddrPeerMap(arg0 common.Address) (struct {
	PeerId string
	Valid  bool
}, error) {
	return _Contracts.Contract.AddrPeerMap(&_Contracts.CallOpts, arg0)
}

// AddrPeerMap is a free data retrieval call binding the contract method 0x9d6c39a2.
//
// Solidity: function addrPeerMap(address ) view returns(string PeerId, bool valid)
func (_Contracts *ContractsCallerSession) AddrPeerMap(arg0 common.Address) (struct {
	PeerId string
	Valid  bool
}, error) {
	return _Contracts.Contract.AddrPeerMap(&_Contracts.CallOpts, arg0)
}

// FileMap is a free data retrieval call binding the contract method 0x167e3841.
//
// Solidity: function fileMap(string ) view returns(address owner, uint256 size, uint256 expireBlock)
func (_Contracts *ContractsCaller) FileMap(opts *bind.CallOpts, arg0 string) (struct {
	Owner       common.Address
	Size        *big.Int
	ExpireBlock *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "fileMap", arg0)

	outstruct := new(struct {
		Owner       common.Address
		Size        *big.Int
		ExpireBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Size = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExpireBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FileMap is a free data retrieval call binding the contract method 0x167e3841.
//
// Solidity: function fileMap(string ) view returns(address owner, uint256 size, uint256 expireBlock)
func (_Contracts *ContractsSession) FileMap(arg0 string) (struct {
	Owner       common.Address
	Size        *big.Int
	ExpireBlock *big.Int
}, error) {
	return _Contracts.Contract.FileMap(&_Contracts.CallOpts, arg0)
}

// FileMap is a free data retrieval call binding the contract method 0x167e3841.
//
// Solidity: function fileMap(string ) view returns(address owner, uint256 size, uint256 expireBlock)
func (_Contracts *ContractsCallerSession) FileMap(arg0 string) (struct {
	Owner       common.Address
	Size        *big.Int
	ExpireBlock *big.Int
}, error) {
	return _Contracts.Contract.FileMap(&_Contracts.CallOpts, arg0)
}

// GetPeerList is a free data retrieval call binding the contract method 0x21d47c2e.
//
// Solidity: function getPeerList(uint256 num) view returns((string,string[],bool)[])
func (_Contracts *ContractsCaller) GetPeerList(opts *bind.CallOpts, num *big.Int) ([]IPFSPeer, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPeerList", num)

	if err != nil {
		return *new([]IPFSPeer), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPFSPeer)).(*[]IPFSPeer)

	return out0, err

}

// GetPeerList is a free data retrieval call binding the contract method 0x21d47c2e.
//
// Solidity: function getPeerList(uint256 num) view returns((string,string[],bool)[])
func (_Contracts *ContractsSession) GetPeerList(num *big.Int) ([]IPFSPeer, error) {
	return _Contracts.Contract.GetPeerList(&_Contracts.CallOpts, num)
}

// GetPeerList is a free data retrieval call binding the contract method 0x21d47c2e.
//
// Solidity: function getPeerList(uint256 num) view returns((string,string[],bool)[])
func (_Contracts *ContractsCallerSession) GetPeerList(num *big.Int) ([]IPFSPeer, error) {
	return _Contracts.Contract.GetPeerList(&_Contracts.CallOpts, num)
}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_Contracts *ContractsCaller) Head(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "head")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_Contracts *ContractsSession) Head() (common.Address, error) {
	return _Contracts.Contract.Head(&_Contracts.CallOpts)
}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_Contracts *ContractsCallerSession) Head() (common.Address, error) {
	return _Contracts.Contract.Head(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// PeerList is a free data retrieval call binding the contract method 0xe96e501a.
//
// Solidity: function peerList(address ) view returns(address)
func (_Contracts *ContractsCaller) PeerList(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "peerList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PeerList is a free data retrieval call binding the contract method 0xe96e501a.
//
// Solidity: function peerList(address ) view returns(address)
func (_Contracts *ContractsSession) PeerList(arg0 common.Address) (common.Address, error) {
	return _Contracts.Contract.PeerList(&_Contracts.CallOpts, arg0)
}

// PeerList is a free data retrieval call binding the contract method 0xe96e501a.
//
// Solidity: function peerList(address ) view returns(address)
func (_Contracts *ContractsCallerSession) PeerList(arg0 common.Address) (common.Address, error) {
	return _Contracts.Contract.PeerList(&_Contracts.CallOpts, arg0)
}

// PeerNum is a free data retrieval call binding the contract method 0xf9c20442.
//
// Solidity: function peerNum() view returns(uint256)
func (_Contracts *ContractsCaller) PeerNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "peerNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeerNum is a free data retrieval call binding the contract method 0xf9c20442.
//
// Solidity: function peerNum() view returns(uint256)
func (_Contracts *ContractsSession) PeerNum() (*big.Int, error) {
	return _Contracts.Contract.PeerNum(&_Contracts.CallOpts)
}

// PeerNum is a free data retrieval call binding the contract method 0xf9c20442.
//
// Solidity: function peerNum() view returns(uint256)
func (_Contracts *ContractsCallerSession) PeerNum() (*big.Int, error) {
	return _Contracts.Contract.PeerNum(&_Contracts.CallOpts)
}

// PidAddrMap is a free data retrieval call binding the contract method 0xad284b46.
//
// Solidity: function pidAddrMap(string ) view returns(address)
func (_Contracts *ContractsCaller) PidAddrMap(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "pidAddrMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PidAddrMap is a free data retrieval call binding the contract method 0xad284b46.
//
// Solidity: function pidAddrMap(string ) view returns(address)
func (_Contracts *ContractsSession) PidAddrMap(arg0 string) (common.Address, error) {
	return _Contracts.Contract.PidAddrMap(&_Contracts.CallOpts, arg0)
}

// PidAddrMap is a free data retrieval call binding the contract method 0xad284b46.
//
// Solidity: function pidAddrMap(string ) view returns(address)
func (_Contracts *ContractsCallerSession) PidAddrMap(arg0 string) (common.Address, error) {
	return _Contracts.Contract.PidAddrMap(&_Contracts.CallOpts, arg0)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Contracts *ContractsCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Contracts *ContractsSession) Price() (*big.Int, error) {
	return _Contracts.Contract.Price(&_Contracts.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Contracts *ContractsCallerSession) Price() (*big.Int, error) {
	return _Contracts.Contract.Price(&_Contracts.CallOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x167d3e9c.
//
// Solidity: function SetOwner(address _to) returns()
func (_Contracts *ContractsTransactor) SetOwner(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetOwner", _to)
}

// SetOwner is a paid mutator transaction binding the contract method 0x167d3e9c.
//
// Solidity: function SetOwner(address _to) returns()
func (_Contracts *ContractsSession) SetOwner(_to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetOwner(&_Contracts.TransactOpts, _to)
}

// SetOwner is a paid mutator transaction binding the contract method 0x167d3e9c.
//
// Solidity: function SetOwner(address _to) returns()
func (_Contracts *ContractsTransactorSession) SetOwner(_to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetOwner(&_Contracts.TransactOpts, _to)
}

// SetPrice is a paid mutator transaction binding the contract method 0x4f5539c0.
//
// Solidity: function SetPrice(uint256 _price) returns()
func (_Contracts *ContractsTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x4f5539c0.
//
// Solidity: function SetPrice(uint256 _price) returns()
func (_Contracts *ContractsSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetPrice(&_Contracts.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x4f5539c0.
//
// Solidity: function SetPrice(uint256 _price) returns()
func (_Contracts *ContractsTransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetPrice(&_Contracts.TransactOpts, _price)
}

// SetToken is a paid mutator transaction binding the contract method 0xefc1fd16.
//
// Solidity: function SetToken(address _token) returns()
func (_Contracts *ContractsTransactor) SetToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetToken", _token)
}

// SetToken is a paid mutator transaction binding the contract method 0xefc1fd16.
//
// Solidity: function SetToken(address _token) returns()
func (_Contracts *ContractsSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetToken(&_Contracts.TransactOpts, _token)
}

// SetToken is a paid mutator transaction binding the contract method 0xefc1fd16.
//
// Solidity: function SetToken(address _token) returns()
func (_Contracts *ContractsTransactorSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetToken(&_Contracts.TransactOpts, _token)
}

// AddFile is a paid mutator transaction binding the contract method 0x7f7847fe.
//
// Solidity: function addFile(string uid, string cid, uint256 size, uint256 wad) payable returns()
func (_Contracts *ContractsTransactor) AddFile(opts *bind.TransactOpts, uid string, cid string, size *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addFile", uid, cid, size, wad)
}

// AddFile is a paid mutator transaction binding the contract method 0x7f7847fe.
//
// Solidity: function addFile(string uid, string cid, uint256 size, uint256 wad) payable returns()
func (_Contracts *ContractsSession) AddFile(uid string, cid string, size *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddFile(&_Contracts.TransactOpts, uid, cid, size, wad)
}

// AddFile is a paid mutator transaction binding the contract method 0x7f7847fe.
//
// Solidity: function addFile(string uid, string cid, uint256 size, uint256 wad) payable returns()
func (_Contracts *ContractsTransactorSession) AddFile(uid string, cid string, size *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddFile(&_Contracts.TransactOpts, uid, cid, size, wad)
}

// AddPeer is a paid mutator transaction binding the contract method 0x6ff31945.
//
// Solidity: function addPeer(string uid, (string,string[],bool) _peer) returns()
func (_Contracts *ContractsTransactor) AddPeer(opts *bind.TransactOpts, uid string, _peer IPFSPeer) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addPeer", uid, _peer)
}

// AddPeer is a paid mutator transaction binding the contract method 0x6ff31945.
//
// Solidity: function addPeer(string uid, (string,string[],bool) _peer) returns()
func (_Contracts *ContractsSession) AddPeer(uid string, _peer IPFSPeer) (*types.Transaction, error) {
	return _Contracts.Contract.AddPeer(&_Contracts.TransactOpts, uid, _peer)
}

// AddPeer is a paid mutator transaction binding the contract method 0x6ff31945.
//
// Solidity: function addPeer(string uid, (string,string[],bool) _peer) returns()
func (_Contracts *ContractsTransactorSession) AddPeer(uid string, _peer IPFSPeer) (*types.Transaction, error) {
	return _Contracts.Contract.AddPeer(&_Contracts.TransactOpts, uid, _peer)
}

// PeerHeartbeat is a paid mutator transaction binding the contract method 0xf71e3e11.
//
// Solidity: function peerHeartbeat(string uid) returns()
func (_Contracts *ContractsTransactor) PeerHeartbeat(opts *bind.TransactOpts, uid string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "peerHeartbeat", uid)
}

// PeerHeartbeat is a paid mutator transaction binding the contract method 0xf71e3e11.
//
// Solidity: function peerHeartbeat(string uid) returns()
func (_Contracts *ContractsSession) PeerHeartbeat(uid string) (*types.Transaction, error) {
	return _Contracts.Contract.PeerHeartbeat(&_Contracts.TransactOpts, uid)
}

// PeerHeartbeat is a paid mutator transaction binding the contract method 0xf71e3e11.
//
// Solidity: function peerHeartbeat(string uid) returns()
func (_Contracts *ContractsTransactorSession) PeerHeartbeat(uid string) (*types.Transaction, error) {
	return _Contracts.Contract.PeerHeartbeat(&_Contracts.TransactOpts, uid)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x99f553dc.
//
// Solidity: function removeFile(string uid, string cid) payable returns()
func (_Contracts *ContractsTransactor) RemoveFile(opts *bind.TransactOpts, uid string, cid string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeFile", uid, cid)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x99f553dc.
//
// Solidity: function removeFile(string uid, string cid) payable returns()
func (_Contracts *ContractsSession) RemoveFile(uid string, cid string) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveFile(&_Contracts.TransactOpts, uid, cid)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x99f553dc.
//
// Solidity: function removeFile(string uid, string cid) payable returns()
func (_Contracts *ContractsTransactorSession) RemoveFile(uid string, cid string) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveFile(&_Contracts.TransactOpts, uid, cid)
}

// RemovePeer is a paid mutator transaction binding the contract method 0xf683d96c.
//
// Solidity: function removePeer(string uid) returns()
func (_Contracts *ContractsTransactor) RemovePeer(opts *bind.TransactOpts, uid string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removePeer", uid)
}

// RemovePeer is a paid mutator transaction binding the contract method 0xf683d96c.
//
// Solidity: function removePeer(string uid) returns()
func (_Contracts *ContractsSession) RemovePeer(uid string) (*types.Transaction, error) {
	return _Contracts.Contract.RemovePeer(&_Contracts.TransactOpts, uid)
}

// RemovePeer is a paid mutator transaction binding the contract method 0xf683d96c.
//
// Solidity: function removePeer(string uid) returns()
func (_Contracts *ContractsTransactorSession) RemovePeer(uid string) (*types.Transaction, error) {
	return _Contracts.Contract.RemovePeer(&_Contracts.TransactOpts, uid)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x5ee96147.
//
// Solidity: function updateAddress(string uid, string[] addressList) returns()
func (_Contracts *ContractsTransactor) UpdateAddress(opts *bind.TransactOpts, uid string, addressList []string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateAddress", uid, addressList)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x5ee96147.
//
// Solidity: function updateAddress(string uid, string[] addressList) returns()
func (_Contracts *ContractsSession) UpdateAddress(uid string, addressList []string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateAddress(&_Contracts.TransactOpts, uid, addressList)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x5ee96147.
//
// Solidity: function updateAddress(string uid, string[] addressList) returns()
func (_Contracts *ContractsTransactorSession) UpdateAddress(uid string, addressList []string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateAddress(&_Contracts.TransactOpts, uid, addressList)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string uid, uint256 wad) returns()
func (_Contracts *ContractsTransactor) Withdraw(opts *bind.TransactOpts, uid string, wad *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdraw", uid, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string uid, uint256 wad) returns()
func (_Contracts *ContractsSession) Withdraw(uid string, wad *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, uid, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string uid, uint256 wad) returns()
func (_Contracts *ContractsTransactorSession) Withdraw(uid string, wad *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, uid, wad)
}

// ContractsSuccessIterator is returned from FilterSuccess and is used to iterate over the raw logs and unpacked data for Success events raised by the Contracts contract.
type ContractsSuccessIterator struct {
	Event *ContractsSuccess // Event containing the contract specifics and raw log

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
func (it *ContractsSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSuccess)
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
		it.Event = new(ContractsSuccess)
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
func (it *ContractsSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSuccess represents a Success event raised by the Contracts contract.
type ContractsSuccess struct {
	Uid common.Hash
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSuccess is a free log retrieval operation binding the contract event 0x166b96a382dc3abc86ba052ca9c020cfd7022d2d3059e868918a3d20779e41cc.
//
// Solidity: event Success(string indexed uid)
func (_Contracts *ContractsFilterer) FilterSuccess(opts *bind.FilterOpts, uid []string) (*ContractsSuccessIterator, error) {

	var uidRule []interface{}
	for _, uidItem := range uid {
		uidRule = append(uidRule, uidItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Success", uidRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSuccessIterator{contract: _Contracts.contract, event: "Success", logs: logs, sub: sub}, nil
}

// WatchSuccess is a free log subscription operation binding the contract event 0x166b96a382dc3abc86ba052ca9c020cfd7022d2d3059e868918a3d20779e41cc.
//
// Solidity: event Success(string indexed uid)
func (_Contracts *ContractsFilterer) WatchSuccess(opts *bind.WatchOpts, sink chan<- *ContractsSuccess, uid []string) (event.Subscription, error) {

	var uidRule []interface{}
	for _, uidItem := range uid {
		uidRule = append(uidRule, uidItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Success", uidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSuccess)
				if err := _Contracts.contract.UnpackLog(event, "Success", log); err != nil {
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

// ParseSuccess is a log parse operation binding the contract event 0x166b96a382dc3abc86ba052ca9c020cfd7022d2d3059e868918a3d20779e41cc.
//
// Solidity: event Success(string indexed uid)
func (_Contracts *ContractsFilterer) ParseSuccess(log types.Log) (*ContractsSuccess, error) {
	event := new(ContractsSuccess)
	if err := _Contracts.contract.UnpackLog(event, "Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
