package implement

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-ipfs-auth/auth-source-eth/contract/ipfs"
	"github.com/ipfs/go-ipfs-auth/standard/model"
	"math/big"
	"time"
)

func (a *peerImpl) AddFile(info model.IpfsFileInfo) error {
	ctx := context.Background()
	// 调用SCToken合约的Approval方法
	cli, scTokenContra, err := GetSCTokenContract(a.Client.SocketUrl, a.ContractMap[contractToken].ContractAddr)
	if err != nil {
		return err
	}
	defer cli.Close()

	ipfsContr, err := ipfs.NewIpfs(common.HexToAddress(a.ContractMap[contractIpfs].ContractAddr), cli)
	if err != nil {
		return err
	}

	// 查询单价
	price, err := ipfsContr.Price(nil)
	if err != nil {
		return err
	}
	// 计算应付价格
	wad := big.NewInt(price.Int64() * info.StoreDays * info.Size)
	opts, err := a.GenTransactOpts(ctx, cli, defaultGasLimit)
	if err != nil {
		return err
	}
	// 查询余额是否足够
	balance, err := scTokenContra.BalanceOf(nil, a.address)
	if err != nil {
		return err
	}

	if balance.Int64() < wad.Int64() {
		return fmt.Errorf("余额不足，当前余额：%v，存储当前文件需要：%v", balance.Int64(), wad.Int64())
	}

	// 给与权限
	tx, err := scTokenContra.Approve(opts, common.HexToAddress(a.ContractMap[contractIpfs].ContractAddr), wad)
	ctx1, _ := context.WithTimeout(ctx, a.Variable.RequestTimeout*time.Second)
	// 等待转账权限结果
	err = waitResultWithoutEvent(ctx1, cli, tx.Hash())
	if err != nil {
		return err
	}

	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {

		return contract.AddFile(opts, uid, info.Cid, big.NewInt(info.Size), wad)
	}

	// 执行交易
	return a.ExecuteIpfsTransact(ctx, f)
}

func (a *peerImpl) DeleteFile(cid string) error {
	ctx := context.Background()

	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.RemoveFile(opts, uid, cid)
	}
	// 执行交易
	return a.ExecuteIpfsTransact(ctx, f)
}
