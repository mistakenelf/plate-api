package controllers

import (
	"net/http"

	"github.com/knipferrc/plate-api/mocks"
	"github.com/labstack/echo"
)

// GetTodoLists returns a users todo lists
func GetTodoLists(c echo.Context) error {
	return c.JSON(http.StatusOK, mocks.TodoListsMock)
}
