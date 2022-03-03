package main

import (
	"encoding/json"
	"flag"
	"fmt"
	impl "github.com/bdengine/go-ipfs-blockchain-eth/implement"
	"io/ioutil"
)

func main() {

	// 读取命令和参数
	httpUrl, socketUrl, centralServerUrl, ipfsAddr, priKey, chainID, timeout, gasLimit, root :=
		flag.String("http", "", "the httpUrl to call blockchain peer"),
		flag.String("socket", "", "the socketUrl to call blockchain peer"),
		flag.String("central", "", "the central mine Url to call"),
		flag.String("ipfsAddr", "", "ipfs contract address"),
		flag.String("pk", "", "privateKey use for blockchain"),
		flag.Int64("chainID", 1111, "chain id"),
		flag.Duration("timeout", 90, "request time out"),
		flag.Uint64("gasLimit", 0, "gas limit"),
		flag.String("pathRoot", ".", "the root path you want to gen this config file")

	flag.Parse()

	// 生成默认的配置文件
	cfg, err := impl.NewConfig(*httpUrl, *socketUrl, *centralServerUrl, *ipfsAddr, *priKey, *chainID, *timeout, *gasLimit)

	filename := *root + impl.ConfigFileName

	marshal, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(filename, marshal, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
}
