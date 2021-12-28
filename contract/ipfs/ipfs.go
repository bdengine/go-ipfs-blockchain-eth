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

// IpfsMetaData contains all meta data concerning the Ipfs contract.
var IpfsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"Success\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"SetOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"SetPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"SetToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"addFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"peerId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"addressList\",\"type\":\"string[]\"}],\"name\":\"addPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addrPeerMap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"PeerId\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"fileMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"}],\"name\":\"getPeerByPid\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"AddressList\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structIPFS.Peer\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getPeerList\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"AddressList\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structIPFS.Peer[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"head\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"peerHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"peerList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peerNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"pidAddrMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"rechargeFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"removeFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"removePeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"addressList\",\"type\":\"string[]\"}],\"name\":\"updateAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IpfsABI is the input ABI used to generate the binding from.
// Deprecated: Use IpfsMetaData.ABI instead.
var IpfsABI = IpfsMetaData.ABI

// Ipfs is an auto generated Go binding around an Ethereum contract.
type Ipfs struct {
	IpfsCaller     // Read-only binding to the contract
	IpfsTransactor // Write-only binding to the contract
	IpfsFilterer   // Log filterer for contract events
}

// IpfsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpfsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpfsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpfsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpfsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpfsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpfsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpfsSession struct {
	Contract     *Ipfs             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpfsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpfsCallerSession struct {
	Contract *IpfsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IpfsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpfsTransactorSession struct {
	Contract     *IpfsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpfsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpfsRaw struct {
	Contract *Ipfs // Generic contract binding to access the raw methods on
}

// IpfsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpfsCallerRaw struct {
	Contract *IpfsCaller // Generic read-only contract binding to access the raw methods on
}

// IpfsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpfsTransactorRaw struct {
	Contract *IpfsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpfs creates a new instance of Ipfs, bound to a specific deployed contract.
func NewIpfs(address common.Address, backend bind.ContractBackend) (*Ipfs, error) {
	contract, err := bindIpfs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipfs{IpfsCaller: IpfsCaller{contract: contract}, IpfsTransactor: IpfsTransactor{contract: contract}, IpfsFilterer: IpfsFilterer{contract: contract}}, nil
}

// NewIpfsCaller creates a new read-only instance of Ipfs, bound to a specific deployed contract.
func NewIpfsCaller(address common.Address, caller bind.ContractCaller) (*IpfsCaller, error) {
	contract, err := bindIpfs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpfsCaller{contract: contract}, nil
}

// NewIpfsTransactor creates a new write-only instance of Ipfs, bound to a specific deployed contract.
func NewIpfsTransactor(address common.Address, transactor bind.ContractTransactor) (*IpfsTransactor, error) {
	contract, err := bindIpfs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpfsTransactor{contract: contract}, nil
}

// NewIpfsFilterer creates a new log filterer instance of Ipfs, bound to a specific deployed contract.
func NewIpfsFilterer(address common.Address, filterer bind.ContractFilterer) (*IpfsFilterer, error) {
	contract, err := bindIpfs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpfsFilterer{contract: contract}, nil
}

// bindIpfs binds a generic wrapper to an already deployed contract.
func bindIpfs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IpfsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipfs *IpfsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipfs.Contract.IpfsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipfs *IpfsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipfs.Contract.IpfsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipfs *IpfsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipfs.Contract.IpfsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipfs *IpfsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipfs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipfs *IpfsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipfs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipfs *IpfsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipfs.Contract.contract.Transact(opts, method, params...)
}

// AddrPeerMap is a free data retrieval call binding the contract method 0x9d6c39a2.
//
// Solidity: function addrPeerMap(address ) view returns(string PeerId, bool valid)
func (_Ipfs *IpfsCaller) AddrPeerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	PeerId string
	Valid  bool
}, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "addrPeerMap", arg0)

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
func (_Ipfs *IpfsSession) AddrPeerMap(arg0 common.Address) (struct {
	PeerId string
	Valid  bool
}, error) {
	return _Ipfs.Contract.AddrPeerMap(&_Ipfs.CallOpts, arg0)
}

