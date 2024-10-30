package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type PlacesService struct {
	ApiKey string
	Client *http.Client
}

func NewPlacesService(apiKey string) *PlacesService {
	return &PlacesService{
		ApiKey: apiKey,
		Client: &http.Client{
			Timeout: 10 * time.Second, // Set a reasonable timeout for the HTTP requests
		},
	}
}

// getPlaces fetches place autocomplete results from Google Maps Places API
func (service *PlacesService) GetPlaces(input string) (map[string]interface{}, error) {
	if input == "" {
		return nil, errors.New("input cannot be empty")
	}

	// Define the Google Maps Places API URL
	placeQuery := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/place/autocomplete/json?key=%s&input=%s&components=country:tn",
		service.ApiKey, input,
	)

	// Make the GET request to the Google API
	resp, err := service.Client.Get(placeQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get places: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 response codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("places api returned non-200 status: %d", resp.StatusCode)
	}

	// Decode the response into a map to handle it as JSON
	var places map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&places); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return places, nil
}
