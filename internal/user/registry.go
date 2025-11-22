package user

import "sync"

type Registry struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewRegistry() *Registry {
	return &Registry{
		users: make(map[string]*User),
	}
}

func (r *Registry) Register(username string) (*User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[username]; exists {
		return nil, ErrUserExists
	}

	user := NewUser(username)
	r.users[username] = user
	return user, nil
}

func (r *Registry) Get(username string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[username]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *Registry) GetOnlineUsers() []*User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	online := make([]*User, 0)
	for _, user := range r.users {
		if user.IsOnline() {
			online = append(online, user)
		}
	}

	return online
}

func (r *Registry) Remove(username string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.users, username)
}
