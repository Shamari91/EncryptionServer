package main

import (
    "fmt"
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
)
 
func main() {
 
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/encrypt", Encrypt)
    router.HandleFunc("/decrypt", Decrypt)
 
    log.Fatal(http.ListenAndServe(":8080", router))
}
 
func Encrypt(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Encrypting!")
}
 
func Decrypt(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Decrypting!")
}
 