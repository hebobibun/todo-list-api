package main

import (
	"log"
	"todo-api/activity/data"
	"todo-api/activity/handler"
	"todo-api/activity/service"
	"todo-api/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	actData := data.New(db)
	actSrv := service.New(actData)
	actHdl := handler.New(actSrv)

	act := e.Group("/activity-groups")

	act.POST("", actHdl.Create())

	if err := e.Start(":3030"); err != nil {
		log.Println(err.Error())
	}
}
