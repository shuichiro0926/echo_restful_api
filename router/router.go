package router

import (
	"awesomeProject/controller"
	"github.com/labstack/echo"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init() *echo.Echo {
	e := echo.New()

	api := e.Group("/api")

	{
		api.GET("/swagger/*", echoSwagger.WrapHandler)
		api.GET("/todo", controller.GetTodoList())
		api.GET("/todo/:id", controller.GetTodo())
	}

	return e
}