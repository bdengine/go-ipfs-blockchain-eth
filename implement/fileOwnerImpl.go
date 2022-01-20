package implement

import (
	"context"
	"fmt"
	"github.com/bdengine/go-ipfs-blockchain-eth/contract/ipfs"
	"github.com/bdengine/go-ipfs-blockchain-standard/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (p *peerImpl) AddFile(info model.IpfsFileInfo) error {
	ctx := context.Background()

	blockNum := big.NewInt(info.StoreDays * 24 * 60 * 60)

	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		canAdd, err := contract.CheckFile(nil, info.Cid)
		if err != nil {
			return nil, err
		}
		if !canAdd {
			return nil, fmt.Errorf("文件已存在")
		}
		var owner common.Address
		if info.Owner == "self" || info.Owner == "" {
			owner = p.address
		} else {
			owner = common.HexToAddress(info.Owner)
		}
		// todo 检查balance和代扣额度
		return contract.AddFile(opts, uid, info.Cid, big.NewInt(info.Size), blockNum, owner)
	}

	// 执行交易
	return p.ExecuteIpfsTransact(ctx, f)
}

func (p *peerImpl) DeleteFile(cid string) error {
	ctx := context.Background()

	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		file, err := contract.GetFile(nil, cid)
		if err != nil {
			return nil, err
		}
		if file.Owner == common.HexToAddress("") {
			return nil, fmt.Errorf("文件不存在")
		}
		return contract.RemoveFile(opts, uid, cid)
	}
	// 执行交易
	return p.ExecuteIpfsTransact(ctx, f)
}

func (p *peerImpl) RechargeFile(cid string, days int64) error {
	ctx := context.Background()
	blockNum := big.NewInt(days * 24 * 60 * 60)

	// 执行充值方法
	var rechargeFunc ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		file, err := contract.GetFile(nil, cid)
		if err != nil {
			return nil, err
		}
		if file.Owner == common.HexToAddress("") {
			return nil, fmt.Errorf("文件不存在")
		}
		return contract.RechargeFile(opts, uid, cid, blockNum)
	}

	// 执行交易
	return p.ExecuteIpfsTransact(ctx, rechargeFunc)
}
