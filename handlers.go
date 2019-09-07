package main
 
import (
	"fmt"
    "log"
    "net/http"
    "io/ioutil"
	"encoding/json"
)
 
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
  
  type decryptionReq struct {
	  ID  string
	  Key []byte
  }
  
  type decryptionResponse struct {
	  Result string
	  Data   string
  }
  
  func Decrypt(w http.ResponseWriter, r *http.Request) {
	  reqBody, err := ioutil.ReadAll(r.Body)
	  defer r.Body.Close()
	  if err != nil {
		  http.Error(w, err.Error(), 500)
		  return
	  }
  
	  var reqPayload decryptionReq
	  err = json.Unmarshal(reqBody, &reqPayload)
	  if err != nil {
		  http.Error(w, err.Error(), 500)
		  return
	  }
  
	  resultPayload := decryptionResponse{}
	  if reqPayload.ID == "" {
		  resultPayload.Result = "ID parameter is missing!"
	  } else if reqPayload.Key == nil {
		  resultPayload.Result = "Key paramter is missing!"
	  } else {
		  encryptedData := retrieveEncryptionData(reqPayload.ID)
		  if encryptedData == nil {
			  resultPayload.Data = ""
		  } else {
			  decryptedData, err := decrypt(encryptedData, reqPayload.Key)
			  if err != nil {
				  fmt.Printf("Error occured while trying to decrypt the data: %v \n", err)
				  resultPayload.Data = ""
			  } else {
				  resultPayload.Data = decryptedData
			  }
		  }
		  resultPayload.Result = "Data decrypted succesfully!"
	  }
  
	  responseBody, err := json.Marshal(resultPayload)
	  w.Header().Set("Content-Type", "application/json")
	  w.Write(responseBody)
  }
   