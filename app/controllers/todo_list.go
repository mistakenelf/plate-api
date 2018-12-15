package controllers

import (
	"net/http"

	"github.com/knipferrc/plate-api/mocks"
	"github.com/labstack/echo"
)

// GetTodoList returns a users todo list
func GetTodoList(c echo.Context) error {
	return c.JSON(http.StatusOK, mocks.TodoListMock)
}
