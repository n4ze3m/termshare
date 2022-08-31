package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func socketReader(ws *websocket.Conn) {
	for {
		message, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(string(message))
	}
}

func (h *Handler) WebsocketHandler(w http.ResponseWriter, r *http.Request) {

}
