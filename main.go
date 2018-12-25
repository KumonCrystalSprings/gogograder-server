package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./auth"
	"./util"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// LoginData is the data that should be given when a user tries to login
type LoginData struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var l LoginData
	err := decoder.Decode(&l)
	if util.CheckError(err, "Error parsing login data") {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	id, ok, err := auth.Login(l.Name, l.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
	} else if !ok {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		json.NewEncoder(w).Encode(id)
	}
}
