package controllers

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
	"todo-server/util"
	"todo-server/models"
)

func HandleGetOneEvnet(context echo.Context) error{
	
	eventId,err := strconv.Atoi(context.Param("id"))

	if err != nil{
		return context.NoContent(http.StatusNotFound)
	}

	var db *gorm.DB = util.GetDB()
	defer db.Close()

	var event []models.Todo
	db.Where("id = ?",eventId).Find(&event)

	if len(event) == 0{
		return context.NoContent(http.StatusNotFound)
	}

	respContent := models.NewResponseTodo(event[0])
	
	return context.JSON(http.StatusOK,respContent)
}