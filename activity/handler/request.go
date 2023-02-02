package handler

import "todo-api/activity"

type ActivityRequest struct {
	Title string `validate:"required" json:"title"`
	Email string `validate:"required" json:"email"`
}

type UpdateRequest struct {
	Title string `validate:"required" json:"title"`
}

func ToCore(data interface{}) *activity.Core {
	res := activity.Core{}

	switch data.(type) {
	case ActivityRequest:
		cnv := data.(ActivityRequest)
		res.Title = cnv.Title
		res.Email = cnv.Email
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Title = cnv.Title
	default:
		return nil
	}

	return &res
}
