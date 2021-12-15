package implement

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"github.com/ipfs/go-ipfs-auth/auth-source-eth/contract/ipfs"
	"github.com/ipfs/go-ipfs-auth/standard/model"
	"testing"
)

const (
	configRoot = "D:/IPFSSTORAGE"
)

func TestNewApi(t *testing.T) {
	indent, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(indent))
}

var a *peerImpl
var peerId string = "test1"

func init() {
	test, err := NewApi(configRoot, peerId)
	if err != nil {
		panic(err)
	}
	a = test
}

func TestApi_InitPeer(t *testing.T) {
	err := a.InitPeer(model.CorePeer{
		PeerId:    a.ID.Pid,
		Addresses: []string{"address1", "address2"},
	})
	if err != nil {
		t.Fatal(err)
	}

	peer, err := a.GetPeer(peerId)
	fmt.Println(peer)

	/*err = a.RemovePeer()
	if err != nil {
		t.Fatal(err)
	}*/
}

func TestApi_RemovePeer(t *testing.T) {
	err := a.RemovePeer()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPeerImpl_AddFile(t *testing.T) {
	info := model.IpfsFileInfo{
		Cid:       "cid3",
		Size:      100,
		StoreDays: 10,
	}
	err := a.AddFile(info)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPeerImpl_RechargeFile(t *testing.T) {
	err := a.RechargeFile("cid3", 30)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEvent(t *testing.T) {
	_, contract, err := GetIpfsContract(a.Client.SocketUrl, a.ContractMap[contractIpfs].ContractAddr)
	if err != nil {
		t.Fatal(err)
	}
	sch := make(chan *ipfs.IpfsSuccess)
	sub, err := contract.WatchSuccess(nil, sch, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case success := <-sch:
			fmt.Println("成功事件触发")
			fmt.Println(success)
		}
	}
}

func TestClientEvent(t *testing.T) {
	client, _, err := GetIpfsContract(a.Client.SocketUrl, a.ContractMap[contractIpfs].ContractAddr)
	if err != nil {
		t.Fatal(err)
	}
	sch := make(chan types.Log)
	query := ethereum.FilterQuery{}
	sub, err := client.SubscribeFilterLogs(context.TODO(), query, sch)
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case success := <-sch:
			fmt.Println("成功事件触发")
			fmt.Println(success)
		}
	}
}

func TestApi_AddFile(t *testing.T) {
	cid := "testFile1"
	err := a.AddFile(model.IpfsFileInfo{
		Cid: cid,
		Uid: uuid.New().String(),
	})
	if err != nil {
		t.Fatal(err)
	}
	err = a.DeleteFile(cid)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApi_RemoveFile(t *testing.T) {
	err := a.DeleteFile("testFile1")
	if err != nil {
		t.Fatal(err)
	}
}
