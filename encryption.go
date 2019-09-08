package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

var IV = []byte("1837435475840435")

func createKey() ([]byte, error) {
	genkey := make([]byte, 16)
	_, err := rand.Read(genkey)
	if err != nil {
		return nil, err
	}
	return genkey, nil
}

func createCipher(key []byte) (cipher.Block, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return blockCipher, nil
}

func encrypt(data []byte) ([]byte, []byte, error) {
	generatedKey, err := createKey()
	if err != nil {
		return nil, nil, err
	}

	blockCipher, err := createCipher(generatedKey)
	if err != nil {
		return nil, nil, err
	}

	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(data, data)
	return data, generatedKey, nil
}

func decrypt(cipherText []byte, key []byte) (string, error) {
	blockCipher, err := createCipher(key)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}
