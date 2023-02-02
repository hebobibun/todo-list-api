package data

import (
	"errors"
	"fmt"
	"log"
	"todo-api/features/todo"

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

	return ToCores(todo), nil
}

func (q *todoQuery) GetOne(id uint) (todo.Core, error) {
	act := Todo{}

	err := q.db.Where("id = ?", id).First(&act).Error
	if err != nil {
		log.Println("Query get todo by ID error : ", err.Error())
		return todo.Core{}, err
	}

	return ToCores(act), nil
}

func (q *todoQuery) GetAll(actID uint) ([]todo.Core, error) {
	allTodo := []Todo{}

	if actID <= 0 {
		err := q.db.Find(&allTodo).Error
		if err != nil {
			log.Println("Query get All activities error : ", err.Error())
			return []todo.Core{}, err
		}
	} else {
		err := q.db.Where("activity_group_id = ?", actID).Find(&allTodo).Error
		if err != nil {
			log.Println("Query get All todo error : ", err.Error())
			return []todo.Core{}, err
		}
	}

	return ToCoreArr(allTodo), nil
}

func (q *todoQuery) Update(id uint, updatedTodo todo.Core) (todo.Core, error) {
	cnvUpdated := CoreToData(updatedTodo)

	qry := q.db.Model(Todo{}).Where("id = ?", id).Updates(&cnvUpdated)
	toggle := q.db.Model(&cnvUpdated).Where("id = ?", id).Update("is_active", updatedTodo.IsActive)
	err := qry.Error

	affRow := toggle.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		msg := fmt.Sprintf("Todo with ID %d Not Found", id)
		return todo.Core{}, errors.New(msg)
	}

	affRow = qry.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		msg := fmt.Sprintf("Todo with ID %d Not Found", id)
		return todo.Core{}, errors.New(msg)
	}

	if err != nil {
		log.Println("Query update todo by ID error : ", err.Error())
		return todo.Core{}, errors.New("Error")
	}

	var updatedRow Todo
	q.db.First(&updatedRow, "id = ?", id)

	return ToCores(updatedRow), nil
}

func (q *todoQuery) Delete(id uint) error {
	qryDelete := q.db.Delete(&Todo{}, id)

	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		msg := fmt.Sprintf("Todo with ID %d Not Found", id)
		return errors.New(msg)
	}

	return nil
}
