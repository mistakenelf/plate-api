package pg

import (
	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/gorm"
)

// DeleteTodo removes a todo item from a list
func DeleteTodo(todo *models.Todo) {
	gorm.DBCon().Delete(&todo)
}

// UpdateTodo updates a todo item
func UpdateTodo(todo *models.Todo) {
	gorm.DBCon().Save(&todo)
}
