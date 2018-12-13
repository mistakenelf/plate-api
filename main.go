package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/knipferrc/plate-api/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	e := echo.New()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseName := os.Getenv("DB_NAME")
	databaseUsername := os.Getenv("DB_USERNAME")
	databasePassword := os.Getenv("DB_PASSWORD")

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

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s.documents.azure.com:10255", databaseName)},
		Timeout:  60 * time.Second,
		Database: databaseName,
		Username: databaseUsername,
		Password: databasePassword,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	db, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	db.SetMode(mgo.Monotonic, true)

	h := &handlers.Handler{DBSession: db}

	// Public routes
	e.POST("/api/login", handlers.Login)
	e.POST("/api/addTodo", h.AddTodo)

	// Restricted routes
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/todo-lists", handlers.GetTodoLists)
	r.GET("/todo-lists/:id", handlers.GetTodoList)
	r.GET("/me", handlers.GetUser)

	e.Logger.Fatal(e.Start(":5000"))
}
