package app

import (
	"github.com/knipferrc/plate-api/app/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server main app server
var Server *echo.Echo

// Init starts the server
func Init() {
	Server = echo.New()

	// Middleware
	Server.Use(middleware.Secure())
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())
	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowOrigins:     []string{"http://localhost:8080", "https://plate-app.azurewebsites.net"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Public routes
	Server.POST("/api/v1/login", controllers.Login)
	Server.POST("/api/v1/register", controllers.Register)

	// Restricted routes
	r := Server.Group("/api/v1")
	r.Use(middleware.JWT([]byte("secret")))
	r.POST("/todo-lists/add", controllers.AddTodoList)
	r.GET("/todo-lists", controllers.GetTodoLists)
	r.GET("/todo-lists/:id", controllers.GetTodoList)
	r.GET("/me", controllers.GetUser)
	r.GET("/dashboard", controllers.Dashboard)
}
