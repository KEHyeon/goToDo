package models

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"column:email;not null" json:"email"`
	Password string `gorm:"column:password;not null" json:"password"`
	Name     string `gorm:"column:name;not null" json:"name"`
}

func (User) TableName() string {
	return "user"
}