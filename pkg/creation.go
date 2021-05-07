package pkg

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
)

type BrokerKey struct {
	Fbpie string
	Fb1   string
	Fb2   string
}

var (
	crowdsourcingCategory = []string{
		"Accounting", "Financial_Planning", "Human_Resource",
		"Management Consulting", "Data Entry", "Project Management",
		"Transcription", "Web Research", "Customer Service", "Technical Support",
		"Data Extraction", "Data Visualization", "Machine Learning", "Animation",
		"Audio Production", "Motion Graphics", "Photography", "Information Security",
		"Contract Law", "Criminal Law",
	}
)

// GenerateBrokers generate a list broker map, which contains numberOfBrokers brokers, each of
// contains the numberOfTasks task index.
func GenerateBrokers(numberOfBrokers, numberOfTasks int) ([]map[string][][]byte, error) {
	if numberOfBrokers <= 0 || numberOfTasks <= 0 {
		return nil, errors.New("brokers or tasks must greater than 0")
	}
	brokersMap := make([]map[string][][]byte, numberOfBrokers)
	index := new(big.Int)
	for i := 0; i < numberOfBrokers; i++ {
		brokersMap[i] = make(map[string][][]byte)
		for _, key := range crowdsourcingCategory {
			index.SetBytes([]byte(key))
			tasks := make([][]byte, numberOfTasks)
			for j := 0; j < numberOfTasks; j++ {
				tasks[j] = make([]byte, 32)
				_, err := rand.Read(tasks[j])
				if err != nil {
					return nil, err
				}
			}
			brokersMap[i][index.String()] = tasks
		}
	}
	return brokersMap, nil
}

// GenerateBrokerKeys generates the numberOfBrokers keys, each contains
// Fb1, Fb2 and Fbpie, which used to construct trapdoor and used as index
func GenerateBrokerKeys(numberOfBrokers int) []BrokerKey {
	brokerKeys := make([]BrokerKey, numberOfBrokers)
	for i := 1; i <= numberOfBrokers; i++ {
		brokerKeys[i-1] = getKeys(i)
	}
	return brokerKeys
}

func getKeys(id int) BrokerKey {
	masterKey1 := hmac.New(md5.New, []byte(fmt.Sprintf("b%vmaster1", id))).Sum(nil)
	masterKey2 := hmac.New(md5.New, []byte(fmt.Sprintf("b%vmaster2", id))).Sum(nil)
	secretKey := hmac.New(md5.New, []byte(fmt.Sprintf("b%vsecret", id))).Sum(nil)
	brokerID := []byte(fmt.Sprintf("broker%vID0000000", id))
	Fbpie := hmac.New(sha256.New, masterKey1)
	Fbpie.Write(brokerID)
	Fb1 := hmac.New(sha256.New, masterKey2)
	Fb1.Write(brokerID)
	Fb2 := hmac.New(sha256.New, secretKey)
	Fb2.Write(brokerID)
	tmp := new(big.Int)
	tmp.SetBytes(Fb1.Sum(nil))
	res := BrokerKey{}
	res.Fb1 = tmp.String()
	tmp.SetBytes(Fb2.Sum(nil))
	res.Fb2 = tmp.String()
	tmp.SetBytes(Fbpie.Sum(nil))
	res.Fbpie = tmp.String()
	return res
}