// AddrPeerMap is a free data retrieval call binding the contract method 0x9d6c39a2.
//
// Solidity: function addrPeerMap(address ) view returns(string PeerId, bool valid)
func (_Ipfs *IpfsCallerSession) AddrPeerMap(arg0 common.Address) (struct {
	PeerId string
	Valid  bool
}, error) {
	return _Ipfs.Contract.AddrPeerMap(&_Ipfs.CallOpts, arg0)
}

// FileMap is a free data retrieval call binding the contract method 0x167e3841.
//
// Solidity: function fileMap(string ) view returns(address owner, uint256 size, uint256 expireBlock, address uploader)
func (_Ipfs *IpfsCaller) FileMap(opts *bind.CallOpts, arg0 string) (struct {
	Owner       common.Address
	Size        *big.Int
	ExpireBlock *big.Int
	Uploader    common.Address
}, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "fileMap", arg0)

	outstruct := new(struct {
		Owner       common.Address
		Size        *big.Int
		ExpireBlock *big.Int
		Uploader    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Size = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExpireBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Uploader = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// FileMap is a free data retrieval call binding the contract method 0x167e3841.
//
// Solidity: function fileMap(string ) view returns(address owner, uint256 size, uint256 expireBlock, address uploader)
func (_Ipfs *IpfsSession) FileMap(arg0 string) (struct {
	Owner       common.Address
	Size        *big.Int
	ExpireBlock *big.Int
	Uploader    common.Address
}, error) {
	return _Ipfs.Contract.FileMap(&_Ipfs.CallOpts, arg0)
}

// FileMap is a free data retrieval call binding the contract method 0x167e3841.
//
// Solidity: function fileMap(string ) view returns(address owner, uint256 size, uint256 expireBlock, address uploader)
func (_Ipfs *IpfsCallerSession) FileMap(arg0 string) (struct {
	Owner       common.Address
	Size        *big.Int
	ExpireBlock *big.Int
	Uploader    common.Address
}, error) {
	return _Ipfs.Contract.FileMap(&_Ipfs.CallOpts, arg0)
}

// GetPeerByPid is a free data retrieval call binding the contract method 0xfd993f96.
//
// Solidity: function getPeerByPid(string pid) view returns((string,string[],bool))
func (_Ipfs *IpfsCaller) GetPeerByPid(opts *bind.CallOpts, pid string) (IPFSPeer, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getPeerByPid", pid)

	if err != nil {
		return *new(IPFSPeer), err
	}

	out0 := *abi.ConvertType(out[0], new(IPFSPeer)).(*IPFSPeer)

	return out0, err

}

// GetPeerByPid is a free data retrieval call binding the contract method 0xfd993f96.
//
// Solidity: function getPeerByPid(string pid) view returns((string,string[],bool))
func (_Ipfs *IpfsSession) GetPeerByPid(pid string) (IPFSPeer, error) {
	return _Ipfs.Contract.GetPeerByPid(&_Ipfs.CallOpts, pid)
}

// GetPeerByPid is a free data retrieval call binding the contract method 0xfd993f96.
//
// Solidity: function getPeerByPid(string pid) view returns((string,string[],bool))
func (_Ipfs *IpfsCallerSession) GetPeerByPid(pid string) (IPFSPeer, error) {
	return _Ipfs.Contract.GetPeerByPid(&_Ipfs.CallOpts, pid)
}

// GetPeerList is a free data retrieval call binding the contract method 0x21d47c2e.
//
// Solidity: function getPeerList(uint256 num) view returns((string,string[],bool)[])
func (_Ipfs *IpfsCaller) GetPeerList(opts *bind.CallOpts, num *big.Int) ([]IPFSPeer, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getPeerList", num)

	if err != nil {
		return *new([]IPFSPeer), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPFSPeer)).(*[]IPFSPeer)

	return out0, err

}

// GetPeerList is a free data retrieval call binding the contract method 0x21d47c2e.
//
// Solidity: function getPeerList(uint256 num) view returns((string,string[],bool)[])
func (_Ipfs *IpfsSession) GetPeerList(num *big.Int) ([]IPFSPeer, error) {
	return _Ipfs.Contract.GetPeerList(&_Ipfs.CallOpts, num)
}

// GetPeerList is a free data retrieval call binding the contract method 0x21d47c2e.
//
// Solidity: function getPeerList(uint256 num) view returns((string,string[],bool)[])
func (_Ipfs *IpfsCallerSession) GetPeerList(num *big.Int) ([]IPFSPeer, error) {
	return _Ipfs.Contract.GetPeerList(&_Ipfs.CallOpts, num)
}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_Ipfs *IpfsCaller) Head(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "head")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_Ipfs *IpfsSession) Head() (common.Address, error) {
	return _Ipfs.Contract.Head(&_Ipfs.CallOpts)
}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_Ipfs *IpfsCallerSession) Head() (common.Address, error) {
	return _Ipfs.Contract.Head(&_Ipfs.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ipfs *IpfsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ipfs *IpfsSession) Owner() (common.Address, error) {
	return _Ipfs.Contract.Owner(&_Ipfs.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ipfs *IpfsCallerSession) Owner() (common.Address, error) {
	return _Ipfs.Contract.Owner(&_Ipfs.CallOpts)
}

// PeerList is a free data retrieval call binding the contract method 0xe96e501a.
//
// Solidity: function peerList(address ) view returns(address)
func (_Ipfs *IpfsCaller) PeerList(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "peerList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PeerList is a free data retrieval call binding the contract method 0xe96e501a.
//
// Solidity: function peerList(address ) view returns(address)
func (_Ipfs *IpfsSession) PeerList(arg0 common.Address) (common.Address, error) {
	return _Ipfs.Contract.PeerList(&_Ipfs.CallOpts, arg0)
}

// PeerList is a free data retrieval call binding the contract method 0xe96e501a.
//
// Solidity: function peerList(address ) view returns(address)
func (_Ipfs *IpfsCallerSession) PeerList(arg0 common.Address) (common.Address, error) {
	return _Ipfs.Contract.PeerList(&_Ipfs.CallOpts, arg0)
}

// PeerNum is a free data retrieval call binding the contract method 0xf9c20442.
//
// Solidity: function peerNum() view returns(uint256)
func (_Ipfs *IpfsCaller) PeerNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "peerNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeerNum is a free data retrieval call binding the contract method 0xf9c20442.
//
// Solidity: function peerNum() view returns(uint256)
func (_Ipfs *IpfsSession) PeerNum() (*big.Int, error) {
	return _Ipfs.Contract.PeerNum(&_Ipfs.CallOpts)
}

// PeerNum is a free data retrieval call binding the contract method 0xf9c20442.
//
// Solidity: function peerNum() view returns(uint256)
func (_Ipfs *IpfsCallerSession) PeerNum() (*big.Int, error) {
	return _Ipfs.Contract.PeerNum(&_Ipfs.CallOpts)
}

// PidAddrMap is a free data retrieval call binding the contract method 0xad284b46.
//
// Solidity: function pidAddrMap(string ) view returns(address)
func (_Ipfs *IpfsCaller) PidAddrMap(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "pidAddrMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PidAddrMap is a free data retrieval call binding the contract method 0xad284b46.
//
// Solidity: function pidAddrMap(string ) view returns(address)
func (_Ipfs *IpfsSession) PidAddrMap(arg0 string) (common.Address, error) {
	return _Ipfs.Contract.PidAddrMap(&_Ipfs.CallOpts, arg0)
}

// PidAddrMap is a free data retrieval call binding the contract method 0xad284b46.
//
// Solidity: function pidAddrMap(string ) view returns(address)
func (_Ipfs *IpfsCallerSession) PidAddrMap(arg0 string) (common.Address, error) {
	return _Ipfs.Contract.PidAddrMap(&_Ipfs.CallOpts, arg0)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Ipfs *IpfsCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Ipfs *IpfsSession) Price() (*big.Int, error) {
	return _Ipfs.Contract.Price(&_Ipfs.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Ipfs *IpfsCallerSession) Price() (*big.Int, error) {
	return _Ipfs.Contract.Price(&_Ipfs.CallOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x167d3e9c.
//
// Solidity: function SetOwner(address _to) returns()
func (_Ipfs *IpfsTransactor) SetOwner(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "SetOwner", _to)
}

// SetOwner is a paid mutator transaction binding the contract method 0x167d3e9c.
//
// Solidity: function SetOwner(address _to) returns()
func (_Ipfs *IpfsSession) SetOwner(_to common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.SetOwner(&_Ipfs.TransactOpts, _to)
}

// SetOwner is a paid mutator transaction binding the contract method 0x167d3e9c.
//
// Solidity: function SetOwner(address _to) returns()
func (_Ipfs *IpfsTransactorSession) SetOwner(_to common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.SetOwner(&_Ipfs.TransactOpts, _to)
}

// SetPrice is a paid mutator transaction binding the contract method 0x4f5539c0.
//
// Solidity: function SetPrice(uint256 _price) returns()
func (_Ipfs *IpfsTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "SetPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x4f5539c0.
//
// Solidity: function SetPrice(uint256 _price) returns()
func (_Ipfs *IpfsSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.SetPrice(&_Ipfs.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x4f5539c0.
//
// Solidity: function SetPrice(uint256 _price) returns()
func (_Ipfs *IpfsTransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.SetPrice(&_Ipfs.TransactOpts, _price)
}

// SetToken is a paid mutator transaction binding the contract method 0xefc1fd16.
//
// Solidity: function SetToken(address _token) returns()
func (_Ipfs *IpfsTransactor) SetToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "SetToken", _token)
}

// SetToken is a paid mutator transaction binding the contract method 0xefc1fd16.
//
// Solidity: function SetToken(address _token) returns()
func (_Ipfs *IpfsSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.SetToken(&_Ipfs.TransactOpts, _token)
}

// SetToken is a paid mutator transaction binding the contract method 0xefc1fd16.
//
// Solidity: function SetToken(address _token) returns()
func (_Ipfs *IpfsTransactorSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.SetToken(&_Ipfs.TransactOpts, _token)
}

// AddFile is a paid mutator transaction binding the contract method 0xc0f645b2.
//
// Solidity: function addFile(string uid, string cid, uint256 size, uint256 blockNum, address _owner) returns()
func (_Ipfs *IpfsTransactor) AddFile(opts *bind.TransactOpts, uid string, cid string, size *big.Int, blockNum *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "addFile", uid, cid, size, blockNum, _owner)
}

// AddFile is a paid mutator transaction binding the contract method 0xc0f645b2.
//
// Solidity: function addFile(string uid, string cid, uint256 size, uint256 blockNum, address _owner) returns()
func (_Ipfs *IpfsSession) AddFile(uid string, cid string, size *big.Int, blockNum *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.AddFile(&_Ipfs.TransactOpts, uid, cid, size, blockNum, _owner)
}

// AddFile is a paid mutator transaction binding the contract method 0xc0f645b2.
//
// Solidity: function addFile(string uid, string cid, uint256 size, uint256 blockNum, address _owner) returns()
func (_Ipfs *IpfsTransactorSession) AddFile(uid string, cid string, size *big.Int, blockNum *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.AddFile(&_Ipfs.TransactOpts, uid, cid, size, blockNum, _owner)
}

// AddPeer is a paid mutator transaction binding the contract method 0x5124bd12.
//
// Solidity: function addPeer(string uid, string peerId, string[] addressList) returns()
func (_Ipfs *IpfsTransactor) AddPeer(opts *bind.TransactOpts, uid string, peerId string, addressList []string) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "addPeer", uid, peerId, addressList)
}

// AddPeer is a paid mutator transaction binding the contract method 0x5124bd12.
//
// Solidity: function addPeer(string uid, string peerId, string[] addressList) returns()
func (_Ipfs *IpfsSession) AddPeer(uid string, peerId string, addressList []string) (*types.Transaction, error) {
	return _Ipfs.Contract.AddPeer(&_Ipfs.TransactOpts, uid, peerId, addressList)
}

// AddPeer is a paid mutator transaction binding the contract method 0x5124bd12.
//
// Solidity: function addPeer(string uid, string peerId, string[] addressList) returns()
func (_Ipfs *IpfsTransactorSession) AddPeer(uid string, peerId string, addressList []string) (*types.Transaction, error) {
	return _Ipfs.Contract.AddPeer(&_Ipfs.TransactOpts, uid, peerId, addressList)
}

// PeerHeartbeat is a paid mutator transaction binding the contract method 0xf71e3e11.
//
// Solidity: function peerHeartbeat(string uid) returns()
func (_Ipfs *IpfsTransactor) PeerHeartbeat(opts *bind.TransactOpts, uid string) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "peerHeartbeat", uid)
}

// PeerHeartbeat is a paid mutator transaction binding the contract method 0xf71e3e11.
//
// Solidity: function peerHeartbeat(string uid) returns()
func (_Ipfs *IpfsSession) PeerHeartbeat(uid string) (*types.Transaction, error) {
	return _Ipfs.Contract.PeerHeartbeat(&_Ipfs.TransactOpts, uid)
}

// PeerHeartbeat is a paid mutator transaction binding the contract method 0xf71e3e11.
//
// Solidity: function peerHeartbeat(string uid) returns()
func (_Ipfs *IpfsTransactorSession) PeerHeartbeat(uid string) (*types.Transaction, error) {
	return _Ipfs.Contract.PeerHeartbeat(&_Ipfs.TransactOpts, uid)
}

// RechargeFile is a paid mutator transaction binding the contract method 0x39429038.
//
// Solidity: function rechargeFile(string uid, string cid, uint256 blockNum) returns()
func (_Ipfs *IpfsTransactor) RechargeFile(opts *bind.TransactOpts, uid string, cid string, blockNum *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "rechargeFile", uid, cid, blockNum)
}

// RechargeFile is a paid mutator transaction binding the contract method 0x39429038.
//
// Solidity: function rechargeFile(string uid, string cid, uint256 blockNum) returns()
func (_Ipfs *IpfsSession) RechargeFile(uid string, cid string, blockNum *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.RechargeFile(&_Ipfs.TransactOpts, uid, cid, blockNum)
}

// RechargeFile is a paid mutator transaction binding the contract method 0x39429038.
//
// Solidity: function rechargeFile(string uid, string cid, uint256 blockNum) returns()
func (_Ipfs *IpfsTransactorSession) RechargeFile(uid string, cid string, blockNum *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.RechargeFile(&_Ipfs.TransactOpts, uid, cid, blockNum)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x99f553dc.
//
// Solidity: function removeFile(string uid, string cid) returns()
func (_Ipfs *IpfsTransactor) RemoveFile(opts *bind.TransactOpts, uid string, cid string) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "removeFile", uid, cid)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x99f553dc.
//
// Solidity: function removeFile(string uid, string cid) returns()
func (_Ipfs *IpfsSession) RemoveFile(uid string, cid string) (*types.Transaction, error) {
	return _Ipfs.Contract.RemoveFile(&_Ipfs.TransactOpts, uid, cid)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x99f553dc.
//
// Solidity: function removeFile(string uid, string cid) returns()
func (_Ipfs *IpfsTransactorSession) RemoveFile(uid string, cid string) (*types.Transaction, error) {
	return _Ipfs.Contract.RemoveFile(&_Ipfs.TransactOpts, uid, cid)
}

// RemovePeer is a paid mutator transaction binding the contract method 0xf683d96c.
//
// Solidity: function removePeer(string uid) returns()
func (_Ipfs *IpfsTransactor) RemovePeer(opts *bind.TransactOpts, uid string) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "removePeer", uid)
}

