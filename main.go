package main

import (
	"log"

	"github.com/OperationQuasarFire/cmd/config/db"
	server "github.com/OperationQuasarFire/internal/infrastructure/api/starshipCommsResolver"
)

func main() {
	dbInstance, err := db.NewPostgreSQLDB()
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	server.RunServer(dbInstance)
}
