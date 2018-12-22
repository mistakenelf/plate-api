package controllers

import (
	"net/http"

	"github.com/knipferrc/plate-api/db/pg"

	"github.com/labstack/echo"
)

// GetTodoLists returns a users todo lists
func GetTodoLists(c echo.Context) error {
	todoLists := pg.GetTodoLists()
	return c.JSON(http.StatusOK, todoLists)
}
