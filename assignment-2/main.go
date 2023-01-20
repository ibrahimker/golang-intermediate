package main

import (
	"context"
	"log"

	"github.com/ibrahimker/golang-intermediate/assignment-2/driver"
	"github.com/ibrahimker/golang-intermediate/assignment-2/handler"
	"github.com/ibrahimker/golang-intermediate/assignment-2/repository"
	"github.com/ibrahimker/golang-intermediate/assignment-2/router"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Renderer = driver.NewRenderer("template/*.html", true)
	ctx := context.Background()
	store := driver.NewPostgresStore()
	pgPool, err := driver.NewPostgresConn(ctx)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUser(pgPool)

	loginHandler := handler.NewLoginHandler(nil, store, userRepo)
	registerHandler := handler.NewRegisterHandler(nil, store, userRepo)

	router.SetupRouter(e, loginHandler, registerHandler)

	e.Start(":8080")
}
