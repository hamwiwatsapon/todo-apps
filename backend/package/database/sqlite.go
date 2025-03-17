package database

import (
	"github.com/hamwiwatsapon/todo-projects/backend/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&domain.Todo{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
