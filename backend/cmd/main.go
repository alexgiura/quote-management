package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"quote-management/internal/app"
	"quote-management/internal/config"
)

func main() {
	// Load configurations
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize the cmd
	application := app.NewApp(cfg)

	// Start the server
	err = application.Run(context.Background())
	if err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}

	// Graceful shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	<-stopChan

	log.Println("Shutting down server...")
	if err := application.StopServer(); err != nil {
		log.Fatalf("Error while stopping server: %v", err)
	}

	log.Println("Server stopped gracefully")
}
