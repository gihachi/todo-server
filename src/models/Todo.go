package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct{
	gorm.Model
	Deadline string `json:"deadline" validate:"required"`
	Title string `json:"title" validate:"required"`
	Memo string `json:"memo" validate:"required"`
}

func (todo Todo) NewTodo(deadline, title ,memo string) Todo{
	return Todo{Deadline:deadline,Title:title,Memo:memo}
}