package pkg

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"nju/cosec/heng/pkg/utils"

	"github.com/ethereum/go-ethereum/crypto"
)

func padding(number int) string {
	return fmt.Sprintf("%064x", number)
}

// BuildIndex builds task index according to the description of paper for each broker
func BuildIndex(brokers *[]map[string][][]byte, brokerLocalState *[]map[string]int, brokerKeys []BrokerKey) map[string][]byte {
	numberOfBrokers := len(*brokers)
	g, p, q := new(big.Int), new(big.Int), new(big.Int)
	g.SetString(gStr, base)
	p.SetString(pStr, base)
	q.SetString(qStr, base)
	taskIndex := make(map[string][]byte)
	intOp, tmp := new(big.Int), new(big.Int)
	for i := 0; i < numberOfBrokers; i++ {
		for key, val := range (*brokers)[i] {
			keyNumber, Fb1 := new(big.Int), new(big.Int)
			keyNumber.SetString(key, base)
			Fb1.SetString(brokerKeys[i].Fb1, base)
			tmp.Mul(keyNumber, Fb1)
			tmp.Mod(tmp, q)
			backdoor := intOp.Exp(g, tmp, p)
			if backdoor == nil {
				log.Fatalln("Exp error")
			}
			trap := fmt.Sprintf("%x", backdoor)
			for _, block := range val {
				trapdoor := trap + padding((*brokerLocalState)[i][key])
				(*brokerLocalState)[i][key]++
				if len(trapdoor) % 2 == 1 {
					trapdoor = "0" + trapdoor
				}
				hexTrapDoor, err := hex.DecodeString(trapdoor)
				if err != nil {
					log.Fatalf("Build index encounter hex decode error: %v\n", err)
				}
				label := crypto.Keccak256(hexTrapDoor)
				cipher, err := aes.NewCipher([]byte("1234512345123451"))
				if err != nil {
					log.Fatalln("AES create new cipher error")
				}
				C := utils.AESECBEncryption(cipher, block)
				minLen := min(len(C), len(label))
				P := make([]byte, minLen)
				for v := 0; v < minLen; v++ {
					P[v] = C[v] ^ label[v]
				}
				taskIndex["0x"+hex.EncodeToString(label)] = P
			}
		}
	}
	return taskIndex
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
