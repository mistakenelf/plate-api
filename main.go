package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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
				ID:          "2",
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

type userLoginData struct {
	Email    string
	Password string
}

func login(c echo.Context) error {
	user := new(userLoginData)
	c.Bind(&user)
	if user.Email == "test@email.com" && user.Password == "password" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Boss Man"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

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
		return c.JSON(http.StatusOK, todoLists)
	})

	e.GET("/api/todo-lists/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todoList)
	})

	e.POST("/api/login", login)

	e.Logger.Fatal(e.Start(":5000"))
}
