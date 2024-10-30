package handler

import (
	"net/http"
	"server/internal/places/service"

	"github.com/labstack/echo/v4"
)

type PlacesHandler struct {
	Service *service.PlacesService
}

// GetPlaces is an HTTP handler that retrieves place suggestions from the Google Places API
func (handler *PlacesHandler) GetPlaces(c echo.Context) error {
	// Get the 'input' query parameter from the request
	input := c.QueryParam("input")
	if input == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Input parameter is required"})
	}

	// Call the service to fetch place suggestions
	places, err := handler.Service.GetPlaces(input)
	if err != nil {
		// Log the error and return an internal server error
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return the places response as JSON
	return c.JSON(http.StatusOK, places)
}
