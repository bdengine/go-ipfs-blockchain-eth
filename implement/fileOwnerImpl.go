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
)

func (a *peerImpl) AddFile(info model.IpfsFileInfo) error {
	ctx := context.Background()

	blockNum := big.NewInt(info.StoreDays * 24 * 60 * 60)

	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		file, err := contract.FileMap(nil, info.Cid)
		if err != nil {
			return nil, err
		}
		if file.Owner != common.HexToAddress("") {
			return nil, fmt.Errorf("文件已存在")
		}
		var owner common.Address
		if info.Owner == "self" || info.Owner == "" {
			owner = a.address
		} else {
			owner = common.HexToAddress(info.Owner)
		}
		return contract.AddFile(opts, uid, info.Cid, big.NewInt(info.Size), blockNum, owner)
	}

	// 执行交易
	return a.ExecuteIpfsTransact(ctx, f)
}

func (a *peerImpl) DeleteFile(cid string) error {
	ctx := context.Background()

	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		file, err := contract.FileMap(nil, cid)
		if err != nil {
			return nil, err
		}
		if file.Owner == common.HexToAddress("") {
			return nil, fmt.Errorf("文件不存在")
		}
		return contract.RemoveFile(opts, uid, cid)
	}
	// 执行交易
	return a.ExecuteIpfsTransact(ctx, f)
}

func (a *peerImpl) RechargeFile(cid string, days int64) error {
	ctx := context.Background()
	blockNum := big.NewInt(days * 24 * 60 * 60)

	// 执行充值方法
	var rechargeFunc ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		file, err := contract.FileMap(nil, cid)
		if err != nil {
			return nil, err
		}
		if file.Owner == common.HexToAddress("") {
			return nil, fmt.Errorf("文件不存在")
		}
		return contract.RechargeFile(opts, uid, cid, blockNum)
	}

	// 执行交易
	return a.ExecuteIpfsTransact(ctx, rechargeFunc)
}
