package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"./routes"
	"./util"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/login", routes.LoginHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/worksheets/{ws}/{page:[0-9]+}", routes.WorksheetPageHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/worksheets", routes.WorksheetListHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/submit", routes.WorksheetActivityWriteHandler).Methods("POST", "OPTIONS")

	http.Handle("/", r)

	fmt.Println("GoGoGrader server started!")

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(util.Config.Port), nil))
}
