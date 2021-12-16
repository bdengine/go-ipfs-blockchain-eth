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
	"math/bits"
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
var peerId string = "12D3KooWLUTBUkLnfdcJbyV1C7ZRsdFcDoN89mmYknFH5ef9pTyM"

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

func TestPeerImpl_DeleteFile(t *testing.T) {
	err := a.DeleteFile("cid3")
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

func TestDefaultConfig(t *testing.T) {
	c := config{
		Client: clientInfo{
			HttpUrl:   "http://39.108.194.46:8545",
			SocketUrl: "ws://39.108.194.46:8546",
		},
		CentralServerUrl: "",
		ContractMap:      map[string]contractInfo{contractIpfs: {ContractAddr: "0x7334B7D92fF253F4462eCA1629581B241455b482"}, contractToken: {ContractAddr: "0x9AD3748905722C6a1b00F7A2400582b1fC0b67A3"}},
		Chain:            chainInfo{ChainId: 188888},
		Variable:         configInfo{RequestTimeout: 30},
		ID: Identity{
			Pid:    "test",
			PriKey: "268a086d68c0ca3181e2726f9f3878add1af94a18a58a181b075114396f7f8be",
		},
	}

	indent, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(indent))

}

func ZeroPrefixLen(id []byte) int {
	for i, b := range id {
		if b != 0 {
			return i*8 + bits.LeadingZeros8(uint8(b))
		}
	}
	return len(id) * 8
}

func XOR(a, b []byte) []byte {
	c := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] ^ b[i]
	}
	return c
}

func CommonPrefixLen(a, b []byte) int {
	c := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] ^ b[i]
	}
	for i, by := range c {
		if by != 0 {
			return i*8 + bits.LeadingZeros8(uint8(by))
		}
	}
	return len(c) * 8
}

func Test_CommonPrefixLen(t *testing.T) {
	ba, bb := []byte{0, 2, 3}, []byte{127, 2, 4}
	fmt.Println(CommonPrefixLen(ba, bb))
}
