package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todo-api/features/activity"
	"todo-api/helper"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type actHandler struct {
	srv activity.ActivityService
}

func New(srv activity.ActivityService) activity.ActivityHandler {
	return &actHandler{
		srv: srv,
	}
}

func (h *actHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := ActivityRequest{}
		if err := c.Bind(&input); err != nil {
			response := helper.APIResponseNoData("Bad Request", "Bad Request")
			return c.JSON(http.StatusBadRequest, response)
		}

		validate := validator.New()
		if err := validate.Struct(input); err != nil {
			msg := ""
			fmt.Println(err.Error())
			if strings.Contains(err.Error(), "Title") {
				msg = "title cannot be null"
			} else if strings.Contains(err.Error(), "Email") {
				msg = "email cannot be null"
			} else {
				msg = "request body cannot be null"
			}

			response := helper.APIResponseNoData("Bad Request", msg)
			return c.JSON(http.StatusBadRequest, response)
		}

		res, err := h.srv.Create(*ToCore(input))
		if err != nil {
			response := helper.APIResponseNoData("Error", "Error")
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := helper.APIResponse("Success", "Success", ToResponse(res))
		return c.JSON(http.StatusCreated, response)
	}
}

func (h *actHandler) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := helper.APIResponseNoData("Error", "Error")
			return c.JSON(http.StatusBadRequest, response)
		}

		res, err := h.srv.GetOne(uint(id))
		if err != nil {
			msg := fmt.Sprintf("Activity with ID %d Not Found", id)
			response := helper.APIResponseNoData("Not Found", msg)
			return c.JSON(http.StatusNotFound, response)
		}

		response := helper.APIResponse("Success", "Success", ToResponse(res))
		return c.JSON(http.StatusOK, response)
	}
}

func (h *actHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.srv.GetAll()
		if err != nil {
			response := helper.APIResponseNoData("Error", "Error")
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := helper.APIResponse("Success", "Success", ToResponseArr(res))
		return c.JSON(http.StatusOK, response)
	}
}

func (h *actHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := helper.APIResponseNoData("Error", "Error")
			return c.JSON(http.StatusBadRequest, response)
		}

		err = h.srv.Delete(uint(id))
		if err != nil {
			msg := fmt.Sprintf("Activity with ID %d Not Found", id)
			response := helper.APIResponseNoData("Not Found", msg)
			return c.JSON(http.StatusNotFound, response)
		}

		response := helper.APIResponse("Success", "Success", helper.NoData{})
		return c.JSON(http.StatusOK, response)
	}
}

func (h *actHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := helper.APIResponseNoData("Error", "Error")
			return c.JSON(http.StatusNotFound, response)
		}

		input := UpdateRequest{}

		if err := c.Bind(&input); err != nil {
			response := helper.APIResponseNoData("Bad Request", "Bad request")
			return c.JSON(http.StatusBadRequest, response)
		}

		validate := validator.New()
		if err := validate.Struct(input); err != nil {
			response := helper.APIResponseNoData("Bad Request", "title cannot be null")
			return c.JSON(http.StatusBadRequest, response)
		}

		res, err := h.srv.Update(uint(id), *ToCore(input))
		if err != nil {
			msg := fmt.Sprintf("Activity with ID %d Not Found", id)
			response := helper.APIResponseNoData("Not Found", msg)
			return c.JSON(http.StatusNotFound, response)
		}

		response := helper.APIResponse("Success", "Success", ToResponse(res))
		return c.JSON(http.StatusOK, response)
	}
}
