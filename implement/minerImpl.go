package implement

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-ipfs-auth/auth-source-eth/contract/ipfs"
	"github.com/ipfs/go-ipfs-auth/standard/model"
	"github.com/ipfs/go-ipfs-auth/standard/standardConst"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

func (a *peerImpl) GetChallenge() (string, error) {
	challenge, _, err := a.tokenContract.GetChallenge(nil)
	if challenge == "" {
		return challenge, standardConst.ChallengeError
	}
	return challenge, err
}

type miningResponse struct {
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success,omitempty"`
}

func (a *peerImpl) Mining(m model.IpfsMining) error {
	if m.Address == "" {
		m.Address = a.address.String()
	}
	marshal, err := json.Marshal(m)
	if err != nil {
		return err
	}
	res, err := sendRequest(string(marshal), a.CentralServerUrl)
	if err != nil {
		return err
	}
	var r miningResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return err
	}
	if !r.Success {
		return fmt.Errorf(r.Message)
	}
	return nil
}

func sendRequest(body string, url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 根据返回内容处理
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *peerImpl) UpdateAddress(addrList []string) error {
	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		peer, err := a.ipfsContract.AddrPeerMap(nil, a.address)
		if err != nil {
			return nil, err
		}
		if !peer.Valid {
			return nil, fmt.Errorf("节点不存在")
		}
		return contract.UpdateAddress(opts, uid, addrList)
	}
	return a.ExecuteIpfsTransact(context.Background(), f)
}

func (a *peerImpl) Heartbeat() error {
	ctx := context.Background()

	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.PeerHeartbeat(opts, uid)
	}
	return a.ExecuteIpfsTransact(ctx, f)
}

func (a *peerImpl) GetFileList(n int64) ([]string, error) {
	return a.ipfsContract.GetFileList(nil, big.NewInt(n))
}
