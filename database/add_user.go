package database

import (
	"github.com/knipferrc/plate-api/models"
)

// AddUser adds a new user to the database
func AddUser(user *models.User) {
	DBCon.Create(&user)
}
