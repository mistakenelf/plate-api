package pg

import (
	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/gorm"
)

// CreateUser creates a new user to the database
func CreateUser(user *models.User) {
	if err := gorm.DBCon().Create(&user).Error; err != nil {
		panic(err)
	}
}
