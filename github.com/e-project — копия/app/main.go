package main

import (
	"e-project/config"
	"e-project/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig() // parsing and writing data from config file
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
