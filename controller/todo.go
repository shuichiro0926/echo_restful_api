package controller

import (
	"awesomeProject/db"
	"github.com/labstack/echo"
	"net/http"
)

func GetTodoList() echo.HandlerFunc {
	return func(c echo.Context) error {
		todoList := []db.Todo{}
		db := db.GetConnection()
		data := db.Find(&todoList)

		return c.JSON(http.StatusOK, data)
	}
}

func GetTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := db.Todo{}
		db := db.GetConnection()
		data := db.Find(&todo)

		return c.JSON(http.StatusOK, data)
	}
}
