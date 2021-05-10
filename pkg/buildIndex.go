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
func BuildIndex(brokers *map[string][][]byte, brokerLocalState *[]map[string]int, brokerKeys []BrokerKey) map[string][]byte {
	numberOfBrokers := len(*brokers)
	g, p, q := new(big.Int), new(big.Int), new(big.Int)
	g.SetString(gStr, 10)
	p.SetString(pStr, 10)
	q.SetString(qStr, 10)
	taskIndex := make(map[string][]byte)
	intOp, tmp := new(big.Int), new(big.Int)
	for i := 0; i < numberOfBrokers; i++ {
		for key, val := range *brokers {
			keyNumber, Fb1 := new(big.Int), new(big.Int)
			keyNumber.SetString(key, 10)
			Fb1.SetString(brokerKeys[i].Fb1, 10)
			tmp.Mul(keyNumber, Fb1)
			tmp.Mod(tmp, q)
			backdoor := intOp.Exp(g, tmp, p)
			trap := hex.EncodeToString([]byte(backdoor.String()))
			for _, block := range val {
				trapdoor := trap + padding((*brokerLocalState)[i][key])
				(*brokerLocalState)[i][key]++
				label := crypto.Keccak256([]byte(trapdoor))
				cipher, err := aes.NewCipher([]byte("1234512345123451"))
				if err != nil {
					log.Fatalln("AES create new cipher error")
				}
				C := utils.AESECBEncryption(cipher, block)
				minLen := min(len(C), len(label))
				P := make([]byte, minLen)
				for i := 0; i < minLen; i++ {
					P[i] = C[i] ^ label[i]
				}
				taskIndex[hex.EncodeToString(label)] = P
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
