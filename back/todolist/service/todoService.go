package service

import (
	"database/sql"
	"errors"
	"fmt"
	"go-api/todolist/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func GetTodoService(dbengine string, dsn string) (*DBORM, error) {
	sqlDB, err := sql.Open(dbengine, dsn)

	gormDB, err := gorm.Open(
		mysql.New(mysql.Config{Conn: sqlDB}),
			&gorm.Config{})
	
	gormDB.AutoMigrate(&models.Todo{})
	return &DBORM{
		DB: gormDB,
	}, err
}

func (db *DBORM) GetAllTodo() ([]models.Todo, error) {
	var todoList []models.Todo
	results := db.Find(&todoList);
	return todoList, results.Error
}

func (db *DBORM) CreateTodo(todo models.Todo) (models.Todo, error) {
	return todo, db.Create(&todo).Error
}

func (db *DBORM) GetOneTodo(id uint) (models.Todo, error) {
	var todo models.Todo
	result := db.First(&todo, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return todo, fmt.Errorf("Todo with ID %d not found", id)
		}
		return todo, result.Error
	}

	return todo, nil
}

func (db *DBORM) ToggleTodo(id uint) (string, error) {
	todo, err := db.GetOneTodo(id)
	if err != nil {
		return "Fail", fmt.Errorf("Failed to toggle todo: %v", err)
	}

	todo.IsChecked = !todo.IsChecked

	if err := db.Save(&todo).Error; err != nil {
		return "Fail", fmt.Errorf("Failed to save todo: %v", err)
	}

	return "Ok", nil
}