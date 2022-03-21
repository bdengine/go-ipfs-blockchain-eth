package implement

import (
	"context"
	"fmt"
	"github.com/bdengine/go-ipfs-blockchain-eth/contract/ipfs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/prometheus/common/log"
	"math/big"
	"time"
)

const (
	defaultGasLimit = 300000
	txStatusFail    = 0
)

type ExecuteIpfsFunc func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error)

// GenTransactOpts 生成默认交易配置项
func (p *peerImpl) GenTransactOpts(ctx context.Context, sCli *ethclient.Client, gasLimit uint64) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(p.priKey, p.chainId)
	if err != nil {
		return nil, err
	}
	nonce, err := sCli.PendingNonceAt(ctx, p.address)
	if err != nil {
		return nil, err
	}
	gasPrice, err := sCli.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)

	if gasLimit == 0 {
		gasLimit = defaultGasLimit
	}
	opts.GasLimit = gasLimit // in units
	opts.GasPrice = gasPrice
	return opts, nil
}

// ExecuteIpfsTransact 执行交易,等待返回成功与否
func (p *peerImpl) ExecuteIpfsTransact(ctx context.Context, f ExecuteIpfsFunc) error {
	// 获取客户端和合约实例
	uid := uuid.New().String()
	tx, err := p.syncExecuteIpfsContract(ctx, f, uid)
	if err != nil {
		return err
	}
	//ctx, _ = context.WithTimeout(ctx, p.Variable.RequestTimeout*time.Second)
	// todo 日志
	log.Debugf("%v交易结果查询：", tx.Hash())
	receipt := p.WatchResult(tx.Hash())
	if receipt.Status == txStatusFail {
		return fmt.Errorf("交易失败")
	}
	return nil
}

func (p *peerImpl) syncExecuteIpfsContract(ctx context.Context, f ExecuteIpfsFunc, uid string) (*types.Transaction, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 生成配置项
	opts, err := p.GenTransactOpts(ctx, p.httpClient, p.config.Variable.GasLimit)
	if err != nil {
		return nil, err
	}
	// 执行交易方法
	return f(uid, p.ipfsContract, opts)
}


func (p *peerImpl) WatchResult(hash common.Hash) *types.Receipt {
	t0 := time.Now()
	ch := make(chan *types.Receipt)
	p.txResultRegister.Store(hash.String(),ch)
	p.transHahCh<- hash
	log.Debugf("交易%s确认时长:%v",hash.String(),time.Now().Sub(t0))
	return <-ch
}

