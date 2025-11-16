package main

import (
	"context"
	"log"
	"microserver/internal/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	addr := getAddr()

	// Creiamo il nostro server applicativo
	srv := server.New(addr)

	// Avvio del server in una goroutine
	go func() {
		log.Printf("Starting server on %s ...", addr)
		if err := srv.Start(); err != nil {
			log.Printf("server error: %v", err)
		}
	}()

	// Canale per intercettare i segnali di shutdown (Ctrl+C, kill, ecc.)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Bloccante finché non arriva un segnale
	<-sigChan
	log.Println("Shutdown signal received")

	// Contesto con timeout per lo shutdown "gentile"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	} else {
		log.Println("Server stopped cleanly")
	}
}

// getAddr legge PORT dall'ambiente, oppure usa la default :8080
func getAddr() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	// Bind esplicito su 0.0.0.0 (tutte le interfacce)
	return "0.0.0.0:8080"
}
