package data

import (
	"todo-api/activity"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Title string
	Email string
}

func CoreToData(data activity.Core) Activity {
	return Activity{
		Model: gorm.Model{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
		Title: data.Title,
		Email: data.Email,
	}
}

func ToCores(data Activity) activity.Core {
	return activity.Core{
		ID:        data.ID,
		Title:     data.Title,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToCoresArr(data []Activity) []activity.Core {
	arrRes := []activity.Core{}
	for _, v := range data {
		tmp := ToCores(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}
