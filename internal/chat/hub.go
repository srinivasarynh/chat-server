package chat

import (
	"log"
	"sync"
)

type Hub struct {
	rooms      map[string]*Room
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
	mu         sync.RWMutex
	once       sync.Once
}

func NewHub() *Hub {
	return &Hub{
		rooms:      make(map[string]*Room),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}
}

func (h *Hub) GetOrCreateRoom(name string) *Room {
	h.mu.Lock()
	defer h.mu.Unlock()

	room, exists := h.rooms[name]
	if !exists {
		room = NewRoom(name)
		h.rooms[name] = room
	}
	return room
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			room := h.GetOrCreateRoom(client.room)
			room.AddClient(client)

			joinMsg := NewMessage(JoinMessage, client.username, "joined the room", client.room)
			if data, err := joinMsg.ToJSON(); err == nil {
				room.BroadCast(data)
			}
			log.Printf("%s joined room %s", client.username, client.room)

		case client := <-h.unregister:
			h.mu.RLock()
			room, exists := h.rooms[client.room]
			h.mu.RUnlock()

			if exists {
				room.RemoveClient(client)
				close(client.send)

				leaveMsg := NewMessage(LeaveMessage, client.username, "left the room", client.room)
				if data, err := leaveMsg.ToJSON(); err == nil {
					room.BroadCast(data)
				}
				log.Printf("%s left room %s", client.username, client.room)
			}

		case message := <-h.broadcast:
			h.mu.RLock()
			room, exists := h.rooms[message.Room]
			h.mu.RUnlock()

			if exists {
				if data, err := message.ToJSON(); err == nil {
					room.BroadCast(data)
				}
			}
		}
	}
}
