package service

import (
	"database/sql"
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

// func (db *DBORM) CreateTodo(todo models.Todo)  models.Todo, error) {
// 	return todo, db.Create(todo).Error
// }