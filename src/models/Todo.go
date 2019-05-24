package models

import "github.com/jinzhu/gorm"

type Todo struct{
	gorm.Model
	Deadline string `json:"deadline"`
	Title string `json:"title"`
	Memo string `json:"memo"`
}

func (todo Todo) NewTodo(deadline, title ,memo string) Todo{
	return Todo{Deadline:deadline,Title:title,Memo:memo}
}