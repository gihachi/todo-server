package util

import "github.com/jinzhu/gorm"

func CheckConnectError(err error){
	if err != nil{
		panic(err)
	}
}

func GetDB() *gorm.DB{
	db, err := gorm.Open("sqlite3","./db/todo.db")
	CheckConnectError(err)
	return db
}