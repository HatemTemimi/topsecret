package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
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

// haversine function to calculate distance between two latitude/longitude pairs
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371e3 // Earth's radius in meters
	phi1 := lat1 * math.Pi / 180
	phi2 := lat2 * math.Pi / 180
	deltaPhi := (lat2 - lat1) * math.Pi / 180
	deltaLambda := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
		math.Cos(phi1)*math.Cos(phi2)*math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func (service *PlacesService) GetAddressFromLatLng(latitude, longitude string) (map[string]interface{}, error) {
	if latitude == "" || longitude == "" {
		return nil, errors.New("both latitude and longitude are required")
	}

	query := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/geocode/json?latlng=%s,%s&key=%s",
		latitude, longitude, service.ApiKey,
	)

	resp, err := service.Client.Get(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get address from geocoding")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("address geocode API returned non-200 status: %d", resp.StatusCode)
	}

	var response struct {
		Results []struct {
			FormattedAddress string `json:"formatted_address"`
			Geometry         struct {
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
			} `json:"geometry"`
			AddressComponents []struct {
				LongName string   `json:"long_name"`
				Types    []string `json:"types"`
			} `json:"address_components"`
		} `json:"results"`
		Status string `json:"status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if response.Status != "OK" || len(response.Results) == 0 {
		return nil, errors.New("no results found")
	}

	// Convert input latitude and longitude to float64 for distance calculation
	lat, _ := strconv.ParseFloat(latitude, 64)
	lon, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid latitude or longitude format: %v", err)
	}

	var closestRouteResult, closestNonRouteResult map[string]interface{}
	minRouteDistance := math.MaxFloat64
	minNonRouteDistance := math.MaxFloat64

	for _, result := range response.Results {
		// Check if the result is within Tunisia
		inTunisia := false
		for _, component := range result.AddressComponents {
			if component.Types[0] == "country" && component.LongName == "Tunisia" {
				inTunisia = true
				break
			}
		}

		if !inTunisia {
			continue
		}

		// Check if the result contains "route" in its address components
		hasRouteType := false
		for _, component := range result.AddressComponents {
			for _, t := range component.Types {
				if t == "route" {
					hasRouteType = true
					break
				}
			}
			if hasRouteType {
				break
			}
		}

		// Calculate the distance to the provided coordinates
		resultLat := result.Geometry.Location.Lat
		resultLng := result.Geometry.Location.Lng
		distance := haversine(lat, lon, resultLat, resultLng)

		// Prioritize results with "route" type, but keep track of both types
		if hasRouteType {
			if distance < minRouteDistance {
				minRouteDistance = distance
				closestRouteResult = map[string]interface{}{
					"formatted_address": result.FormattedAddress,
					"location": map[string]float64{
						"lat": resultLat,
						"lng": resultLng,
					},
					"address_components": result.AddressComponents,
				}
			}
		} else {
			if distance < minNonRouteDistance {
				minNonRouteDistance = distance
				closestNonRouteResult = map[string]interface{}{
					"formatted_address": result.FormattedAddress,
					"location": map[string]float64{
						"lat": resultLat,
						"lng": resultLng,
					},
					"address_components": result.AddressComponents,
				}
			}
		}
	}

	// Return the closest result with "route" type if available, otherwise fallback to the closest non-route result
	if closestRouteResult != nil {
		return closestRouteResult, nil
	} else if closestNonRouteResult != nil {
		return closestNonRouteResult, nil
	}

	return nil, errors.New("no results found in Tunisia")
}
