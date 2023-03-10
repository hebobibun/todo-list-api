package service

import (
	"todo-api/features/todo"
	"todo-api/helper"

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

func (s *todoService) GetOne(id uint) (todo.Core, error) {
	res, err := s.qry.GetOne(id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *todoService) GetAll(actID uint) ([]todo.Core, error) {
	res, err := s.qry.GetAll(actID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *todoService) Update(id uint, updatedTodo todo.Core) (todo.Core, error) {
	res, err := s.qry.Update(id, updatedTodo)
	if err != nil {
		return updatedTodo, err
	}

	return res, nil
}

func (s *todoService) Delete(id uint) error {
	err := s.qry.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
