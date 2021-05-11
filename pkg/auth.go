package pkg

import (
	"fmt"
	"math/big"
)

// GetAuthorizationIndex generates the authorization list of each broker, whose authorize
// other brokers in authorizationTable
func GetAuthorizationIndex(brokers, authorizationTable []BrokerKey) map[string][]string {
	operates := new(big.Int)
	p, q, g := new(big.Int), new(big.Int), new(big.Int)
	p.SetString(pStr, base)
	q.SetString(qStr, base)
	g.SetString(gStr, base)
	authIndex := make([][]string, len(authorizationTable))
	for i := range authIndex {
		authIndex[i] = make([]string, len(brokers))
	}
	rev := new(big.Int)
	for j, authorized := range authorizationTable {
		b := new(big.Int)
		b.SetString(authorized.Fb2, base)
		rev.SetBytes(operates.ModInverse(b, q).Bytes())
		for i, broker := range brokers {
			a := new(big.Int)
			a.SetString(broker.Fb1, base)
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
