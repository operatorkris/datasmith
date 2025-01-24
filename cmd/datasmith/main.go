package main

import (
	"fmt"
	"log"
	"net/http"

	"datasmith/internal/api"
	"datasmith/internal/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Start HTTP server
	router := api.SetupRouter()
	log.Printf("Starting server on port %s...", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
