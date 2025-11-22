# Real-Time Chat Server

A WebSocket-based chat server built with Go to learn concurrency patterns.

## Quick Start

```bash
# Install dependencies
go mod tidy

# Run server
go run cmd/server/main.go

# Open browser → http://localhost:8080
```

## Usage

1. Enter username and room name
2. Open multiple browser tabs
3. Join the same room with different usernames
4. Start chatting in real-time!

## Features

- Multiple chat rooms
- Real-time messaging via WebSocket
- Concurrent client handling (2 goroutines per client)
- Thread-safe with mutexes
- Graceful shutdown

## Project Structure

```
chatserver/
├── cmd/server/          # Entry point
├── internal/
│   ├── chat/           # Core chat logic (Hub, Client, Room)
│   ├── user/           # User management
│   └── server/         # HTTP handlers
└── web/static/         # HTML frontend
```

## Tech Stack

- Go 1.21+
- gorilla/websocket
- Channels & Goroutines
- sync.Mutex / sync.RWMutex
