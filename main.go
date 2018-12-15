package main

import (
	"log"

	"github.com/knipferrc/plate-api/app"
	"github.com/knipferrc/plate-api/db/gorm"

	"github.com/joho/godotenv"
)

func main() {
	app.Init()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	gorm.Init()

	// Start server
	app.Server.Logger.Fatal(app.Server.Start(":5000"))
}
