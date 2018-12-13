package main

import (
	"github.com/knipferrc/plate-api/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowOrigins:     []string{"http://localhost:8080", "https://plate-app.azurewebsites.net"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Public routes
	e.POST("/api/login", handlers.Login)

	// Restricted routes
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/todo-lists", handlers.GetTodoLists)
	r.GET("/todo-lists/:id", handlers.GetTodoList)
	r.GET("/me", handlers.GetUser)

	e.Logger.Fatal(e.Start(":5000"))
}
