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

	// Me Routes
	r.GET("/me", controllers.GetUser)

	// Dashboard Routes
	r.GET("/dashboard", controllers.Dashboard)

	// TodoList Routes
	r.GET("/todo-lists", controllers.GetTodoLists)
	r.POST("/todo-lists", controllers.AddTodoList)
	r.DELETE("/todo-lists", controllers.DeleteTodoList)

	// TodoList Details Routes
	r.GET("/todo-lists/:id", controllers.GetTodoList)

	// Todos Routes
	r.DELETE("/todos", controllers.DeleteTodo)
	r.PUT("/todos", controllers.UpdateTodo)
}
