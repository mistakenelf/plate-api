package gorm

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	// Needed in here for abstraction
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

// Init initialize database
func Init() {
	var err error

	connStr := os.Getenv("DB_CONNECTION_STRING")
	db, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

// DBCon export DB
func DBCon() *gorm.DB {
	return db
}
