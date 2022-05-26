package marlin

// #cgo LDFLAGS: -lmarlin_zsk -L${SRCDIR}/../lib
// #include <stdbool.h>
// #include <stdlib.h>
// #include "./../lib/marlin_zsk.h"
// DataEvaluationResult newDataEvaluationResult(unsigned long long add, unsigned long long minus) {
//		return (DataEvaluationResult) {add, minus};
// }
import "C"
import (
	"encoding/hex"
	"log"
)

// ZebraLancerGenerateProofAndVerifyKey generates the zero-knowledge proof desribed by
// ICDCS paper ZebraLancer
func ZebraLancerGenerateProofAndVerifyKey(prefix, msg, sk, pk, cert, mpk, t1, t2 []byte) (proof Proof, verifyKey VerifyKey) {
	hexPrefix := C.CString(hex.EncodeToString(prefix))
	hexMsg := C.CString(hex.EncodeToString(msg))
	hexSk := C.CString(hex.EncodeToString(sk))
	hexPk := C.CString(hex.EncodeToString(pk))
	hexCert := C.CString(hex.EncodeToString(cert))
	hexMpk := C.CString(hex.EncodeToString(mpk))
	hexT1 := C.CString(hex.EncodeToString(t1))
	hexT2 := C.CString(hex.EncodeToString(t2))
	witness := C.ZebraLancerWitness{
		sk:   hexSk,
		pk:   hexPk,
		cert: hexCert,
	}
	proofAndKey := C.generate_proof_zebralancer(hexPrefix, hexMsg, hexMpk, hexT1, hexT2, witness)
	defer C.free_proof_and_verify(proofAndKey.proof, proofAndKey.vk)
	encodedProof, encodedVerifyKey := C.GoString(proofAndKey.proof), C.GoString(proofAndKey.vk)
	var err error
	if proof, err = hex.DecodeString(encodedProof); err != nil {
		log.Fatalln("Hex decode error")
	}
	if verifyKey, err = hex.DecodeString(encodedVerifyKey); err != nil {
		log.Fatalln("Hex decode error")
	}
	return proof, verifyKey
}

// ZebraLancerVerifyProof verifys the zero knowledge proof generated by requester with
// its public inputs
func ZebraLancerVerifyProof(prefix, msg, mpk, t1, t2 []byte, proof Proof, verifyKey VerifyKey) bool {
	hexT1, hexT2 := C.CString(hex.EncodeToString(t1)), C.CString(hex.EncodeToString(t2))
	hexPrefix, hexMsg, hexMpk := C.CString(hex.EncodeToString(prefix)), C.CString(hex.EncodeToString(msg)), C.CString(hex.EncodeToString(mpk))
	proofHex, verifyKeyHex := C.CString(hex.EncodeToString(proof)), C.CString(hex.EncodeToString(verifyKey))
	return bool(C.verify_proof_zebralancer(hexPrefix, hexMsg, hexMpk, hexT1, hexT2, proofHex, verifyKeyHex))
}

// GenerateEncryptionZKProofAndVerifyKey generates the zero-knowledge proof of encrypted data
func GenerateEncryptionZKProofAndVerifyKey(mu, sigmaSquare uint, data []uint,
	publicKey, privateKey []byte, encryptedData, rawData [][]byte, dataSize uint) (proof Proof, verifyKey VerifyKey) {
	/****************************************
			Convert to C String
	****************************************/
	size := len(rawData)
	hexRawData, hexEncryptedData := make([]*C.char, size), make([]*C.char, size)
	hexPublicKey := C.CString(hex.EncodeToString(publicKey))
	hexPrivateKey := C.CString(hex.EncodeToString(privateKey))
	uData := make([]C.uint, size)
	for i := 0; i < size; i++ {
		hexRawData[i] = C.CString(hex.EncodeToString(rawData[i]))
		hexEncryptedData[i] = C.CString(hex.EncodeToString(encryptedData[i]))
		uData[i] = C.uint(data[i])
	}

	proofAndKey := C.generate_proof_zebralancer_rewarding(C.uint(mu), C.uint(sigmaSquare), &uData[0], C.uint(size),
		hexPublicKey, hexPrivateKey, &hexEncryptedData[0], &hexRawData[0], C.uint(dataSize))
	defer C.free_proof_and_verify(proofAndKey.proof, proofAndKey.vk)
	encodedProof, encodedVerifyKey := C.GoString(proofAndKey.proof), C.GoString(proofAndKey.vk)
	var err error
	if proof, err = hex.DecodeString(encodedProof); err != nil {
		log.Fatalln("Hex decode error")
	}
	if verifyKey, err = hex.DecodeString(encodedVerifyKey); err != nil {
		log.Fatalln("Hex decode error")
	}
	return proof, verifyKey
}

// VerifyEncryptionZKProof verifies the proof generated by the requester
func VerifyEncryptionZKProof(evals []EvaluationResults, ciphertext [][]byte, proof, vk []byte) bool {
	/****************************************
			Convert to C String
	****************************************/
	size := len(ciphertext)
	dataEvals := make([]C.DataEvaluationResult, size)
	hexCiphertext := make([]*C.char, size)
	for i := 0; i < size; i++ {
		dataEvals[i] = C.newDataEvaluationResult(C.ulonglong(evals[i][0]), C.ulonglong(evals[i][1]))
		hexCiphertext[i] = C.CString(hex.EncodeToString(ciphertext[i]))
	}
	hexProof := C.CString(hex.EncodeToString(proof))
	hexVK := C.CString(hex.EncodeToString(vk))
	return bool(C.verify_proof_zebralancer_rewarding(&dataEvals[0], C.uint(size), &hexCiphertext[0], hexProof, hexVK))
}