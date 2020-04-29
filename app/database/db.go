package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres
)

// ConnDB represents an instance of the database
var ConnDB *gorm.DB

// Connect opens a connection to the database
func Connect() (*gorm.DB, error) {
	dbConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	return gorm.Open("postgres", dbConfig)
}

// SetConnDatabase sets the actual instance of the database
func SetConnDatabase(db *gorm.DB) {
	ConnDB = db
}
