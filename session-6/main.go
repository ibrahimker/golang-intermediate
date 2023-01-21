package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const SESSION_ID = "test-session-id"

func newPostgresStore() *pgstore.PGStore {
	url := "postgres://postgresuser:postgrespassword@127.0.0.1:5432/postgres?sslmode=disable"
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	return store
}

func newCookieStore() *sessions.CookieStore {
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true

	return store
}

func main() {
	e := echo.New()
	store := newCookieStore()

	root := e.Group("/")        // /
	todo := root.Group("/todo") // /todo
	todo.Use()
	_ = todo.Group("/user", middleware.Logger()) // /todo/user
	e.Group("/user")
	e.GET("/", func(ctx echo.Context) error {
		data := "Hello from /"
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
