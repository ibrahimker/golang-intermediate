package main

import (
	"net/http"

	"github.com/ibrahimker/golang-intermediate/session-4/router"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

const (
	port    = "9000"
	baseURL = "0.0.0.0:" + port
)

func main() {
	log.SetReportCaller(true)
	r := mux.NewRouter()
	r.HandleFunc("/", router.ServeHTML)
	r.HandleFunc("/login", router.HandleLogin)

	log.Infoln("Listening at ", baseURL)
	log.Fatal(http.ListenAndServe(baseURL, r))
}
