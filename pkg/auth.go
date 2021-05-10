package pkg

import (
	"fmt"
	"log"
	"math/big"
)

// GetAuthorizationIndex generates the authorization list of each broker, whose authorize
// other brokers in authorizationTable
func GetAuthorizationIndex(brokers, authorizationTable []BrokerKey) map[string][]string {
	operates := new(big.Int)
	p, q, g := new(big.Int), new(big.Int), new(big.Int)
	p.SetString(pStr, 10)
	q.SetString(qStr, 10)
	g.SetString(gStr, 10)
	authIndex := make([][]string, len(authorizationTable))
	for i := range authIndex {
		authIndex[i] = make([]string, len(brokers))
	}
	for j, authorized := range authorizationTable {
		b := new(big.Int)
		b.SetString(authorized.Fb2, 10)
		for i, broker := range brokers {
			a := new(big.Int)
			a.SetString(broker.Fb1, 10)
			rev := operates.ModInverse(a, q)
			if rev == nil {
				log.Fatalln("Error, a and q must be relatively prime.")
			}
			tmp := operates.Mul(a, rev)
			tmp.Mod(tmp, q)
			tmp.Exp(g, tmp, p)
			hexTmp := fmt.Sprintf("%#0768x", tmp)
			authIndex[j][i] = hexTmp
		}
	}
	auth := make(map[string][]string)
	for i, broker := range authorizationTable {
		auth[broker.Fbpie] = authIndex[i]
	}
	return auth
}
