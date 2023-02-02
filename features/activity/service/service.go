package service

import (
	"todo-api/features/activity"
	"todo-api/helper"

	"github.com/go-playground/validator"
)

type actService struct {
	qry activity.ActivityData
	vld *validator.Validate
}

func New(ad activity.ActivityData) activity.ActivityService {
	return &actService{
		qry: ad,
		vld: validator.New(),
	}
}

func (s *actService) Create(newActivity activity.Core) (activity.Core, error) {
	err := helper.Validation(newActivity)
	if err != nil {
		return newActivity, err
	}

	res, err := s.qry.Create(newActivity)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *actService) GetOne(id uint) (activity.Core, error) {
	res, err := s.qry.GetOne(id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *actService) GetAll() ([]activity.Core, error) {
	res, err := s.qry.GetAll()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *actService) Delete(id uint) error {
	err := s.qry.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *actService) Update(id uint, updatedActivity activity.Core) (activity.Core, error) {
	res, err := s.qry.Update(id, updatedActivity)
	if err != nil {
		return updatedActivity, err
	}

	return res, nil
}
