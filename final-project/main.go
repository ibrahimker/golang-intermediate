package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ibrahimker/golang-intermediate/final-project/client"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		content, err := ioutil.ReadFile("template/chat.html")
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "could not open html")
		}

		return ctx.HTML(http.StatusOK, string(content))
	})
	e.Static("/template", "template")

	e.Any("/ws", func(ctx echo.Context) error {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		currentGorillaConn, err := upgrader.Upgrade(ctx.Response().Writer, ctx.Request(), nil)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "Could not open websocket connection")
		}

		username := ctx.Request().URL.Query().Get("username")
		age := ctx.Request().URL.Query().Get("age")

		go client.HandleIO(&client.WebSocketConnection{
			Conn:     currentGorillaConn,
			Username: username,
			Age:      age,
		})
		return nil
	})

	e.Start(":8080")
}
