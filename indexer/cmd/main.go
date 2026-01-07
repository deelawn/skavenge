package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/deelawn/skavenge/indexer"
)

func main() {
	// Parse command line flags
	configPath := flag.String("config", "config.json", "Path to configuration file")
	dbPath := flag.String("db", "indexer.db", "Path to SQLite database file")
	flag.Parse()

	// Load configuration
	config, err := indexer.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create logger
	logger := log.New(os.Stdout, "[INDEXER] ", log.LstdFlags)

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

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		logger.Println("Received shutdown signal")
		cancel()
	}()

	// Start indexer
	logger.Println("Starting blockchain indexer...")
	if err := idx.Start(ctx); err != nil && err != context.Canceled {
		log.Fatalf("Indexer error: %v", err)
	}

	logger.Println("Indexer shutdown complete")
}
