package main

import (
	"github.com/Shamari91/EncryptionServer/internal/datastore"
	"github.com/Shamari91/EncryptionServer/internal/server"
)

func main() {
	datastore.EncryptionMap = make(map[string][]byte)

	server.Init()
	server.Run()
}
