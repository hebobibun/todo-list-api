package main

import (
	"fmt"
	"todo-api/config"
)

func main() {
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	fmt.Println(db)
	fmt.Println("Database connection is good")
}
