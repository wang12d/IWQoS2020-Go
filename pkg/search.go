package pkg

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

// GetSearchToken generates the token for onchain search
func GetSearchToken(keyWord, Fb2 *big.Int) *big.Int {
	opts := new(big.Int)
	q := new(big.Int)
	q.SetString(qStr, base)
	opts.Mul(keyWord, Fb2)
	opts.Mod(opts, q)
	return opts
}

// SearchKeyTokenLocally using tbw and fbpie to search the task broker interested
func SearchKeyTokenLocally(tbw, fbpie *big.Int, authorizationTable map[string][]string, taskIndex map[string][]byte) []string {
	interestedTask := make([]string, 0)
	opterators, p := new(big.Int), new(big.Int)
	p.SetString(pStr, base)
	for _, auth := range authorizationTable[fbpie.String()] {
		authIndex := new(big.Int)
		authIndex.SetString(auth, base)
		expt := opterators.Exp(authIndex, tbw, p)
		c := 0
		trap := fmt.Sprintf("%#v", expt)
		trapdoor := trap + padding(c)
		label := hex.EncodeToString(crypto.Keccak256([]byte(trapdoor)))
		for _, ok := taskIndex[label]; ok; {
			G2 := []byte(label)
			minLen := min(len(G2), len(taskIndex[label]))
			ciphertext := make([]byte, minLen)
			for i := 0; i < minLen; i++ {
				ciphertext[i] = G2[i] ^ taskIndex[label][i]
			}
			interestedTask = append(interestedTask, hex.EncodeToString(ciphertext))
			c++
			trapdoor = trap + padding(c)
			label = hex.EncodeToString(crypto.Keccak256([]byte(trapdoor)))
			_, ok = taskIndex[label]
		}
	}
	return interestedTask
}
