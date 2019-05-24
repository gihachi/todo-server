package controllers

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"todo-server/src/util"
	"todo-server/src/models"
)

func HandleGetAllEvents(context echo.Context) error{

	db, err := gorm.Open("sqlite3","./db/todo.db")
	util.CheckConnectError(err)
	defer db.Close()

	var todos []models.Todo
	db.Find(&todos)

	var returnContent struct{
		Events []models.Todo `json:"events"`
	}
	returnContent.Events = todos

	return context.JSON(http.StatusOK,returnContent)

}