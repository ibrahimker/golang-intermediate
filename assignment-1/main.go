package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/ibrahimker/golang-intermediate/assignment-1/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

const baseURL = "0.0.0.0:8080"

// @title Todo Application
// @version 1.0
// @description This is a todo list test management application
// @contact.name Ibrahim Nurandita Isbandiputra
// @contact.email ibrahimker@gmail.com
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", Create).Methods(http.MethodPost)
	r.HandleFunc("/todos", Get).Methods(http.MethodGet)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// serve http server
	log.Fatal(http.ListenAndServe(baseURL, r))
}

// Create is a handler for create todos API
// @Summary Create new todos
// @Description get string by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos [post]
func Create(w http.ResponseWriter, r *http.Request) {

}

// Get is a handler for get todos API
// @Summary Get new todos
// @Description get all todo list
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos [get]
func Get(w http.ResponseWriter, r *http.Request) {

}
