package server

import (
	"net/http"
	placesHandler "server/internal/places/handler"

	rentalHandler "server/internal/rental/handler"

	"github.com/labstack/echo/v4"
)

// Router struct with a field for the places handler
// More handlers will be added
type Router struct {
	PlacesHandler *placesHandler.PlacesHandler
	RentalHandler *rentalHandler.RentalHandler
}

// Init initializes the routes by assigning the handlers
func (router *Router) Init(e *echo.Echo) {

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "All Good!")
	})

	e.GET("/api/placeDetails", router.PlacesHandler.GetPlaceDetails)
	e.GET("/api/places", router.PlacesHandler.GetPlaces)
	e.GET("/api/address/lookup", router.PlacesHandler.GetAddressFromLatLng)
	e.POST("/api/rental/add", router.RentalHandler.AddRental)

}
