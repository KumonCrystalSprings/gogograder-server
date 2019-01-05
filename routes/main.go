package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../auth"
	"../db"
	"../util"

	"github.com/gorilla/mux"
)

func setDefaultHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// LoginData is the data that should be given when a user tries to login
type LoginData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(&w)
	decoder := json.NewDecoder(r.Body)
	var l LoginData
	err := decoder.Decode(&l)
	if util.CheckError(err, "Error parsing login data") {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[ERROR] Error parsing login data")
		return
	}

	id, ok, err := auth.Login(l.Name, l.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[ERROR] %v", err)
	} else if !ok {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		json.NewEncoder(w).Encode(id)
	}
}

func authSession(w http.ResponseWriter, r *http.Request) bool {
	id := r.URL.Query().Get("id")

	if id == "" || !auth.VerifySession(id) {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}

func WorksheetListHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(&w)
	if !authSession(w, r) {
		return
	}

	sheets, err := db.FetchWorksheetNames()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[ERROR] %v", err)
		return
	}

	json.NewEncoder(w).Encode(sheets)
}

func WorksheetPageHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(&w)
	if !authSession(w, r) {
		return
	}

	vars := mux.Vars(r)
	page, ok, err := db.FetchWorksheetPage(vars["ws"], vars["page"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[ERROR] %v", err)
		return
	} else if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		json.NewEncoder(w).Encode(page)
	}
}

func WorksheetActivityWriteHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(&w)
	if !authSession(w, r) {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var a db.StudentActivity
	err := decoder.Decode(&a)
	if util.CheckError(err, "Error parsing student activity data") {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[ERROR] Error parsing student activity data")
		return
	}

	name, _ := auth.GetSession(r.URL.Query().Get("id"))

	a.Name = name
	a.Date = db.JSONDate(time.Now())

	err = db.WriteStudentActivity(&a)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[ERROR] Error parsing student activity data")
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
