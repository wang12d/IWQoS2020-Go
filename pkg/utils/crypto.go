package utils

import (
	"crypto/cipher"
	"log"
)

func AESECBEncryption(cipher cipher.Block, data []byte) []byte {
	if len(data)%cipher.BlockSize() != 0 {
		log.Fatalln("AES encryption data must be a multiple of block size.")
	}
	blockSize := cipher.BlockSize()
	encrypted := make([]byte, len(data))
	for start, end := 0, blockSize; start < len(data); start, end = start+blockSize, end+blockSize {
		cipher.Encrypt(encrypted[start:end], data[start:end])
	}
	return encrypted
}

func AESECBDecryption(cipher cipher.Block, data []byte) []byte {
	if len(data)%cipher.BlockSize() != 0 {
		log.Fatalln("AES encryption data must be a multiple of block size.")
	}
	blockSize := cipher.BlockSize()
	decrypted := make([]byte, len(data))
	for start, end := 0, blockSize; start < len(data); start, end = start+blockSize, end+blockSize {
		cipher.Decrypt(decrypted[start:end], data[start:end])
	}
	return decrypted
}
