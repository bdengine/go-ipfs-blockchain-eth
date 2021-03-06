package implement

import (
	"context"
	"fmt"
	"github.com/bdengine/go-ipfs-blockchain-eth/contract/ipfs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
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
	sch := make(chan *ipfs.IpfsSuccess)
	sub, err := p.ipfsContract.IpfsFilterer.WatchSuccess(nil, sch, []string{uid})
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	tx, err := p.syncExecuteIpfsContract(ctx, f, uid)
	if err != nil {
		return err
	}
	ctx, _ = context.WithTimeout(ctx, p.Variable.RequestTimeout*time.Second)
	// todo 日志
	log.Debugf("%v交易结果查询：", tx.Hash())
	return waitResult(ctx, p.socketClient, sch, sub, tx.Hash(), uid)
}

func (p *peerImpl) syncExecuteIpfsContract(ctx context.Context, f ExecuteIpfsFunc, uid string) (*types.Transaction, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 生成配置项
	opts, err := p.GenTransactOpts(ctx, p.socketClient, p.config.Variable.GasLimit)
	if err != nil {
		return nil, err
	}
	// 执行交易方法
	return f(uid, p.ipfsContract, opts)
}

func waitResult(ctx context.Context, sCli *ethclient.Client, sch chan *ipfs.IpfsSuccess, sub event.Subscription, txId common.Hash, uid string) error {
	tick := time.Tick(2 * time.Second)
	var err error
	for {
		select {
		case s := <-sch:
			log.Debugf("成功回调:%s", uid)
			fmt.Println(s.Raw.BlockNumber)
			return nil
		case err = <-sub.Err():
			return err
		case <-tick:
			log.Debugln("未收到成功事件,轮询回执")
			receipt, err := sCli.TransactionReceipt(context.Background(), txId)
			if err != nil {
				if err.Error() == "not found" {
					log.Debugf("交易结果不确定,继续轮询")
				}
			} else if receipt.Status == txStatusFail {
				fmt.Printf("%+v\n", receipt)
				return fmt.Errorf("交易失败")
			} else {
				log.Debugln("交易成功")
				return nil
			}
		case <-ctx.Done():
			log.Debugln("交易超时，最后一次查询回执")
			receipt, err := sCli.TransactionReceipt(context.Background(), txId)
			if err != nil {
				if err.Error() == "not found" {
					return fmt.Errorf("交易执行结果不确定,请稍候查询")
				}
				return err
			}
			if receipt.Status == txStatusFail {
				fmt.Printf("%+v\n", receipt)
				return fmt.Errorf("交易失败")
			}
			log.Debugln("交易成功")
			return nil
		}
	}
}
