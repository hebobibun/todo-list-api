package service

import (
	"todo-api/helper"
	"todo-api/todo"

	"github.com/go-playground/validator"
)

type todoService struct {
	qry todo.TodoData
	vld *validator.Validate
}

func New(td todo.TodoData) todo.TodoService {
	return &todoService{
		qry: td,
		vld: validator.New(),
	}
}

func (s *todoService) Create(newTodo todo.Core) (todo.Core, error) {
	err := helper.Validation(newTodo)
	if err != nil {
		return newTodo, err
	}

	res, err := s.qry.Create(newTodo)
	if err != nil {
		return res, err
	}

	return res, nil
}
