package controllers

import (
	"net/http"

	"github.com/knipferrc/plate-api/db/gorm"
	"github.com/labstack/echo"
)

type TodoItem struct {
	Name        string
	Description string
	Completed   bool
}

func AddTodo(c echo.Context) error {
	todo := new(TodoItem)
	c.Bind(&todo)

	if err := gorm.DBCon().Create(&todo).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, todo)
}
