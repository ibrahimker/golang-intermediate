package router

import (
	"net/http"

	"github.com/ibrahimker/golang-intermediate/assignment-2/handler"
	"github.com/labstack/echo"
)

func SetupRouter(e *echo.Echo, loginHandler *handler.Login, registerHandler *handler.Register) {
	e.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/home")
	})

	e.GET("/login", loginHandler.LoginHandler)
	e.POST("/login", loginHandler.LoginHandler)

	e.GET("/home", loginHandler.HomeHandler)
	e.POST("/home", loginHandler.HomeHandler)

	e.POST("/register", registerHandler.RegisterHandler)

	e.POST("/logout", loginHandler.LogoutHandler)
}
