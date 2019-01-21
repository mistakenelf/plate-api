package controllers

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/pg"

	"github.com/labstack/echo"
)

// CreateTask creates a new task in the db
func CreateTask(c echo.Context) error {
	task := new(models.Task)
	c.Bind(&task)
	pg.CreateTask(task)
	return c.JSON(http.StatusOK, task)
}

// GetTasks returns a users tasks
func GetTasks(c echo.Context) error {
	jwtToken := c.Get("user").(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)
	tasks := pg.GetTasks(id)

	return c.JSON(http.StatusOK, tasks)
}

// GetTaskDetails returns a tasks details
func GetTaskDetails(c echo.Context) error {
	taskID := c.Param("id")
	task := pg.GetTaskDetails(taskID)
	return c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task by id
func DeleteTask(c echo.Context) error {
	taskID := c.Param("id")
	task := pg.DeleteTask(taskID)
	return c.JSON(http.StatusOK, task)
}
