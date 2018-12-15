package controllers

import (
	"net/http"
	"time"

	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/pg"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Register handles user registration and returns a JWT
func Register(c echo.Context) error {
	user := new(models.User)
	c.Bind(&user)

	pg.AddUser(user)

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = "123lkj234324"
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
