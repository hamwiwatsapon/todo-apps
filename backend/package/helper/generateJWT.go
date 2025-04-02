package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT generates both an access and a refresh token.
func GenerateJWT(userID, email, role string) (string, string, error) {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		return "", "", errors.New("JWT secret is not set")
	}

	// Convert secret key to []byte
	secretKey := []byte(jwtSecret)

	// Access token expires in 24 hours
	accessTokenExp := time.Now().Add(24 * time.Hour)

	// Create claims for access token
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     accessTokenExp.Unix(), // Expiration time
		"iat":     time.Now().Unix(),     // Issued at
	}

	// Create access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", errors.New("failed to sign access token")
	}

	// Refresh token expires in 7 days
	refreshTokenExp := time.Now().Add(7 * 24 * time.Hour)

	// Create claims for refresh token
	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     refreshTokenExp.Unix(),
		"iat":     time.Now().Unix(),
	}

	// Create refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", "", errors.New("failed to sign refresh token")
	}

	return accessTokenString, refreshTokenString, nil
}
