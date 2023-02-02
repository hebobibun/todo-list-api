package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todo-api/helper"
	"todo-api/todo"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type todoHandler struct {
	srv todo.TodoService
}

func New(srv todo.TodoService) todo.TodoHandler {
	return &todoHandler{
		srv: srv,
	}
}

func (h *todoHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := TodoRequest{}
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
			} else if strings.Contains(err.Error(), "Priority") {
				msg = "priority cannot be null"
			} else if strings.Contains(err.Error(), "Activity") {
				msg = "activity_group_id cannot be null"
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

func (h *todoHandler) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := helper.APIResponseNoData("Error", "Error")
			return c.JSON(http.StatusNotFound, response)
		}

		res, err := h.srv.GetOne(uint(id))
		if err != nil {
			msg := fmt.Sprintf("Todo with ID %d Not Found", id)
			response := helper.APIResponseNoData("Not Found", msg)
			return c.JSON(http.StatusNotFound, response)
		}

		response := helper.APIResponse("Success", "Success", ToResponse(res))
		return c.JSON(http.StatusOK, response)
	}
}

func (h *todoHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		actGroupID := c.QueryParam("activity_group_id")

		actID := 0

		if actGroupID != "" {
			id, err := strconv.Atoi(actGroupID)
			if err != nil {
				response := helper.APIResponseNoData("Bad Request", "Bad Request")
				return c.JSON(http.StatusBadRequest, response)
			}

			actID = id
		}

		res, err := h.srv.GetAll(uint(actID))
		if err != nil {
			response := helper.APIResponseNoData("Error", "Error")
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := helper.APIResponse("Success", "Success", ToResponseArr(res))
		return c.JSON(http.StatusOK, response)
	}
}

func (h *todoHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := helper.APIResponseNoData("Error", "Error")
			return c.JSON(http.StatusBadRequest, response)
		}

		input := TodoUpdateRequest{}

		if err := c.Bind(&input); err != nil {
			response := helper.APIResponseNoData("Bad Request", "Bad request")
			return c.JSON(http.StatusBadRequest, response)
		}

		res, err := h.srv.Update(uint(id), *ToCore(input))
		if err != nil {
			msg := fmt.Sprintf("Todo with ID %d Not Found", id)
			response := helper.APIResponseNoData("Not Found", msg)
			return c.JSON(http.StatusNotFound, response)
		}

		response := helper.APIResponse("Success", "Success", ToResponse(res))
		return c.JSON(http.StatusOK, response)
	}
}

func (h *todoHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := helper.APIResponseNoData("Error", "Error")
			return c.JSON(http.StatusBadRequest, response)
		}

		err = h.srv.Delete(uint(id))
		if err != nil {
			msg := fmt.Sprintf("Todo with ID %d Not Found", id)
			response := helper.APIResponseNoData("Not Found", msg)
			return c.JSON(http.StatusNotFound, response)
		}

		response := helper.APIResponse("Success", "Success", helper.NoData{})
		return c.JSON(http.StatusOK, response)
	}
}
