package main

import(
	"fmt"
	"log"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"todo-server/models"
	"todo-server/util"
	"todo-server/controllers"
)

func main(){
	initMigration()
	handleRequest()
}

func initMigration(){
	db,err := gorm.Open("sqlite3","./db/todo.db")
	util.CheckConnectError(err)
	defer db.Close()
	db.AutoMigrate(&models.Todo{})
}

func handleRequest(){

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// end point
	e.POST("/api/v1/event",controllers.HandlePostTodo)
	e.GET("/api/v1/event",controllers.HandleGetAllEvents)
	e.GET("/api/v1/event/:id",controllers.HandleGetOneEvnet)

	fmt.Println("Run server localhost : 8080")
	log.Fatal(e.Start(":8080"))

}

