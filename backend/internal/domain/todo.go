package domain

import (
	"time"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey" example:"1"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Priority    uint      `json:"priority" gorm:"default:1;check:priority >= 1 AND priority <= 3" example:"1" enums:"1,2,3" description:"Priority level: 1 (low), 2 (medium), 3 (high)"`          // 1:low, 2:medium, 3:high
	Difficulty  uint      `json:"difficulty" gorm:"default:1;check:difficulty >= 1 AND difficulty <= 3" example:"1" enums:"1,2,3" description:"Difficulty level: 1 (easy), 2 (medium), 3 (hard)"` // 1:easy, 2:medium, 3:hard
	Completed   bool      `json:"completed" gorm:"default:false"`
	EndDate     time.Time `json:"end_date" gorm:"default:NULL"`
	UserID      string    `json:"user_id" gorm:"not null;index"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateTodoDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    uint   `json:"priority" example:"1" enums:"1,2,3" description:"Priority level: 1 (low), 2 (medium), 3 (high)"`
	Difficulty  uint   `json:"difficulty" example:"1" enums:"1,2,3" description:"Difficulty level: 1 (easy), 2 (medium), 3 (hard)"`
	EndDate     string `json:"end_date" example:"2024-03-20T10:00:00Z" description:"End date in RFC3339 format"`
	UserID      string `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000" description:"User ID"`
}

type GetAllTodosResponse struct {
	Data []Todo `json:"data" example:"[{\"id\":1,\"title\":\"Learn Go\",\"description\":\"Study Golang programming\",\"priority\":2,\"difficulty\":1,\"completed\":false,\"created_at\":\"2024-03-20T10:00:00Z\",\"updated_at\":\"2024-03-20T10:00:00Z\"}]"`
}

type TodoRepository interface {
	Create(todo *CreateTodoDTO) (*Todo, error)
	GetByID(id uint) (*Todo, error)
	GetAllByUserID(userID string) ([]Todo, error)
	Update(todo *Todo) error
	Delete(id uint) error
}

type TodoUsecase interface {
	Create(todo *CreateTodoDTO) (*Todo, error)
	GetByID(id uint) (*Todo, error)
	GetAllByUserID(userID string) ([]Todo, error)
	Update(todo *Todo) error
	Delete(id uint) error
}

type ErrorResponse400 struct {
	Error string `json:"error" example:"Invalid request body"`
}

type ErrorResponse401 struct {
	Error string `json:"error" example:"Unauthorized"`
}

type ErrorResponse403 struct {
	Error string `json:"error" example:"Forbidden"`
}

type ErrorResponse404 struct {
	Error string `json:"error" example:"Todo not found"`
}

type ErrorResponse500 struct {
	Error string `json:"error" example:"Internal server error"`
}
