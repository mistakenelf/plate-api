package main

import (
	"log"
	"os"

    "github.com/knipferrc/plate-api/database"
    "github.com/knipferrc/plate-api/handlers"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

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

	// Connect to postgres
	connStr := os.Getenv("DB_CONNECTION_STRING")
	database.DBCon, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Public routes
	e.POST("/api/v1/login", handlers.Login)

	// Restricted routes
	r := e.Group("/api/v1")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/todo-lists", handlers.GetTodoLists)
	r.GET("/todo-lists/:id", handlers.GetTodoList)
	r.GET("/me", handlers.GetUser)

	e.Logger.Fatal(e.Start(":5000"))
}
