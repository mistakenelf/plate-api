package controllers

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/pg"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// Login handles user login and returns a JWT
func Login(c echo.Context) error {
	user := new(models.LoginRequest)
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

// Register handles user registration and returns a JWT
func Register(c echo.Context) error {
	user := new(models.RegisterRequest)
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

// GetUser returns the current user
func GetUser(c echo.Context) error {
	jwtToken := c.Get("user").(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)

	user := pg.GetUserByID(id)

	return c.JSON(http.StatusOK, user)
}
