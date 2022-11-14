package handler

import (
	"net/http"

	"github.com/antonlindstrom/pgstore"
	"github.com/ibrahimker/golang-intermediate/assignment-2/driver"
	"github.com/labstack/echo"
)

func storeSessionHelper(c echo.Context, pgs *pgstore.PGStore, username string) error {
	session, err := pgs.Get(c.Request(), driver.SESSION_ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	session.Values["username"] = username
	if err = session.Save(c.Request(), c.Response()); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}
