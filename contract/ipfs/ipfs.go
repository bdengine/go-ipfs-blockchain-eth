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

// ipfsFile is an auto generated low-level Go binding around an user-defined struct.
type ipfsFile struct {
	Owner       common.Address
	SliceNum    *big.Int
	Size        *big.Int
	ExpireBlock *big.Int
	Uploader    common.Address
	Id          *big.Int
}

// ipfsPeer is an auto generated low-level Go binding around an user-defined struct.
type ipfsPeer struct {
	PeerId      string
	AddressList []string
	PeerAddress common.Address
	Valid       bool
}

// IpfsMetaData contains all meta data concerning the Ipfs contract.
var IpfsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cge\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cgeSeq\",\"type\":\"uint256\"}],\"name\":\"SetChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cge\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cgeSeq\",\"type\":\"uint256\"}],\"name\":\"SetStoreChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cge\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cgeSeq\",\"type\":\"uint256\"}],\"name\":\"SetWinner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Start\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"SetPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sliceNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"addFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"peerId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"addressList\",\"type\":\"string[]\"}],\"name\":\"addPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burnSelf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"checkFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChallenge\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChallengeStage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"getFile\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sliceNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"internalType\":\"structipfs.File\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getFileList\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFileStatistic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIpfsMintNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"}],\"name\":\"getPeerByPid\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"AddressList\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"peerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structipfs.Peer\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getPeerList\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"AddressList\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"peerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structipfs.Peer[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStoreChallenge\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWinner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ipfsNode\",\"type\":\"address\"}],\"name\":\"mintForIpfs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mintSelf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"peerHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peerStore\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"head\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"peerNum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"push\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"rechargeFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"removeFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"removePeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cyclePeriod\",\"type\":\"uint256\"}],\"name\":\"resetChallengeRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_challenge\",\"type\":\"string\"}],\"name\":\"setChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setIpfsMintNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_storeChallenge\",\"type\":\"string\"}],\"name\":\"setStoreChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_winner\",\"type\":\"string\"}],\"name\":\"setWinner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"addressList\",\"type\":\"string[]\"}],\"name\":\"updateAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Ipfs *IpfsCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Ipfs *IpfsSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ipfs.Contract.Allowance(&_Ipfs.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Ipfs *IpfsCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ipfs.Contract.Allowance(&_Ipfs.CallOpts, arg0, arg1)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Ipfs *IpfsCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Ipfs *IpfsSession) Authority() (common.Address, error) {
	return _Ipfs.Contract.Authority(&_Ipfs.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Ipfs *IpfsCallerSession) Authority() (common.Address, error) {
	return _Ipfs.Contract.Authority(&_Ipfs.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Ipfs *IpfsCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Ipfs *IpfsSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Ipfs.Contract.BalanceOf(&_Ipfs.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Ipfs *IpfsCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Ipfs.Contract.BalanceOf(&_Ipfs.CallOpts, arg0)
}

// CheckFile is a free data retrieval call binding the contract method 0x12df278f.
//
// Solidity: function checkFile(string cid) view returns(bool)
func (_Ipfs *IpfsCaller) CheckFile(opts *bind.CallOpts, cid string) (bool, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "checkFile", cid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckFile is a free data retrieval call binding the contract method 0x12df278f.
//
// Solidity: function checkFile(string cid) view returns(bool)
func (_Ipfs *IpfsSession) CheckFile(cid string) (bool, error) {
	return _Ipfs.Contract.CheckFile(&_Ipfs.CallOpts, cid)
}

// CheckFile is a free data retrieval call binding the contract method 0x12df278f.
//
// Solidity: function checkFile(string cid) view returns(bool)
func (_Ipfs *IpfsCallerSession) CheckFile(cid string) (bool, error) {
	return _Ipfs.Contract.CheckFile(&_Ipfs.CallOpts, cid)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ipfs *IpfsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ipfs *IpfsSession) Decimals() (uint8, error) {
	return _Ipfs.Contract.Decimals(&_Ipfs.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ipfs *IpfsCallerSession) Decimals() (uint8, error) {
	return _Ipfs.Contract.Decimals(&_Ipfs.CallOpts)
}

// GetChallenge is a free data retrieval call binding the contract method 0x759014f0.
//
// Solidity: function getChallenge() view returns(string, uint256)
func (_Ipfs *IpfsCaller) GetChallenge(opts *bind.CallOpts) (string, *big.Int, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getChallenge")

	if err != nil {
		return *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetChallenge is a free data retrieval call binding the contract method 0x759014f0.
//
// Solidity: function getChallenge() view returns(string, uint256)
func (_Ipfs *IpfsSession) GetChallenge() (string, *big.Int, error) {
	return _Ipfs.Contract.GetChallenge(&_Ipfs.CallOpts)
}

// GetChallenge is a free data retrieval call binding the contract method 0x759014f0.
//
// Solidity: function getChallenge() view returns(string, uint256)
func (_Ipfs *IpfsCallerSession) GetChallenge() (string, *big.Int, error) {
	return _Ipfs.Contract.GetChallenge(&_Ipfs.CallOpts)
}

// GetChallengeStage is a free data retrieval call binding the contract method 0x018904bd.
//
// Solidity: function getChallengeStage() view returns(uint256, string, uint256)
func (_Ipfs *IpfsCaller) GetChallengeStage(opts *bind.CallOpts) (*big.Int, string, *big.Int, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getChallengeStage")

	if err != nil {
		return *new(*big.Int), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetChallengeStage is a free data retrieval call binding the contract method 0x018904bd.
//
// Solidity: function getChallengeStage() view returns(uint256, string, uint256)
func (_Ipfs *IpfsSession) GetChallengeStage() (*big.Int, string, *big.Int, error) {
	return _Ipfs.Contract.GetChallengeStage(&_Ipfs.CallOpts)
}

// GetChallengeStage is a free data retrieval call binding the contract method 0x018904bd.
//
// Solidity: function getChallengeStage() view returns(uint256, string, uint256)
func (_Ipfs *IpfsCallerSession) GetChallengeStage() (*big.Int, string, *big.Int, error) {
	return _Ipfs.Contract.GetChallengeStage(&_Ipfs.CallOpts)
}

// GetFile is a free data retrieval call binding the contract method 0xe0876aa8.
//
// Solidity: function getFile(string cid) view returns((address,uint256,uint256,uint256,address,uint256))
func (_Ipfs *IpfsCaller) GetFile(opts *bind.CallOpts, cid string) (ipfsFile, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getFile", cid)

	if err != nil {
		return *new(ipfsFile), err
	}

	out0 := *abi.ConvertType(out[0], new(ipfsFile)).(*ipfsFile)

	return out0, err

}

// GetFile is a free data retrieval call binding the contract method 0xe0876aa8.
//
// Solidity: function getFile(string cid) view returns((address,uint256,uint256,uint256,address,uint256))
func (_Ipfs *IpfsSession) GetFile(cid string) (ipfsFile, error) {
	return _Ipfs.Contract.GetFile(&_Ipfs.CallOpts, cid)
}

// GetFile is a free data retrieval call binding the contract method 0xe0876aa8.
//
// Solidity: function getFile(string cid) view returns((address,uint256,uint256,uint256,address,uint256))
func (_Ipfs *IpfsCallerSession) GetFile(cid string) (ipfsFile, error) {
	return _Ipfs.Contract.GetFile(&_Ipfs.CallOpts, cid)
}

// GetFileList is a free data retrieval call binding the contract method 0x31260d1e.
//
// Solidity: function getFileList(uint256 num) view returns(string[])
func (_Ipfs *IpfsCaller) GetFileList(opts *bind.CallOpts, num *big.Int) ([]string, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getFileList", num)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetFileList is a free data retrieval call binding the contract method 0x31260d1e.
//
// Solidity: function getFileList(uint256 num) view returns(string[])
func (_Ipfs *IpfsSession) GetFileList(num *big.Int) ([]string, error) {
	return _Ipfs.Contract.GetFileList(&_Ipfs.CallOpts, num)
}

// GetFileList is a free data retrieval call binding the contract method 0x31260d1e.
//
// Solidity: function getFileList(uint256 num) view returns(string[])
func (_Ipfs *IpfsCallerSession) GetFileList(num *big.Int) ([]string, error) {
	return _Ipfs.Contract.GetFileList(&_Ipfs.CallOpts, num)
}

// GetFileStatistic is a free data retrieval call binding the contract method 0x2bbbe9e7.
//
// Solidity: function getFileStatistic() view returns(uint256, uint256)
func (_Ipfs *IpfsCaller) GetFileStatistic(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getFileStatistic")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFileStatistic is a free data retrieval call binding the contract method 0x2bbbe9e7.
//
// Solidity: function getFileStatistic() view returns(uint256, uint256)
func (_Ipfs *IpfsSession) GetFileStatistic() (*big.Int, *big.Int, error) {
	return _Ipfs.Contract.GetFileStatistic(&_Ipfs.CallOpts)
}

// GetFileStatistic is a free data retrieval call binding the contract method 0x2bbbe9e7.
//
// Solidity: function getFileStatistic() view returns(uint256, uint256)
func (_Ipfs *IpfsCallerSession) GetFileStatistic() (*big.Int, *big.Int, error) {
	return _Ipfs.Contract.GetFileStatistic(&_Ipfs.CallOpts)
}

// GetIpfsMintNum is a free data retrieval call binding the contract method 0xa77e1a93.
//
// Solidity: function getIpfsMintNum() view returns(uint256)
func (_Ipfs *IpfsCaller) GetIpfsMintNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getIpfsMintNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIpfsMintNum is a free data retrieval call binding the contract method 0xa77e1a93.
//
// Solidity: function getIpfsMintNum() view returns(uint256)
func (_Ipfs *IpfsSession) GetIpfsMintNum() (*big.Int, error) {
	return _Ipfs.Contract.GetIpfsMintNum(&_Ipfs.CallOpts)
}

// GetIpfsMintNum is a free data retrieval call binding the contract method 0xa77e1a93.
//
// Solidity: function getIpfsMintNum() view returns(uint256)
func (_Ipfs *IpfsCallerSession) GetIpfsMintNum() (*big.Int, error) {
	return _Ipfs.Contract.GetIpfsMintNum(&_Ipfs.CallOpts)
}

// GetPeerByPid is a free data retrieval call binding the contract method 0xfd993f96.
//
// Solidity: function getPeerByPid(string pid) view returns((string,string[],address,bool))
func (_Ipfs *IpfsCaller) GetPeerByPid(opts *bind.CallOpts, pid string) (ipfsPeer, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getPeerByPid", pid)

	if err != nil {
		return *new(ipfsPeer), err
	}

	out0 := *abi.ConvertType(out[0], new(ipfsPeer)).(*ipfsPeer)

	return out0, err

}

// GetPeerByPid is a free data retrieval call binding the contract method 0xfd993f96.
//
// Solidity: function getPeerByPid(string pid) view returns((string,string[],address,bool))
func (_Ipfs *IpfsSession) GetPeerByPid(pid string) (ipfsPeer, error) {
	return _Ipfs.Contract.GetPeerByPid(&_Ipfs.CallOpts, pid)
}

// GetPeerByPid is a free data retrieval call binding the contract method 0xfd993f96.
//
// Solidity: function getPeerByPid(string pid) view returns((string,string[],address,bool))
func (_Ipfs *IpfsCallerSession) GetPeerByPid(pid string) (ipfsPeer, error) {
	return _Ipfs.Contract.GetPeerByPid(&_Ipfs.CallOpts, pid)
}

// GetPeerList is a free data retrieval call binding the contract method 0x21d47c2e.
//
// Solidity: function getPeerList(uint256 num) view returns((string,string[],address,bool)[])
func (_Ipfs *IpfsCaller) GetPeerList(opts *bind.CallOpts, num *big.Int) ([]ipfsPeer, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getPeerList", num)

	if err != nil {
		return *new([]ipfsPeer), err
	}

	out0 := *abi.ConvertType(out[0], new([]ipfsPeer)).(*[]ipfsPeer)

	return out0, err

}

// GetPeerList is a free data retrieval call binding the contract method 0x21d47c2e.
//
// Solidity: function getPeerList(uint256 num) view returns((string,string[],address,bool)[])
func (_Ipfs *IpfsSession) GetPeerList(num *big.Int) ([]ipfsPeer, error) {
	return _Ipfs.Contract.GetPeerList(&_Ipfs.CallOpts, num)
}

// GetPeerList is a free data retrieval call binding the contract method 0x21d47c2e.
//
// Solidity: function getPeerList(uint256 num) view returns((string,string[],address,bool)[])
func (_Ipfs *IpfsCallerSession) GetPeerList(num *big.Int) ([]ipfsPeer, error) {
	return _Ipfs.Contract.GetPeerList(&_Ipfs.CallOpts, num)
}

// GetStoreChallenge is a free data retrieval call binding the contract method 0xe2838648.
//
// Solidity: function getStoreChallenge() view returns(string, uint256)
func (_Ipfs *IpfsCaller) GetStoreChallenge(opts *bind.CallOpts) (string, *big.Int, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getStoreChallenge")

	if err != nil {
		return *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStoreChallenge is a free data retrieval call binding the contract method 0xe2838648.
//
// Solidity: function getStoreChallenge() view returns(string, uint256)
func (_Ipfs *IpfsSession) GetStoreChallenge() (string, *big.Int, error) {
	return _Ipfs.Contract.GetStoreChallenge(&_Ipfs.CallOpts)
}

// GetStoreChallenge is a free data retrieval call binding the contract method 0xe2838648.
//
// Solidity: function getStoreChallenge() view returns(string, uint256)
func (_Ipfs *IpfsCallerSession) GetStoreChallenge() (string, *big.Int, error) {
	return _Ipfs.Contract.GetStoreChallenge(&_Ipfs.CallOpts)
}

// GetWinner is a free data retrieval call binding the contract method 0x8e7ea5b2.
//
// Solidity: function getWinner() view returns(string, uint256)
func (_Ipfs *IpfsCaller) GetWinner(opts *bind.CallOpts) (string, *big.Int, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "getWinner")

	if err != nil {
		return *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetWinner is a free data retrieval call binding the contract method 0x8e7ea5b2.
//
// Solidity: function getWinner() view returns(string, uint256)
func (_Ipfs *IpfsSession) GetWinner() (string, *big.Int, error) {
	return _Ipfs.Contract.GetWinner(&_Ipfs.CallOpts)
}

// GetWinner is a free data retrieval call binding the contract method 0x8e7ea5b2.
//
// Solidity: function getWinner() view returns(string, uint256)
func (_Ipfs *IpfsCallerSession) GetWinner() (string, *big.Int, error) {
	return _Ipfs.Contract.GetWinner(&_Ipfs.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ipfs *IpfsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ipfs *IpfsSession) Name() (string, error) {
	return _Ipfs.Contract.Name(&_Ipfs.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ipfs *IpfsCallerSession) Name() (string, error) {
	return _Ipfs.Contract.Name(&_Ipfs.CallOpts)
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

// PeerStore is a free data retrieval call binding the contract method 0x3a621d2a.
//
// Solidity: function peerStore() view returns(address head, uint256 peerNum)
func (_Ipfs *IpfsCaller) PeerStore(opts *bind.CallOpts) (struct {
	Head    common.Address
	PeerNum *big.Int
}, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "peerStore")

	outstruct := new(struct {
		Head    common.Address
		PeerNum *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Head = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PeerNum = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PeerStore is a free data retrieval call binding the contract method 0x3a621d2a.
//
// Solidity: function peerStore() view returns(address head, uint256 peerNum)
func (_Ipfs *IpfsSession) PeerStore() (struct {
	Head    common.Address
	PeerNum *big.Int
}, error) {
	return _Ipfs.Contract.PeerStore(&_Ipfs.CallOpts)
}

// PeerStore is a free data retrieval call binding the contract method 0x3a621d2a.
//
// Solidity: function peerStore() view returns(address head, uint256 peerNum)
func (_Ipfs *IpfsCallerSession) PeerStore() (struct {
	Head    common.Address
	PeerNum *big.Int
}, error) {
	return _Ipfs.Contract.PeerStore(&_Ipfs.CallOpts)
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

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(bool)
func (_Ipfs *IpfsCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "stopped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(bool)
func (_Ipfs *IpfsSession) Stopped() (bool, error) {
	return _Ipfs.Contract.Stopped(&_Ipfs.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(bool)
func (_Ipfs *IpfsCallerSession) Stopped() (bool, error) {
	return _Ipfs.Contract.Stopped(&_Ipfs.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ipfs *IpfsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ipfs *IpfsSession) Symbol() (string, error) {
	return _Ipfs.Contract.Symbol(&_Ipfs.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ipfs *IpfsCallerSession) Symbol() (string, error) {
	return _Ipfs.Contract.Symbol(&_Ipfs.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ipfs *IpfsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ipfs.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ipfs *IpfsSession) TotalSupply() (*big.Int, error) {
	return _Ipfs.Contract.TotalSupply(&_Ipfs.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ipfs *IpfsCallerSession) TotalSupply() (*big.Int, error) {
	return _Ipfs.Contract.TotalSupply(&_Ipfs.CallOpts)
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

// AddFile is a paid mutator transaction binding the contract method 0xf73aa6fb.
//
// Solidity: function addFile(string uid, string cid, uint256 size, uint256 sliceNum, uint256 blockNum, address _owner) returns()
func (_Ipfs *IpfsTransactor) AddFile(opts *bind.TransactOpts, uid string, cid string, size *big.Int, sliceNum *big.Int, blockNum *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "addFile", uid, cid, size, sliceNum, blockNum, _owner)
}

// AddFile is a paid mutator transaction binding the contract method 0xf73aa6fb.
//
// Solidity: function addFile(string uid, string cid, uint256 size, uint256 sliceNum, uint256 blockNum, address _owner) returns()
func (_Ipfs *IpfsSession) AddFile(uid string, cid string, size *big.Int, sliceNum *big.Int, blockNum *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.AddFile(&_Ipfs.TransactOpts, uid, cid, size, sliceNum, blockNum, _owner)
}

// AddFile is a paid mutator transaction binding the contract method 0xf73aa6fb.
//
// Solidity: function addFile(string uid, string cid, uint256 size, uint256 sliceNum, uint256 blockNum, address _owner) returns()
func (_Ipfs *IpfsTransactorSession) AddFile(uid string, cid string, size *big.Int, sliceNum *big.Int, blockNum *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.AddFile(&_Ipfs.TransactOpts, uid, cid, size, sliceNum, blockNum, _owner)
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

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Ipfs *IpfsTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Ipfs *IpfsSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Approve(&_Ipfs.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Ipfs *IpfsTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Approve(&_Ipfs.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_Ipfs *IpfsTransactor) Burn(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "burn", guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_Ipfs *IpfsSession) Burn(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Burn(&_Ipfs.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_Ipfs *IpfsTransactorSession) Burn(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Burn(&_Ipfs.TransactOpts, guy, wad)
}

// BurnSelf is a paid mutator transaction binding the contract method 0xb8b12a1e.
//
// Solidity: function burnSelf(uint256 wad) returns()
func (_Ipfs *IpfsTransactor) BurnSelf(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "burnSelf", wad)
}

// BurnSelf is a paid mutator transaction binding the contract method 0xb8b12a1e.
//
// Solidity: function burnSelf(uint256 wad) returns()
func (_Ipfs *IpfsSession) BurnSelf(wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.BurnSelf(&_Ipfs.TransactOpts, wad)
}

// BurnSelf is a paid mutator transaction binding the contract method 0xb8b12a1e.
//
// Solidity: function burnSelf(uint256 wad) returns()
func (_Ipfs *IpfsTransactorSession) BurnSelf(wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.BurnSelf(&_Ipfs.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address guy, uint256 wad) returns()
func (_Ipfs *IpfsTransactor) Mint(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "mint", guy, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address guy, uint256 wad) returns()
func (_Ipfs *IpfsSession) Mint(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Mint(&_Ipfs.TransactOpts, guy, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address guy, uint256 wad) returns()
func (_Ipfs *IpfsTransactorSession) Mint(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Mint(&_Ipfs.TransactOpts, guy, wad)
}

// MintForIpfs is a paid mutator transaction binding the contract method 0x551a1ccf.
//
// Solidity: function mintForIpfs(address ipfsNode) returns()
func (_Ipfs *IpfsTransactor) MintForIpfs(opts *bind.TransactOpts, ipfsNode common.Address) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "mintForIpfs", ipfsNode)
}

// MintForIpfs is a paid mutator transaction binding the contract method 0x551a1ccf.
//
// Solidity: function mintForIpfs(address ipfsNode) returns()
func (_Ipfs *IpfsSession) MintForIpfs(ipfsNode common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.MintForIpfs(&_Ipfs.TransactOpts, ipfsNode)
}

// MintForIpfs is a paid mutator transaction binding the contract method 0x551a1ccf.
//
// Solidity: function mintForIpfs(address ipfsNode) returns()
func (_Ipfs *IpfsTransactorSession) MintForIpfs(ipfsNode common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.MintForIpfs(&_Ipfs.TransactOpts, ipfsNode)
}

// MintSelf is a paid mutator transaction binding the contract method 0x720eacad.
//
// Solidity: function mintSelf(uint256 wad) returns()
func (_Ipfs *IpfsTransactor) MintSelf(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "mintSelf", wad)
}

// MintSelf is a paid mutator transaction binding the contract method 0x720eacad.
//
// Solidity: function mintSelf(uint256 wad) returns()
func (_Ipfs *IpfsSession) MintSelf(wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.MintSelf(&_Ipfs.TransactOpts, wad)
}

// MintSelf is a paid mutator transaction binding the contract method 0x720eacad.
//
// Solidity: function mintSelf(uint256 wad) returns()
func (_Ipfs *IpfsTransactorSession) MintSelf(wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.MintSelf(&_Ipfs.TransactOpts, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Ipfs *IpfsTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "move", src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Ipfs *IpfsSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Move(&_Ipfs.TransactOpts, src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Ipfs *IpfsTransactorSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Move(&_Ipfs.TransactOpts, src, dst, wad)
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

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_Ipfs *IpfsTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_Ipfs *IpfsSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Pull(&_Ipfs.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_Ipfs *IpfsTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Pull(&_Ipfs.TransactOpts, src, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_Ipfs *IpfsTransactor) Push(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "push", dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_Ipfs *IpfsSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Push(&_Ipfs.TransactOpts, dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_Ipfs *IpfsTransactorSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Push(&_Ipfs.TransactOpts, dst, wad)
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

// ResetChallengeRule is a paid mutator transaction binding the contract method 0xedb4ce39.
//
// Solidity: function resetChallengeRule(uint256 _startBlock, uint256 _cyclePeriod) returns()
func (_Ipfs *IpfsTransactor) ResetChallengeRule(opts *bind.TransactOpts, _startBlock *big.Int, _cyclePeriod *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "resetChallengeRule", _startBlock, _cyclePeriod)
}

// ResetChallengeRule is a paid mutator transaction binding the contract method 0xedb4ce39.
//
// Solidity: function resetChallengeRule(uint256 _startBlock, uint256 _cyclePeriod) returns()
func (_Ipfs *IpfsSession) ResetChallengeRule(_startBlock *big.Int, _cyclePeriod *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.ResetChallengeRule(&_Ipfs.TransactOpts, _startBlock, _cyclePeriod)
}

// ResetChallengeRule is a paid mutator transaction binding the contract method 0xedb4ce39.
//
// Solidity: function resetChallengeRule(uint256 _startBlock, uint256 _cyclePeriod) returns()
func (_Ipfs *IpfsTransactorSession) ResetChallengeRule(_startBlock *big.Int, _cyclePeriod *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.ResetChallengeRule(&_Ipfs.TransactOpts, _startBlock, _cyclePeriod)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Ipfs *IpfsTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Ipfs *IpfsSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.SetAuthority(&_Ipfs.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Ipfs *IpfsTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.SetAuthority(&_Ipfs.TransactOpts, authority_)
}

// SetChallenge is a paid mutator transaction binding the contract method 0x941d5b9c.
//
// Solidity: function setChallenge(string _challenge) returns()
func (_Ipfs *IpfsTransactor) SetChallenge(opts *bind.TransactOpts, _challenge string) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "setChallenge", _challenge)
}

// SetChallenge is a paid mutator transaction binding the contract method 0x941d5b9c.
//
// Solidity: function setChallenge(string _challenge) returns()
func (_Ipfs *IpfsSession) SetChallenge(_challenge string) (*types.Transaction, error) {
	return _Ipfs.Contract.SetChallenge(&_Ipfs.TransactOpts, _challenge)
}

// SetChallenge is a paid mutator transaction binding the contract method 0x941d5b9c.
//
// Solidity: function setChallenge(string _challenge) returns()
func (_Ipfs *IpfsTransactorSession) SetChallenge(_challenge string) (*types.Transaction, error) {
	return _Ipfs.Contract.SetChallenge(&_Ipfs.TransactOpts, _challenge)
}

// SetIpfsMintNum is a paid mutator transaction binding the contract method 0xa7e31509.
//
// Solidity: function setIpfsMintNum(uint256 num) returns()
func (_Ipfs *IpfsTransactor) SetIpfsMintNum(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "setIpfsMintNum", num)
}

// SetIpfsMintNum is a paid mutator transaction binding the contract method 0xa7e31509.
//
// Solidity: function setIpfsMintNum(uint256 num) returns()
func (_Ipfs *IpfsSession) SetIpfsMintNum(num *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.SetIpfsMintNum(&_Ipfs.TransactOpts, num)
}

// SetIpfsMintNum is a paid mutator transaction binding the contract method 0xa7e31509.
//
// Solidity: function setIpfsMintNum(uint256 num) returns()
func (_Ipfs *IpfsTransactorSession) SetIpfsMintNum(num *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.SetIpfsMintNum(&_Ipfs.TransactOpts, num)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_Ipfs *IpfsTransactor) SetName(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_Ipfs *IpfsSession) SetName(name_ string) (*types.Transaction, error) {
	return _Ipfs.Contract.SetName(&_Ipfs.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_Ipfs *IpfsTransactorSession) SetName(name_ string) (*types.Transaction, error) {
	return _Ipfs.Contract.SetName(&_Ipfs.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Ipfs *IpfsTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Ipfs *IpfsSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.SetOwner(&_Ipfs.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Ipfs *IpfsTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Ipfs.Contract.SetOwner(&_Ipfs.TransactOpts, owner_)
}

// SetStoreChallenge is a paid mutator transaction binding the contract method 0x58b66e00.
//
// Solidity: function setStoreChallenge(string _storeChallenge) returns()
func (_Ipfs *IpfsTransactor) SetStoreChallenge(opts *bind.TransactOpts, _storeChallenge string) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "setStoreChallenge", _storeChallenge)
}

// SetStoreChallenge is a paid mutator transaction binding the contract method 0x58b66e00.
//
// Solidity: function setStoreChallenge(string _storeChallenge) returns()
func (_Ipfs *IpfsSession) SetStoreChallenge(_storeChallenge string) (*types.Transaction, error) {
	return _Ipfs.Contract.SetStoreChallenge(&_Ipfs.TransactOpts, _storeChallenge)
}

// SetStoreChallenge is a paid mutator transaction binding the contract method 0x58b66e00.
//
// Solidity: function setStoreChallenge(string _storeChallenge) returns()
func (_Ipfs *IpfsTransactorSession) SetStoreChallenge(_storeChallenge string) (*types.Transaction, error) {
	return _Ipfs.Contract.SetStoreChallenge(&_Ipfs.TransactOpts, _storeChallenge)
}

// SetWinner is a paid mutator transaction binding the contract method 0xc35dee34.
//
// Solidity: function setWinner(string _winner) returns()
func (_Ipfs *IpfsTransactor) SetWinner(opts *bind.TransactOpts, _winner string) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "setWinner", _winner)
}

// SetWinner is a paid mutator transaction binding the contract method 0xc35dee34.
//
// Solidity: function setWinner(string _winner) returns()
func (_Ipfs *IpfsSession) SetWinner(_winner string) (*types.Transaction, error) {
	return _Ipfs.Contract.SetWinner(&_Ipfs.TransactOpts, _winner)
}

// SetWinner is a paid mutator transaction binding the contract method 0xc35dee34.
//
// Solidity: function setWinner(string _winner) returns()
func (_Ipfs *IpfsTransactorSession) SetWinner(_winner string) (*types.Transaction, error) {
	return _Ipfs.Contract.SetWinner(&_Ipfs.TransactOpts, _winner)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Ipfs *IpfsTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Ipfs *IpfsSession) Start() (*types.Transaction, error) {
	return _Ipfs.Contract.Start(&_Ipfs.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Ipfs *IpfsTransactorSession) Start() (*types.Transaction, error) {
	return _Ipfs.Contract.Start(&_Ipfs.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Ipfs *IpfsTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Ipfs *IpfsSession) Stop() (*types.Transaction, error) {
	return _Ipfs.Contract.Stop(&_Ipfs.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Ipfs *IpfsTransactorSession) Stop() (*types.Transaction, error) {
	return _Ipfs.Contract.Stop(&_Ipfs.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Ipfs *IpfsTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Ipfs *IpfsSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Transfer(&_Ipfs.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Ipfs *IpfsTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.Transfer(&_Ipfs.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Ipfs *IpfsTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Ipfs *IpfsSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.TransferFrom(&_Ipfs.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Ipfs *IpfsTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Ipfs.Contract.TransferFrom(&_Ipfs.TransactOpts, src, dst, wad)
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

// IpfsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ipfs contract.
type IpfsApprovalIterator struct {
	Event *IpfsApproval // Event containing the contract specifics and raw log

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
func (it *IpfsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsApproval)
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
		it.Event = new(IpfsApproval)
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
func (it *IpfsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsApproval represents a Approval event raised by the Ipfs contract.
type IpfsApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Ipfs *IpfsFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*IpfsApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &IpfsApprovalIterator{contract: _Ipfs.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Ipfs *IpfsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IpfsApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsApproval)
				if err := _Ipfs.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Ipfs *IpfsFilterer) ParseApproval(log types.Log) (*IpfsApproval, error) {
	event := new(IpfsApproval)
	if err := _Ipfs.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IpfsBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Ipfs contract.
type IpfsBurnIterator struct {
	Event *IpfsBurn // Event containing the contract specifics and raw log

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
func (it *IpfsBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsBurn)
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
		it.Event = new(IpfsBurn)
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
func (it *IpfsBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsBurn represents a Burn event raised by the Ipfs contract.
type IpfsBurn struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_Ipfs *IpfsFilterer) FilterBurn(opts *bind.FilterOpts, guy []common.Address) (*IpfsBurnIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "Burn", guyRule)
	if err != nil {
		return nil, err
	}
	return &IpfsBurnIterator{contract: _Ipfs.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_Ipfs *IpfsFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IpfsBurn, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "Burn", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsBurn)
				if err := _Ipfs.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_Ipfs *IpfsFilterer) ParseBurn(log types.Log) (*IpfsBurn, error) {
	event := new(IpfsBurn)
	if err := _Ipfs.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IpfsLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Ipfs contract.
type IpfsLogSetAuthorityIterator struct {
	Event *IpfsLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *IpfsLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsLogSetAuthority)
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
		it.Event = new(IpfsLogSetAuthority)
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
func (it *IpfsLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsLogSetAuthority represents a LogSetAuthority event raised by the Ipfs contract.
type IpfsLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Ipfs *IpfsFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*IpfsLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &IpfsLogSetAuthorityIterator{contract: _Ipfs.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Ipfs *IpfsFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *IpfsLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsLogSetAuthority)
				if err := _Ipfs.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// ParseLogSetAuthority is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Ipfs *IpfsFilterer) ParseLogSetAuthority(log types.Log) (*IpfsLogSetAuthority, error) {
	event := new(IpfsLogSetAuthority)
	if err := _Ipfs.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IpfsLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Ipfs contract.
type IpfsLogSetOwnerIterator struct {
	Event *IpfsLogSetOwner // Event containing the contract specifics and raw log

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
func (it *IpfsLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsLogSetOwner)
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
		it.Event = new(IpfsLogSetOwner)
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
func (it *IpfsLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsLogSetOwner represents a LogSetOwner event raised by the Ipfs contract.
type IpfsLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Ipfs *IpfsFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*IpfsLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &IpfsLogSetOwnerIterator{contract: _Ipfs.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Ipfs *IpfsFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *IpfsLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsLogSetOwner)
				if err := _Ipfs.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Ipfs *IpfsFilterer) ParseLogSetOwner(log types.Log) (*IpfsLogSetOwner, error) {
	event := new(IpfsLogSetOwner)
	if err := _Ipfs.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IpfsMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Ipfs contract.
type IpfsMintIterator struct {
	Event *IpfsMint // Event containing the contract specifics and raw log

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
func (it *IpfsMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsMint)
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
		it.Event = new(IpfsMint)
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
func (it *IpfsMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsMint represents a Mint event raised by the Ipfs contract.
type IpfsMint struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_Ipfs *IpfsFilterer) FilterMint(opts *bind.FilterOpts, guy []common.Address) (*IpfsMintIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "Mint", guyRule)
	if err != nil {
		return nil, err
	}
	return &IpfsMintIterator{contract: _Ipfs.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_Ipfs *IpfsFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IpfsMint, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "Mint", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsMint)
				if err := _Ipfs.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_Ipfs *IpfsFilterer) ParseMint(log types.Log) (*IpfsMint, error) {
	event := new(IpfsMint)
	if err := _Ipfs.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IpfsSetChallengeIterator is returned from FilterSetChallenge and is used to iterate over the raw logs and unpacked data for SetChallenge events raised by the Ipfs contract.
type IpfsSetChallengeIterator struct {
	Event *IpfsSetChallenge // Event containing the contract specifics and raw log

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
func (it *IpfsSetChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsSetChallenge)
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
		it.Event = new(IpfsSetChallenge)
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
func (it *IpfsSetChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsSetChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsSetChallenge represents a SetChallenge event raised by the Ipfs contract.
type IpfsSetChallenge struct {
	Cge    string
	CgeSeq *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetChallenge is a free log retrieval operation binding the contract event 0x9bf2941793efdd6f6157187773c27e7ead5a4c62a79e7434488ef13d3b21c856.
//
// Solidity: event SetChallenge(string cge, uint256 cgeSeq)
func (_Ipfs *IpfsFilterer) FilterSetChallenge(opts *bind.FilterOpts) (*IpfsSetChallengeIterator, error) {

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "SetChallenge")
	if err != nil {
		return nil, err
	}
	return &IpfsSetChallengeIterator{contract: _Ipfs.contract, event: "SetChallenge", logs: logs, sub: sub}, nil
}

// WatchSetChallenge is a free log subscription operation binding the contract event 0x9bf2941793efdd6f6157187773c27e7ead5a4c62a79e7434488ef13d3b21c856.
//
// Solidity: event SetChallenge(string cge, uint256 cgeSeq)
func (_Ipfs *IpfsFilterer) WatchSetChallenge(opts *bind.WatchOpts, sink chan<- *IpfsSetChallenge) (event.Subscription, error) {

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "SetChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsSetChallenge)
				if err := _Ipfs.contract.UnpackLog(event, "SetChallenge", log); err != nil {
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

// ParseSetChallenge is a log parse operation binding the contract event 0x9bf2941793efdd6f6157187773c27e7ead5a4c62a79e7434488ef13d3b21c856.
//
// Solidity: event SetChallenge(string cge, uint256 cgeSeq)
func (_Ipfs *IpfsFilterer) ParseSetChallenge(log types.Log) (*IpfsSetChallenge, error) {
	event := new(IpfsSetChallenge)
	if err := _Ipfs.contract.UnpackLog(event, "SetChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IpfsSetStoreChallengeIterator is returned from FilterSetStoreChallenge and is used to iterate over the raw logs and unpacked data for SetStoreChallenge events raised by the Ipfs contract.
type IpfsSetStoreChallengeIterator struct {
	Event *IpfsSetStoreChallenge // Event containing the contract specifics and raw log

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
func (it *IpfsSetStoreChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsSetStoreChallenge)
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
		it.Event = new(IpfsSetStoreChallenge)
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
func (it *IpfsSetStoreChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsSetStoreChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsSetStoreChallenge represents a SetStoreChallenge event raised by the Ipfs contract.
type IpfsSetStoreChallenge struct {
	Cge    string
	CgeSeq *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetStoreChallenge is a free log retrieval operation binding the contract event 0x7e93e5083d0db25c0944a3c885d3f98d9e8457ee5c05bef46721333cd4ee7c6d.
//
// Solidity: event SetStoreChallenge(string cge, uint256 cgeSeq)
func (_Ipfs *IpfsFilterer) FilterSetStoreChallenge(opts *bind.FilterOpts) (*IpfsSetStoreChallengeIterator, error) {

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "SetStoreChallenge")
	if err != nil {
		return nil, err
	}
	return &IpfsSetStoreChallengeIterator{contract: _Ipfs.contract, event: "SetStoreChallenge", logs: logs, sub: sub}, nil
}

// WatchSetStoreChallenge is a free log subscription operation binding the contract event 0x7e93e5083d0db25c0944a3c885d3f98d9e8457ee5c05bef46721333cd4ee7c6d.
//
// Solidity: event SetStoreChallenge(string cge, uint256 cgeSeq)
func (_Ipfs *IpfsFilterer) WatchSetStoreChallenge(opts *bind.WatchOpts, sink chan<- *IpfsSetStoreChallenge) (event.Subscription, error) {

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "SetStoreChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsSetStoreChallenge)
				if err := _Ipfs.contract.UnpackLog(event, "SetStoreChallenge", log); err != nil {
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

// ParseSetStoreChallenge is a log parse operation binding the contract event 0x7e93e5083d0db25c0944a3c885d3f98d9e8457ee5c05bef46721333cd4ee7c6d.
//
// Solidity: event SetStoreChallenge(string cge, uint256 cgeSeq)
func (_Ipfs *IpfsFilterer) ParseSetStoreChallenge(log types.Log) (*IpfsSetStoreChallenge, error) {
	event := new(IpfsSetStoreChallenge)
	if err := _Ipfs.contract.UnpackLog(event, "SetStoreChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IpfsSetWinnerIterator is returned from FilterSetWinner and is used to iterate over the raw logs and unpacked data for SetWinner events raised by the Ipfs contract.
type IpfsSetWinnerIterator struct {
	Event *IpfsSetWinner // Event containing the contract specifics and raw log

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
func (it *IpfsSetWinnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsSetWinner)
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
		it.Event = new(IpfsSetWinner)
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
func (it *IpfsSetWinnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsSetWinnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsSetWinner represents a SetWinner event raised by the Ipfs contract.
type IpfsSetWinner struct {
	Cge    string
	CgeSeq *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetWinner is a free log retrieval operation binding the contract event 0x903ea6ee4909f001da5cb3427de2578796f6b0a50dd2f05c4053cf92a70b9e1d.
//
// Solidity: event SetWinner(string cge, uint256 cgeSeq)
func (_Ipfs *IpfsFilterer) FilterSetWinner(opts *bind.FilterOpts) (*IpfsSetWinnerIterator, error) {

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "SetWinner")
	if err != nil {
		return nil, err
	}
	return &IpfsSetWinnerIterator{contract: _Ipfs.contract, event: "SetWinner", logs: logs, sub: sub}, nil
}

// WatchSetWinner is a free log subscription operation binding the contract event 0x903ea6ee4909f001da5cb3427de2578796f6b0a50dd2f05c4053cf92a70b9e1d.
//
// Solidity: event SetWinner(string cge, uint256 cgeSeq)
func (_Ipfs *IpfsFilterer) WatchSetWinner(opts *bind.WatchOpts, sink chan<- *IpfsSetWinner) (event.Subscription, error) {

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "SetWinner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsSetWinner)
				if err := _Ipfs.contract.UnpackLog(event, "SetWinner", log); err != nil {
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

// ParseSetWinner is a log parse operation binding the contract event 0x903ea6ee4909f001da5cb3427de2578796f6b0a50dd2f05c4053cf92a70b9e1d.
//
// Solidity: event SetWinner(string cge, uint256 cgeSeq)
func (_Ipfs *IpfsFilterer) ParseSetWinner(log types.Log) (*IpfsSetWinner, error) {
	event := new(IpfsSetWinner)
	if err := _Ipfs.contract.UnpackLog(event, "SetWinner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IpfsStartIterator is returned from FilterStart and is used to iterate over the raw logs and unpacked data for Start events raised by the Ipfs contract.
type IpfsStartIterator struct {
	Event *IpfsStart // Event containing the contract specifics and raw log

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
func (it *IpfsStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsStart)
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
		it.Event = new(IpfsStart)
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
func (it *IpfsStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsStart represents a Start event raised by the Ipfs contract.
type IpfsStart struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStart is a free log retrieval operation binding the contract event 0x1b55ba3aa851a46be3b365aee5b5c140edd620d578922f3e8466d2cbd96f954b.
//
// Solidity: event Start()
func (_Ipfs *IpfsFilterer) FilterStart(opts *bind.FilterOpts) (*IpfsStartIterator, error) {

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return &IpfsStartIterator{contract: _Ipfs.contract, event: "Start", logs: logs, sub: sub}, nil
}

// WatchStart is a free log subscription operation binding the contract event 0x1b55ba3aa851a46be3b365aee5b5c140edd620d578922f3e8466d2cbd96f954b.
//
// Solidity: event Start()
func (_Ipfs *IpfsFilterer) WatchStart(opts *bind.WatchOpts, sink chan<- *IpfsStart) (event.Subscription, error) {

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsStart)
				if err := _Ipfs.contract.UnpackLog(event, "Start", log); err != nil {
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

// ParseStart is a log parse operation binding the contract event 0x1b55ba3aa851a46be3b365aee5b5c140edd620d578922f3e8466d2cbd96f954b.
//
// Solidity: event Start()
func (_Ipfs *IpfsFilterer) ParseStart(log types.Log) (*IpfsStart, error) {
	event := new(IpfsStart)
	if err := _Ipfs.contract.UnpackLog(event, "Start", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IpfsStopIterator is returned from FilterStop and is used to iterate over the raw logs and unpacked data for Stop events raised by the Ipfs contract.
type IpfsStopIterator struct {
	Event *IpfsStop // Event containing the contract specifics and raw log

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
func (it *IpfsStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsStop)
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
		it.Event = new(IpfsStop)
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
func (it *IpfsStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsStop represents a Stop event raised by the Ipfs contract.
type IpfsStop struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStop is a free log retrieval operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Ipfs *IpfsFilterer) FilterStop(opts *bind.FilterOpts) (*IpfsStopIterator, error) {

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return &IpfsStopIterator{contract: _Ipfs.contract, event: "Stop", logs: logs, sub: sub}, nil
}

// WatchStop is a free log subscription operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Ipfs *IpfsFilterer) WatchStop(opts *bind.WatchOpts, sink chan<- *IpfsStop) (event.Subscription, error) {

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsStop)
				if err := _Ipfs.contract.UnpackLog(event, "Stop", log); err != nil {
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

// ParseStop is a log parse operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Ipfs *IpfsFilterer) ParseStop(log types.Log) (*IpfsStop, error) {
	event := new(IpfsStop)
	if err := _Ipfs.contract.UnpackLog(event, "Stop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// IpfsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ipfs contract.
type IpfsTransferIterator struct {
	Event *IpfsTransfer // Event containing the contract specifics and raw log

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
func (it *IpfsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsTransfer)
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
		it.Event = new(IpfsTransfer)
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
func (it *IpfsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsTransfer represents a Transfer event raised by the Ipfs contract.
type IpfsTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Ipfs *IpfsFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*IpfsTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Ipfs.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &IpfsTransferIterator{contract: _Ipfs.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Ipfs *IpfsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IpfsTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Ipfs.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsTransfer)
				if err := _Ipfs.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Ipfs *IpfsFilterer) ParseTransfer(log types.Log) (*IpfsTransfer, error) {
	event := new(IpfsTransfer)
	if err := _Ipfs.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
