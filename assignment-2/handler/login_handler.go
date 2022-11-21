package handler

import (
	"net/http"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/ibrahimker/golang-intermediate/assignment-2/driver"
	"github.com/ibrahimker/golang-intermediate/assignment-2/entity"
	"github.com/ibrahimker/golang-intermediate/assignment-2/repository"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type HTMLTemplate struct {
	User entity.User
	Err  string
}

type Login struct {
	cs      *sessions.CookieStore
	pgStore *pgstore.PGStore
	ur      *repository.User
}

func NewLoginHandler(cs *sessions.CookieStore, pgs *pgstore.PGStore, ur *repository.User) *Login {
	return &Login{
		cs:      cs,
		pgStore: pgs,
		ur:      ur,
	}
}

func (l *Login) LoginHandler(c echo.Context) error {
	log.Info("Start Login Handler")
	// authenticate in db
	user, err := l.ur.GetByUsernamePassword(c.Request().Context(), c.FormValue("username"), c.FormValue("password"))
	if err != nil {
		log.Warn(err)
		return c.Render(http.StatusOK, "login.html", HTMLTemplate{
			User: entity.User{},
			Err:  "Error when login",
		})
	}
	log.Info("Successfully authenticate user", user)

	// store username in session
	if err = storeSessionHelper(c, l.pgStore, user.Username); err != nil {
		log.Warn("error when store session")
		return c.Render(http.StatusOK, "login.html", HTMLTemplate{
			User: entity.User{},
			Err:  "Error when store session",
		})
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/home")
}

func (l *Login) HomeHandler(c echo.Context) error {
	log.Info("Start Home Handler")

	session, err := l.pgStore.Get(c.Request(), driver.SESSION_ID)
	if err != nil {
		log.WithError(err).Warn("error when retrieve session")
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	if len(session.Values) == 0 {
		log.Info("no session values available")
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	log.Info("Session exists, render html")
	return c.Render(http.StatusOK, "home.html", HTMLTemplate{
		User: entity.User{Username: session.Values["username"].(string)},
		Err:  "",
	})
}

func (l *Login) LogoutHandler(c echo.Context) error {
	log.Info("Start Logout Handler")
	session, _ := l.pgStore.Get(c.Request(), driver.SESSION_ID)
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "/home")
}
