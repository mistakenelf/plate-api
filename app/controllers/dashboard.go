package controllers

import (
	"net/http"

	"github.com/knipferrc/plate-api/app/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/knipferrc/plate-api/db/pg"
	"github.com/labstack/echo"
)

// Dashboard returns data for dashboard
func Dashboard(c echo.Context) error {
	jwtToken := c.Get("user").(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)

	counts := pg.GetTotals(id)
	todoLists := pg.GetDashboardTodoLists(id)

	dashboard := models.Dashboard{
		DashboardCounts: counts,
		TodoLists:       todoLists,
	}

	return c.JSON(http.StatusOK, dashboard)
}
