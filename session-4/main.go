package main

import (
	"fmt"
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/ibrahimker/golang-intermediate/session-4/ldap"
)

const (
	port    = "9000"
	baseURL = "0.0.0.0:" + port
)

func main() {
	log.SetReportCaller(true)
	r := mux.NewRouter()
	r.HandleFunc("/", handleRoute)
	r.HandleFunc("/login", handleLogin)

	log.Infoln("Listening at ", baseURL)
	log.Fatal(http.ListenAndServe(baseURL, r))
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("login.html")
	if err := parsedTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	ok, data, err := ldap.AuthUsingLDAP(username, password)
	if !ok {
		log.Error("auth using ldap not ok")
		http.Error(w, "invalid username/password", http.StatusUnauthorized)
		return
	}
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	message := fmt.Sprintf("Welcome %s\n", data.FullName)
	w.Write([]byte(message))
}
