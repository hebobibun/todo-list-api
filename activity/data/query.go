package data

import (
	"todo-api/activity"

	"gorm.io/gorm"
)

type activityQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) activity.ActivityData {
	return &activityQuery{
		db: db,
	}
}

func (q *activityQuery) Create(newActivity activity.Core) (activity.Core, error) {
	activity := CoreToData(newActivity)

	err := q.db.Create(&activity).Error
	if err != nil {
		return newActivity, err
	}

	return ToCores(activity), nil
}
