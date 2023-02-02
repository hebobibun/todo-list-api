package handler

import "todo-api/todo"

type TodoRequest struct {
	Title           string `validate:"required" json:"title"`
	Priority        string `json:"priority"`
	ActivityGroupID uint   `json:"activity_group_id"`
}

func ToCore(data interface{}) *todo.Core {
	res := todo.Core{}

	switch data.(type) {
	case TodoRequest:
		cnv := data.(TodoRequest)
		res.Title = cnv.Title
		res.Priority = cnv.Priority
		res.ActivityGroupID = cnv.ActivityGroupID
	default:
		return nil
	}

	return &res
}
