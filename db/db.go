package db

import (
	"database/sql"
	"fmt"
	"log"

	"go-tutuplapak-user/config"

	_ "github.com/lib/pq"
)

var DBConnection *sql.DB

func InitDB(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	var err error
	DBConnection, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	if err := DBConnection.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	log.Println("Database connection established successfully")
	return DBConnection, nil
}

func CloseDB() {
	if DBConnection != nil {
		if err := DBConnection.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}
}
