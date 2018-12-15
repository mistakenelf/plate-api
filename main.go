package main

import (
	"log"

	"github.com/knipferrc/plate-api/app"
	"github.com/knipferrc/plate-api/db/gorm"

	"github.com/joho/godotenv"
)

func main() {
	// Initialize app
	app.Init()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize database
	gorm.Init()

	// Start server
	app.Server.Logger.Fatal(app.Server.Start(":5000"))
}
