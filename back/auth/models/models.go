package models

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	email    string `gorm:"column:email" json:"email"`
	password string `gorm:"column:password" json:"password"`
	name     string `gorm:"column:name" json:"name"`
}

func (User) TableName() string {
	return "todo"
}