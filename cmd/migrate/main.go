package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/Nasa28/hotel-room-reservation/config"
)

func main() {
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

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error running migration: %v", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error running migration: %v", err)
	}
	fmt.Println("Migrations ran successfully!")
}
