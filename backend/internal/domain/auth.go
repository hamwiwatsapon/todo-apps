package domain

type User struct {
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Role      string `json:"role" gorm:"default:'user'"` // user, admin
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}

type AuthRepository interface {
	CreateUser(user *CreateUserDTO) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(email string) error
}

type AuthUsecase interface {
	CreateUser(user *CreateUserDTO) (*User, error)
	Login(email, password string) (string, error) // returns token
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(email string) error
}

type CreateUserDTO struct {
	Email    string `json:"email" example:"test@email.user.com" description:"User email"`
	Password string `json:"password" example:"passwordconfig" description:"User password"`
	Role     string `json:"role" example:"user" description:"User role"` // user, admin
}
