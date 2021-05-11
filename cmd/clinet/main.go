package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"nju/cosec/heng/pkg"
)

func main() {
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
	fmt.Println(len(onChainIndex))
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
	a := pkg.SearchKeyTokenLocally(tbw1, fbpie, authTable, onChainIndex)
	fmt.Println(a)

}
