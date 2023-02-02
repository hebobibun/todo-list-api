package handler

import (
	"fmt"
	"net/http"
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
