package handler

import (
	"net/http"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/ibrahimker/golang-intermediate/assignment-2/entity"
	"github.com/ibrahimker/golang-intermediate/assignment-2/repository"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type Register struct {
	cs  *sessions.CookieStore
	pgs *pgstore.PGStore
	ur  *repository.User
}

func NewRegisterHandler(cs *sessions.CookieStore, pgs *pgstore.PGStore, ur *repository.User) *Register {
	return &Register{
		cs:  cs,
		pgs: pgs,
		ur:  ur,
	}
}

func (r *Register) RegisterHandler(c echo.Context) error {
	data := "Hello from /html"
	u := &entity.User{
		Username:  c.FormValue("username"),
		FirstName: c.FormValue("firstname"),
		LastName:  c.FormValue("lastname"),
		Password:  c.FormValue("password"),
	}
	log.Info("user", u)
	if u.Username == "" || u.Password == "" {
		log.Warn("username password cannot be empty")
		return c.Redirect(http.StatusTemporaryRedirect, "/home")
	}
	if err := r.ur.Insert(c.Request().Context(), u); err != nil {
		return c.Render(http.StatusOK, "login.html", data)
	}

	// store username in session
	if err := storeSessionHelper(c, r.pgs, u.Username); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/home")
}
