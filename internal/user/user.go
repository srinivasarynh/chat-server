package user

import (
	"errors"
	"time"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Username  string
	JoinedAt  time.Time
	LastSeen  time.Time
	isOnline  bool
	roomCount int
}

func NewUser(username string) *User {
	return &User{
		Username: username,
		JoinedAt: time.Now(),
		LastSeen: time.Now(),
		isOnline: true,
	}
}

func (u *User) MarkOnline() {
	u.isOnline = true
	u.LastSeen = time.Now()
}

func (u *User) MarkOffline() {
	u.isOnline = false
	u.LastSeen = time.Now()
}

func (u *User) IsOnline() bool {
	return u.isOnline
}

func (u *User) JoinRoom() {
	u.roomCount++
}

func (u *User) LeaveRoom() {
	if u.roomCount > 0 {
		u.roomCount--
	}
}
