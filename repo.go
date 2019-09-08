package main

var encryptionMap map[string][]byte

func persistEncryptionData(id string, data []byte) {
	encryptionMap[id] = data
}

func retrieveEncryptionData(id string) ([]byte, bool) {
	value, ok := encryptionMap[id]
	return value, ok
}
