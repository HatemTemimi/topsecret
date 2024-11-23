package handler

import (
	"net/http"
	"server/config"
	"server/internal/user/service"
	"server/internal/user/types"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{service: userService}
}

// CreateUser handles the POST request to create a new user.
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user types.User

	// Bind and validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	// Delegate to the service layer
	id, err := h.service.CreateUser(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully", "id": id})
}

// Authenticate handles the POST request for user authentication and issues a JWT token.
func (h *UserHandler) Authenticate(c echo.Context) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind and validate the request body
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	// Delegate to the service layer
	user, err := h.service.Authenticate(c.Request().Context(), credentials.Email, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	// Create JWT token
	claims := &config.JWTClaims{
		UserID: user.ID.Hex(),
		Email:  user.Email,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.TokenExpiry).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(config.SecretKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	// Return the token
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Authentication successful",
		"token":   t,
		"user": map[string]interface{}{
			"id":        user.ID.Hex(),
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
			"role":      user.Role,
		},
	})
}

// GetUserByID handles the GET request to retrieve a user by their ID.
func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	// Get user info from JWT claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTClaims)

	// Allow admin or the owner to access the resource
	if claims.Role != "admin" && claims.UserID != id {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Access denied"})
	}

	// Delegate to the service layer
	fetchedUser, err := h.service.GetUserByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if fetchedUser == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, fetchedUser)
}

// UpdateUser handles the PUT request to update a user's details.
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	var updateData map[string]interface{}
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	// Get user info from JWT claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTClaims)

	// Allow admin or the owner to update the resource
	if claims.Role != "admin" && claims.UserID != id {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Access denied"})
	}

	// Delegate to the service layer
	if err := h.service.UpdateUser(c.Request().Context(), id, updateData); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User updated successfully"})
}

// DeleteUser handles the DELETE request to remove a user.
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	// Get user info from JWT claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTClaims)

	// Allow admin or the owner to delete the resource
	if claims.Role != "admin" && claims.UserID != id {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Access denied"})
	}

	// Delegate to the service layer
	if err := h.service.DeleteUser(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
