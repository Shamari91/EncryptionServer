package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type encryptionRequest struct {
	ID   string
	Data string
}

type encryptionResponse struct {
	Result string
	Key    []byte
}

func Encrypt(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestData encryptionRequest
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData := encryptionResponse{}
	if requestData.ID == "" {
		responseData.Result = "ID parameter is missing!"
	} else {
		encryptedData, key, err := encrypt([]byte(requestData.Data))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		responseData.Result = "Data encrypted succesfully!"
		responseData.Key = key

		persistEncryptionData(requestData.ID, encryptedData)
	}

	responseBody, err := json.Marshal(responseData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

type decryptionRequest struct {
	ID  string
	Key []byte
}

type decryptionResponse struct {
	Result string
	Data   string
}

func Decrypt(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var reqPayload decryptionRequest
	err = json.Unmarshal(reqBody, &reqPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData := decryptionResponse{}
	responseData.Result = "Decryption failed!"

	if reqPayload.ID == "" {
		responseData.Result = "ID parameter is missing!"
	} else if reqPayload.Key == nil {
		responseData.Result = "Key paramter is missing!"
	} else {
		encryptedData, idValid := retrieveEncryptionData(reqPayload.ID)
		if idValid {
			decryptedData, err := decrypt(encryptedData, reqPayload.Key)
			if err == nil {
				responseData.Data = decryptedData
			}
			responseData.Result = "Data decrypted succesfully!"
		} else {
			responseData.Result = "ID is invalid!"
		}
	}

	responseBody, err := json.Marshal(responseData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
