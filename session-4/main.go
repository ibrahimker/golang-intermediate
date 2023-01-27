package main

import (
	"fmt"
	"net/http"

	"github.com/go-ldap/ldap"
	"github.com/ibrahimker/golang-intermediate/session-4/config"
	"github.com/ibrahimker/golang-intermediate/session-4/repository"
	"github.com/ibrahimker/golang-intermediate/session-4/router"
	"github.com/ibrahimker/golang-intermediate/session-4/service"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

const (
	port    = "9000"
	baseURL = "0.0.0.0:" + port
)

func main() {
	log.SetReportCaller(true)

	// init ldap connection
	ldapConn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", config.LdapServer, config.LdapPort))
	if err != nil {
		panic(err)
	}
	defer ldapConn.Close()
	// bind to ldap server
	if err = ldapConn.Bind(config.LdapBindDN, config.LdapPassword); err != nil {
		panic(err)
	}

	ldapRepo := repository.NewLDAPRepo(ldapConn)
	loginService := service.NewLoginService(ldapRepo)

	r := mux.NewRouter()
	routerHandler := router.NewRouterHandler(loginService)
	r.HandleFunc("/", routerHandler.ServeHTML)
	r.HandleFunc("/login", routerHandler.HandleLogin)

	log.Infoln("Listening at ", baseURL)
	log.Fatal(http.ListenAndServe(baseURL, r))
}
