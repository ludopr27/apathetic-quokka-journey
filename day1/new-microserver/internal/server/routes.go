package server

import (
    "encoding/json"
    "log"
    "net/http"
)

type HealthResponse struct {
    Status string `json:"status"`
}

func (s *Server) routes() {
    s.mux.HandleFunc("/health", s.handleHealth)
    s.mux.HandleFunc("/hello", s.handleHello)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
    log.Println("GET /health")
    resp := HealthResponse{Status: "ok"}
    json.NewEncoder(w).Encode(resp)
}

func (s *Server) handleHello(w http.ResponseWriter, r *http.Request) {
    log.Println("GET /hello")
    w.Write([]byte("Hello from Go Microserver!"))
}
