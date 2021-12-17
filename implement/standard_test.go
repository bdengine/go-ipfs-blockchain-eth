package implement

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/google/uuid"
	"github.com/ipfs/go-ipfs-auth/auth-source-eth/contract/scToken"
	"github.com/ipfs/go-ipfs-auth/standard/model"
	"github.com/prometheus/common/log"
	"io/ioutil"
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

func (a *peerImpl) testToken() error {

	ctx := context.Background()
	var f ExecuteTokenFunc = func(uid string, contract *scToken.ScToken, opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.Approve(opts, common.HexToAddress("0xEA2AC66C8dcC73cA0599255841e5Cef5c0A0ff5F"), big.NewInt(1))
	}
	// 执行交易
	return a.ExecuteTokenTransact(ctx, f)
}

func (a *peerImpl) ExecuteTokenTransact(ctx context.Context, f ExecuteTokenFunc) error {
	// 获取客户端和合约实例
	cli, tknC, err := GenNewSCTokenContract(a.Client.SocketUrl, a.ContractMap[contractToken].ContractAddr)
	if err != nil {
		return err
	}
	defer cli.Close()
	uid := uuid.New().String()
	sch := make(chan *scToken.ScTokenApproval)
	sub, err := tknC.ScTokenFilterer.WatchApproval(nil, sch, []common.Address{}, []common.Address{})
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	// 生成配置项
	opts, err := a.GenTransactOpts(ctx, cli, a.config.Variable.GasLimit)
	if err != nil {
		return err
	}
	// 执行交易方法
	c, t2, err := GenNewSCTokenContract(a.Client.HttpUrl, a.ContractMap[contractToken].ContractAddr)
	if err != nil {
		return err
	}
	defer c.Close()
	tx, err := f(uid, t2, opts)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, a.Variable.RequestTimeout*time.Second)
	// todo 日志
	log.Debugf("%v交易结果查询：", tx.Hash())
	a.lock.Unlock()
	err = waitResult1(ctx, cli, sch, sub, tx.Hash(), uid)
	cancel()
	if err != nil {
		return err
	}
	return nil
}

func waitResult1(ctx context.Context, sCli *ethclient.Client, sch chan *scToken.ScTokenApproval, sub event.Subscription, txId common.Hash, uid string) error {
	tick := time.Tick(10 * time.Second)
	var err error
	for {
		select {
		case s := <-sch:
			log.Debugf("成功回调:%s", uid)
			fmt.Println(s.Raw.BlockNumber)
			return nil
		case err = <-sub.Err():
			return err
		case <-tick:
			log.Debugln("未收到成功事件,轮询回执")
			receipt, err := sCli.TransactionReceipt(context.Background(), txId)
			if err != nil {
				if err.Error() == "not found" {
					log.Debugf("交易结果不确定,继续轮询")
				}
			} else if receipt.Status == txStatusFail {
				return fmt.Errorf("交易失败")
			} else {
				fmt.Println(receipt.BlockNumber)
				log.Debugln("交易成功")
				fmt.Println(receipt.BlockNumber)
				return nil
			}
		case <-ctx.Done():
			log.Debugln("交易超时，最后一次查询回执")
			receipt, err := sCli.TransactionReceipt(context.Background(), txId)
			if err != nil {
				if err.Error() == "not found" {
					return fmt.Errorf("交易执行结果不确定,请稍候查询")
				}
				return err
			}
			if receipt.Status == txStatusFail {
				return fmt.Errorf("交易失败")
			}
			log.Debugln("交易成功")
			return nil
		}
	}
}

func Test_quickTen(t *testing.T) {
	pk := []string{
		"68d2cf469d2638a727079a63bc5307a99889b580f41f04a34a09f59fdfc79832",
		"a5e8ac9f7c4500beb0cce0c27dc15ff38ac7dbc2009438a93c3ac3600f336d2b",
		"29da35b897f0ccfc385decbf6093f06cbba6c8c27182966c62bdeecc0748ba9b",
		"d40b837f3988666b70d3232007785324a9a5ef8471707dbaa188a931f8de97f9",
		"223805033fe03d4a84c2a524c6b9fa0c459b65baa82e83af0e6ff504f7f92dfd",
		"234d7249ea68e55aece81bfc4d008a2f8b37b2796e0e026b1d03d101c1f5c396",
	}
	apiList, err := BatchNewApi(configRoot+"/test", pk)
	if err != nil {
		t.Fatal(err)
	}
	times := 10
	eCh := make(chan bool, len(apiList)*times)
	start := time.Now().Unix()
	for _, impl := range apiList {
		go func() {
			for i := 0; i < 1; i++ {
				go func() {
					e := impl.testToken()
					eCh <- e == nil
				}()
			}
		}()
	}

	fail := 0
	success := 0
	for i := 0; i < len(apiList)*times; i++ {
		e := <-eCh
		fmt.Printf("调用结果%v\n", i)
		if e {
			success++
		} else {
			fail++
		}

	}
	fmt.Printf("成功:%v,失败:%v,耗时:%v", success, fail, (time.Now().Unix()-start)/int64(time.Second))
}

func BatchNewApi(configRoot string, pk []string) ([]*peerImpl, error) {
	l := len(pk)
	var res = make([]*peerImpl, l)
	for i, s := range pk {
		fileName := configRoot + configFileName
		readFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
		var cfg config
		var r peerImpl
		err = json.Unmarshal(readFile, &cfg)
		if err != nil {
			return nil, err
		}

		// 检查配置项是否完整
		err = checkConfig(&cfg)
		if err != nil {
			panic(err)
			return nil, err
		}

		// 设置参数
		priKey, err := crypto.HexToECDSA(s)
		if err != nil {
			return nil, err
		}
		pubKey := priKey.Public()
		pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("公钥无法断言为类型 *ecdsa.PublicKey")
		}
		address := crypto.PubkeyToAddress(*pubKeyECDSA)

		r.priKey = priKey
		r.address = address
		r.chainId = big.NewInt(cfg.Chain.ChainId)
		r.config = &cfg
		r.lock = &sync.Mutex{}
		res[i] = &r
	}
	return res, nil
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
