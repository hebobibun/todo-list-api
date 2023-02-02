package activity

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ActivityHandler interface {
	Create() echo.HandlerFunc
	GetOne() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ActivityService interface {
	Create(newActivity Core) (Core, error)
	GetOne(id uint) (Core, error)
	GetAll() ([]Core, error)
	Update(id uint, updatedActivity Core) (Core, error)
	Delete(id uint) error
}

type ActivityData interface {
	Create(newActivity Core) (Core, error)
	GetOne(id uint) (Core, error)
	GetAll() ([]Core, error)
	Update(id uint, updatedActivity Core) (Core, error)
	Delete(id uint) error
}
