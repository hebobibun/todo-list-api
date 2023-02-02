package handler

import "todo-api/todo"

type TodoRequest struct {
	Title           string `validate:"required" json:"title"`
	Priority        string `json:"priority"`
	ActivityGroupID uint   `json:"activity_group_id"`
}

type TodoUpdateRequest struct {
	Title    string `json:"title"`
	Priority string `json:"priority"`
	IsActive bool   `json:"is_active"`
}

func ToCore(data interface{}) *todo.Core {
	res := todo.Core{}

	switch data.(type) {
	case TodoRequest:
		cnv := data.(TodoRequest)
		res.Title = cnv.Title
		res.Priority = cnv.Priority
		res.ActivityGroupID = cnv.ActivityGroupID
	case TodoUpdateRequest:
		cnv := data.(TodoUpdateRequest)
		res.Title = cnv.Title
		res.Priority = cnv.Priority
		res.IsActive = cnv.IsActive
	default:
		return nil
	}

	return &res
}
