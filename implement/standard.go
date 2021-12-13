package implement

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-ipfs-auth/auth-source-eth/contract/ipfs"
	"github.com/ipfs/go-ipfs-auth/standard/model"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

const (
	configFileName = "/cosmosConfig"
	contractIpfs   = "ipfs"
	contractToken  = "token"
)

type clientInfo struct {
	HttpUrl   string
	SocketUrl string
}

type contractInfo struct {
	ContractAddr string
}

type chainInfo struct {
	ChainId int64
	chainId *big.Int
}

type Identity struct {
	Pid     string
	PriKey  string
	priKey  *ecdsa.PrivateKey
	address common.Address
}

type configInfo struct {
	RequestTimeout time.Duration
}

type api struct {
	Client           clientInfo
	CentralServerUrl string
	ContractMap      map[string]contractInfo
	Chain            chainInfo
	Config           configInfo
	ID               Identity
}

func NewApi(configRoot string, peerId string) (*api, error) {
	fileName := configRoot + configFileName
	readFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var a api
	err = json.Unmarshal(readFile, &a)
	if err != nil {
		return nil, err
	}

	// 检查配置项是否完整  todo 检查配置项是否满足条件
	if a.Chain.ChainId == 0 || a.ID.Pid == "" {
		return nil, fmt.Errorf("配置文件配置项不完整")
	}

	_, exist := a.ContractMap[contractIpfs]
	if !exist {
		return nil, fmt.Errorf("配置文件配置项不完整")
	}
	_, exist = a.ContractMap[contractToken]
	if !exist {
		return nil, fmt.Errorf("配置文件配置项不完整")
	}

	priKey, err := crypto.HexToECDSA(a.ID.PriKey)
	if err != nil {
		return nil, err
	}
	pubKey := priKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("公钥无法断言为类型 *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*pubKeyECDSA)

	a.ID.priKey = priKey
	a.ID.address = address
	a.Chain.chainId = big.NewInt(a.Chain.ChainId)
	// 更新pid
	if a.ID.Pid != peerId {
		a.ID.Pid = peerId
		marshal, err := json.MarshalIndent(a, "", "  ")
		if err != nil {
			return nil, err
		}
		err = os.Remove(fileName)
		if err != nil {
			return nil, err
		}

		err = ioutil.WriteFile(fileName, marshal, 0600)
		if err != nil {
			return nil, err
		}
	}

	return &a, err
}

func (a *api) ReportContribute(num int64) error {
	return nil
}

func (a *api) InitPeer(peer model.CorePeer) error {
	ctx := context.Background()
	var f executeFunc = func(uid string, contract *ipfs.Contracts, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.AddPeer(opts, uid, ipfs.IPFSPeer{
			PeerId:      peer.PeerId,
			AddressList: peer.Addresses,
			Valid:       true,
		})
	}
	// 执行交易
	return a.ExecuteTransact(ctx, f)
}

func (a *api) RemovePeer(pid string) error {
	ctx := context.Background()
	var f executeFunc = func(uid string, contract *ipfs.Contracts, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.RemovePeer(opts, uid)
	}
	// 执行交易
	return a.ExecuteTransact(ctx, f)
}

func (a *api) DaemonPeer() error {
	return nil
}

func (a *api) ApplyLocal(cid string) error {
	return nil
}

func (a *api) ApplyRemote(cid string) error {
	return nil
}

func (a *api) AddFile(info model.IpfsFileInfo) error {
	ctx := context.Background()
	var f executeFunc = func(uid string, contract *ipfs.Contracts, opts *bind.TransactOpts) (*types.Transaction, error) {

		return contract.AddFile(opts, uid, info.Cid, big.NewInt(0), big.NewInt(0))
	}
	// 执行交易
	return a.ExecuteTransact(ctx, f)
}

func (a *api) DeleteFile(cid string) error {
	ctx := context.Background()

	var f executeFunc = func(uid string, contract *ipfs.Contracts, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.RemoveFile(opts, uid, cid)
	}
	// 执行交易
	return a.ExecuteTransact(ctx, f)
}

func (a *api) GetPeerList() ([]model.CorePeer, error) {
	httpClient, contract, err := GetIpfsContract(a.Client.HttpUrl, a.ContractMap[contractIpfs].ContractAddr)
	if err != nil {
		return nil, err
	}
	defer httpClient.Close()

	list, err := contract.GetPeerList(nil, big.NewInt(-1))
	res := make([]model.CorePeer, len(list))
	for i, peer := range list {
		res[i] = model.CorePeer{
			PeerId:    peer.PeerId,
			Addresses: peer.AddressList,
		}
	}

	return res, nil
}

func (a *api) GetUserCode() (string, error) {
	return a.ID.address.String(), nil
}

func (a *api) GetPeer(id string) (model.CorePeer, error) {
	res := model.CorePeer{}
	return res, nil
}

func (a *api) GetChallenge() (string, error) {
	return "", nil
}

func (a *api) Mining([]model.IpfsMining) error {
	return nil
}
