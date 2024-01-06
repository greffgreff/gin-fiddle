package databases

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var DB *gorm.DB

func ConnectToDB() {
	connectionString := os.Getenv("DB_DSN")
	if len(connectionString) == 0 {
		log.Fatal("Could not fetch `DB_DSN` in environment variables.")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error occurred while connecting to the employee database: " + err.Error())
	}
}
