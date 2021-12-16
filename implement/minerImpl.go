package implement

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-ipfs-auth/auth-source-eth/contract/ipfs"
	"github.com/ipfs/go-ipfs-auth/standard/model"
	"github.com/ipfs/go-ipfs-auth/standard/standardConst"
)

func (a *peerImpl) GetChallenge() (string, error) {
	cli, contr, err := GetSCTokenContract(a.Client.HttpUrl, a.ContractMap[contractToken].ContractAddr)
	if err != nil {
		return "", err
	}
	defer cli.Close()
	challenge, _, err := contr.GetChallenge(nil)
	if challenge == "" {
		return challenge, standardConst.ChallengeError
	}
	return challenge, err
}

func (a *peerImpl) Mining([]model.IpfsMining) error {
	return nil
}

func (a *peerImpl) UpdateAddress(addrList []string) error {
	ctx := context.Background()
	cli, ipfsContra, err := GetIpfsContract(a.Client.SocketUrl, a.ContractMap[contractIpfs].ContractAddr)
	if err != nil {
		return err
	}
	defer cli.Close()
	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		peer, err := ipfsContra.AddrPeerMap(nil, a.address)
		if err != nil {
			return nil, err
		}
		if !peer.Valid {
			return nil, fmt.Errorf("节点不存在")
		}
		return ipfsContra.UpdateAddress(opts, uid, addrList)
	}
	return a.ExecuteIpfsTransact(ctx, f)
}

func (a *peerImpl) Heartbeat() error {
	ctx := context.Background()
	cli, ipfsContra, err := GetIpfsContract(a.Client.SocketUrl, a.ContractMap[contractIpfs].ContractAddr)
	if err != nil {
		return err
	}
	defer cli.Close()
	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		return ipfsContra.PeerHeartbeat(opts, uid)
	}
	return a.ExecuteIpfsTransact(ctx, f)
}
