package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Инициализация логгера
	logger := log.New(os.Stdout, "", log.LstdFlags)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v", err)
		return
	}
	defer client.Disconnect(context.Background())

	studentRepo := repository.NewStudentRepository()

	studentUsecase := usecase.NewStudentUsecase(studentRepo)

	studentHandler := handler.NewStudentHandler(studentUsecase, logger)

	httpServer := http.NewServer()

	httpServer.RegisterHandler(studentHandler)

	err := httpServer.Start(":8080")
	if err != nil {
		fmt.Printf("Failed to start HTTP server: %v", err)
	}
}
