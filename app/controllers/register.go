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

// RegisterRequest regisger input
type RegisterRequest struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

// Register handles user registration and returns a JWT
func Register(c echo.Context) error {
	user := new(RegisterRequest)
	c.Bind(&user)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &models.User{
		Email:        user.Email,
		PasswordHash: string(passwordHash),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
	}

	pg.CreateUser(newUser)

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = newUser.ID
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
