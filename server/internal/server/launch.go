package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/internal/places/handler"
	"server/internal/places/service"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Server struct {
	router *Router
	echo   *echo.Echo
}

func (s *Server) SetupAndLaunch() {

	s.echo = echo.New()

	// Get the API key from environment variables or configuration
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}

	envFile := "../config/.env." + env

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	apiKey := os.Getenv("GOOGLE_PLACES_API_KEY")

	if apiKey == "" {
		log.Fatal("API key for Google Places is not set")
	}

	port := os.Getenv("MAIN_PORT")

	if port == "" {
		log.Fatal("PORT is not set")
	}

	// Create the PlacesService
	placesService := service.NewPlacesService(apiKey)

	// Create the PlacesHandler with the service injected
	placesHandler := &handler.PlacesHandler{
		Service: placesService,
	}

	// Initialize the Router with the PlacesHandler
	s.router = &Router{
		PlacesHandler: placesHandler,
	}

	go func() {
		if err := s.echo.Start(fmt.Sprintf(":%s", port)); err != nil && err != http.ErrServerClosed {
			s.echo.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.echo.Shutdown(ctx); err != nil {
		s.echo.Logger.Fatal("Error during shutdown:", err)
	}
}