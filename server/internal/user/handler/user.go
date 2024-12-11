package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"server/config"
	"server/internal/user/service"
	"server/internal/user/types"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service  service.UserService
	validate *validator.Validate
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		service:  userService,
		validate: validator.New(),
	}
}

// CreateUser handles the POST request to create a new user.
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user types.User

	// Bind and validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	// Validate the user struct
	if err := h.validate.Struct(user); err != nil {
		validationErrors := map[string]string{}
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors[e.Field()] = e.Tag()
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Validation failed", "details": validationErrors})
	}

	// Delegate to the service layer
	id, err := h.service.CreateUser(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully", "id": id})
}

// Authenticate handles the POST request for user authentication and issues a JWT token.
// Deprecated
func (h *UserHandler) Authenticate(c echo.Context) error {
	var credentials struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	// Bind and validate the request body
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	// Validate the credentials struct
	if err := h.validate.Struct(credentials); err != nil {
		validationErrors := map[string]string{}
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors[e.Field()] = e.Tag()
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Validation failed", "details": validationErrors})
	}

	// Delegate to the service layer
	user, err := h.service.Authenticate(c.Request().Context(), credentials.Email, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	// Create JWT token
	token, err := config.GenerateToken(user.ID.Hex(), user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	// Return the token
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Authentication successful",
		"token":   token,
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

// Authenticate handles the POST request for user authentication and issues a JWT token.
func (h *UserHandler) AuthenticateWithCookie(c echo.Context) error {
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
	token, err := config.GenerateToken(user.ID.Hex(), user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	// Set token as cookie
	c.SetCookie(&http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,                  // Set to true in production with HTTPS
		SameSite: http.SameSiteNoneMode, // Use "None" for third-party context
		Expires:  time.Now().Add(config.TokenExpiry),
	})

	// Create a JSON representation of the user
	userJSON, err := json.Marshal(map[string]string{
		"id":        user.ID.Hex(),
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"role":      user.Role,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate user info"})
	}

	// URL encode the JSON string
	encodedUserInfo := url.QueryEscape(string(userJSON))

	c.SetCookie(&http.Cookie{
		Name:     "user_info",
		Value:    encodedUserInfo,
		Path:     "/",
		HttpOnly: false,                 // Allow JS access
		Secure:   true,                  // Set to true in production with HTTPS
		SameSite: http.SameSiteNoneMode, // Use "None" for third-party context
		Expires:  time.Now().Add(config.TokenExpiry),
	})

	return c.JSON(http.StatusOK, map[string]string{"message": "Authentication successful"})
}

// GetAuthUser handles the GET request to verify authentication and return the authenticated user's details.
func (h *UserHandler) GetAuthUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTClaims)

	// Fetch user details from the database
	authenticatedUser, err := h.service.GetUserByID(c.Request().Context(), claims.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user"})
	}
	if authenticatedUser == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, authenticatedUser)
}

// Logout handles the POST request to log out the user by clearing the auth cookie.
func (h *UserHandler) Logout(c echo.Context) error {
	// Clear the auth_token cookie
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0), // Expire immediately
		HttpOnly: true,
		Secure:   true, // Use HTTPS in production
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully logged out"})
}
