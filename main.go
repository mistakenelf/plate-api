package main

import (
	"github.com/knipferrc/plate-api/app"
	"github.com/knipferrc/plate-api/db/gorm"
	"github.com/knipferrc/plate-api/env"
)

func main() {
	app.Init()
	env.Init()
	gorm.Init()

	// Start server
	app.Server.Logger.Fatal(app.Server.Start(":5000"))
}