// RemovePeer is a paid mutator transaction binding the contract method 0xf683d96c.
//
// Solidity: function removePeer(string uid) returns()
func (_Ipfs *IpfsSession) RemovePeer(uid string) (*types.Transaction, error) {
	return _Ipfs.Contract.RemovePeer(&_Ipfs.TransactOpts, uid)
}

// RemovePeer is a paid mutator transaction binding the contract method 0xf683d96c.
//
// Solidity: function removePeer(string uid) returns()
func (_Ipfs *IpfsTransactorSession) RemovePeer(uid string) (*types.Transaction, error) {
	return _Ipfs.Contract.RemovePeer(&_Ipfs.TransactOpts, uid)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x5ee96147.
//
// Solidity: function updateAddress(string uid, string[] addressList) returns()
func (_Ipfs *IpfsTransactor) UpdateAddress(opts *bind.TransactOpts, uid string, addressList []string) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "updateAddress", uid, addressList)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x5ee96147.
//
// Solidity: function updateAddress(string uid, string[] addressList) returns()
func (_Ipfs *IpfsSession) UpdateAddress(uid string, addressList []string) (*types.Transaction, error) {
	return _Ipfs.Contract.UpdateAddress(&_Ipfs.TransactOpts, uid, addressList)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x5ee96147.
//
// Solidity: function updateAddress(string uid, string[] addressList) returns()
func (_Ipfs *IpfsTransactorSession) UpdateAddress(uid string, addressList []string) (*types.Transaction, error) {
	return _Ipfs.Contract.UpdateAddress(&_Ipfs.TransactOpts, uid, addressList)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string uid, uint256 wad) returns()
func (_Ipfs *IpfsTransactor) Withdraw(opts *bind.TransactOpts, uid string, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "withdraw", uid, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string uid, uint256 wad) returns()
func (_Ipfs *IpfsSession) Withdraw(uid string, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Withdraw(&_Ipfs.TransactOpts, uid, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string uid, uint256 wad) returns()
func (_Ipfs *IpfsTransactorSession) Withdraw(uid string, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Withdraw(&_Ipfs.TransactOpts, uid, wad)
}

// IpfsSuccessIterator is returned from FilterSuccess and is used to iterate over the raw logs and unpacked data for Success events raised by the Ipfs contract.
type IpfsSuccessIterator struct {
	Event *IpfsSuccess // Event containing the contract specifics and raw log

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
func (it *IpfsSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsSuccess)
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
		it.Event = new(IpfsSuccess)
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
func (it *IpfsSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsSuccess represents a Success event raised by the Ipfs contract.
type IpfsSuccess struct {
	Uid common.Hash
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSuccess is a free log retrieval operation binding the contract event 0x166b96a382dc3abc86ba052ca9c020cfd7022d2d3059e868918a3d20779e41cc.
//
// Solidity: event Success(string indexed uid)
func (_Ipfs *IpfsFilterer) FilterSuccess(opts *bind.FilterOpts, uid []string) (*IpfsSuccessIterator, error) {

	var uidRule []interface{}
	for _, uidItem := range uid {
		uidRule = append(uidRule, uidItem)
	}

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "Success", uidRule)
	if err != nil {
		return nil, err
	}
	return &IpfsSuccessIterator{contract: _Ipfs.contract, event: "Success", logs: logs, sub: sub}, nil
}

// WatchSuccess is a free log subscription operation binding the contract event 0x166b96a382dc3abc86ba052ca9c020cfd7022d2d3059e868918a3d20779e41cc.
//
// Solidity: event Success(string indexed uid)
func (_Ipfs *IpfsFilterer) WatchSuccess(opts *bind.WatchOpts, sink chan<- *IpfsSuccess, uid []string) (event.Subscription, error) {

	var uidRule []interface{}
	for _, uidItem := range uid {
		uidRule = append(uidRule, uidItem)
	}

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "Success", uidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsSuccess)
				if err := _Ipfs.contract.UnpackLog(event, "Success", log); err != nil {
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
func (_Ipfs *IpfsFilterer) ParseSuccess(log types.Log) (*IpfsSuccess, error) {
	event := new(IpfsSuccess)
	if err := _Ipfs.contract.UnpackLog(event, "Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
