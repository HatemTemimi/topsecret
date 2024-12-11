package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"server/internal/rental/service"

	"server/internal/rental/types"
	"server/internal/rental/utils"
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

func (h *RentalHandler) AddRental(c echo.Context) error {
	var rental types.Rental

	// Get user ID from form and validate it
	userID, err := primitive.ObjectIDFromHex(c.FormValue("createdBy"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Verify that the user exists
	user, err := h.userService.GetUserByID(c.Request().Context(), userID.Hex())
	if err != nil || user == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User does not exist"})
	}

	// Parse rental details from form
	rental.Name = c.FormValue("name")
	rental.StreetNumber = c.FormValue("streetNumber")
	rental.Street = c.FormValue("street")
	rental.City = c.FormValue("city")
	rental.Country = c.FormValue("country")
	rental.FullAddress = c.FormValue("fullAddress")
	rental.Description = c.FormValue("description")

	// Parse numeric fields with fallback to 0
	rental.Price, _ = strconv.ParseInt(c.FormValue("price"), 10, 64)
	rental.Bedrooms, _ = strconv.ParseInt(c.FormValue("bedrooms"), 10, 64)
	rental.Bathrooms, _ = strconv.ParseInt(c.FormValue("bathrooms"), 10, 64)
	rental.AreaSize, _ = strconv.ParseInt(c.FormValue("areaSize"), 10, 64)

	// Parse boolean fields
	rental.Available = c.FormValue("available") == "true" || c.FormValue("available") == "1"
	rental.Agree = c.FormValue("agree") == "true" || c.FormValue("agree") == "1"
	rental.Status = c.FormValue("status") == "true" || c.FormValue("status") == "1"

	// Parse location details
	rental.Lat = c.FormValue("lat")
	rental.Lng = c.FormValue("lng")

	// Set createdBy and updatedBy fields
	rental.CreatedBy = userID
	rental.UpdatedBy = userID

	// Add tags if provided
	tags := c.FormValue("tags")
	if tags != "" {
		rental.Tags = strings.Split(tags, ",") // Split tags by comma
	}

	// Generate a unique rental ID
	rental.ID = primitive.NewObjectID()

	// Create directories for rental images
	rentalFolder := filepath.Join(utils.GetBasePath(), rental.ID.Hex(), "images")
	err = os.MkdirAll(rentalFolder, os.ModePerm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create directory for rental images"})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse form data"})
	}

	// WaitGroup to track goroutines
	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex to safely append to the slice
	files := form.File["images"]

	// Process images concurrently
	if err := utils.ProcessImages(files, rentalFolder, &rental, &wg, &mu); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Validate that at least one image was uploaded
	if len(rental.Images) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "At least one image is required for the rental"})
	}

	// Set timestamps
	rental.CreatedAt = time.Now()
	rental.UpdatedAt = time.Now()

	// Call the service to add the rental
	if err := h.service.AddRental(c.Request().Context(), rental); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add rental"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Rental added successfully"})
}

// GetAllRentals handles the GET request to retrieve all rentals
func (h *RentalHandler) GetAllRentals(c echo.Context) error {
	// Fetch all rentals
	rentals, err := h.service.GetAllRentals(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve rentals"})
	}

	// Convert image file paths to public URLs for each rental
	for i := range rentals {
		rentals[i].Images = utils.MapImagePathsToURLs(c, rentals[i].Images)
	}

	// Return the modified rentals
	return c.JSON(http.StatusOK, rentals)
}

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
		return c.JSON(http.StatusOK, map[string]string{"message": "Rental not found", "status": "empty"})
	}

	// Convert image file paths to public URLs using the helper
	rental.Images = utils.MapImagePathsToURLs(c, rental.Images)

	return c.JSON(http.StatusOK, rental)
}

func (h *RentalHandler) UpdateRental(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Rental ID is required"})
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Rental ID format"})
	}

	// Retrieve the existing rental to update its fields
	existingRental, err := h.service.GetRentalByID(c.Request().Context(), objectID.Hex())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch rental"})
	}
	if existingRental == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Rental not found"})
	}

	// Parse and map form fields to update the rental
	existingRental.Name = c.FormValue("name")
	existingRental.StreetNumber = c.FormValue("streetNumber")
	existingRental.Street = c.FormValue("street")
	existingRental.City = c.FormValue("city")
	existingRental.Country = c.FormValue("country")
	existingRental.FullAddress = c.FormValue("fullAddress")
	existingRental.Lat = c.FormValue("lat")
	existingRental.Lng = c.FormValue("lng")

	// Parse numeric fields
	if price := c.FormValue("price"); price != "" {
		existingRental.Price, _ = strconv.ParseInt(price, 10, 64)
	}
	if bedrooms := c.FormValue("bedrooms"); bedrooms != "" {
		existingRental.Bedrooms, _ = strconv.ParseInt(bedrooms, 10, 64)
	}
	if bathrooms := c.FormValue("bathrooms"); bathrooms != "" {
		existingRental.Bathrooms, _ = strconv.ParseInt(bathrooms, 10, 64)
	}
	if areaSize := c.FormValue("areaSize"); areaSize != "" {
		existingRental.AreaSize, _ = strconv.ParseInt(areaSize, 10, 64)
	}

	// Parse boolean fields
	available := c.FormValue("available")
	existingRental.Available = available == "true" || available == "1"

	agree := c.FormValue("agree")
	existingRental.Agree = agree == "true" || agree == "1"

	// Update description
	existingRental.Description = c.FormValue("description")

	// Handle new image uploads
	form, err := c.MultipartForm()
	if err == nil && form != nil {
		files := form.File["images"]
		if len(files) > 0 {
			// Create a folder for the rental if it doesn't exist
			rentalFolder := filepath.Join(utils.GetBasePath(), objectID.Hex(), "images")
			err := os.MkdirAll(rentalFolder, os.ModePerm)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create directory for rental images"})
			}

			// Process new images
			var wg sync.WaitGroup
			var mu sync.Mutex
			//newImages := []string{}
			if err := utils.ProcessImages(files, rentalFolder, existingRental, &wg, &mu); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

			// Append new image paths to existing ones
			//existingRental.Images = append(existingRental.Images, newImages...)
		}
	}

	// Update the `UpdatedAt` timestamp
	existingRental.UpdatedAt = time.Now()

	// Call the service to update the rental using objectID
	if err := h.service.UpdateRental(c.Request().Context(), objectID.Hex(), *existingRental); err != nil {
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

	// Convert image file paths to public URLs using the helper

	for i := range rentals {
		rentals[i].Images = utils.MapImagePathsToURLs(c, rentals[i].Images)
	}

	return c.JSON(http.StatusOK, rentals)
}
