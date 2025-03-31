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
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateTodoDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    uint   `json:"priority" example:"1" enums:"1,2,3" description:"Priority level: 1 (low), 2 (medium), 3 (high)"`
	Difficulty  uint   `json:"difficulty" example:"1" enums:"1,2,3" description:"Difficulty level: 1 (easy), 2 (medium), 3 (hard)"`
}

type GetAllTodosResponse struct {
	Data []Todo `json:"data" example:"[{\"id\":1,\"title\":\"Learn Go\",\"description\":\"Study Golang programming\",\"priority\":2,\"difficulty\":1,\"completed\":false,\"created_at\":\"2024-03-20T10:00:00Z\",\"updated_at\":\"2024-03-20T10:00:00Z\"}]"`
}

type TodoRepository interface {
	Create(todo *CreateTodoDTO) (*Todo, error)
	GetByID(id uint) (*Todo, error)
	GetAll() ([]Todo, error)
	Update(todo *Todo) error
	Delete(id uint) error
}

type TodoUsecase interface {
	Create(todo *CreateTodoDTO) (*Todo, error)
	GetByID(id uint) (*Todo, error)
	GetAll() ([]Todo, error)
	Update(todo *Todo) error
	Delete(id uint) error
}

type ErrorResponse400 struct {
	Error string `json:"error" example:"Invalid request body"`
}
type ErrorResponse404 struct {
	Error string `json:"error" example:"Todo not found"`
}

type ErrorResponse500 struct {
	Error string `json:"error" example:"Internal server error"`
}
