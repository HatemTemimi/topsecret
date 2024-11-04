// handler/rental_handler.go
package handler

import (
	"net/http"
	"server/internal/rental/service"
	"server/internal/rental/types"

	"github.com/labstack/echo/v4"
)

type RentalHandler struct {
	service service.RentalService
}

func NewRentalHandler(service service.RentalService) *RentalHandler {
	return &RentalHandler{service: service}
}

// AddRental handles the POST request to add a new rental.
func (h *RentalHandler) AddRental(c echo.Context) error {
	var rental types.Rental

	// Bind the request body to the rental struct
	if err := c.Bind(&rental); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Call the service to add the rental
	if err := h.service.AddRental(c.Request().Context(), rental); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add rental"})
	}

	// Return a success response
	return c.JSON(http.StatusCreated, map[string]string{"message": "Rental added successfully"})
}
