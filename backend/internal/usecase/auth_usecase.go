package usecase

import (
	"errors"

	"github.com/hamwiwatsapon/todo-projects/backend/internal/domain"
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
	panic("unimplemented")
}

// Login implements domain.AuthUsecase.
func (a *authUsecase) Login(email string, password string) (string, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.AuthUsecase.
func (a *authUsecase) UpdateUser(user *domain.User) error {
	panic("unimplemented")
}
