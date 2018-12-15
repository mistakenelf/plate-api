package env

import (
	"log"

	"github.com/joho/godotenv"
)

// Init initialize env vars
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
