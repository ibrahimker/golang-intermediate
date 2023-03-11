package client

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

var connections []*WebSocketConnection

type M map[string]interface{}

const MESSAGE_NEW_USER = "New User"
const MESSAGE_CHAT = "Chat"
const MESSAGE_LEAVE = "Leave"

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	From    string
	Type    string
	Message string
}

type WebSocketConnection struct {
	*websocket.Conn
	Username string
	Age      string
}

func HandleIO(currentConn *WebSocketConnection) {
	connections = append(connections, currentConn)
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	broadcastMessage(currentConn, MESSAGE_NEW_USER, "")

	for {
		payload := SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(currentConn, MESSAGE_LEAVE, "")
				ejectConnection(currentConn)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		broadcastMessage(currentConn, MESSAGE_CHAT, payload.Message)
	}
}

func ejectConnection(currentConn *WebSocketConnection) {
	var newConn []*WebSocketConnection
	for _, conn := range connections {
		if conn != currentConn {
			newConn = append(newConn, conn)
		}
	}
	connections = newConn
}

func broadcastMessage(currentConn *WebSocketConnection, kind, message string) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		eachConn.WriteJSON(SocketResponse{
			From:    fmt.Sprintf(currentConn.Username + " Age: " + currentConn.Age),
			Type:    kind,
			Message: message,
		})
	}
}
