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
			Timeout: 10 * time.Second,
		},
	}
}

func (service *PlacesService) GetPlaces(input string) (map[string]interface{}, error) {

	if input == "" {
		return nil, errors.New("input cannot be empty")
	}

	// Define Query
	placeQuery := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/place/autocomplete/json?key=%s&input=%s&components=country:tn",
		service.ApiKey, input,
	)

	// Make request
	resp, err := service.Client.Get(placeQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get places: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("places api returned non-200 status: %d", resp.StatusCode)
	}

	// Decode the response
	var places map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&places); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return places, nil
}

// receives a placeID, and returns details such as lat, lng
func (service *PlacesService) GetPlaceDetails(placeID string) (map[string]interface{}, error) {

	if placeID == "" {
		return nil, errors.New("no placeID provided")
	}

	query := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/place/details/json?fields=geometry%%2Cname&place_id=%s&key=%s",
		placeID, service.ApiKey,
	)

	resp, err := service.Client.Get(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get place_details: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("place_details api returned non-200 status: %d", resp.StatusCode)
	}

	var placeDetails map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&placeDetails); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return placeDetails, nil
}

// returns address
func (service *PlacesService) GetAddress(placeID string) (map[string]interface{}, error) {
	return nil, nil
}
