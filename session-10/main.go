package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crewjam/saml/samlsp"
	"github.com/labstack/echo"
)

func landingHandler(w http.ResponseWriter, r *http.Request) {
	name := samlsp.AttributeFromContext(r.Context(), "displayName")
	w.Write([]byte(fmt.Sprintf("Welcome, %s!", name)))
}

func main() {
	sp, err := newSamlMiddleware()
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/index")
	})
	e.GET("/index", echo.WrapHandler(sp.RequireAccount(
		http.HandlerFunc(landingHandler),
	)))
	e.Any("/saml/", echo.WrapHandler(sp))

	portString := fmt.Sprintf(":%d", webserverPort)
	fmt.Println("server started at", portString)
	e.Start(portString)
}
