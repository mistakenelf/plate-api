package controllers

import (
	"net/http"

	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/pg"

	"github.com/labstack/echo"
)

// DeleteTodo removes an todo from the database
func DeleteTodo(c echo.Context) error {
	todo := new(models.Todo)
	c.Bind(&todo)
	pg.DeleteTodo(todo)
	return c.JSON(http.StatusNoContent, models.Todo{})
}

// UpdateTodo updates a todo item
func UpdateTodo(c echo.Context) error {
	todo := new(models.Todo)
	c.Bind(&todo)
	pg.UpdateTodo(todo)
	return c.JSON(http.StatusOK, todo)
}
