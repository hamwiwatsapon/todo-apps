package domain

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Role      string    `json:"role" gorm:"default:'user'"` // user, admin
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type AuthRepository interface {
	CreateUser(user *CreateUserDTO) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(email string) error
}

type AuthUsecase interface {
	CreateUser(user *CreateUserDTO) (*User, error)
	Login(email, password string) (string, string, error) // returns token
	RefreshToken(token string) (string, string, error)    // returns new token
	GetUserByEmail(email string) (*User, error)
	GetUserByToken(token string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(email string) error
}

type CreateUserDTO struct {
	Email    string `json:"email" example:"test@email.user.com" description:"User email"`
	Password string `json:"password" example:"passwordconfig" description:"User password"`
	Role     string `json:"role" example:"user" description:"User role"` // user, admin
}

type LoginDTO struct {
	Email    string `json:"email" example:"test@email.user.com" description:"User email"`
	Password string `json:"password" example:"passwordconfig" description:"User password"`
}

type LoginReturn struct {
	Message      string `json:"message" example:"Login successful"`
	Token        string `json:"token" example:"your_token_here"`
	RefreshToken string `json:"refresh_token" example:"your_refresh_token_here"`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token" example:"your_token_here"`
}

type RefreshTokenReturn struct {
	Message      string `json:"message" example:"Token refreshed successfully"`
	Token        string `json:"token" example:"your_new_token_here"`
	RefreshToken string `json:"refresh_token" example:"your_new_refresh_token_here"`
}
