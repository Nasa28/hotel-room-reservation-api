package db

import (
	"fmt"
	"log"

	"github.com/Nasa28/hotel-room-reservation/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func MyNewSQLDB(conf config.Config) (*sqlx.DB, error) {
	// Format the connection string
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.Host, conf.DBPort, conf.User, conf.Password, conf.DBName, conf.SSLMode,
	)

	// Open the database connection
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Ping the database to ensure connection is established
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	log.Println("Successfully connected to the database!")
	return db, nil
}
