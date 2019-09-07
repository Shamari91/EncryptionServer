package main
 
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "log"
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
		log.Fatalf("Failed to generate AES key: %s", err)
	}

	blockCipher, err := createCipher(generatedKey)
	if err != nil {
		log.Fatalf("Failed to create the AES cipher: %s", err)
	}

    stream := cipher.NewCTR(blockCipher, IV)
    stream.XORKeyStream(data, data)
	
	return data, generatedKey, nil
}
