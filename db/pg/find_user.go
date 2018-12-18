package pg

import (
	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/gorm"
)

// FindUser finds a user based on id
func FindUser(id float64) models.User {
	var user models.User
	gorm.DBCon().Where("ID = ?", id).First(&user)
	return user
}
