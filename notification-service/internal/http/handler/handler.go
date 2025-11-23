package handler

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Log *logrus.Logger
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	clients    = make(map[*websocket.Conn]bool)
	clientsMux sync.Mutex
)

func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.Log.Error("Failed to upgrade WebSocket connection: ", err)
		return
	}
	defer func() {
		conn.Close()
		h.Log.Info("WebSocket connection closed")
	}()

	clientsMux.Lock()
	clients[conn] = true
	clientsMux.Unlock()
	h.Log.Info("New WebSocket client connected")

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			h.Log.Error("Failed to read message from WebSocket: ", err)
			clientsMux.Lock()
			delete(clients, conn)
			clientsMux.Unlock()
			h.Log.Info("WebSocket client disconnected")
			break
		}
	}
}

func (h *Handler) BroadcastToClients(message string) {

	clientsMux.Lock()
	defer clientsMux.Unlock()

	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			h.Log.Error("Failed to send message to WebSocket client: ", err)
			client.Close()
			delete(clients, client)
		} else {
			h.Log.Info("Message sent to WebSocket client")
		}
	}
}
