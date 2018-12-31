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

// GetUserByEmail finds a user by email
func GetUserByEmail(email string) models.User {
	var user models.User
	res := gorm.DBCon().Find(&user, &models.User{Email: email})
	if res.RecordNotFound() {
		panic(nil)
	}
	return user
}

// GetUserByID finds a user based on id
func GetUserByID(id string) models.User {
	var user models.User
	if err := gorm.DBCon().Where("ID = ?", id).First(&user).Error; err != nil {
		panic(err)
	}
	return user
}
