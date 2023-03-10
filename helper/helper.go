package helper

import (
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseNoData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type NoData struct {
}

func APIResponse(status string, message string, data interface{}) Response {
	res := Response{
		Message: message,
		Status:  status,
		Data:    data,
	}

	return res
}

func APIResponseNoData(status string, message string) ResponseNoData {
	res := ResponseNoData{
		Message: message,
		Status:  status,
	}

	return res
}

var validate *validator.Validate

func Validation(data interface{}) error {
	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		log.Println(err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		msg := ""
		if strings.Contains(err.Error(), "required") {
			msg = "Body request cannot be blank"
		} else if strings.Contains(err.Error(), "title") {
			msg = "title cannot be null"
		} else if strings.Contains(err.Error(), "email") {
			msg = "email cannot be null"
		}
		return errors.New(msg)
	}
	return nil
}
