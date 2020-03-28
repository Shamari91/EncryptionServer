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

type encryptionRequest struct {
	ID   string
	Data string
}

type encryptionResponse struct {
	Result string
	Key    []byte
}

func Encrypt(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body :: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := &encryptionRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("failed to unmarshal request :: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateEncryptBody(req)
	if err != nil {
		log.Printf("failed to validate request :: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encryptedData, key, err := encryption.Encrypt([]byte(req.Data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	datastore.PersistEncryptionData(req.ID, encryptedData)

	responseData := &encryptionResponse{}
	responseData.Result = "Data encrypted succesfully!"
	responseData.Key = key

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

func validateEncryptBody(req *encryptionRequest) error {
	if req.ID == "" {
		return errors.New("missing identification number")
	}

	if req.Data == "" {
		return errors.New("missing data to encrypt")
	}

	return nil
}
