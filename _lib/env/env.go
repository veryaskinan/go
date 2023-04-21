package env

import (
	"github.com/joho/godotenv"
	"main/_lib/logger"
	"os"
	"strconv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		logger.Info("No .env file found")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) int {
	value, _ := strconv.Atoi(os.Getenv(key))
	return value
}
