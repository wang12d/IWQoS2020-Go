package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"nju/cosec/heng/contracts/iwqos2020"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/requester"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/task"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/encoder"
	ethutil "github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/utils/ethereum"
	"github.com/wang12d/Go-Crowdsourcing-DApp/pkg/crowdsourcing/worker"

	"nju/cosec/heng/pkg"
)

const (
	URL               = "http://localhost:8545"
	opts              = "97e0ba71d9ba8a41bf31c74b80484d58aa400ef2084a5ae15174d2394ccb3124"
	numberOfIteration = 10
	reward            = 5000
	mu                = 1000
	sigma             = 250
	byteSize          = 2048
)

var ()

func main() {
	var err error

	// Local experiment code
	numberOfBrokers := 1
	numberOfTasks := 1
	brokers, err := pkg.GenerateBrokers(numberOfBrokers, numberOfTasks)
	if err != nil {
		log.Fatalf("generate brokers error: %v", err)
	}
	brokerKeys := pkg.GenerateBrokerKeys(numberOfBrokers)
	brokerLocalState := make([]map[string]int, numberOfBrokers)
	for i := 0; i < numberOfBrokers; i++ {
		brokerLocalState[i] = make(map[string]int)
		for key := range brokers[i] {
			brokerLocalState[i][key] = 0
		}
	}
	onChainIndex := pkg.BuildIndex(&brokers, &brokerLocalState, brokerKeys)
	authTable := pkg.GetAuthorizationIndex(brokerKeys, brokerKeys)
	token, pp, fbpie := new(big.Int), new(big.Int), new(big.Int)
	token.SetString("308787214124473614167655", 10)
	pp.SetString(brokerKeys[0].Fb2, 10)
	fbpie.SetString(brokerKeys[0].Fbpie, 10)
	tbw1 := pkg.GetSearchToken(token, pp)
	pkg.SearchKeyTokenLocally(tbw1, fbpie, authTable, onChainIndex)

	// On chain experiment code
	// First connect to the ganache experiment network
	client, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("connect to ganache client error: %v\n", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("get chain id error: %v\n", err)
	}
	optsPrivateKey, optsAddress := ethutil.PrivateKeyAndAddress(opts)
	transactor := ethutil.KeyedTransactor(client, optsPrivateKey, optsAddress, chainID, big.NewInt(0))
	transactor.GasLimit = 67219750000
	transactor.GasPrice = big.NewInt(200000)
	// Deploy contract
	_, _, instance, err := iwqos2020.DeployIwqos2020(transactor, client)
	if err != nil {
		log.Fatalf("iwqos2020 deployed error: %v\n", err)
	}
	ethutil.UpdateNonce(client, transactor, optsAddress)
	pkg.SetP(client, instance, transactor, optsAddress)
	ethutil.UpdateNonce(client, transactor, optsAddress)
	pkg.UploadAuthorizationIndex(client, instance, transactor, optsAddress, authTable)
	// Experiment evalutions
	numberOfWorkers := make([]int, 0)
	for _, str := range os.Args[1:] {
		nw, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Convert to number err: %v\n", err)
		}
		numberOfWorkers = append(numberOfWorkers, nw)
	}
	length := len(numberOfWorkers)
	registerationTimeCost, taskPublicationTimeCost := make([]time.Duration, length), make([]time.Duration, length)
	taskParticipationTimeCost, dataCollectionTimeCost := make([]time.Duration, length), make([]time.Duration, length)
	rewardingTimeCost := make([]time.Duration, length)
	onChainStorage := make([]int, length)
	communication := make([]int, length)

	var registerationCost, taskPublicationCost, taskParticipationCost, dataCollectionCost, rewardingCost time.Duration
	var onChainBytes int

	var timeStart time.Time
	var communicationCost int

	for i := 0; i < int(numberOfIteration); i++ {
		for n, workerRequired := range numberOfWorkers {
			onChainBytes, registerationCost, taskPublicationCost, taskParticipationCost, dataCollectionCost, rewardingCost = 0, 0, 0, 0, 0, 0
			communicationCost = 0
			r := requester.NewRequester(byteSize)
			timeStart = time.Now()
			r.Register()
			registerationCost += time.Since(timeStart)
			taskDescription := "Collecting the time of daliy smartphone usage"

			// In PFCrowd, requester must upload the task index and authentication table before posting task
			// Evaluating the time cost of posting task
			timeStart = time.Now()
			pkg.UploadAuthorizationIndex(client, instance, transactor, optsAddress, authTable)
			pkg.UploadTaskIndexBatch(client, instance, transactor, optsAddress, &onChainIndex, 1000)
			taskPublicationCost += time.Since(timeStart)
			timeStart = time.Now()
			r.PostTask(workerRequired, reward, taskDescription)
			taskPublicationCost += time.Since(timeStart)

			taskPublicationTimeCost[n] += taskPublicationCost

			workers := make([]*worker.Worker, workerRequired)
			for ii := 0; ii < workerRequired; ii++ {
				workers[ii] = worker.NewWorker()
				timeStart = time.Now()
				workers[ii].Register(1)
				registerationCost += time.Since(timeStart)
			}
			registerationTimeCost[n] += registerationCost / time.Duration(workerRequired+1)

			for _, w := range workers {
				daliyTime := mu + 3*sigma + rand.Uint64()%5000
				data := encoder.Uint64ToBytes(daliyTime)

				// In PFCrowd, worker must search task index beofer participating the task
				// Evaluating the time cost of task participation
				timeStart = time.Now()
				pkg.SearchTokenOnChain(client, instance, transactor, optsAddress, token, fbpie)
				seatoken, err := instance.GetReturnC(nil)
				if err != nil {
					log.Fatalln("Get search token error")
				}
				for _, t := range seatoken {
					instance.GetTaskindex(nil, t)
				}
				taskParticipationCost += time.Since(timeStart)
				timeStart = time.Now()
				w.ParticipantTask(r.Task())
				taskParticipationCost += time.Since(timeStart)

				timeStart = time.Now()
				w.CollectData(0, data)
				w.SubmitData(0)
				dataCollectionCost += time.Since(timeStart)
			}
			// Besides encrypted data, PFCrowd also uploads to the blockchain
			for _, index := range onChainIndex {
				onChainBytes += len(index)
				communicationCost += len(index)
			}
			for _, istring := range authTable {
				for _, str := range istring {
					onChainBytes += len([]byte(str))
					communicationCost += len([]byte(str))
				}
			}
			encryptedData := r.Task().Data()
			communicationCost += len(encryptedData[0])
			for _, d := range encryptedData {
				onChainBytes += len(d)
			}

			taskParticipationTimeCost[n] += taskParticipationCost / time.Duration(workerRequired)
			dataCollectionTimeCost[n] += dataCollectionCost / time.Duration(workerRequired)

			timeStart = time.Now()
			r.Rewarding(&PC{})
			rewardingCost = time.Since(timeStart)
			rewardingTimeCost[n] += rewardingCost

			onChainStorage[n] += onChainBytes
			communication[n] += communicationCost
		}
	}

	for i := 0; i < length; i++ {
		fmt.Printf("%v,%v,%v,%v,%v,%v,%v\n", registerationTimeCost[i]/time.Duration(numberOfIteration),
			taskPublicationTimeCost[i]/time.Duration(numberOfIteration), taskParticipationTimeCost[i]/time.Duration(numberOfIteration),
			dataCollectionTimeCost[i]/time.Duration(numberOfIteration), rewardingTimeCost[i]/time.Duration(numberOfIteration),
			float64(onChainStorage[i])/1024.0/float64(numberOfIteration), float64(communication[i])/1024.0/float64(numberOfIteration),
		)
	}
}

type PC struct {
}

func (cp *PC) CalculateRewards(t *task.Task, reward *big.Int, workerID int) *big.Int {
	return reward
}
