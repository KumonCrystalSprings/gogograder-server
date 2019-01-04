package main

import (
	"fmt"
	"log"
	"net/http"

	"./routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", routes.LoginHandler).Methods("POST")
	r.HandleFunc("/worksheets/{ws}/{page:[0-9]+}", routes.WorksheetPageHandler).Methods("GET")
	r.HandleFunc("/worksheets", routes.WorksheetListHandler).Methods("GET")
	r.HandleFunc("/submit", routes.WorksheetActivityWriteHandler).Methods("POST")

	http.Handle("/", r)

	fmt.Println("GoGoGrader server started!")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
