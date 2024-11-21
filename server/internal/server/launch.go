package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/config"
	"server/internal/places/handler"
	"server/internal/places/service"

	rentalHandler "server/internal/rental/handler"
	rentalRepository "server/internal/rental/repository"
	rentalService "server/internal/rental/service"

	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	router *Router
	Db     *DB
}

// SetupRouter initializes routing handlers using services and repositories
func (s *Server) SetupRouter(e *echo.Echo, cfg *config.Config) {
	// Initialize the rental repository, service, and handler
	rentalRepo := rentalRepository.NewRentalRepository(s.Db.database)
	rentalService := rentalService.NewRentalService(rentalRepo)
	rentalHandler := rentalHandler.NewRentalHandler(rentalService)

	// Create the PlacesService using the API key from config
	placesService := service.NewPlacesService(cfg.GooglePlacesAPIKey)

	// Create the PlacesHandler with the injected service
	placesHandler := &handler.PlacesHandler{
		Service: placesService,
	}

	// Initialize the Router with both handlers
	s.router = &Router{
		PlacesHandler: placesHandler,
		RentalHandler: rentalHandler,
	}

	// Initialize routes
	s.router.Init(e)
}

// SetupAndLaunch launches the server
func (s *Server) SetupAndLaunch(e *echo.Echo, cfg *config.Config) {
	// Use the config values directly instead of reading environment variables
	port := cfg.MainPort
	apiKey := cfg.GooglePlacesAPIKey

	// Validate that necessary config values are set
	if apiKey == "" {
		log.Fatal("API key for Google Places is not set")
	}
	if port == 0 {
		log.Fatal("PORT is not set")
	}

	// Initialize MongoDB connection
	mongoDB, err := NewDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer mongoDB.Close()

	s.Db = mongoDB
	s.Db.InitMockRentals()

	// Initialize router with configuration
	s.SetupRouter(e, cfg)

	// Start the server in a goroutine
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for an interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// Context for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Error during shutdown:", err)
	}
}
