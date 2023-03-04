package middleware

import (
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func BasicAuthMiddleware(username, password string, c echo.Context) (bool, error) {
	// Be careful to use constant time comparison to prevent timing attacks
	if username == viper.GetString("BasicAuthUser") && password == viper.GetString("BasicAuthPass") {
		return true, nil
	}
	return false, nil
}
