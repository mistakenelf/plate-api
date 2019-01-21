package pg

import (
	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/gorm"
)

// CreateTask puts a new task in the DB
func CreateTask(task *models.Task) {
	if err := gorm.DBCon().Create(&task).Error; err != nil {
		panic(err)
	}
}

// GetTasks returns a users tasks
func GetTasks(id string) []models.Task {
	var tasks []models.Task
	res := gorm.DBCon().Where("created_by = ?", id).Find(&tasks, &models.Task{})
	if res.RecordNotFound() {
		panic(nil)
	}
	return tasks
}

// GetTaskDetails returns a tasks deatils from the Db
func GetTaskDetails(id string) models.Task {
	var task models.Task
	res := gorm.DBCon().Where("ID = ?", id).First(&task)
	if res.RecordNotFound() {
		panic(nil)
	}
	return task
}

// DeleteTask removes a task from the DB
func DeleteTask(id string) models.Task {
	var task models.Task
	res := gorm.DBCon().Where("ID = ?", id).Delete(&task)
	if res.RecordNotFound() {
		panic(nil)
	}
	return task
}
