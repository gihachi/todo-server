package models

import "github.com/jinzhu/gorm"

type Todo struct{
	gorm.Model
	Deadline string
	Title string
	Memo string
}