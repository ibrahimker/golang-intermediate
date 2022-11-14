package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

const SESSION_ID = "test-session-id"

func main() {
	e := echo.New()
	store := sessions.NewCookieStore([]byte("test-session-key"))

	e.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	e.GET("/set", func(c echo.Context) error {
		session, err := store.Get(c.Request(), SESSION_ID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		session.Values["message1"] = "hello"
		session.Values["message2"] = "world"
		if err := session.Save(c.Request(), c.Response()); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.GET("/get", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)

		if len(session.Values) == 0 {
			return c.String(http.StatusOK, "empty result")
		}

		return c.String(http.StatusOK, fmt.Sprintf(
			"%s %s",
			session.Values["message1"],
			session.Values["message2"],
		))
	})

	e.GET("/delete", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Options.MaxAge = -1
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.Start(":9000")
}
