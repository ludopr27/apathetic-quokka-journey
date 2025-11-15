package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"microserver/internal/server"
)

func main() {
	srv := server.New("0.0.0.0:8080")

	// Start server
	go func() {
		log.Println("Server starting on http://localhost:8080")
		if err := srv.Start(); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	// OS signal handling (CTRL+C, docker stop, ecc.)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}

	log.Println("Server stopped cleanly.")
}
