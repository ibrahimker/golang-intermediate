package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/ibrahimker/golang-intermediate/assignment-1/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Todo struct {
	Name string `json:"name"`
}

var Todos []*Todo

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
	log.Println("Listening in url " + baseURL)
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
	var t Todo
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&t)
	Todos = append(Todos, &t)
	w.Write([]byte("Success add todo " + t.Name))
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
	todosRes, _ := json.Marshal(Todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(todosRes)
}
