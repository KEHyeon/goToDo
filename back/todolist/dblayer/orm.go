package dblayer

import (
	"database/sql"
	"go-api/todolist/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbengine string, dsn string) (*DBORM, error) {
	sqlDB, err := sql.Open(dbengine, dsn)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	
	return &DBORM{
		DB: gormDB,
	}, err
}

func (db *DBORM) GetAllContents() (todoList []models.Todo, err error) {
	return todoList, db.Find(&todoList).Error
}