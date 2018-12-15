package handlers

import (
	"net/http"

	"github.com/knipferrc/plate-api/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GetUser returns the current user
func GetUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)

	u := &models.User{
		ID:        id,
		FirstName: "Larkin",
		LastName:  "Jones",
		Email:     "jlarkin@gmail.com",
	}

	return c.JSON(http.StatusOK, u)
}
