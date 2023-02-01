package main

import (
	"fmt"
	"log"
	"todo-api/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	fmt.Println(db, "\nconnection is good")

	if err := e.Start(":3030"); err != nil {
		log.Println(err.Error())
	}
}
