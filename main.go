package main

import (
	"net/http"

	"github.com/knipferrc/plate-api/handlers"
	"github.com/knipferrc/plate-api/mocks"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowOrigins:     []string{"http://localhost:8080", "https://plate-app.azurewebsites.net"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/api/todo-lists", func(c echo.Context) error {
		return c.JSON(http.StatusOK, mocks.TodoListsMock)
	})

	e.GET("/api/todo-lists/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, mocks.TodoListMock)
	})

	e.POST("/api/login", handlers.LoginHandler)

	e.Logger.Fatal(e.Start(":5000"))
}
