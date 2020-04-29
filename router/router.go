package router

import (
	"awesomeProject/controller"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	api := e.Group("/api")

	{
		api.GET("/todo", controller.GetTodoList())
		api.GET("/todo/:id", controller.GetTodo())
	}

	return e
}