package main

import (
    "net/http"
	"net/http/httptest"
    "testing"
	"bytes"
	"encoding/json"
)

// This file will be a integration/Acceptence test where we test the full functionlty of encrypting/decrypting

func TestEncryptionService(t *testing.T) {
	encryptionMap = make(map[string][]byte)

	responseRecorder := SendEncryptionRequest(t)

	response, err := ConvertEncryptionResponseToStructure(responseRecorder)
	if err != nil {
		t.Error("Error parsing the response")
		return
	}

	decryptionRequestBody, err := ConvertDecryptionRequestToJSON(response)
	if err != nil {
		t.Error("Error converting decryption request to JSON")
		return
	}

	responseRecorder = SendDecryptionRequest(t, decryptionRequestBody)

	CheckResult(t, responseRecorder)
}

func SendEncryptionRequest(t *testing.T) (*httptest.ResponseRecorder) {
	var reqBody = []byte(`{"ID": "1", "Data": "Yoti"}`)
    request, err := http.NewRequest("POST", "/encrypt", bytes.NewBuffer(reqBody))
    if err != nil {
        t.Fatal(err)
    }

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(Encrypt)
    handler.ServeHTTP(responseRecorder, request)

    if status := responseRecorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
	}

	return responseRecorder
}

func SendDecryptionRequest(t *testing.T, requestBody []byte) (*httptest.ResponseRecorder) {
	request, err := http.NewRequest("POST", "/decrypt", bytes.NewBuffer(requestBody))
    if err != nil {
        t.Fatal(err)
    }

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(Decrypt)
    handler.ServeHTTP(responseRecorder, request)

    if status := responseRecorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
	}
	
	return responseRecorder
}

func CheckResult(t *testing.T, responseRecorder *httptest.ResponseRecorder){
	expected := `{"Result":"Data decrypted succesfully!","Data":"Yoti"}`
    if responseRecorder.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
			responseRecorder.Body.String(), expected)
    }
}

func ConvertEncryptionResponseToStructure(responseRecorder *httptest.ResponseRecorder) (encryptionResponse, error) {
	var response encryptionResponse
	encryptionResponseBody := responseRecorder.Body.String()
	err := json.Unmarshal([]byte(encryptionResponseBody), &response)
	if err != nil {
		return encryptionResponse{}, err
	}
	return response, nil
}

func ConvertDecryptionRequestToJSON(response encryptionResponse) ([]byte, error){
	decryptionRequest := decryptionRequest{}
	decryptionRequest.ID = "1"
	decryptionRequest.Key = response.Key

	decryptionRequestBody, err := json.Marshal(decryptionRequest)
	if err != nil {
		return nil, err	
	}
    return decryptionRequestBody, nil
}
