package data

import (
	"errors"
	"fmt"
	"log"
	"todo-api/features/activity"

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
		log.Println("Query create a new activity error : ", err.Error())
		return newActivity, err
	}

	fmt.Println(activity)
	fmt.Println(ToCores(activity))

	return ToCores(activity), nil
}

func (q *activityQuery) GetOne(id uint) (activity.Core, error) {
	act := Activity{}

	err := q.db.Where("id = ?", id).First(&act).Error
	if err != nil {
		log.Println("Query get activity by ID error : ", err.Error())
		return activity.Core{}, err
	}

	return ToCores(act), nil
}

func (q *activityQuery) GetAll() ([]activity.Core, error) {
	allAct := []Activity{}

	err := q.db.Find(&allAct).Error
	if err != nil {
		log.Println("Query get All activities error : ", err.Error())
		return []activity.Core{}, err
	}

	return ToCoresArr(allAct), nil
}

func (q *activityQuery) Delete(id uint) error {
	qryDelete := q.db.Delete(&Activity{}, id)

	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		msg := fmt.Sprintf("Activity with ID %d Not Found", id)
		return errors.New(msg)
	}

	return nil
}

func (q *activityQuery) Update(id uint, updatedActivity activity.Core) (activity.Core, error) {
	cnvUpdated := CoreToData(updatedActivity)
	qry := q.db.Model(Activity{}).Where("id = ?", id).Updates(&cnvUpdated)
	err := qry.Error

	affRow := qry.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		msg := fmt.Sprintf("Activity with ID %d Not Found", id)
		return activity.Core{}, errors.New(msg)
	}

	if err != nil {
		log.Println("Query update activity by ID error : ", err.Error())
		return activity.Core{}, errors.New("Error")
	}

	var updatedRow Activity
	q.db.First(&updatedRow, "id = ?", id)

	return ToCores(updatedRow), nil
}
