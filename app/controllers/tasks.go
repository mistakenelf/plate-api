package controllers

import (
	"net/http"

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
