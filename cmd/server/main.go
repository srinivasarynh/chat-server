package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/srinivasarynh/chatserver/internal/chat"
	"github.com/srinivasarynh/chatserver/internal/server"
	"github.com/srinivasarynh/chatserver/internal/user"
)

func main() {
	hub := chat.NewHub()
	registry := user.NewRegistry()

	go hub.Run()
	mux := http.NewServeMux()
	wsHandler := server.NewHandler(hub, registry)
	mux.Handle("/ws", server.Logger(wsHandler))

	mux.Handle("/", http.FileServer(http.Dir("./web/static")))

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("server starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown: %v", err)
	}
	log.Println("server exited")
}
