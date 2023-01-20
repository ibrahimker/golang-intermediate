package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/ibrahimker/golang-intermediate/session-5/driver"
	"github.com/labstack/echo"
)

type M map[string]interface{}

type User struct {
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
}

type Company struct {
	Name string `json:"name" form:"name" query:"name" validate:"required"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type Template struct {
	Nama string
}

func main() {
	r := echo.New()
	r.Validator = &CustomValidator{validator: validator.New()}
	r.Renderer = driver.NewRenderer("*.html", true)

	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/index", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/html/:name", func(ctx echo.Context) error {
		//data := "<html>" +
		//	"<head></head>" +
		//	"<body>" +
		//	"<h1>Hai</h1>" +
		//	"<img src='https://lh3.googleusercontent.com/ogw/AAEL6sgkeRbd97SwMEJKmC1wbv_HZYYEVSHwnPBX46Ppcw=s32-c-mo'/></body>" +
		//	"</html>"

		//data := `<html>
		//	<head></head>
		//	<body>
		//		<h1>Hai</h1>
		//		<img src='https://lh3.googleusercontent.com/ogw/AAEL6sgkeRbd97SwMEJKmC1wbv_HZYYEVSHwnPBX46Ppcw=s32-c-mo'/></body>
		//	</html>`

		//return ctx.HTML(http.StatusOK, data)
		name := ctx.Param("name")
		return ctx.Render(http.StatusOK, "test.html", Template{Nama: name})
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

	r.GET("/company", func(ctx echo.Context) error {
		name := ctx.Param("name")
		u := &Company{Name: name}

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
