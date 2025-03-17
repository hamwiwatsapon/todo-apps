package domain

import "time"

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Priority    uint      `json:"priority" gorm:"default:1;check:priority >= 1 AND priority <= 3"`       // 1:low, 2:medium, 3:high
	Difficulty  uint      `json:"difficulty" gorm:"default:1;check:difficulty >= 1 AND difficulty <= 3"` // 1:easy, 2:medium, 3:hard
	Completed   bool      `json:"completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoRepository interface {
	Create(todo *Todo) error
	GetByID(id uint) (*Todo, error)
	GetAll() ([]Todo, error)
	Update(todo *Todo) error
	Delete(id uint) error
}

type TodoUsecase interface {
	Create(todo *Todo) error
	GetByID(id uint) (*Todo, error)
	GetAll() ([]Todo, error)
	Update(todo *Todo) error
	Delete(id uint) error
}
