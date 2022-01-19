package implement

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

func (p *peerImpl) GetChallenge() (string, error) {
	challenge, _, err := p.tokenContract.GetChallenge(nil)
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

func (p *peerImpl) Mining(m model.IpfsMining) error {
	if m.Address == "" {
		m.Address = p.address.String()
	}
	marshal, err := json.Marshal(m)
	if err != nil {
		return err
	}
	res, err := sendRequest(string(marshal), p.CentralServerUrl)
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

func (p *peerImpl) UpdateAddress(addrList []string) error {
	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		peer, err := p.ipfsContract.AddrPeerMap(nil, p.address)
		if err != nil {
			return nil, err
		}
		if !peer.Valid {
			return nil, fmt.Errorf("节点不存在")
		}
		return contract.UpdateAddress(opts, uid, addrList)
	}
	return p.ExecuteIpfsTransact(context.Background(), f)
}

func (p *peerImpl) Heartbeat() error {
	ctx := context.Background()

	var f ExecuteIpfsFunc = func(uid string, contract *ipfs.Ipfs, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.PeerHeartbeat(opts, uid)
	}
	return p.ExecuteIpfsTransact(ctx, f)
}

func (p *peerImpl) GetFileList(n int64) ([]string, error) {
	return p.ipfsContract.GetFileList(nil, big.NewInt(n))
}
