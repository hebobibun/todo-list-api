package handler

import (
	"fmt"
	"net/http"
	"strconv"
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

func (h *acthandler) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			msg := fmt.Sprintf("Activity with ID %d Not Found", id)
			response := helper.APIResponse("Not Found", msg, nil)
			return c.JSON(http.StatusNotFound, response)
		}

		res, err := h.srv.GetOne(uint(id))
		if err != nil {
			response := helper.APIResponse("Error", "Error", nil)
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := helper.APIResponse("Success", "Success", ToResponse(res))
		return c.JSON(http.StatusOK, response)
	}
}

func (h *acthandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.srv.GetAll()
		if err != nil {
			response := helper.APIResponse("Error", "Error", nil)
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := helper.APIResponse("Success", "Success", ToResponseArr(res))
		return c.JSON(http.StatusOK, response)
	}
}

func (h *acthandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := helper.APIResponse("Error", "Error", nil)
			return c.JSON(http.StatusNotFound, response)
		}

		err = h.srv.Delete(uint(id))
		if err != nil {
			msg := fmt.Sprintf("Activity with ID %d Not Found", id)
			response := helper.APIResponse("Not found", msg, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := helper.APIResponse("Success", "Success", nil)
		return c.JSON(http.StatusOK, response)
	}
}
