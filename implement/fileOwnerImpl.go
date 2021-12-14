package implement

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-ipfs-auth/auth-source-eth/contract/ipfs"
	"github.com/ipfs/go-ipfs-auth/standard/model"
	"math/big"
)

func (a *peerImpl) AddFile(info model.IpfsFileInfo) error {
	ctx := context.Background()

	blockNum := big.NewInt(info.StoreDays * 24 * 60 * 60)

	cli, scTokenContra, err := GetSCTokenContract(a.Client.SocketUrl, a.ContractMap[contractToken].ContractAddr)
	if err != nil {
		return err
	}
	defer cli.Close()

	// todo 检查余额
	// 调用SCToken合约的Approval方法 给与权限
	var approveFunc ExecuteFunc = func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return scTokenContra.Approve(opts, common.HexToAddress(a.ContractMap[contractIpfs].ContractAddr), blockNum)
	}
	err = a.ExecuteTransact(ctx, cli, approveFunc)
	if err != nil {
		return err
	}

	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.AddFile(opts, uid, info.Cid, big.NewInt(info.Size), blockNum)
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

func (a *peerImpl) RechargeFile(cid string, days int64) error {
	ctx := context.Background()
	blockNum := big.NewInt(days * 24 * 60 * 60)

	cli, scTokenContra, err := GetSCTokenContract(a.Client.SocketUrl, a.ContractMap[contractToken].ContractAddr)
	if err != nil {
		return err
	}
	defer cli.Close()
	// todo 查询余额是否足够

	// 调用SCToken合约的Approval方法 给与权限
	var approveFunc ExecuteFunc = func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return scTokenContra.Approve(opts, common.HexToAddress(a.ContractMap[contractIpfs].ContractAddr), blockNum)
	}
	err = a.ExecuteTransact(ctx, cli, approveFunc)
	if err != nil {
		return err
	}

	// 执行充值方法
	var rechargeFunc ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.RechargeFile(opts, uid, cid, blockNum)
	}
	// 执行交易
	return a.ExecuteIpfsTransact(ctx, rechargeFunc)
}
