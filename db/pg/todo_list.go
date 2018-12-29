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
func GetTodoLists(id string) []models.TodoList {
	var todoLists []models.TodoList
	res := gorm.DBCon().Where("created_by = ?", id).Preload("Todos").Find(&todoLists, &models.TodoList{})
	if res.RecordNotFound() {
		panic(nil)
	}
	return todoLists
}

// GetTodoList returns a single todo list
func GetTodoList(id string) models.TodoList {
	var todoList models.TodoList
	res := gorm.DBCon().Where("ID = ?", id).Preload("Todos").First(&todoList)
	if res.RecordNotFound() {
		panic(nil)
	}
	return todoList
}

// DeleteTodoList removes a todo list
func DeleteTodoList(todoList *models.TodoList) {
	gorm.DBCon().Where("todo_list_id = ?", todoList.ID).Delete(&models.Todo{})
	gorm.DBCon().Delete(&todoList)
}

// UpdateTodoList updates a todo list
func UpdateTodoList(todoList *models.TodoList) {
	gorm.DBCon().Save(&todoList)
}
