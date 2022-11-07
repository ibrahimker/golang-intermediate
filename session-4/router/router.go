package router

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ibrahimker/golang-intermediate/session-4/service"
)

func ServeHTML(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("../login.html")
	if err := parsedTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	data, err := service.Authenticate(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	message := fmt.Sprintf("Welcome Name: %s, FullName: %s, Email: %s\n", data.Name, data.FullName, data.Email)
	w.Write([]byte(message))
}
