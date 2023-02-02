package data

import (
	"todo-api/todo"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title           string
	IsActive        bool
	Priority        string
	ActivityGroupID uint
}

func CoreToData(data todo.Core) Todo {
	return Todo{
		Model: gorm.Model{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
		Title:           data.Title,
		IsActive:        data.IsActive,
		Priority:        data.Priority,
		ActivityGroupID: data.ActivityGroupID,
	}
}

func ToCores(data Todo) todo.Core {
	return todo.Core{
		ID:              data.ID,
		Title:           data.Title,
		IsActive:        data.IsActive,
		Priority:        data.Priority,
		ActivityGroupID: data.ActivityGroupID,
	}
}

func ToCoreArr(data []Todo) []todo.Core {
	arrRes := []todo.Core{}
	for _, v := range data {
		tmp := ToCores(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}
