package server

import (
	"net/http"
	"server/internal/places/handler"

	"github.com/labstack/echo/v4"
)

// Router struct with a field for the places handler
type Router struct {
	PlacesHandler *handler.PlacesHandler
}

// Init initializes the routes and assigns the handlers
func (router *Router) Init(e *echo.Echo) {

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "All Good!")
	})

	e.GET("/api/places", router.PlacesHandler.GetPlaces)
	e.GET("/api/placesDetails", router.PlacesHandler.GetPlaceDetails)

}
