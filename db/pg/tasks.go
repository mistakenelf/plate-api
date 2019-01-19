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
