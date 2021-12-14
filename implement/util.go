package implement

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/google/uuid"
	"github.com/ipfs/go-ipfs-auth/auth-source-eth/contract/ipfs"
	"github.com/ipfs/go-ipfs-auth/auth-source-eth/contract/scToken"
	"github.com/prometheus/common/log"
	"math/big"
	"time"
)

type ExecuteFunc func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error)

func GetIpfsContract(url string, addr string) (*ethclient.Client, *ipfs.Ipfs, error) {
	httpClient, err := ethclient.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	// 获取合约
	contract, err := ipfs.NewIpfs(common.HexToAddress(addr), httpClient)
	if err != nil {
		return nil, nil, err
	}
	return httpClient, contract, nil
}

func GetSCTokenContract(url string, addr string) (*ethclient.Client, *scToken.ScToken, error) {
	httpClient, err := ethclient.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	// 获取合约
	contract, err := scToken.NewScToken(common.HexToAddress(addr), httpClient)
	if err != nil {
		return nil, nil, err
	}
	return httpClient, contract, nil
}

// GenTransactOpts 生成默认交易配置项
func (a *peerImpl) GenTransactOpts(ctx context.Context, sCli *ethclient.Client) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(a.priKey, a.chainId)
	if err != nil {
		return nil, err
	}
	nonce, err := sCli.PendingNonceAt(ctx, a.address)
	if err != nil {
		return nil, err
	}
	gasPrice, err := sCli.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)
	opts.GasLimit = uint64(300000) // in units
	opts.GasPrice = gasPrice
	return opts, nil
}

// ExecuteIpfsTransact 执行交易,等待返回成功与否
func (a *peerImpl) ExecuteIpfsTransact(ctx context.Context, f ExecuteFunc) error {
	// 获取客户端和合约实例
	sCli, contract, err := GetIpfsContract(a.Client.SocketUrl, a.ContractMap[contractIpfs].ContractAddr)
	if err != nil {
		return err
	}
	defer sCli.Close()
	uid := uuid.New().String()
	sch := make(chan *ipfs.IpfsSuccess)
	sub, err := contract.IpfsFilterer.WatchSuccess(nil, sch, []string{})
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	// 生成配置项
	opts, err := a.GenTransactOpts(ctx, sCli)
	if err != nil {
		return err
	}
	// 执行交易方法
	tx, err := f(uid, contract, opts)
	if err != nil {
		return err
	}

	ctx, _ = context.WithTimeout(ctx, a.Variable.RequestTimeout*time.Second)

	// todo 日志
	log.Debugf("%v交易结果查询：", tx.Hash())
	return WaitResult(ctx, sCli, sch, sub, tx.Hash(), uid)
}

func WaitResult(ctx context.Context, sCli *ethclient.Client, sch chan *ipfs.IpfsSuccess, sub event.Subscription, txId common.Hash, uid string) error {
	tick := time.Tick(2 * time.Second)
	var err error
	for {
		select {
		case <-sch:
			log.Debugf("成功回调:%s", uid)
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
			} else if receipt.Status == 0 {
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
			if receipt.Status == 0 {
				return fmt.Errorf("交易失败")
			}
			log.Debugln("交易成功")
			return nil
		}
	}
}
