package server

import (
	"net/http"
	placesHandler "server/internal/places/handler"

	rentalHandler "server/internal/rental/handler"

	userHandler "server/internal/user/handler"

	authHandler "server/internal/auth/handler"

	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

// Router struct with a field for the places handler
// More handlers will be added
type Router struct {
	PlacesHandler *placesHandler.PlacesHandler
	RentalHandler *rentalHandler.RentalHandler
	UserHandler   *userHandler.UserHandler
	AuthHandler   *authHandler.OAuthHandler
}

func (router *Router) Init(e *echo.Echo) {

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "All Good!")
	})

	// Protected Routes
	// Middleware for JWT protection (to be implemented)
	apiGroup := e.Group("/api")
	// apiGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	//	SigningKey: []byte("your-secret-key"),
	//	Claims:     &config.JWTClaims{},
	// }))

	apiGroup.POST("/auth/google/login", router.AuthHandler.GoogleLogin)
	apiGroup.GET("/auth/google/callback", router.AuthHandler.GoogleCallback)
	apiGroup.POST("/auth/login", router.AuthHandler.AuthenticateWithCookie) // Authenticate and get a token
	apiGroup.GET("/auth/me", router.AuthHandler.GetAuthUser)                // Check auth and return user
	apiGroup.POST("/auth/logout", router.AuthHandler.Logout)                // Logout the user

	apiGroup.GET("/users/:id", router.UserHandler.GetUserByID)       // Get user by ID
	apiGroup.POST("/users/create", router.UserHandler.CreateUser)    // Create a new user
	apiGroup.PUT("/users/update/:id", router.UserHandler.UpdateUser) // Update user by ID
	apiGroup.DELETE("/users/:id", router.UserHandler.DeleteUser)

	// Rental endpoints
	apiGroup.POST("/rental/add", router.RentalHandler.AddRental)
	apiGroup.GET("/rental/list", router.RentalHandler.GetAllRentals)
	apiGroup.GET("/rental/:id", router.RentalHandler.GetRentalByID)
	apiGroup.PUT("/rental/:id", router.RentalHandler.UpdateRental)
	apiGroup.DELETE("/rental/:id", router.RentalHandler.DeleteRental)
	apiGroup.GET("/rental/user/:id", router.RentalHandler.GetRentalsByUserID)

	// Places endpoints
	apiGroup.GET("/placeDetails", router.PlacesHandler.GetPlaceDetails)
	apiGroup.GET("/places", router.PlacesHandler.GetPlaces)
	apiGroup.GET("/address/lookup", router.PlacesHandler.GetAddressFromLatLng)

	// Static files for assets
	e.Static("/assets", "../assets")
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
}
