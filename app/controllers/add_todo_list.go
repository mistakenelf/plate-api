package controllers

import (
	"net/http"

	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/pg"
	"github.com/labstack/echo"
)

// AddTodoList adds a new todo list to the db
func AddTodoList(c echo.Context) error {
	todoList := new(models.TodoList)
	c.Bind(&todoList)
	pg.AddTodoList(todoList)
	return c.JSON(http.StatusOK, todoList)
}
