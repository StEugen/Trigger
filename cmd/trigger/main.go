package main

import (
	"fmt"
	"os"

	"github.com/steugen/trigger/internal/config"
	"github.com/steugen/trigger/internal/event"
)

func main() {
	fmt.Println("Starting Trigger...")

	// Here will be load of conf
	cfg, err := config.Load("config.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading configuration: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Configuration loaded:", cfg)
	// Placeholder here for now
	err = event.Start(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting event listener: %v\n", err)
		os.Exit(1)
	}

}
