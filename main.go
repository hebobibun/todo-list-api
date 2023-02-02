package main

import (
	"log"
	actD "todo-api/activity/data"
	actH "todo-api/activity/handler"
	actS "todo-api/activity/service"
	"todo-api/config"
	todoD "todo-api/todo/data"
	todoH "todo-api/todo/handler"
	todoS "todo-api/todo/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	actData := actD.New(db)
	actSrv := actS.New(actData)
	actHdl := actH.New(actSrv)

	todoData := todoD.New(db)
	todoSrv := todoS.New(todoData)
	todoHdl := todoH.New(todoSrv)

	act := e.Group("/activity-groups")

	act.POST("", actHdl.Create())
	act.GET("", actHdl.GetAll())
	act.GET("/:id", actHdl.GetOne())
	act.PATCH("/:id", actHdl.Update())
	act.DELETE("/:id", actHdl.Delete())

	todo := e.Group("/todo-items")

	todo.POST("", todoHdl.Create())
	todo.GET("/:id", todoHdl.GetOne())
	todo.GET("", todoHdl.GetAll())
	todo.PATCH("/:id", todoHdl.Update())
	todo.DELETE("/:id", todoHdl.Delete())

	if err := e.Start(":3030"); err != nil {
		log.Println(err.Error())
	}
}
