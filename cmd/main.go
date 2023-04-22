package main

import (
	"log"
	"task-backend/internal/config"
	"task-backend/internal/data/database"
)

func main() {
	cfg, err := config.LoadConfig("configs/configs.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.NewDatabase(cfg.DatabaseUrl, cfg.DatabaseName, cfg.CollectionName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

}
