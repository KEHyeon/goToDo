package service

import (
	"database/sql"
	"errors"
	"fmt"
	"go-api/todolist/models"
	"time"

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

func (db *DBORM) GetAllTodoWithDate(date string) ([]models.Todo, error) {
	var todoList []models.Todo

	// 파싱할 때 KST 시간대를 명시적으로 지정
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return todoList, err
	}

	startDate, err := time.ParseInLocation("2006-01-02", date, location)
	if err != nil {
		return todoList, err
	}

	// endDate를 계산할 때 startDate에 하루를 더합니다.
	endDate := startDate.AddDate(0, 0, 1)
	fmt.Print(startDate, endDate)
	results := db.Where("created_at BETWEEN ? AND ?", startDate.String(), endDate.String()).Find(&todoList)

	// 에러를 확인하고 반환합니다.
	if results.Error != nil {
		return todoList, results.Error
	}

	return todoList, nil
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