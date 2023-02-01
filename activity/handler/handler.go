package handler

import (
	"net/http"
	"todo-api/activity"
	"todo-api/helper"

	"github.com/labstack/echo/v4"
)

type acthandler struct {
	srv activity.ActivityService
}

func New(srv activity.ActivityService) activity.ActivityHandler {
	return &acthandler{
		srv: srv,
	}
}

func (h *acthandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := ActivityRequest{}
		if err := c.Bind(&input); err != nil {
			response := helper.APIResponse("Bad Request", "Bad Request", nil)
			return c.JSON(http.StatusBadRequest, response)
		}

		res, err := h.srv.Create(*ToCore(input))
		if err != nil {
			response := helper.APIResponse("Error", "Error", nil)
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := helper.APIResponse("Success", "Success", ToResponse(res))

		return c.JSON(http.StatusCreated, response)
	}
}
