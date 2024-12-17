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

	types "server/internal/rental/types"
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

// AddRental handles adding a new rental
func (h *RentalHandler) AddRental(c echo.Context) error {
	var rental types.Rental

	// Validate user ID
	userID, err := primitive.ObjectIDFromHex(c.FormValue("createdBy"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Verify user exists
	user, err := h.userService.GetUserByID(c.Request().Context(), userID.Hex())
	if err != nil || user == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User does not exist"})
	}

	// Populate rental details
	rental.Name = c.FormValue("name")
	rental.Description = c.FormValue("description")
	rental.Price, _ = strconv.ParseInt(c.FormValue("price"), 10, 64)
	rental.Bedrooms, _ = strconv.ParseInt(c.FormValue("bedrooms"), 10, 64)
	rental.Bathrooms, _ = strconv.ParseInt(c.FormValue("bathrooms"), 10, 64)
	rental.AreaSize, _ = strconv.ParseInt(c.FormValue("areaSize"), 10, 64)
	rental.Available = c.FormValue("available") == "true"

	// Address and Geometry
	rental.Address = types.Address{
		StreetNumber: c.FormValue("address.streetNumber"),
		Street:       c.FormValue("address.street"),
		City:         c.FormValue("address.city"),
		Country:      c.FormValue("address.country"),
	}

	rental.Geometry = types.Geometry{
		Lat: c.FormValue("geometry.lat"),
		Lng: c.FormValue("geometry.lng"),
	}

	// Amenities
	rental.Amenities = types.Amenities{
		AirConditioning: c.FormValue("amenities.airConditioning") == "true",
		Heating:         c.FormValue("amenities.heating") == "true",
		Refrigerator:    c.FormValue("amenities.refrigerator") == "true",
		Parking:         c.FormValue("amenities.parking") == "true",
	}

	rental.Rules = types.Rules{
		PetsAllowed:    c.FormValue("rules.petsAllowed") == "true",
		SmokingAllowed: c.FormValue("rules.partiesAllowed") == "true",
		PartiesAllowed: c.FormValue("rules.smokingAllowed") == "true",
	}

	// Parse Tags
	tags := c.FormValue("tags")
	if tags != "" {
		rental.Tags = strings.Split(tags, ",")
	}

	// Default values
	rental.Status = types.Pending
	rental.Currency = "TND"
	rental.Standing = types.Standing("standing")
	rental.CreatedBy = userID
	rental.UpdatedBy = userID
	rental.CreatedAt = time.Now()
	rental.UpdatedAt = time.Now()

	// Process Images
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse form data"})
	}

	// Create image folder
	rentalFolder := filepath.Join(utils.GetBasePath(), rental.ID.Hex(), "images")
	if err := os.MkdirAll(rentalFolder, os.ModePerm); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create directory for rental images"})
	}

	files := form.File["images"]
	var wg sync.WaitGroup
	var mu sync.Mutex
	if err := utils.ProcessImages(files, rentalFolder, &rental, &wg, &mu); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Validate at least one image is uploaded
	if len(rental.Images) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "At least one image is required"})
	}

	// Call the service to add the rental
	if err := h.service.AddRental(c.Request().Context(), rental); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
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

	// Validate Rental ID format
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Rental ID format"})
	}

	// Retrieve the existing rental
	existingRental, err := h.service.GetRentalByID(c.Request().Context(), objectID.Hex())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch rental"})
	}
	if existingRental == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Rental not found"})
	}

	// Update rental fields from form input
	existingRental.Name = c.FormValue("name")
	existingRental.Description = c.FormValue("description")
	existingRental.Price, _ = strconv.ParseInt(c.FormValue("price"), 10, 64)
	existingRental.Bedrooms, _ = strconv.ParseInt(c.FormValue("bedrooms"), 10, 64)
	existingRental.Bathrooms, _ = strconv.ParseInt(c.FormValue("bathrooms"), 10, 64)
	existingRental.AreaSize, _ = strconv.ParseInt(c.FormValue("areaSize"), 10, 64)

	// Update address and geometry
	existingRental.Address = types.Address{
		StreetNumber: c.FormValue("streetNumber"),
		Street:       c.FormValue("street"),
		City:         c.FormValue("city"),
		Country:      c.FormValue("country"),
	}
	existingRental.Geometry = types.Geometry{
		Lat: c.FormValue("lat"),
		Lng: c.FormValue("lng"),
	}

	// Update amenities
	existingRental.Amenities = types.Amenities{
		AirConditioning: c.FormValue("airConditioning") == "true",
		Heating:         c.FormValue("heating") == "true",
		Refrigerator:    c.FormValue("refrigerator") == "true",
		Parking:         c.FormValue("parking") == "true",
	}

	// Update availability and tags
	existingRental.Available = c.FormValue("available") == "true"
	tags := c.FormValue("tags")
	if tags != "" {
		existingRental.Tags = strings.Split(tags, ",")
	}

	// Update standing and status if provided
	if standing := c.FormValue("standing"); standing != "" {
		existingRental.Standing = types.Standing(standing)
	}
	if status := c.FormValue("status"); status != "" {
		existingRental.Status = types.Status(status)
	}

	// Handle new image uploads
	form, err := c.MultipartForm()
	if err == nil && form != nil {
		files := form.File["images"]
		if len(files) > 0 {
			// Create a folder for the rental if it doesn't exist
			rentalFolder := filepath.Join(utils.GetBasePath(), objectID.Hex(), "images")
			if err := os.MkdirAll(rentalFolder, os.ModePerm); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create directory for rental images"})
			}

			// Process new images
			var wg sync.WaitGroup
			var mu sync.Mutex
			if err := utils.ProcessImages(files, rentalFolder, existingRental, &wg, &mu); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}
	}

	// Update the `UpdatedAt` timestamp
	existingRental.UpdatedAt = time.Now()

	// Call the service to update the rental
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Convert image file paths to public URLs using the helper
	for i := range rentals {
		rentals[i].Images = utils.MapImagePathsToURLs(c, rentals[i].Images)
	}

	return c.JSON(http.StatusOK, rentals)
}
