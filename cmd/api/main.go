package main

import (
	"fmt"
	"log"
	"os"

	"github.com/OderoCeasar/goapi/internal/config"
	"github.com/OderoCeasar/goapi/internal/server"
)

func main() {
	// load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	// create and start server
	srv := server.New(cfg)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := srv.Start(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}