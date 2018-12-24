package controllers

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/knipferrc/plate-api/db/pg"

	"github.com/labstack/echo"
)

// GetTodoLists returns a users todo lists
func GetTodoLists(c echo.Context) error {
	jwtToken := c.Get("user").(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)

	todoLists := pg.GetTodoLists(id)
	return c.JSON(http.StatusOK, todoLists)
}
