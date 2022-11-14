package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type M map[string]interface{}

type User struct {
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	r := echo.New()
	r.Validator = &CustomValidator{validator: validator.New()}

	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/index", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/html", func(ctx echo.Context) error {
		data := "Hello from /html"
		return ctx.HTML(http.StatusOK, data)
	})

	r.GET("/index", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	r.GET("/json", func(ctx echo.Context) error {
		data := M{"Message": "Hello", "Counter": 2}
		return ctx.JSON(http.StatusOK, data)
	})

	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")

		data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)

		return ctx.String(http.StatusOK, data)
	})

	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		data := fmt.Sprintf(
			"Hello %s, I have message for you: %s",
			name,
			strings.Replace(message, "/", "", 1),
		)

		return ctx.String(http.StatusOK, data)
	})

	r.POST("/user", func(ctx echo.Context) error {
		u := new(User)

		if err := ctx.Bind(u); err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		if err := ctx.Validate(u); err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}

		return ctx.JSON(http.StatusOK, u)
	})

	r.Start(":9000")
}
