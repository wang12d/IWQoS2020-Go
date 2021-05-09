package pkg

import (
	"log"
	"math/big"
	"nju/cosec/heng/contracts/iwqos2020"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// TaskUploadSingle upload task index to smart contract one by one
func TaskUploadSingle(instance *iwqos2020.Iwqos2020, opts *bind.TransactOpts, taskIndex *map[string][]byte) {
	for token, index := range *taskIndex {
		var tokenArray, indexArray [32]byte
		copy(tokenArray[:], []byte(token))
		copy(indexArray[:], index)
		_, err := instance.Settask(opts, tokenArray, indexArray)
		if err != nil {
			log.Fatalf("upload single task failed, %v\n", err)
		}
	}
}

// TaskUploadBatch upload task index by batch
func TaskUploadBatch(instance *iwqos2020.Iwqos2020, opts *bind.TransactOpts, taskIndex *map[string][]byte, batchSize int) {
	batches, remaining := len(*taskIndex)/batchSize, len(*taskIndex)%batchSize
	tokens := make([][32]byte, batchSize)
	indexes := make([][32]byte, batchSize)
	length := big.NewInt(int64(batchSize))
	rem := big.NewInt(int64(remaining))
	cnt, currentBatch := 0, 0
	for token, index := range *taskIndex {
		copy(tokens[cnt][:], []byte(token))
		copy(indexes[cnt][:], []byte(index))
		cnt++
		if cnt == batchSize {
			cnt, currentBatch = 0, currentBatch+1
			instance.SetTaskindex(opts, tokens, indexes, length)
		}
		if currentBatch == batches && remaining == cnt {
			instance.SetTaskindex(opts, tokens[:remaining], indexes[:remaining], rem)
		}
	}
}
