package controllers

import (
	"net/http"
	"time"

	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/pg"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// LoginRequest login input
type LoginRequest struct {
	Email    string
	Password string
}

// Login handles user login and returns a JWT
func Login(c echo.Context) error {
	user := new(LoginRequest)
	c.Bind(&user)

	foundUser := pg.GetUserByEmail(user.Email)

	passwordMatch := bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(user.Password))

	if (foundUser != (models.User{})) && passwordMatch == nil {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["ID"] = foundUser.ID
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
