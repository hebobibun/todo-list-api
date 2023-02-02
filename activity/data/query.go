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

func (q *activityQuery) GetOne(id uint) (activity.Core, error) {
	act := Activity{}

	err := q.db.Where("id = ?", id).First(&act).Error
	if err != nil {
		return activity.Core{}, err
	}

	return ToCores(act), nil
}

func (q *activityQuery) GetAll() ([]activity.Core, error) {
	allAct := []Activity{}

	err := q.db.Find(&allAct).Error
	if err != nil {
		return []activity.Core{}, err
	}

	return ToCoresArr(allAct), nil
}
