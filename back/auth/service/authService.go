package service

import (
	"database/sql"
	"go-api/auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func GetAuthService(dbengine string, dsn string) (*DBORM, error) {
	sqlDB, err := sql.Open(dbengine, dsn)

	gormDB, err := gorm.Open(
		mysql.New(mysql.Config{Conn: sqlDB}),
			&gorm.Config{})
	
	gormDB.AutoMigrate(&models.User{})
	return &DBORM{
		DB: gormDB,
	}, err
}


func Register(user models.User) (models.User, error){}
func Login(email string, password string) (models.User, error){}