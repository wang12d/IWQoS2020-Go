package pkg

import "math/big"

// SearchToken generates the token for onchain search
func SearchToken(keyWord, Fb2 *big.Int) *big.Int {
	opts := new(big.Int)
	q := new(big.Int)
	q.SetString(qStr, base)
	opts.Mul(keyWord, Fb2)
	opts.Mod(opts, q)
	return opts
}
