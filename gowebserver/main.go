package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("starting service...")

	// Simple HTTP server to receive support requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", logHeaders)
	mux.HandleFunc("/.well-known/ready", handleHealth)
	mux.HandleFunc("/.well-known/live", handleHealth)

	err := http.ListenAndServe(":9898", mux)
	log.Fatalf("error serving: %v", err)
}

func logHeaders(w http.ResponseWriter, r *http.Request) {
	log.Printf("forwarded is %s", r.Header.Get("X-Forwarded-Proto"))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("ready and live")
}
