package models

import (
       "github.com/jinzhu/gorm"
)

type Todo struct{
       gorm.Model
       Deadline string `json:"deadline" validate:"is_date,required"`
       Title string `json:"title" validate:"required"`
       Memo string `json:"memo"`
}

func (todo Todo) NewTodo(deadline, title ,memo string) Todo{
       return Todo{Deadline:deadline,Title:title,Memo:memo}
}

type ResponseTodo struct{
       ID uint
       Deadline string
       Title string
       Memo string
}

func NewResponseTodo(todo Todo) *ResponseTodo{
       respTodo := &ResponseTodo{}
       respTodo.ID = todo.ID
       respTodo.Deadline = todo.Deadline
       respTodo.Title = todo.Title
       respTodo.Memo = todo.Memo
       return respTodo
}
