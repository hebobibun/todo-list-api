package data

import (
	"fmt"
	"log"
	"todo-api/todo"

	"gorm.io/gorm"
)

type todoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) todo.TodoData {
	return &todoQuery{
		db: db,
	}
}

func (q *todoQuery) Create(newTodo todo.Core) (todo.Core, error) {
	todo := CoreToData(newTodo)

	err := q.db.Create(&todo).Error
	if err != nil {
		log.Println("Query create a new todo error : ", err.Error())
		return newTodo, err
	}

	fmt.Println(todo)
	fmt.Println(ToCores(todo))

	return ToCores(todo), nil
}
