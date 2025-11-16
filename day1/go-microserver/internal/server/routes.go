package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/hello", handleHello)
	mux.HandleFunc("/time", handleTime)
	mux.HandleFunc("/info", handleInfo)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	// Leggi i query param
	query := r.URL.Query()
	nome := query.Get("name")
	//ruolo := query.Get("ruolo")

	// Risposta dinamica
	if nome != "" {
		msg := fmt.Sprintf("Salve %s ", nome)
		w.Write([]byte(msg))
	} else {
		w.Write([]byte("Salve straniero"))
	}
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"server_time": time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"app":         "microserver",
		"version":     "1.0.0",
		"server_time": time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Errore nella codifica JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
