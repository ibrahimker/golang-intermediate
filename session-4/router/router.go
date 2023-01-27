package router

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ibrahimker/golang-intermediate/session-4/service"
)

type RouterHandler struct {
	loginService *service.LoginSvc
}

func NewRouterHandler(loginService *service.LoginSvc) *RouterHandler {
	return &RouterHandler{loginService: loginService}
}

func (rh *RouterHandler) ServeHTML(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("login.html")
	if err := parsedTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rh *RouterHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	data, err := rh.loginService.Authenticate(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	message := fmt.Sprintf("Welcome Name: %s, FullName: %s, Email: %s\n", data.Name, data.FullName, data.Email)
	w.Write([]byte(message))
}
