package router

import (
	"awesomeProject/infrastructure/injector"
	"github.com/labstack/echo"
)

func LoggerRouter() *echo.Echo {
	e := echo.New()

	loggerApi := e.Group("/api/logger/search")
	{
		loggerApi.GET("/deploy",
			injector.InjectController().SearchDeployHistory(),
		)
	}

	return e
}