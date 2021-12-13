package auth_source_eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"ipfs/go-ipfs-auth/auth-source-eth/testContracts"
	"math/big"
	"sync"
	"testing"
	"time"
)

const (
	clientUrl       = "http://39.108.194.46:8545"
	socketUrl       = "ws://39.108.194.46:8546"
	contractAddress = "0xB2751A2bF4E308d0067dC266Cab5d992714674b6"
	pk              = "268a086d68c0ca3181e2726f9f3878add1af94a18a58a181b075114396f7f8be"
)

func TestCallContract(t *testing.T) {
	logs := make(chan types.Log)
	contract, opts, _, err := GetContract(logs)
	if err != nil {
		t.Fatal(err)
	}

	num := big.NewInt(256)
	tr, err := contract.Store(opts, num)
	if err != nil {
		t.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		contract1, opts1, _, err := GetContract(logs)
		if err != nil {
			t.Fatal(err)
		}
		tr1, err := contract1.Store(opts1, num)
		if err != nil {
			t.Fatal(err)
		}
		conn, err := ethclient.Dial(clientUrl)
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		time.Sleep(10 * time.Second)
		receipt1, err := conn.TransactionReceipt(context.Background(), tr1.Hash())
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(receipt1.BlockNumber)
		wg.Done()
	}()

	retrieve, err := contract.Retrieve(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(retrieve)

	conn, err := ethclient.Dial(clientUrl)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	time.Sleep(10 * time.Second)
	receipt, err := conn.TransactionReceipt(context.Background(), tr.Hash())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(receipt.BlockNumber)
	wg.Wait()
}

func TestWaitForEvent(t *testing.T) {
	conn, err := ethclient.Dial(socketUrl)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	contractAddr := common.HexToAddress(contractAddress)
	contract, err := testContracts.NewTestContracts(contractAddr, conn)
	if err != nil {
		t.Fatal(err)
	}
	sch := make(chan *testContracts.TestContractsStart)
	sub, err := contract.TestContractsFilterer.WatchStart(nil, sch)
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case start := <-sch:
			fmt.Println(start)
		case err = <-sub.Err():
			t.Fatal(err)
		}
	}
}

func GetContract(logs chan types.Log) (*testContracts.TestContracts, *bind.TransactOpts, ethereum.Subscription, error) {
	conn, err := ethclient.Dial(clientUrl)
	if err != nil {
		return nil, nil, nil, err
	}
	defer conn.Close()
	contractAddr := common.HexToAddress(contractAddress)
	address, err := testContracts.NewTestContracts(contractAddr, conn)
	if err != nil {
		return nil, nil, nil, err
	}
	priKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, nil, nil, err
	}
	pubKey := priKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
	if !ok {
		return nil, nil, nil, fmt.Errorf("公钥无法断言为类型 *ecdsa.PublicKey")
	}
	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, nil, nil, err
	}

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}
	chainId := big.NewInt(18888)
	auth, err := bind.NewKeyedTransactorWithChainID(priKey, chainId)
	if err != nil {
		return nil, nil, nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return address, auth, nil, nil
}
