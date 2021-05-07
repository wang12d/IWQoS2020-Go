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
	for _, broker := range brokers {
		for k, v := range broker {
			fmt.Printf("%v: \n", k)
			for _, b := range v {
				fmt.Printf("\t%v\n", b)
			}
		}
	}
}
