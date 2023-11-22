package service

import (
	"database/sql"
	"errors"
	"go-api/auth/models"

	"golang.org/x/crypto/bcrypt"
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


func (db *DBORM) Signup(user models.User) (models.User, error) {
	// 사용자를 이메일을 통해 조회
	var existingUser models.User
	result := db.Where("email = ?", user.Email).First(&existingUser)

	// 조회 중 에러가 발생한 경우
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return models.User{}, result.Error
	} else if result.Error == nil {
		return models.User{}, errors.New("email already registered")
	}

	// 패스워드 해시 생성
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	// 패스워드 설정 및 사용자 등록
	user.Password = string(hashedPassword)
	result = db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
func (db *DBORM) Signin(email string, password string) (string, error) {
	var user models.User
	result := db.Where("email = ?", email).First(&user);
	if result.Error != nil {
		return "", result.Error
	} 

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	} 
	return "success", nil
}