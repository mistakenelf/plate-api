package handlers

import (
	"net/http"

	"github.com/knipferrc/plate-api/mocks"
	"github.com/labstack/echo"
)

// TodoListsHandler returns a users todo lists
func TodoListsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, mocks.TodoListsMock)
}
