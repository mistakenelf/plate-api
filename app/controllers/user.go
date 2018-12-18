package controllers

import (
	"net/http"

	"github.com/knipferrc/plate-api/db/pg"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GetUser returns the current user
func GetUser(c echo.Context) error {
	jwtToken := c.Get("user").(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	id := claims["ID"].(float64)

	user := pg.FindUser(id)

	return c.JSON(http.StatusOK, user)
}
