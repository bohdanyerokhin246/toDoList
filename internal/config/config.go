package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	var err error

	// Load the .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST_APP"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
	}

	if err = db.Ping(); err != nil {
		fmt.Printf("Pinging is failed. Error: %v\n", err)
	} else {
		fmt.Println("DB connected")
	}
	return db
}
