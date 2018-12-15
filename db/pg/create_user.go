package pg

import (
	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/gorm"
)

// AddUser adds a new user to the database
func AddUser(user *models.User) {
	gorm.DBCon().Create(&user)
}
