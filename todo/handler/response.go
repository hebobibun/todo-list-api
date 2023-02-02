package handler

import (
	"time"
	"todo-api/todo"
)

type TodoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	IsActive  bool      `json:"is_active"`
	Priority  string    `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToResponse(data todo.Core) TodoResponse {
	return TodoResponse{
		ID:        data.ID,
		Title:     data.Title,
		IsActive:  data.IsActive,
		Priority:  data.Priority,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToResponseArr(data []todo.Core) []TodoResponse {
	res := []TodoResponse{}
	for _, v := range data {
		tmp := ToResponse(v)
		res = append(res, tmp)
	}

	return res
}
