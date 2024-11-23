package server

import (
	"net/http"
	placesHandler "server/internal/places/handler"

	rentalHandler "server/internal/rental/handler"

	userHandler "server/internal/user/handler"

	"github.com/labstack/echo/v4"
)

// Router struct with a field for the places handler
// More handlers will be added
type Router struct {
	PlacesHandler *placesHandler.PlacesHandler
	RentalHandler *rentalHandler.RentalHandler
	UserHandler   *userHandler.UserHandler
}

// Init initializes the routes by assigning the handlers
func (router *Router) Init(e *echo.Echo) {

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "All Good!")
	})

	// Public Routes
	e.POST("/users/create", router.UserHandler.CreateUser)         // Create a new user
	e.POST("/users/authenticate", router.UserHandler.Authenticate) // Authenticate and get a token

	// Protected Routes
	/*
	   apiGroup := e.Group("/api")
	   apiGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	   	SigningKey: []byte("your-secret-key"),
	   	Claims:     &config.JWTClaims{},
	   }))
	*/
	e.GET("/users/:id", router.UserHandler.GetUserByID)       // Get user by ID
	e.PUT("/users/update/:id", router.UserHandler.UpdateUser) // Update user by ID
	e.DELETE("/users/:id", router.UserHandler.DeleteUser)

	e.GET("/api/placeDetails", router.PlacesHandler.GetPlaceDetails)
	e.GET("/api/places", router.PlacesHandler.GetPlaces)
	e.GET("/api/address/lookup", router.PlacesHandler.GetAddressFromLatLng)
	e.POST("/api/rental/add", router.RentalHandler.AddRental)
	e.GET("/api/rental/list", router.RentalHandler.GetAllRentals)
	e.GET("/api/rental/:id", router.RentalHandler.GetRentalByID)
	e.PUT("/api/rental/:id", router.RentalHandler.UpdateRental)
	e.DELETE("/api/rental/:id", router.RentalHandler.DeleteRental)

}
