package repository

import (
	"github.com/google/uuid"
	"github.com/hamwiwatsapon/todo-projects/backend/internal/domain"
	"github.com/hamwiwatsapon/todo-projects/backend/package/helper"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &authRepository{db}
}

// CreateUser implements domain.AuthRepository.
func (a *authRepository) CreateUser(userDto *domain.CreateUserDTO) (*domain.User, error) {
	hashedPassword, err := helper.HashPassword(userDto.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:       uuid.NewString(),
		Email:    userDto.Email,
		Password: hashedPassword,
		Role:     userDto.Role,
	}

	if err := a.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser implements domain.AuthRepository.
func (a *authRepository) DeleteUser(email string) error {
	panic("unimplemented")
}

// GetUserByEmail implements domain.AuthRepository.
func (a *authRepository) GetUserByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	if err := a.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (a *authRepository) GetUserByID(id string) (*domain.User, error) {
	user := &domain.User{}
	if err := a.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser implements domain.AuthRepository.
func (a *authRepository) UpdateUser(user *domain.User) error {
	panic("unimplemented")
}
