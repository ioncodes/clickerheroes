package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type TicketResponse struct {
	Success  string `json:"success"`
	Response string `json:"response"`
}

// Spoof response from savedgames.clickerheroes.com
func httpRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	fmt.Printf("Request URL: %s\n", r.URL.Path)

	if body, err := io.ReadAll(r.Body); err == nil {
		fmt.Printf("Request Body: %s\n", string(body))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TicketResponse{Success: "success", Response: "layle was here"})
}

func main() {
	certificate, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		InsecureSkipVerify: true,
		MaxVersion:         tls.VersionTLS12,
	}
	server := http.Server{
		Addr:      ":443",
		Handler:   http.HandlerFunc(httpRequestHandler),
		TLSConfig: tlsConfig,
	}

	defer server.Close()
	log.Fatal(server.ListenAndServeTLS("", ""))
}
