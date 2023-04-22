package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbConfig struct {
	Host       string
	Port       int
	Username   string
	Password   string
	Database   string
	Collection string
}

type MongoDbClient struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewMongoDbClient(config MongoDbConfig) (*MongoDbClient, error) {
	connectionUrl = fmt.Sprintf("mongodb://%s:%s@%s:%d",
		config.Username,
		config.Password,
		config.Host,
		config.Port)

	return config, nil
}
