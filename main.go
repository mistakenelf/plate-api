package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var todoLists = []TodoList{
	TodoList{
		ID:          "1",
		Title:       "Title",
		Description: "Description",
		Todos: []Todo{
			Todo{
				ID:          "1",
				Name:        "Name",
				Description: "Description",
				Completed:   false,
			},
		},
	},
	TodoList{
		ID:          "2",
		Title:       "Title",
		Description: "Description",
		Todos: []Todo{
			Todo{
				ID:          "1",
				Name:        "Name",
				Description: "Description",
				Completed:   false,
			},
		},
	},
	TodoList{
		ID:          "3",
		Title:       "Title",
		Description: "Description",
		Todos: []Todo{
			Todo{
				ID:          "1",
				Name:        "Name",
				Description: "Description",
				Completed:   false,
			},
		},
	},
	TodoList{
		ID:          "4",
		Title:       "Title",
		Description: "Description",
		Todos: []Todo{
			Todo{
				ID:          "1",
				Name:        "Name",
				Description: "Description",
				Completed:   false,
			},
		},
	},
	TodoList{
		ID:          "5",
		Title:       "Title",
		Description: "Description",
		Todos: []Todo{
			Todo{
				ID:          "1",
				Name:        "Name",
				Description: "Description",
				Completed:   false,
			},
		},
	},
}

var todoList = TodoList{
	ID:          "5",
	Title:       "Title",
	Description: "Description",
	Todos: []Todo{
		Todo{
			ID:          "1",
			Name:        "Name",
			Description: "Description",
			Completed:   false,
		},
	},
}

var token = Token{
	Token: "8LSKD8KJSHDFJKH",
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "https://plate-app.azurewebsites.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/api/todo-lists", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todoLists)
	})

	e.GET("/api/todo-lists/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todoList)
	})

	e.POST("/api/login", func(c echo.Context) error {
		return c.JSON(http.StatusOK, token)
	})

	e.Logger.Fatal(e.Start(":5000"))
}
