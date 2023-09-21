package main

import (
	"log"

	"github.com/traningWithPgxPackage/config"
	"github.com/traningWithPgxPackage/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
