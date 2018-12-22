package pg

import (
	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/gorm"
)

// AddTodoList adds a new todo list to the db
func AddTodoList(todoList *models.TodoList) {
	if err := gorm.DBCon().Create(&todoList).Error; err != nil {
		panic(err)
	}
}

// GetTodoLists returns users todo lists
func GetTodoLists() []models.TodoList {
	var todoLists []models.TodoList
	res := gorm.DBCon().Limit(5).Preload("Todos").Find(&todoLists, &models.TodoList{})
	if res.RecordNotFound() {
		panic(nil)
	}
	return todoLists
}
