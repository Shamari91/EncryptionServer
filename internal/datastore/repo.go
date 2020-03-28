package datastore

var EncryptionMap map[string][]byte

func PersistEncryptionData(id string, data []byte) {
	EncryptionMap[id] = data
}

func RetrieveEncryptionData(id string) ([]byte, bool) {
	value, ok := EncryptionMap[id]
	return value, ok
}
