package models

type Todo struct {
	ID        uint
	Title     string
	Content   string
	IsChecked bool
	CreatedAt string
}

func (Todo) TableName() string {
	return "todo"
}