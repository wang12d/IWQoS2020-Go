package pkg

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ethutil "github.com/wang12d/Go-Crowdsourcing-DApp/src/crowdsourcing/utils/ethereum"
	"log"
	"math/big"
	"nju/cosec/heng/contracts/iwqos2020"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// UploadTaskIndexSingle upload task index to smart contract one by one
func UploadTaskIndexSingle(client *ethclient.Client, instance *iwqos2020.Iwqos2020, opts *bind.TransactOpts,
	optsAddress common.Address, taskIndex *map[string][]byte) {
	for token, index := range *taskIndex {
		var tokenArray, indexArray [32]byte
		copy(tokenArray[:], []byte(token))
		copy(indexArray[:], index)
		ethutil.UpdateNonce(client, opts, optsAddress)
		_, err := instance.Settask(opts, tokenArray, indexArray)
		if err != nil {
			log.Fatalf("upload task index single error: %v\n", err)
		}
	}
}

// UploadTaskIndexBatch upload task index by batch
func UploadTaskIndexBatch(client *ethclient.Client, instance *iwqos2020.Iwqos2020, opts *bind.TransactOpts,
	optsAddress common.Address, taskIndex *map[string][]byte, batchSize int) {
	batches, remaining := len(*taskIndex)/batchSize, len(*taskIndex)%batchSize
	tokens := make([][32]byte, batchSize)
	indexes := make([][32]byte, batchSize)
	length := big.NewInt(int64(batchSize))
	rem := big.NewInt(int64(remaining))
	cnt, currentBatch := 0, 0
	for token, index := range *taskIndex {
		copy(tokens[cnt][:], token)
		copy(indexes[cnt][:], index)
		cnt++
		if cnt == batchSize {
			cnt, currentBatch = 0, currentBatch+1
			ethutil.UpdateNonce(client, opts, optsAddress)
			_, err := instance.SetTaskindex(opts, tokens, indexes, length)
			if err != nil {
				log.Fatalf("upload task index batch error: %v\n", err)
			}
		}
		if currentBatch == batches && remaining == cnt {
			ethutil.UpdateNonce(client, opts, optsAddress)
			_, err := instance.SetTaskindex(opts, tokens[:remaining], indexes[:remaining], rem)
			if err != nil {
				log.Fatalf("upload task index batch error: %v\n", err)
			}
		}
	}
}

// UploadAuthorizationIndex upload the authorization index to smart contract
func UploadAuthorizationIndex(instance *iwqos2020.Iwqos2020, opts *bind.TransactOpts, authIndex map[string][]string) {
	for broker, auth := range authIndex {
		key := new(big.Int)
		key.SetString(broker, base)
		for _, index := range auth {
			_, err := instance.Setauthorize(opts, key, []byte(index))
			if err != nil {
				log.Fatalf("upload authorization index error: %v\n", err)
			}
		}
	}
}

// SetP upload p to blockchain to search for corresponding task index
func SetP(instance *iwqos2020.Iwqos2020, opts *bind.TransactOpts) {
	p := new(big.Int)
	p.SetString(pStr, base)
	_, err := instance.SetP(opts, p.Bytes())
	if err != nil {
		log.Fatalf("set p error: %v\n", err)
	}
}

// SearchTokenOnChain generates the search token on chain
func SearchTokenOnChain(instance *iwqos2020.Iwqos2020, opts *bind.TransactOpts, token, fbpie *big.Int) {
	_, err := instance.GetSearchtoke(opts, token, fbpie)
	if err != nil {
		log.Fatalf("get search token error: %v\n", err)
	}
}
