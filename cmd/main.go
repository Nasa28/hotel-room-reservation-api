package main

import (
	"log"

	"github.com/Nasa28/hotel-room-reservation/cmd/api"
	"github.com/Nasa28/hotel-room-reservation/config"
	"github.com/Nasa28/hotel-room-reservation/db"
)

func main() {
	// Connect to the database
	dbConn, err := db.MyNewSQLDB(config.Env)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer dbConn.Close()

	// Create a new API server
	server := api.NewAPIServer(config.Env.Port, dbConn)

	// Run the server
	if err := server.Run(); err != nil {
		log.Fatal("Failed to run the server:", err)
	}
}
