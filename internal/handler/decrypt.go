package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Shamari91/EncryptionServer/internal/datastore"
	"github.com/Shamari91/EncryptionServer/internal/encryption"
	"github.com/julienschmidt/httprouter"
)

type decryptionRequest struct {
	ID  string
	Key []byte
}

type decryptionResponse struct {
	Result string
	Data   string
}

func Decrypt(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := &decryptionRequest{}
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateDecryptBody(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData := decryptionResponse{}
	encryptedData, idValid := datastore.RetrieveEncryptionData(req.ID)
	if idValid {
		decryptedData, err := encryption.Decrypt(encryptedData, req.Key)
		if err == nil {
			responseData.Data = decryptedData
		}
		responseData.Result = "Data decrypted succesfully!"
	} else {
		responseData.Result = "ID is invalid!"
	}

	responseBody, err := json.Marshal(responseData)
	if err != nil {
		log.Printf("error marshaling response :: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Printf("error writing response :: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func validateDecryptBody(req *decryptionRequest) error {
	if req.ID == "" {
		return errors.New("missing identification number")
	}

	if req.Key == nil {
		return errors.New("missing encrypting key")
	}

	return nil
}
