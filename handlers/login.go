package handlers

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// UserLoginInput form input from login
type UserLoginInput struct {
	Email    string
	Password string
}

// Login handles user login and returns a JWT
func Login(c echo.Context) error {
	user := new(UserLoginInput)
	c.Bind(&user)

	if user.Email == "test@email.com" && user.Password == "password" {
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
	return echo.ErrUnauthorized
}
