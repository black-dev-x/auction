package main

import (
	"github.com/black-dev-x/auction/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application...")
	godotenv.Load()
}
