package controller

import (
	"awesomeProject/application"
	"github.com/labstack/echo"
	"net/http"
)

type SampleController struct {
	sampleService *application.SampleService
}

func NewController(sampleService *application.SampleService) *SampleController {
	return &SampleController{sampleService: sampleService}
}

func (sampleController SampleController) GetTodoList() echo.HandlerFunc {
	return func(c echo.Context) error {
		result := sampleController.sampleService.GetTodoList()
		return c.JSON(http.StatusOK, result)
	}
}