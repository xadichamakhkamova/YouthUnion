package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Client struktura — har bir user WebSocketga ulanganda shu obyekt bo'ladi
// Unda: conn (WebSocket), user_id (kim ekanligi)
type Client struct {
	Conn   *websocket.Conn
	UserID string
}

// Hub — barcha WebSocket clientlarni boshqaradi
// clients[userID] => list of connections
type WSHub struct {
	Clients map[string][]*websocket.Conn
	Mux     sync.Mutex
	Log     *logrus.Logger
}

func NewHub(log *logrus.Logger) *WSHub {
	return &WSHub{
		Clients: make(map[string][]*websocket.Conn),
		Log:     log,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// HandleWebSocket: user WebSocketga ulanadi
// Misol: ws://localhost:9000/ws?user_id=123
func (h *WSHub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.Log.Error("WebSocket upgrade failed: ", err)
		return
	}

	// Clientni ro'yxatga qo'shamiz
	h.Mux.Lock()
	h.Clients[userID] = append(h.Clients[userID], conn)
	h.Mux.Unlock()
	h.Log.Infof("User %s connected", userID)

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			// Ulangan client chiqib ketdi — ro'yxatdan o'chiramiz
			h.removeConnection(userID, conn)
			h.Log.Infof("User %s disconnected", userID)
			break
		}
	}
}

func (h *WSHub) removeConnection(userID string, conn *websocket.Conn) {
	h.Mux.Lock()
	defer h.Mux.Unlock()

	conns := h.Clients[userID]
	for i, c := range conns {
		if c == conn {
			h.Clients[userID] = append(conns[:i], conns[i+1:]...)
			break
		}
	}
}

// SendToUser — 1 ta userga notification yuboradi
func (h *WSHub) SendToUser(userID string, msg interface{}) {
	h.Mux.Lock()
	conns := h.Clients[userID]
	h.Mux.Unlock()

	if len(conns) == 0 {
		h.Log.Infof("User %s online emas", userID)
		return
	}

	data, _ := json.Marshal(msg)

	for _, c := range conns {
		if err := c.WriteMessage(websocket.TextMessage, data); err != nil {
			h.Log.Error("Send error: ", err)
		}
	}
}

func (h *WSHub) BroadcastPublic(msg interface{}) {
	h.Mux.Lock()
	allClients := h.Clients
	h.Mux.Unlock()

	data, _ := json.Marshal(msg)

	for _, conns := range allClients {
		for i := 0; i < len(conns); i++ {
			c := conns[i]
			if err := c.WriteMessage(websocket.TextMessage, data); err != nil {
				h.Log.Error("Broadcast error: ", err)
				c.Close()
				// Slice’dan o‘chirish
				conns = append(conns[:i], conns[i+1:]...)
				i-- // indexni to‘g‘ri ushlab qolish
			}
		}
	}
}
