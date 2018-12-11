package handlers

import (
	"net/http"

	"github.com/knipferrc/plate-api/mocks"
	"github.com/labstack/echo"
)

// TodoListHandler returns a users todo list
func TodoListHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, mocks.TodoListMock)
}
