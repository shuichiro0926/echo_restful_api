package controller

import (
	"awesomeProject/db"
	"github.com/labstack/echo"
	"net/http"
)

// getUsers is getting users.
// @Summary get users
// @Description get users in a group
// @Accept  json
// @Produce  json
// @Param group_id path int true "Group ID"
// @Param gender query string false "Gender" Enum(man, woman)
// @Router /groups/{group_id}/users [get]
func GetTodoList() echo.HandlerFunc {
	return func(c echo.Context) error {
		todoList := []db.Todo{}
		db := db.GetConnection()
		data := db.Find(&todoList)

		return c.JSON(http.StatusOK, data)
	}
}

// getUsers is getting users.
// @Summary get users
// @Description get users in a group
// @Accept  json
// @Produce  json
// @Param group_id path int true "Group ID"
// @Param gender query string false "Gender" Enum(man, woman)
// @Router /groups/{group_id}/users [get]
func GetTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := db.Todo{}
		db := db.GetConnection()
		data := db.Find(&todo)

		return c.JSON(http.StatusOK, data)
	}
}
