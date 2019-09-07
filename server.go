package main

import (
    "log"
    "net/http"
)
 
func main() {
	encryptionMap = make(map[string][]byte)
	
	router := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
 
