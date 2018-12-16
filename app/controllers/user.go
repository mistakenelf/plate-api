package controllers

import (
	"net/http"

	"github.com/knipferrc/plate-api/app/models"

	"github.com/labstack/echo"
)

// GetUser returns the current user
func GetUser(c echo.Context) error {
	u := &models.User{
		FirstName: "Larkin",
		LastName:  "Jones",
		Email:     "jlarkin@gmail.com",
	}

	return c.JSON(http.StatusOK, u)
}
