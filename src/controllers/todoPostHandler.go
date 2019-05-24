package controllers

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"todo-server/src/util"
	"todo-server/src/models"
)

func HandlePostTodo(context echo.Context) error {

	db, err := gorm.Open("sqlite3","./db/todo.db")
	util.CheckConnectError(err)
	defer db.Close()

	todo := new(models.Todo)
	if err = context.Bind(todo);err != nil{

		var returnContent struct {
			Status string `json:"status"`
			Message string `json:"message"`
		}
		returnContent.Status =  "failure"
		returnContent.Message =  "invalid date format"

		return context.JSON(http.StatusBadRequest, returnContent)
	} 

	db.Create(&todo)

	var returnContent struct{
		Status string `json:"status"`
		Message string `json:"message"`
		Id uint `json:"id"`
	}
	returnContent.Status = "success"
	returnContent.Message = "registered"
	returnContent.Id = todo.ID

	return context.JSON(http.StatusOK, returnContent)

}