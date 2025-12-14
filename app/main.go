package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"
)

func getClientIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		return xff
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"ip":        getClientIP(r),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
