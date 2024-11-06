package main

import (
	"log"
	"server/config"
	"server/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo framework
	e := echo.New()

	// Load configuration from the environment or .env file
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize and set up the server
	client := &server.Server{}
	client.SetupAndLaunch(e, cfg)
}
