package todo

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID              uint
	Title           string
	Priority        string
	IsActive        bool
	ActivityGroupID uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TodoHandler interface {
	Create() echo.HandlerFunc
	GetOne() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
type TodoService interface {
	Create(newTodo Core) (Core, error)
	GetOne(id uint) (Core, error)
	GetAll(actID uint) ([]Core, error)
	Update(id uint, updatedTodo Core) (Core, error)
	Delete(id uint) error
}

type TodoData interface {
	Create(newTodo Core) (Core, error)
	GetOne(id uint) (Core, error)
	GetAll(actID uint) ([]Core, error)
	Update(id uint, updatedTodo Core) (Core, error)
	Delete(id uint) error
}
