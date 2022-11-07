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
	view    = `
		<html>
		<head>
			<title>Template</title>
		</head>
		<body>
			<form method="post" action="/login">
				<div>
					<label>username</label>
					<input type="text" name="username" required/>
				</div>
				<div>
					<label>password</label>
					<input type="password" name="password" required/>
				</div>
				<button type="submit">Login</button>
			</form>
		</body>
		</html>
	`
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
	tmpl := template.Must(template.New("main-template").Parse(view))
	if err := tmpl.Execute(w, nil); err != nil {
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
