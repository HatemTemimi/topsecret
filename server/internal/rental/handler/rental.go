package handler

import (
	"net/http"
	service "server/internal/rental/service"
	"strconv"
	"strings"
	"time"

	"server/internal/rental/types"
	userService "server/internal/user/service"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"fmt"
	"io"
	"os"
	"path/filepath"
)

type RentalHandler struct {
	service     service.RentalService
	userService userService.UserService
}

func NewRentalHandler(service service.RentalService, userService userService.UserService) *RentalHandler {
	return &RentalHandler{service: service, userService: userService}
}

func getBasePath() string {
	basePath := os.Getenv("ASSETS_BASE_PATH")
	if basePath == "" {
		basePath = "../assets/rentals"
	}
	return basePath
}

// mapImagePathsToURLs converts file paths to public URLs.
func mapImagePathsToURLs(c echo.Context, imagePaths []string) []string {
	baseURL := fmt.Sprintf("http://%s/", c.Request().Host) // Base URL for assets
	var urls []string
	for _, imagePath := range imagePaths {
		url := fmt.Sprintf("%s%s", baseURL, strings.TrimPrefix(imagePath, "../"))
		urls = append(urls, url)
	}
	return urls
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
	rentalFolder := filepath.Join(getBasePath(), rental.ID.Hex(), "images")
	err = os.MkdirAll(rentalFolder, os.ModePerm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create directory for rental images"})
	}

	// Parse and save uploaded images
	form, err := c.MultipartForm()
	if err == nil && form != nil {
		files := form.File["images"]
		for _, file := range files {
			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read uploaded file"})
			}
			defer src.Close()

			// Create destination file
			dstPath := filepath.Join(rentalFolder, file.Filename)
			dst, err := os.Create(dstPath)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save uploaded file"})
			}
			defer dst.Close()

			// Copy the file content
			if _, err = io.Copy(dst, src); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to write file content"})
			}

			// Add relative path to the rental images
			rental.Images = append(rental.Images, strings.TrimPrefix(dstPath, "../internal/"))
		}
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
		rentals[i].Images = mapImagePathsToURLs(c, rentals[i].Images)
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
	rental.Images = mapImagePathsToURLs(c, rental.Images)

	return c.JSON(http.StatusOK, rental)
}

// UpdateRental handles the PUT request to update a rental
func (h *RentalHandler) UpdateRental(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Rental ID is required"})
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Rental ID format"})
	}

	var updatedRental types.Rental

	// Parse and map form fields
	updatedRental.Name = c.FormValue("name")
	updatedRental.StreetNumber = c.FormValue("streetNumber")
	updatedRental.Street = c.FormValue("street")
	updatedRental.City = c.FormValue("city")
	updatedRental.Country = c.FormValue("country")
	updatedRental.FullAddress = c.FormValue("fullAddress")
	updatedRental.Lat = c.FormValue("lat")
	updatedRental.Lng = c.FormValue("lng")

	// Parse numeric fields and handle conversion
	if price := c.FormValue("price"); price != "" {
		updatedRental.Price, _ = strconv.ParseInt(price, 10, 64)
	}
	if bedrooms := c.FormValue("bedrooms"); bedrooms != "" {
		updatedRental.Bedrooms, _ = strconv.ParseInt(bedrooms, 10, 64)
	}
	if bathrooms := c.FormValue("bathrooms"); bathrooms != "" {
		updatedRental.Bathrooms, _ = strconv.ParseInt(bathrooms, 10, 64)
	}
	if areaSize := c.FormValue("areaSize"); areaSize != "" {
		updatedRental.AreaSize, _ = strconv.ParseInt(areaSize, 10, 64)
	}

	// Parse boolean fields
	available := c.FormValue("available")
	updatedRental.Available = available == "true" || available == "1"

	agree := c.FormValue("agree")
	updatedRental.Agree = agree == "true" || agree == "1"

	// Parse description
	updatedRental.Description = c.FormValue("description")

	// Handle file uploads for images
	form, err := c.MultipartForm()
	if err == nil && form != nil {
		updatedRental.Images = []string{}
		files := form.File["images"]
		for _, file := range files {
			// Add the filename; in production, save the file and store the path
			updatedRental.Images = append(updatedRental.Images, file.Filename)
		}
	}

	// Update the `UpdatedAt` field
	updatedRental.UpdatedAt = time.Now()

	// Call the service to update the rental
	if err := h.service.UpdateRental(c.Request().Context(), objectID.Hex(), updatedRental); err != nil {
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
		rentals[i].Images = mapImagePathsToURLs(c, rentals[i].Images)
	}

	return c.JSON(http.StatusOK, rentals)
}
