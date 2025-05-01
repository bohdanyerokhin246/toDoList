package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Psql struct {
	DB *sql.DB
}

func (p *Psql) Connect() *sql.DB {

	var err error
	psql := new(Psql)
	// Load the .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"))

	psql.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
	}

	if err = psql.DB.Ping(); err != nil {
		fmt.Printf("Pinging is failed. Error: %v\n", err)
	}

	fmt.Println("Psql connected")
	return psql.DB
}
