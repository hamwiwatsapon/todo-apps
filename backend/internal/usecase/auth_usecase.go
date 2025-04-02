package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hamwiwatsapon/todo-projects/backend/internal/domain"
	"github.com/hamwiwatsapon/todo-projects/backend/package/helper"
)

type authUsecase struct {
	authRepo domain.AuthRepository
}

func NewAuthUsecase(repo domain.AuthRepository) domain.AuthUsecase {
	return &authUsecase{
		authRepo: repo,
	}
}

// CreateUser implements domain.AuthUsecase.
func (a *authUsecase) CreateUser(userDto *domain.CreateUserDTO) (*domain.User, error) {
	// Check if the email is already taken
	existingUser, _ := a.authRepo.GetUserByEmail(userDto.Email) // Assume GetUserByEmail exists
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	// Call repository to create a user
	createdUser, err := a.authRepo.CreateUser(userDto)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// DeleteUser implements domain.AuthUsecase.
func (a *authUsecase) DeleteUser(email string) error {
	panic("unimplemented")
}

// GetUserByEmail implements domain.AuthUsecase.
func (a *authUsecase) GetUserByEmail(email string) (*domain.User, error) {
	user, err := a.authRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Login implements domain.AuthUsecase.
func (a *authUsecase) Login(email, password string) (string, string, error) {
	user, err := a.authRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return "", "", errors.New("user not found")
	}

	if !helper.VerifyPassword(password, user.Password) {
		return "", "", errors.New("invalid password")
	}

	// Ideally, generate a JWT token instead of returning the email.
	token, refreshToken, err := helper.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return "", "", errors.New("failed to generate token")
	}

	return token, refreshToken, nil
}

// RefreshToken validates the refresh token and generates a new access token.
func (a *authUsecase) RefreshToken(refreshToken string) (string, string, error) {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		return "", "", errors.New("JWT secret is not set")
	}
	secretKey := []byte(jwtSecret)

	// Parse and validate the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", errors.New("invalid refresh token claims")
	}

	// Check expiration
	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return "", "", errors.New("refresh token expired")
	}

	// Extract user details
	userID, _ := claims["user_id"].(string)

	// Retrieve user details from database (optional)
	user, err := a.authRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return "", "", errors.New("user not found")
	}

	// Generate new access & refresh tokens
	newAccessToken, newRefreshToken, err := helper.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return "", "", errors.New("failed to generate new tokens")
	}

	return newAccessToken, newRefreshToken, nil
}

// GetUserByToken retrieves user information based on the provided token.
func (a *authUsecase) GetUserByToken(tokenString string) (*domain.User, error) {
	// Remove "Bearer " prefix
	tokenString = tokenString[len("Bearer "):]

	// Get the JWT secret key from environment variables
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		return nil, errors.New("JWT secret is not set")
	}
	secretKey := []byte(jwtSecret)

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	// Extract user ID from claims
	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("user ID not found in token claims")
	}

	// Retrieve user details from the repository
	user, err := a.authRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// UpdateUser implements domain.AuthUsecase.
func (a *authUsecase) UpdateUser(user *domain.User) error {
	panic("unimplemented")
}
