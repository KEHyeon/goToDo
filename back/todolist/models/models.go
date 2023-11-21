package models

type Todo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"column:title" json:"title"`
	Content   string `gorm:"column:content" json:"content"`
	IsChecked bool   `gorm:"column:is_checked" json:"isChecked"`
	CreatedAt string `gorm:"column:created_at" json:"createdAt"`
}

func (Todo) TableName() string {
	return "todo"
}