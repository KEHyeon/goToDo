package mmodels

type todo struct {
	ID uint
	title string
	Content string
	isChecked bool
}

func (Todo) TableName() string {
	return "todo"
}