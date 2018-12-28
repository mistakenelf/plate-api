package controllers

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
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

// GetTodoLists returns a users todo lists
func GetTodoLists(c echo.Context) error {
	jwtToken := c.Get("user").(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)

	todoLists := pg.GetTodoLists(id)
	return c.JSON(http.StatusOK, todoLists)
}

// GetTodoList returns a users todo list
func GetTodoList(c echo.Context) error {
	todoListID := c.Param("id")
	todoList := pg.GetTodoList(todoListID)
	return c.JSON(http.StatusOK, todoList)
}

// DeleteTodoList removes an todo from the database
func DeleteTodoList(c echo.Context) error {
	todoList := new(models.TodoList)
	c.Bind(&todoList)
	pg.DeleteTodoList(todoList)
	return c.JSON(http.StatusNoContent, models.TodoList{})
}
