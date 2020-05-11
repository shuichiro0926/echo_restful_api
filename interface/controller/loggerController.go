package controller

import (
	"awesomeProject/application"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type LoggerController struct {
	loggerService application.LoggerService
}

func NewLoggerController(loggerService application.LoggerService) LoggerController {
	return LoggerController{loggerService: loggerService}
}

func (l LoggerController) SearchDeployHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var status, _ = strconv.ParseBool(c.QueryParam("deployStatus"))
		var limit, _ = strconv.Atoi(c.QueryParam("getCount"))

		result := l.loggerService.SearchDeployHistory(
			c.QueryParam("deploySettingType"), c.QueryParam("deployTarget"), status, limit,
		)
		return c.JSON(http.StatusOK, result)
	}
}
