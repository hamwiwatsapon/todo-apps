package repository

import (
	"github.com/hamwiwatsapon/todo-projects/backend/internal/domain"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) domain.TodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) Create(dto *domain.CreateTodoDTO) (*domain.Todo, error) {
	todo := &domain.Todo{
		Title:       dto.Title,
		Description: dto.Description,
		Priority:    dto.Priority,
		Difficulty:  dto.Difficulty,
		Completed:   false,
	}

	if err := r.db.Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepository) GetByID(id uint) (*domain.Todo, error) {
	var todo domain.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) GetAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) Update(todo *domain.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Todo{}, id).Error
}
