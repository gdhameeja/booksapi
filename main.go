package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"gdhameeja/booksapi/models"
	"github.com/gorilla/mux"
)

const somethingWentWrong = "somethingWentWrong"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/books", getBooksHandler).Methods("GET")
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
		http.Error(rw, somethingWentWrong, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(resp)
}

// return all books or a single one according to request params if any
func getBooksHandler(rw http.ResponseWriter, req *http.Request) {
	if !isAuthorized(req) {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var books []models.Book
	param := req.URL.Query()
	if stringId, ok := param["id"]; ok {
		id, err := strconv.Atoi(stringId[0])
		if err != nil {
			http.Error(rw, "Id should be int", http.StatusBadRequest)
		}
		books = append(books, models.GetBook(id))
	} else {
		books = models.GetAllBooks()
	}
	resp, err := json.Marshal(books)
	if err != nil {
		log.Print(err)
		http.Error(rw, somethingWentWrong, http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(resp)
}


func isAuthorized(req *http.Request) bool {
	log.Print(req.Header)
	if token, ok := req.Header["Authorization"]; ok {
		user := models.GetUserForToken(token[0])
		return user != nil
	}
	return false
}
