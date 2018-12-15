package mocks

import "github.com/knipferrc/plate-api/app/models"

// TodoListsMock mock todo lists data
var TodoListsMock = []models.TodoList{
	models.TodoList{
		ID:          "1",
		Title:       "Title",
		Description: "Description",
		Todos: []models.Todo{
			models.Todo{
				ID:          "1",
				Name:        "Name",
				Description: "Description",
				Completed:   false,
			},
		},
	},
	models.TodoList{
		ID:          "2",
		Title:       "Title",
		Description: "Description",
		Todos: []models.Todo{
			models.Todo{
				ID:          "2",
				Name:        "Name",
				Description: "Description",
				Completed:   false,
			},
		},
	},
	models.TodoList{
		ID:          "3",
		Title:       "Title",
		Description: "Description",
		Todos: []models.Todo{
			models.Todo{
				ID:          "1",
				Name:        "Name",
				Description: "Description",
				Completed:   false,
			},
		},
	},
	models.TodoList{
		ID:          "4",
		Title:       "Title",
		Description: "Description",
		Todos: []models.Todo{
			models.Todo{
				ID:          "1",
				Name:        "Name",
				Description: "Description",
				Completed:   false,
			},
		},
	},
	models.TodoList{
		ID:          "5",
		Title:       "Title",
		Description: "Description",
		Todos: []models.Todo{
			models.Todo{
				ID:          "1",
				Name:        "Name",
				Description: "Description",
				Completed:   false,
			},
		},
	},
}
