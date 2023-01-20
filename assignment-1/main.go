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
	ID   string `json:"id"`
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

	r.HandleFunc("/todos", Get).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", GetByID).Methods(http.MethodGet)
	r.HandleFunc("/todos", Create).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", Put).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", Delete).Methods(http.MethodDelete)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	//http.Server{
	//	Addr:    baseURL,
	//	Handler: r,
	//}

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

// Put is a handler for create todos API
// @Summary Update new todos
// @Description Update todos by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos/{id} [put]
func Put(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i := 0; i < len(Todos); i++ {
		if Todos[i].ID == id {
			var t Todo
			decoder := json.NewDecoder(r.Body)
			_ = decoder.Decode(&t)
			Todos[i] = &t
			w.Write([]byte("Success update todo " + t.ID))
			return
		}
	}
}

// Delete is a handler for create todos API
// @Summary Delete new todos
// @Description Delete string by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i := 0; i < len(Todos); i++ {
		if Todos[i].ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			w.Write([]byte("Success delete todo " + id))
			return
		}
	}
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

// GetByID is a handler for get todos API
// @Summary GetByID new todos
// @Description get all todo list
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos/{id} [get]
func GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i := 0; i < len(Todos); i++ {
		if Todos[i].ID == id {
			todosRes, _ := json.Marshal(Todos[i])
			w.Header().Set("Content-Type", "application/json")
			w.Write(todosRes)
		}
	}
}
