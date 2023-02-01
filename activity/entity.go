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
}

type ActivityService interface {
	Create(newActivity Core) (Core, error)
}

type ActivityData interface {
	Create(newActivity Core) (Core, error)
}
