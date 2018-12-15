package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/knipferrc/plate-api/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	// Load .env file
	err := godotenv.Load()

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
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Add db to handlers
	h := &handlers.Handler{DB: db}

	// Public routes
	e.POST("/api/login", h.Login)

	// Restricted routes
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/todo-lists", h.GetTodoLists)
	r.GET("/todo-lists/:id", h.GetTodoList)
	r.GET("/me", h.GetUser)

	e.Logger.Fatal(e.Start(":5000"))
}
