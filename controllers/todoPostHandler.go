package controllers

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/go-playground/validator.v9"
	"todo-server/util"
	"todo-server/models"
)

func HandlePostTodo(context echo.Context) error {

	var db *gorm.DB = util.GetDB()
	defer db.Close()

	todo := new(models.Todo)
	if err := context.Bind(todo);err != nil{
		return context.JSON(http.StatusBadRequest, makeFailureContent())
	} 

	var validate *validator.Validate
	validate = validator.New()
	validate.RegisterValidation("is_date",isDate)

	if err := validate.Struct(todo); err != nil{
		return context.JSON(http.StatusBadRequest, makeFailureContent())
	}

	if !db.NewRecord(todo){
		return context.JSON(http.StatusBadRequest, makeFailureContent())
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

func isDate(fl validator.FieldLevel) bool{

	date := fl.Field().String()
	layout := "2006-01-02T15:04:05Z07:00"
	_,err := time.Parse(layout,date)
	if err != nil{
		return false
	}
	return true
}

type failureContent struct{
	Status string `json:"status"`
	Message string `json:"message"`
}

func makeFailureContent() failureContent{
	var fContent failureContent
	fContent.Status =  "failure"
	fContent.Message =  "invalid date format"
	return fContent
}