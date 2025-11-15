package main

import (
    "log"
    "go-microserver/internal/server"
)

func main() {
    s := server.New()

    log.Println("Starting server on :8080 ...")
    if err := s.Start(":8080"); err != nil {
        log.Fatalf("server error: %v", err)
    }
}
