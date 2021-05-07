package main

import (
	"fmt"
	"log"

	"nju/cosec/heng/pkg"
)

func main() {
	numberOfBrokers := 6
	numberOfTasks := 8
	brokers, err := pkg.GenerateBrokers(numberOfBrokers, numberOfTasks)
	if err != nil {
		log.Fatalf("generate brokers error: %v", err)
	}
	fmt.Println("********************Broker address********************")
	for _, broker := range brokers {
		for k := range broker {
			fmt.Printf("%v\n", k)
		}
	}
	brokerKeys := pkg.GenerateBrokerKeys(numberOfBrokers)
	fmt.Println("********************Broker Keys********************")
	for _, bk := range brokerKeys {
		fmt.Println(bk)
	}
}
