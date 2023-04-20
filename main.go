package main

import (
	"github.com/joho/godotenv"
	"main/_lib/db"
	"main/_lib/httpServer"
	"main/_lib/logger"
	"main/controllers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logger.Info("No .env file found")
	}
	db.Init()
}

func main() {
	httpServer.Start(controllers.GetEnpoints())
}
