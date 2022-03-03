// Package implement /github.com/ipfs/go-ipfs-auth/standard  interface各接口实现
package implement

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/bdengine/go-ipfs-blockchain-eth/contract/ipfs"
	"github.com/bdengine/go-ipfs-blockchain-standard/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"sync"
	"time"
)

const (
	ConfigFileName = "/cosmosConfig"
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
	GasLimit       uint64
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
	lock    sync.Locker

	socketClient *ethclient.Client
	ipfsContract *ipfs.Ipfs

	httpClient *ethclient.Client
}

func NewConfig(httpUrl, socketUrl, centralServerUrl, ipfsAddr, priKey string, chainID int64, timeout time.Duration, gasLimit uint64) (*config, error) {
	if timeout == 0 {
		timeout = defaultTimeout
	}
	if gasLimit == 0 {
		gasLimit = defaultGasLimit
	}
	// todo 检查必要的参数和格式

	return &config{
		Client: clientInfo{
			HttpUrl:   httpUrl,
			SocketUrl: socketUrl,
		},
		CentralServerUrl: centralServerUrl,
		ContractMap:      map[string]contractInfo{contractIpfs: {ipfsAddr}},
		Chain:            chainInfo{ChainId: chainID},
		Variable: configInfo{
			RequestTimeout: timeout * time.Second,
			GasLimit:       gasLimit,
		},
		ID: Identity{PriKey: priKey},
	}, nil
}

// NewApi 初始化一个peerImpl实例
func NewApi(configRoot string, peerId string) (*peerImpl, error) {
	fileName := configRoot + ConfigFileName
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
		panic(err)
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
		marshal, err := json.MarshalIndent(cfg, "", "\t")
		if err != nil {
			return nil, err
		}

		err = ioutil.WriteFile(fileName, marshal, 0600)
		if err != nil {
			return nil, err
		}
	}

	a.config = &cfg
	httpClient, err := ethclient.Dial(a.Client.HttpUrl)
	if err != nil {
		return nil, err
	}
	socketClient, err := ethclient.Dial(a.Client.SocketUrl)
	if err != nil {
		return nil, err
	}

	a.socketClient = socketClient
	a.httpClient = httpClient
	a.lock = &sync.Mutex{}

	ipfsContra, err := ipfs.NewIpfs(common.HexToAddress(a.ContractMap[contractIpfs].ContractAddr), socketClient)
	if err != nil {
		return nil, err
	}
	a.ipfsContract = ipfsContra
	return &a, err
}

func checkConfig(a *config) error {
	if a.Client.HttpUrl == "" || a.Client.SocketUrl == "" {
		return fmt.Errorf("配置文件Client信息不完整")
	}

	if a.ContractMap[contractIpfs].ContractAddr == "" {
		return fmt.Errorf("配置文件ContractMap不完整")
	}

	if a.Chain.ChainId == 0 {
		return fmt.Errorf("配置文件Chain不完整")
	}

	if a.Variable.RequestTimeout == 0 {
		a.Variable.RequestTimeout = defaultTimeout
	}
	if a.Variable.GasLimit == 0 {
		a.Variable.GasLimit = defaultGasLimit
	}

	if a.ID.Pid == "" || a.ID.PriKey == "" {
		return fmt.Errorf("配置文件ID不完整")
	}
	return nil
}

func (p *peerImpl) InitPeer(peer model.CorePeer) error {
	ctx := context.Background()
	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		/*p, err := contract.AddrPeerMap(nil, p.address)
		if err != nil {
			return nil, err
		}
		if p.Valid {
			return nil, fmt.Errorf("节点存在")
		}*/
		return contract.AddPeer(opts, uid, peer.PeerId, peer.Addresses)
	}
	// 执行交易
	return p.ExecuteIpfsTransact(ctx, f)
}

func (p *peerImpl) RemovePeer() error {
	ctx := context.Background()
	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.RemovePeer(opts, uid)
	}
	// 执行交易
	return p.ExecuteIpfsTransact(ctx, f)
}

func (p *peerImpl) DaemonPeer() error {
	return nil
}

func (p *peerImpl) GetSocketClient() (*ethclient.Client, error) {
	if p.socketClient == nil {
		return nil, fmt.Errorf("nil client")
	}
	return p.socketClient, nil
}

func (p *peerImpl) GetHttpClient() (*ethclient.Client, error) {
	if p.httpClient == nil {
		return nil, fmt.Errorf("nil client")
	}
	return p.httpClient, nil
}
