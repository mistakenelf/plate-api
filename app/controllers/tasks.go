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
