package controllers

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
	"todo-server/src/util"
	"todo-server/src/models"
)

func HandleGetOneEvnet(context echo.Context) error{
	
	eventId,err := strconv.Atoi(context.Param("id"))

	if err != nil{
		return context.NoContent(http.StatusNotFound)
	}

	db, err := gorm.Open("sqlite3","./db/todo.db")
	util.CheckConnectError(err)
	defer db.Close()

	var event []models.Todo
	db.Where("id = ?",eventId).Find(&event)

	if len(event) == 0{
		return context.NoContent(http.StatusNotFound)
	}

	var returnContent struct{
		ID uint `json:"id"`
		Deadline string `json:"deadline"`
		Title string `json:"title"`
		Memo string  `json:"memo"`
	}

	returnContent.ID = event[0].ID
	returnContent.Deadline = event[0].Deadline
	returnContent.Title = event[0].Title
	returnContent.Memo = event[0].Memo
	
	return context.JSON(http.StatusOK,returnContent)
}