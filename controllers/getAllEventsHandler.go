package controllers

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"todo-server/util"
	"todo-server/models"
)

func HandleGetAllEvents(context echo.Context) error{

	var db *gorm.DB = util.GetDB()
	defer db.Close()

	var todos []models.Todo
	db.Find(&todos)

	var returnContent struct{
		Events []models.ResponseTodo `json:"events"`
	}
	
	for _,v := range todos{
		returnContent.Events = append(returnContent.Events, *models.NewResponseTodo(v))
	}


	return context.JSON(http.StatusOK,returnContent)

}