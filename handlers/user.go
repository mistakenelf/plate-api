package handlers

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// User current user object
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// GetUser returns the current user
func (h *Handler) GetUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)

	u := &User{
		ID:        id,
		FirstName: "Larkin",
		LastName:  "Jones",
		Email:     "jlarkin@gmail.com",
	}

	return c.JSON(http.StatusOK, u)
}
