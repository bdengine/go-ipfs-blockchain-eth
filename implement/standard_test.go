package implement

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/ipfs/go-ipfs-auth/standard/model"
	"github.com/prometheus/common/log"
	"math/bits"
	"sync"
	"testing"
	"time"
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
	a.lock = &sync.Mutex{}
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

func Test_quick(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			err := a.RechargeFile("cid3", 1)
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestPeerImpl_GetPeer(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			pid, err := a.ipfsContract.GetPeerByPid(nil, peerId)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(pid)
		}()
	}
	time.Sleep(10 * time.Second)
}

func TestPeerImpl_DeleteFile(t *testing.T) {
	err := a.DeleteFile("cid3")
	if err != nil {
		t.Fatal(err)
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

func TestPeerImpl_Mining(t *testing.T) {
	mining := model.IpfsMining{
		Cid:         "testCid",
		Hash:        "testHash",
		Address:     "testAddr",
		LeadingZero: 0,
		Challenge:   "this is a test",
	}
	err := a.Mining(mining)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPeerImpl_GetChallenge(t *testing.T) {
	challenge, err := a.GetChallenge()
	if err != nil {
		t.Fatal(err)
	}

	//a.Mining()
	fmt.Println(challenge)
}

func TestGetTxByHash(t *testing.T) {
	httpClient, err := ethclient.Dial(a.Client.HttpUrl)
	if err != nil {
		t.Fatal(err)
	}

	h := common.HexToHash("0x1d822408666783084ca4d13464a2e805757b701179e7480fa7ddb697e54e8")
	receipt, err := httpClient.TransactionReceipt(context.Background(), h)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(receipt.ContractAddress.String())
}

// CommonPrefixLen 计算两个32位byte数组的共同前导
func CommonPrefixLen(a, b []byte) int {
	if len(a) != 32 || len(b) != 32 {
		log.Error("前导零计算，位数错误")
		return 0
	}
	for i := 0; i < 32; i++ {
		c := a[i] ^ b[i]
		if c != 0 {
			return i*8 + bits.LeadingZeros8(uint8(c))
		}
	}
	return 32 * 8
}

func TestCommon(t *testing.T) {
	decodeString, err := base64.StdEncoding.DecodeString("5z8kbNngAX4KerJ01b13tKB2+SVPxLxbpg4D+pF6GtQ=")
	bytes, err := base64.StdEncoding.DecodeString("Ea/wbSjlv/uvWTpNopS/P2yfg0qS8yKwM/anWgSMgQ0=")
	if err != nil {
		t.Fatal(err)
	}
	prefixLen := CommonPrefixLen(decodeString, bytes)
	fmt.Println(prefixLen)
}
