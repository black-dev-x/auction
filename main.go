package main

import (
	"github.com/black-dev-x/auction/database"
	"github.com/black-dev-x/auction/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application...")
	logger.Info("Loading environment variables...")
	godotenv.Load()

	logger.Info("Connecting to MongoDB...")
	_, err := database.DBConnection()
	if err != nil {
		logger.Error("Failed to connect to MongoDB", err)
		return
	}
}
