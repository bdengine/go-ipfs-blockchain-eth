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
	defaultTimeout = 30
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
}

type Identity struct {
	Pid    string
	PriKey string
}

type configInfo struct {
	RequestTimeout time.Duration
}

type config struct {
	Client           clientInfo
	CentralServerUrl string
	ContractMap      map[string]contractInfo
	Chain            chainInfo
	Variable         configInfo
	ID               Identity
}

type peerImpl struct {
	*config
	priKey  *ecdsa.PrivateKey
	address common.Address
	chainId *big.Int
}

func NewApi(configRoot string, peerId string) (*peerImpl, error) {
	fileName := configRoot + configFileName
	readFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var cfg config
	var a peerImpl
	err = json.Unmarshal(readFile, &cfg)
	if err != nil {
		return nil, err
	}

	// 检查配置项是否完整
	err = checkConfig(&cfg)
	if err != nil {
		return nil, err
	}

	// 设置参数
	priKey, err := crypto.HexToECDSA(cfg.ID.PriKey)
	if err != nil {
		return nil, err
	}
	pubKey := priKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("公钥无法断言为类型 *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*pubKeyECDSA)

	a.priKey = priKey
	a.address = address
	a.chainId = big.NewInt(cfg.Chain.ChainId)
	// 更新pid
	if cfg.ID.Pid != peerId {
		cfg.ID.Pid = peerId
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

	a.config = &cfg
	return &a, err
}

func checkConfig(a *config) error {
	if a.Client.HttpUrl == "" || a.Client.SocketUrl == "" {
		return fmt.Errorf("配置文件Client信息不完整")
	}

	if a.ContractMap[contractIpfs].ContractAddr == "" || a.ContractMap[contractToken].ContractAddr == "" {
		return fmt.Errorf("配置文件ContractMap不完整")
	}

	if a.Chain.ChainId == 0 {
		return fmt.Errorf("配置文件Chain不完整")
	}

	if a.Variable.RequestTimeout == 0 {
		a.Variable.RequestTimeout = defaultTimeout
	}

	if a.ID.Pid == "" || a.ID.PriKey == "" {
		return fmt.Errorf("配置文件ID不完整")
	}
	return nil
}

func (a *peerImpl) InitPeer(peer model.CorePeer) error {
	ctx := context.Background()
	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.AddPeer(opts, uid, peer.PeerId, peer.Addresses)
	}
	// 执行交易
	return a.ExecuteIpfsTransact(ctx, f)
}

func (a *peerImpl) RemovePeer() error {
	ctx := context.Background()
	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.RemovePeer(opts, uid)
	}
	// 执行交易
	return a.ExecuteIpfsTransact(ctx, f)
}

func (a *peerImpl) DaemonPeer() error {
	return nil
}
