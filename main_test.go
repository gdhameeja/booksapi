package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestLogin(t *testing.T) {
	requestData := []byte{`{"username": "gauravdhameeja", "password": "password"}`}
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
