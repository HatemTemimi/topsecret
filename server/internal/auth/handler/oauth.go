package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"server/config"
	"server/internal/user/service"
	"server/internal/user/types"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthHandler struct {
	service     service.UserService
	oauthConfig *oauth2.Config
}

func NewOAuthHandler(userService service.UserService) *OAuthHandler {

	oauthid := config.GetEnv("GOOGLE_OAUTH_CLIENT_ID", "")
	oauthsecret := config.GetEnv("GOOGLE_OAUTH_CLIENT_SECRET", "")

	return &OAuthHandler{
		service: userService,
		oauthConfig: &oauth2.Config{
			ClientID:     oauthid,
			ClientSecret: oauthsecret,
			RedirectURL:  "http://localhost:3001/api/auth/google/callback",
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		},
	}
}

// GenerateRandomPassword generates a random password of the specified length
func GenerateRandomPassword(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	// Use base64 encoding to ensure the password is alphanumeric
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

// Expose Google Login URL
func (h *OAuthHandler) GoogleLogin(c echo.Context) error {
	url := h.oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.JSON(http.StatusOK, map[string]string{"url": url})
}

func (h *OAuthHandler) GoogleCallback(c echo.Context) error {
	// Extract authorization code
	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing authorization code"})
	}

	// Exchange authorization code for access token
	token, err := h.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to exchange token"})
	}

	// Fetch user info from Google
	client := h.oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user info"})
	}
	defer resp.Body.Close()

	var googleUser struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse user info"})
	}

	// Check if user exists or create a new one
	user, err := h.service.GetUserByEmail(c.Request().Context(), googleUser.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user data"})
	}

	if user == nil {

		password, _ := GenerateRandomPassword(8)
		// Create a new user if not found
		user = &types.User{
			Email:     googleUser.Email,
			FirstName: googleUser.GivenName,
			Password:  password,
			LastName:  googleUser.FamilyName,
			Role:      "user",
			Active:    true,
		}
		id, err := h.service.CreateUser(c.Request().Context(), user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
		}
		user.ID, _ = primitive.ObjectIDFromHex(id)
	}

	// Generate JWT token
	tokenString, err := config.GenerateToken(user.ID.Hex(), user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	// Set auth token cookie
	c.SetCookie(&http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(config.TokenExpiry),
	})

	// Prepare user info for cookies
	userJSON, err := json.Marshal(map[string]string{
		"id":        user.ID.Hex(),
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"role":      user.Role,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to serialize user info"})
	}

	encodedUserInfo := url.QueryEscape(string(userJSON))

	c.SetCookie(&http.Cookie{
		Name:     "user_info",
		Value:    encodedUserInfo,
		Path:     "/",
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(config.TokenExpiry),
	})

	// Redirect user to frontend
	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/rentals")
}

// Authenticate handles the POST request for user authentication and issues a JWT token.
func (h *OAuthHandler) AuthenticateWithCookie(c echo.Context) error {
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
func (h *OAuthHandler) GetAuthUser(c echo.Context) error {
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
func (h *OAuthHandler) Logout(c echo.Context) error {
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
