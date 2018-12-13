package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/knipferrc/plate-api/models"
	"github.com/labstack/echo"
)

func (h *Handler) AddTodo(c echo.Context) error {
	todo := new(models.Todo)
	c.Bind(&todo)
	fmt.Println(todo)

	collection := h.DBSession.DB("plate").C("todo")
	err := collection.Insert(todo)

	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, todo)
}
