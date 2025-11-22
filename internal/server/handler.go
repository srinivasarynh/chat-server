package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/srinivasarynh/chatserver/internal/chat"
	"github.com/srinivasarynh/chatserver/internal/user"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Handler struct {
	hub      *chat.Hub
	registry *user.Registry
}

func NewHandler(hub *chat.Hub, registry *user.Registry) *Handler {
	return &Handler{
		hub:      hub,
		registry: registry,
	}
}

func (h *Handler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	room := r.URL.Query().Get("room")

	if username == "" || room == "" {
		http.Error(w, "username and room required", http.StatusBadRequest)
		return
	}

	user, err := h.registry.Register(username)
	if err != nil {
		user, _ = h.registry.Get(username)
	}
	user.MarkOnline()
	user.JoinRoom()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := chat.NewClient(h.hub, conn, username, room)
	h.hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
