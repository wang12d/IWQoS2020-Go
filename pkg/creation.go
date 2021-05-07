package pkg

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

type BrokerKey struct {
	Fb1   string
	Fb2   string
	Fbpie string
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

func GenerateBrokerKeys(numberOfBrokers int)

func getKeys(id int) BrokerKey {
	masterKey1 := hmac.New(md5.New, []byte(fmt.Sprintf("b%vmaster1", id))).Sum(nil)
	masterKey2 := hmac.New(md5.New, []byte(fmt.Sprintf("b%vmaster2", id))).Sum(nil)
	secretKey := hmac.New(md5.New, []byte(fmt.Sprintf("b%vsecret", id))).Sum(nil)
	brokerID := []byte(fmt.Sprintf("broker%vID0000000", id))
	Fb1 := hmac.New(md5.New, []byte(masterKey1)).Sum(brokerID)
	Fb2 := hmac.New(md5.New, []byte(masterKey2)).Sum(brokerID)
	Fbpie := hmac.New(md5.New, []byte(secretKey)).Sum(brokerID)
	tmp := new(big.Int)
	tmp.SetBytes(Fb1)
	res := BrokerKey{}
	res.Fb1 = tmp.String()
	tmp.SetBytes(Fb2)
	res.Fb2 = tmp.String()
	tmp.SetBytes(Fbpie)
	res.Fbpie = tmp.String()
	return res
}
