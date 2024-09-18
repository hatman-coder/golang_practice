package database

import (
	"gin_project/models"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare a global variable to hold the database connection
var DB *gorm.DB

// ConnectDB initializes the connection to the PostgreSQL database
func ConnectDB() {
	// Define the database connection string (Data Source Name or DSN)
	dsn := "host=103.209.41.74 user=postgres password=Techsist@#123 dbname=go_db_v1 port=5432 sslmode=disable"

	var err error

	// Open a connection to the PostgreSQL database using GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}

	log.Println("Database migrated successfully!")

	// Dynamically extract the dbname from the DSN
	parts := strings.Split(dsn, " ")
	dbName := ""
	for _, part := range parts {
		if strings.HasPrefix(part, "dbname=") {
			dbName = strings.TrimPrefix(part, "dbname=")
			break
		}
	}

	// Use log.Printf to format and log the message
	log.Printf("Connected to PostgreSQL database: %s;", dbName)
}
