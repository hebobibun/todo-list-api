package data

import (
	"errors"
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

func (q *todoQuery) GetOne(id uint) (todo.Core, error) {
	act := Todo{}

	err := q.db.Where("id = ?", id).First(&act).Error
	if err != nil {
		log.Println("Query get activity by ID error : ", err.Error())
		return todo.Core{}, err
	}

	return ToCores(act), nil
}

func (q *todoQuery) Delete(id uint) error {
	qryDelete := q.db.Delete(&Todo{}, id)

	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		msg := fmt.Sprintf("Activity with ID %d Not Found", id)
		return errors.New(msg)
	}

	return nil
}
