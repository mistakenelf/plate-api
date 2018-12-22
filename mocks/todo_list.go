package mocks

import "github.com/knipferrc/plate-api/app/models"

// TodoListMock mock todolist data
var TodoListMock = models.TodoList{
	ID:          "5",
	Title:       "Title",
	Description: "Description",
	Todos: []models.Todo{
		models.Todo{
			ID:          "1",
			Title:       "Name",
			Description: "Description",
			Completed:   false,
		},
	},
}
