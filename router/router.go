package router

import (
	"awesomeProject/adpter/controller"
	"awesomeProject/adpter/infrastructure"
	"awesomeProject/application"
	"awesomeProject/db"
	"github.com/labstack/echo"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init() *echo.Echo {
	e := echo.New()

	api := e.Group("/api")

	repositoryImpl := infrastructure.NewRepositoryImpl(db.GetConnection())
	sampleService := application.NewService(repositoryImpl)
	sampleController := controller.NewController(sampleService)

	{
		api.GET("/swagger/*", echoSwagger.WrapHandler)
		api.GET("/todo", sampleController.GetTodoList())
	}

	return e
}