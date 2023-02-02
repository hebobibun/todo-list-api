package handler

import (
	"time"
	"todo-api/activity"
)

type ActivityResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToResponse(data activity.Core) ActivityResponse {
	return ActivityResponse{
		ID:        data.ID,
		Title:     data.Title,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToResponseArr(data []activity.Core) []ActivityResponse {
	res := []ActivityResponse{}
	for _, v := range data {
		tmp := ToResponse(v)
		res = append(res, tmp)
	}
	return res
}
