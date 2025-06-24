package database

import (
	"os"

	"github.com/black-dev-x/auction/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE_NAME = "auction"

var database *mongo.Database

func DBConnection() (*mongo.Database, error) {
	if database != nil {
		return database, nil
	}
	url := os.Getenv("MONGODB_URL")
	client, err := mongo.Connect(nil, options.Client().ApplyURI(url))
	if err != nil {
		logger.Error("Failed to connect to MongoDB", err)
		return nil, err
	}
	logger.Info("Connected to MongoDB successfully")
	database = client.Database(DATABASE_NAME)
	return database, nil
}
