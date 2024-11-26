package handler

import (
	"fmt"
	"net/http"
	service "server/internal/rental/service"

	"server/internal/rental/types"
	userService "server/internal/user/service"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RentalHandler struct {
	service     service.RentalService
	userService userService.UserService
}

func NewRentalHandler(service service.RentalService, userService userService.UserService) *RentalHandler {
	return &RentalHandler{service: service, userService: userService}
}

// AddRental handles the POST request to add a new rental.
func (h *RentalHandler) AddRental(c echo.Context) error {
	var rental types.Rental

	id, err := primitive.ObjectIDFromHex(c.FormValue("createdBy"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	userID, err := primitive.ObjectIDFromHex(id.Hex())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	user, err := h.userService.GetUserByID(c.Request().Context(), userID.Hex())
	if err != nil || user == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "user does not exist"})
	}
	// Parse form fields
	rental.Name = c.FormValue("name")
	rental.StreetNumber = c.FormValue("streetNumber")
	rental.Street = c.FormValue("street")
	rental.City = c.FormValue("city")
	rental.Country = c.FormValue("country")
	rental.FullAddress = c.FormValue("fullAddress")
	rental.Lat = c.FormValue("lat")
	rental.Lng = c.FormValue("lng")
	rental.CreatedBy = userID
	rental.UpdatedBy = userID

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
		rental.Images = append(rental.Images, "https://cdn.vuetifyjs.com/images/cards/hotel.jpg")
	}

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

// GetRentalsByUserID handles the GET request to retrieve rentals by a specific user ID
func (h *RentalHandler) GetRentalsByUserID(c echo.Context) error {
	userID := c.Param("id")
	fmt.Println(userID)
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	// Validate and convert the user ID
	_, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid User ID format"})
	}

	// Fetch rentals for the user
	rentals, err := h.service.GetRentalsByUserID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve rentals"})
	}

	// Return the rentals (empty list if no rentals found)
	return c.JSON(http.StatusOK, rentals)
}
