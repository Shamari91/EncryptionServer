package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleEncryptionStatusOK(t *testing.T) {
	encryptionMap = make(map[string][]byte)

	var reqBody = []byte(`{"ID": "1", "Data": "Yoti"}`)
	req, err := http.NewRequest("POST", "/encrypt", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Encrypt)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleEncryptionbBadRequest(t *testing.T) {
	var reqBody = []byte(`{"ID": , "Data": "Yoti"}`)
	req, err := http.NewRequest("POST", "/encrypt", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Encrypt)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestHandleDecryptionStatusOK(t *testing.T) {
	var reqBody = []byte(`{"ID": "1", "Key": "de4444=="}`)
	req, err := http.NewRequest("POST", "/decrypt", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Decrypt)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleDecryptionBadRequest(t *testing.T) {
	var reqBody = []byte(`{"ID": "1", "Key": `)
	req, err := http.NewRequest("POST", "/decrypt", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Decrypt)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
