package models

import (
	"github.com/jinzhu/gorm"
)

// User current user object
type User struct {
    gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
