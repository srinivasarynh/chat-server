package chat

import (
	"encoding/json"
	"time"
)

type MessageType string

const (
	TextMessage   MessageType = "text"
	JoinMessage   MessageType = "join"
	LeaveMessage  MessageType = "leave"
	SystemMessage MessageType = "system"
)

type Message struct {
	Type      MessageType `json:"type"`
	Username  string      `json:"username"`
	Content   string      `json:"content"`
	Room      string      `json:"room"`
	Timestamp time.Time   `json:"timestamp"`
}

func NewMessage(msgType MessageType, username, content, room string) *Message {
	return &Message{
		Type:      msgType,
		Username:  username,
		Content:   content,
		Room:      room,
		Timestamp: time.Now(),
	}
}

func (m Message) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

func (m Message) IsCommand() bool {
	return len(m.Content) > 0 && m.Content[0] == '/'
}
