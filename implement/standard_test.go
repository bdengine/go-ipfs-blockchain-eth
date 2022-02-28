package implement

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bdengine/go-ipfs-blockchain-eth/contract/ipfs"
	"github.com/bdengine/go-ipfs-blockchain-eth/contract/scToken"
	"github.com/bdengine/go-ipfs-blockchain-standard/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/prometheus/common/log"
	"math/big"
	"math/bits"
	"sync"
	"testing"
	"time"
)

const (
	configRoot = "D:/IPFSSTORAGE"
)

func TestNewApi(t *testing.T) {
	indent, err := json.MarshalIndent(mockPeer, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(indent))
}

var mockPeer *peerImpl = &peerImpl{
	config: &config{
		Client: clientInfo{
			HttpUrl:   "http://192.168.80.113:9545",
			SocketUrl: "ws://192.168.80.113:9546",
		},
		CentralServerUrl: "",
		ContractMap: map[string]contractInfo{
			"ipfs":  {"0x455b97Df329DF0D4bE6b02b23Cf1002C998A5cD6"},
			"token": {"0xF6f682C3d4362A0cE1b5F2B9C381fd0fFe9D7959"},
		},
		Chain: chainInfo{ChainId: 20180518},
		Variable: configInfo{
			RequestTimeout: 30,
			GasLimit:       250000,
		},
		ID: Identity{
			Pid:    "12D3KooWNpzDbLRwxtrw9W8yijeKvSs8epK8yHjEPoiGGj71rF4x",
			PriKey: "3db74108afca5195374fd6ebdcfcdcd4ccc4338d4a70a8da36f693c995345e7e",
		},
	},
	priKey:        nil,
	address:       common.Address{},
	chainId:       nil,
	lock:          nil,
	socketClient:  nil,
	ipfsContract:  nil,
	tokenContract: nil,
	httpClient:    nil,
}

func init() {
	httpClient, err := ethclient.Dial(mockPeer.Client.HttpUrl)
	if err != nil {
		panic(err)
	}
	socketClient, err := ethclient.Dial(mockPeer.Client.SocketUrl)
	if err != nil {
		panic(err)
	}

	mockPeer.socketClient = socketClient
	mockPeer.httpClient = httpClient
	mockPeer.lock = &sync.Mutex{}

	ipfsContra, err := ipfs.NewIpfs(common.HexToAddress(mockPeer.ContractMap[contractIpfs].ContractAddr), socketClient)
	if err != nil {
		panic(err)
	}
	mockPeer.ipfsContract = ipfsContra
	tokenContra, err := scToken.NewScToken(common.HexToAddress(mockPeer.ContractMap[contractToken].ContractAddr), socketClient)
	if err != nil {
		panic(err)
	}
	mockPeer.tokenContract = tokenContra
	mockPeer.chainId = big.NewInt(mockPeer.Chain.ChainId)
	priKey, err := crypto.HexToECDSA(mockPeer.ID.PriKey)
	if err != nil {
		panic(err)
	}
	mockPeer.priKey = priKey
	pubKey := priKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Errorf("公钥无法断言为类型 *ecdsa.PublicKey"))
	}
	address := crypto.PubkeyToAddress(*pubKeyECDSA)
	fmt.Println(address.String())
	mockPeer.address = address
}

func TestApi_InitPeer(t *testing.T) {
	err := mockPeer.InitPeer(model.CorePeer{
		PeerId:    mockPeer.ID.Pid,
		Addresses: []string{"address1", "address2"},
	})
	if err != nil {
		t.Fatal(err)
	}

	peer, err := mockPeer.GetPeer("12D3KooWLUTBUkLnfdcJbyV1C7ZRsdFcDoN89mmYknFH5ef9pTyM")
	fmt.Println(peer)
}

func TestApi_RemovePeer(t *testing.T) {
	err := mockPeer.RemovePeer()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPeerImpl_AddFile(t *testing.T) {
	info := model.IpfsFileInfo{
		Cid:       "2",
		Size:      1,
		StoreDays: 1,
	}
	err := mockPeer.AddFile(info)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFile(t *testing.T) {
	list, err := mockPeer.ipfsContract.GetFileList(nil, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range list {
		file, err := mockPeer.ipfsContract.GetFile(nil, s)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("cid:%v,file:%v\n", s, file)
	}
}

func TestPeerImpl_RechargeFile(t *testing.T) {
	err := mockPeer.RechargeFile("cid3", 30)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_quick(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			err := mockPeer.RechargeFile("cid3", 1)
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
			pid, err := mockPeer.ipfsContract.GetPeerByPid(nil, "12D3KooWLUTBUkLnfdcJbyV1C7ZRsdFcDoN89mmYknFH5ef9pTyM")
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(pid)
		}()
	}
	time.Sleep(10 * time.Second)
}

func TestPeerImpl_DeleteFile(t *testing.T) {
	err := mockPeer.DeleteFile("2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestApi_AddFile(t *testing.T) {
	cid := "testFile1"
	err := mockPeer.AddFile(model.IpfsFileInfo{
		Cid: cid,
		Uid: uuid.New().String(),
	})
	if err != nil {
		t.Fatal(err)
	}
	err = mockPeer.DeleteFile(cid)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApi_RemoveFile(t *testing.T) {
	err := mockPeer.DeleteFile("testFile1")
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
	err := mockPeer.Mining(mining)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPeerImpl_GetChallenge(t *testing.T) {
	challenge, err := mockPeer.GetChallenge()
	if err != nil {
		t.Fatal(err)
	}

	//a.Mining()
	fmt.Println(challenge)
}

func TestGetTxByHash(t *testing.T) {
	httpClient, err := ethclient.Dial(mockPeer.Client.HttpUrl)
	if err != nil {
		t.Fatal(err)
	}

	h := common.HexToHash("0xb8e9167f56d864076f1666bb1e6ee0e7e9eb14e7826233fca4a50c3d8bdcdb86")
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
