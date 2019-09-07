package main

var encryptionMap map[string][]byte

func persistEncryptionData(id string, data []byte) {
	encryptionMap[id] = data
}

func retrieveEncryptionData(id string) []byte {
	return encryptionMap[id]
}
