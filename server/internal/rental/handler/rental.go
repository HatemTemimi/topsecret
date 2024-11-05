// handler/rental_handler.go
package handler

import (
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
	rental.Lat = c.FormValue("lng")
	rental.Lng = c.FormValue("lat")

	// Parse boolean for "agree" field
	agree := c.FormValue("agree")
	rental.Agree = agree == "true" || agree == "1"

	// Parse images if any are uploaded
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid form data"})
	}

	// Collect image files (assuming filenames or other identifiers are sufficient)
	rental.Images = []string{}
	files := form.File["images"]
	for _, file := range files {
		// Here, we just add the filename for demonstration; in practice, you'd save the file
		rental.Images = append(rental.Images, file.Filename)
	}

	// Log rental data for debugging
	log.Printf("Parsed rental data: %+v\n", rental)

	// Call the service to add the rental
	if err := h.service.AddRental(c.Request().Context(), rental); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add rental"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Rental added successfully"})
}
