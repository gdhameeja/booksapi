package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gdhameeja/booksapi/models"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	var reqData map[string]string

	err := json.NewDecoder(req.Body).Decode(&reqData)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	log.Print(reqData)
	var username, password string
	var ok bool
	if username, ok = reqData["username"]; !ok {
		http.Error(rw, "Bad Request: Required username", http.StatusBadRequest)
		return
	}
	if password, ok = reqData["password"]; !ok {
		http.Error(rw, "Bad Request: Required password", http.StatusBadRequest)
		return
	}

	token, err := models.Login(username, password)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if token == "" {
		http.Error(rw, "Wrong username/password", http.StatusUnauthorized)
		return
	}

	response := map[string]string{"token": token}
	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(resp)
}
