package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/Nasa28/hotel-room-reservation/config"
)

func main() {
	// Define the action flag with a default value of "up"
	action := flag.String("action", "", "Define migration action: 'up' or 'down'")
	flag.Parse()

	conf := config.Env

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		conf.User, conf.Password, conf.DB_Host, conf.DBPort, conf.DBName, conf.SSLMode)

	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		dsn,
	)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	// Handle the action based on the flag value
	switch *action {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Error running migration up: %v", err)
		}
		fmt.Println("Migrations applied successfully!")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Error running migration down: %v", err)
		}
		fmt.Println("Migrations reverted successfully!")
	default:
		fmt.Println("Invalid action. Use 'up' or 'down'.")
	}
}
