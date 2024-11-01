package handler

import (
	"net/http"
	"server/internal/places/service"

	"github.com/labstack/echo/v4"
)

type PlacesHandler struct {
	Service *service.PlacesService
}

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

func (handler *PlacesHandler) GetPlaceDetails(c echo.Context) error {
	placeId := c.QueryParam("place_id")
	if placeId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Input parameter is required"})
	}

	// Call the service to fetch place details
	placeDetails, err := handler.Service.GetPlaceDetails(placeId)
	if err != nil {
		// Log the error and return an internal server error
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return the placeDetails response as JSON
	return c.JSON(http.StatusOK, placeDetails)
}

func (handler *PlacesHandler) GetAddressFromLatLng(c echo.Context) error {
	latitude := c.QueryParam("latitude")
	longitude := c.QueryParam("longitude")
	if latitude == "" || longitude == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "both latitude and longitude are required"})
	}

	addressDetails, err := handler.Service.GetAddressFromLatLng(latitude, longitude)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, addressDetails)
}
