package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func NewDatabase(databaseURL string, dbName string, collectionName string) (*MongoDbClient, error) {
	clientOptions := options.Client().ApplyURI(databaseURL)
	clientOptions.SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &MongoDbClient{
		Client:     client,
		Database:   database,
		Collection: collection,
	}, nil
}

// func NewMongoDbClient(config MongoDbConfig) (*MongoDbClient, error) {
// 	connectionURI := fmt.Sprintf("mongodb://%s:%s@%s:%d",
// 		config.Username,
// 		config.Password,
// 		config.Host,
// 		config.Port)

// 	clientOptions := options.Client().ApplyURI(connectionURI)
// 	clientOptions.SetConnectTimeout(10 * time.Second)

// 	client, err := mongo.Connect(context.Background(), clientOptions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = client.Ping(context.Background(), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	database := client.Database(config.Database)
// 	collection := database.Collection(config.Collection)

// 	return &MongoDbClient{
// 		Client:     client,
// 		Database:   database,
// 		Collection: collection,
// 	}, nil

// }

func (c *MongoDbClient) Close() {
	err := c.Client.Disconnect(context.Background())
	if err != nil {
		log.Printf("Failed to close MongoDB connection: %v", err)
	}
}
