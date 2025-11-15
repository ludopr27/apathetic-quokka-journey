package server

import (
    "log"
    "net/http"
)

type Server struct {
    mux *http.ServeMux
}

func New() *Server {
    mux := http.NewServeMux()
    s := &Server{mux: mux}
    s.routes()
    return s
}

func (s *Server) Start(addr string) error {
    log.Printf("Listening on %s", addr)
    return http.ListenAndServe(addr, s.mux)
}
