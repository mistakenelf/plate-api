package database

import (
	"github.com/knipferrc/plate-api/models"
)

// CreateTables sets up DB tables
func CreateTables() {
	DBCon.AutoMigrate(&models.User{})
}
