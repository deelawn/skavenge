package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/deelawn/skavenge/indexer"
)

const defaultDbPath = ":memory:"

func main() {
	// Parse command line flags
	configPath := flag.String("config", "config.json", "Path to configuration file")
	dbPath := flag.String("db", "", "Path to SQLite database file")
	apiPort := flag.Int("api-port", 4040, "Port for API server")
	flag.Parse()

	if *dbPath == "" {
		*dbPath = defaultDbPath
	}

	// Load configuration
	config, err := indexer.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create logger
	logger := log.New(os.Stdout, "[INDEXER] ", log.LstdFlags)
	apiLogger := log.New(os.Stdout, "[API] ", log.LstdFlags)

	// Create storage (using SQLite)
	storage, err := indexer.NewSQLiteStorage(*dbPath)
	if err != nil {
		log.Fatalf("Failed to create storage: %v", err)
	}
	defer storage.Close()

	// Create indexer
	idx, err := indexer.NewIndexer(config, storage, logger)
	if err != nil {
		log.Fatalf("Failed to create indexer: %v", err)
	}
	defer idx.Close()

	// Create API server
	apiServer := indexer.NewAPIServer(storage, apiLogger, *apiPort)

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start API server in a separate goroutine
	go func() {
		if err := apiServer.Start(); err != nil {
			logger.Printf("API server error: %v", err)
		}
	}()

	// Start indexer in a separate goroutine
	indexerDone := make(chan error, 1)
	go func() {
		logger.Println("Starting blockchain indexer...")
		indexerDone <- idx.Start(ctx)
	}()

	// Wait for shutdown signal
	<-sigChan
	logger.Println("Received shutdown signal")
	cancel()

	// Shutdown API server
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := apiServer.Shutdown(shutdownCtx); err != nil {
		logger.Printf("Error shutting down API server: %v", err)
	}

	// Wait for indexer to finish
	if err := <-indexerDone; err != nil && err != context.Canceled {
		log.Fatalf("Indexer error: %v", err)
	}

	logger.Println("Shutdown complete")
}
