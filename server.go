package main

import (
    "log"
    "net/http"
	"github.com/gorilla/mux"
)
 
func main() {
	encryptionMap = make(map[string][]byte)
	
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/encrypt", Encrypt)
    router.HandleFunc("/decrypt", Decrypt)
 
    log.Fatal(http.ListenAndServe(":8080", router))
}
 
