package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./auth"
	"./models"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var l models.LoginModel
	err := decoder.Decode(&l)
	if err != nil {
		log.Println(err)
	}

	if id, ok := auth.Login(l.Name, l.ID); ok {
		json.NewEncoder(w).Encode(models.LoginResponseModel{SessionID: id})
	}
}
