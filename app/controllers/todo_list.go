package controllers

import (
	"net/http"

	"github.com/knipferrc/plate-api/db/pg"

	"github.com/labstack/echo"
)

// GetTodoList returns a users todo list
func GetTodoList(c echo.Context) error {
	todoListID := c.Param("id")
	todoList := pg.GetTodoList(todoListID)
	return c.JSON(http.StatusOK, todoList)
}
