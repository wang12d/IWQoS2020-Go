package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"nju/cosec/heng/contracts/iwqos2020"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	ethutil "github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"

	"nju/cosec/heng/pkg"
)

const (
	URL  = "http://localhost:8545"
	opts = "97e0ba71d9ba8a41bf31c74b80484d58aa400ef2084a5ae15174d2394ccb3124"
)

func main() {
	var err error
	// Local experiment code
	numberOfBrokers := 6
	numberOfTasks := 8
	brokers, err := pkg.GenerateBrokers(numberOfBrokers, numberOfTasks)
	if err != nil {
		log.Fatalf("generate brokers error: %v", err)
	}
	fmt.Println("********************Broker content generated.********************")
	brokerKeys := pkg.GenerateBrokerKeys(numberOfBrokers)
	fmt.Println("********************Broker keys done.********************")
	brokerLocalState := make([]map[string]int, numberOfBrokers)
	for i := 0; i < numberOfBrokers; i++ {
		brokerLocalState[i] = make(map[string]int)
		for key := range brokers[i] {
			brokerLocalState[i][key] = 0
		}
	}
	start := time.Now()
	onChainIndex := pkg.BuildIndex(&brokers, &brokerLocalState, brokerKeys)
	fmt.Println("********************Build on chain index generated.********************")
	fmt.Println("on Chain index: ", onChainIndex)
	fmt.Println(time.Since(start))
	start = time.Now()
	authTable := pkg.GetAuthorizationIndex(brokerKeys, brokerKeys)
	fmt.Println("********************Authorization index generated.********************")
	fmt.Println(time.Since(start))
	token, pp, fbpie := new(big.Int), new(big.Int), new(big.Int)
	token.SetString("308787214124473614167655", 10)
	pp.SetString(brokerKeys[0].Fb2, 10)
	fbpie.SetString(brokerKeys[0].Fbpie, 10)
	tbw1 := pkg.GetSearchToken(token, pp)
	pkg.SearchKeyTokenLocally(tbw1, fbpie, authTable, onChainIndex)
	// On chain experiment code

	// First connect to the ganache experiment network
	client, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("connect to ganache client error: %v\n", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("get chain id error: %v\n", err)
	}
	optsPrivateKey, optsAddress := ethutil.PrivateKeyAndAddress(opts)
	transactor := ethutil.KeyedTransactor(client, optsPrivateKey, optsAddress, chainID, big.NewInt(0))
	fmt.Println("Suggested gas price:", transactor.GasPrice)
	transactor.GasLimit = 67219750000
	transactor.GasPrice = big.NewInt(200000)

	// Deploy contract
	contractAdress, _, instance, err := iwqos2020.DeployIwqos2020(transactor, client)
	fmt.Println("contract deployed at: ", contractAdress)
	if err != nil {
		log.Fatalf("iwqos2020 deployed error: %v\n", err)
	}
	ethutil.UpdateNonce(client, transactor, optsAddress)
	start = time.Now()
	pkg.UploadTaskIndexBatch(client, instance, transactor, optsAddress, &onChainIndex, 1000)
	fmt.Println("upload task index batch time cost: ", time.Since(start))
	start = time.Now()
	pkg.UploadTaskIndexSingle(client, instance, transactor, optsAddress, &onChainIndex)
	fmt.Println("upload task index single time cost: ", time.Since(start))
	fmt.Println("task index uploaded.")
}
