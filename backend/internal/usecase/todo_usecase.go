package usecase

import (
	"github.com/hamwiwatsapon/todo-projects/backend/internal/domain"
)

type todoUsecase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUsecase(repo domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{
		todoRepo: repo,
	}
}

func (u *todoUsecase) Create(todo *domain.CreateTodoDTO) (*domain.Todo, error) {
	return u.todoRepo.Create(todo)
}

func (u *todoUsecase) GetByID(id uint) (*domain.Todo, error) {
	return u.todoRepo.GetByID(id)
}

func (u *todoUsecase) GetAllByUserID(userID string) ([]domain.Todo, error) {
	return u.todoRepo.GetAllByUserID(userID)
}

func (u *todoUsecase) Update(todo *domain.Todo) error {
	return u.todoRepo.Update(todo)
}

func (u *todoUsecase) Delete(id uint) error {
	return u.todoRepo.Delete(id)
}
