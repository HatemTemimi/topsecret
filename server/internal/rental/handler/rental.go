package handler

import (
	"fmt"
	"log"
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

	// Parse form fields
	rental.Name = c.FormValue("name")
	rental.StreetNumber = c.FormValue("streetNumber")
	rental.Street = c.FormValue("street")
	rental.City = c.FormValue("city")
	rental.Country = c.FormValue("country")
	rental.FullAddress = c.FormValue("fullAddress")
	rental.Lat = c.FormValue("lat")
	rental.Lng = c.FormValue("lng")

	// Parse boolean for "agree" field
	agree := c.FormValue("agree")
	rental.Agree = agree == "true" || agree == "1"

	// Parse images if any are uploaded
	form, err := c.MultipartForm()
	if err == nil && form != nil {
		rental.Images = []string{}
		files := form.File["images"]
		for _, file := range files {
			// Here, we just add the filename for demonstration; in practice, you'd save the file
			rental.Images = append(rental.Images, file.Filename)
		}
	}

	// Log rental data for debugging
	log.Printf("Parsed rental data: %+v\n", rental)

	// Call the service to add the rental
	if err := h.service.AddRental(c.Request().Context(), rental); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add rental"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Rental added successfully"})
}

// GetAllRentals handles the GET request to retrieve all rentals
func (h *RentalHandler) GetAllRentals(c echo.Context) error {
	rentals, err := h.service.GetAllRentals(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve rentals"})
	}
	return c.JSON(http.StatusOK, rentals)
}

// GetRentalByID handles the GET request to retrieve a single rental by ID
func (h *RentalHandler) GetRentalByID(c echo.Context) error {
	fmt.Println("correct handler")
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Rental ID is required"})
	}

	rental, err := h.service.GetRentalByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve rental"})
	}
	if rental == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Rental not found"})
	}
	return c.JSON(http.StatusOK, rental)
}

// UpdateRental handles the PUT request to update a rental
func (h *RentalHandler) UpdateRental(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Rental ID is required"})
	}

	var updatedData types.Rental
	if err := c.Bind(&updatedData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	if err := h.service.UpdateRental(c.Request().Context(), id, updatedData); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update rental"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Rental updated successfully"})
}

// DeleteRental handles the DELETE request to remove a rental
func (h *RentalHandler) DeleteRental(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Rental ID is required"})
	}

	if err := h.service.DeleteRental(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete rental"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Rental deleted successfully"})
}
