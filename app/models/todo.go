package models

import (
	"github.com/jinzhu/gorm"
)

// Todo an todo item
type Todo struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(255)"`
	Completed   bool   `gorm:"type:boolean"`
}
