package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
	"github.com/gorilla/mux"
	"encoding/json"
)
 
func main() {
	encryptionMap = make(map[string][]byte)
	
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/encrypt", Encrypt)
    router.HandleFunc("/decrypt", Decrypt)
 
    log.Fatal(http.ListenAndServe(":8080", router))
}
 

type encryptionReq struct {
  ID string
  Data string
}

type encryptionResponse struct {
	Result string
	Key []byte 
}

func Encrypt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Encrypting!")
	
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
      http.Error(w, "can't read body", http.StatusBadRequest)
      return
	}

	var requestData encryptionReq
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	responseData := encryptionResponse{}
    if requestData.ID == "" {
		responseData.Result = "ID parameter is missing!"
	} else {
		dataBytes := []byte(requestData.Data)
		encryptedData, key, err := encrypt(dataBytes)
		if err != nil {
			log.Printf("Error when encrypting: %v", err)
			http.Error(w, err.Error(), 500)
			return
		}

		log.Printf("Encrypted data: %v, Key Is: %v", encryptedData, key)

		responseData.Result = "Data encrypted succesfully!"
		responseData.Key = key

		persistEncryptionData(requestData.ID, encryptedData)
	}

	responseBody, err := json.Marshal(responseData)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
 
func Decrypt(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Decrypting!")
}
 