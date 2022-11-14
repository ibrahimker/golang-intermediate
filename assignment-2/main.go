package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ibrahimker/golang-intermediate/assignment-2/driver"
	"github.com/ibrahimker/golang-intermediate/assignment-2/handler"
	"github.com/ibrahimker/golang-intermediate/assignment-2/repository"
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

	e.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/home")
	})

	e.GET("/login", loginHandler.LoginHandler)
	e.POST("/login", loginHandler.LoginHandler)

	e.GET("/home", loginHandler.HomeHandler)
	e.POST("/home", loginHandler.HomeHandler)

	e.POST("/register", registerHandler.RegisterHandler)

	e.POST("/logout", loginHandler.LogoutHandler)

	e.Start(":8080")
}
