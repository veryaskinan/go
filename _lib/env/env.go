package env

import (
	"os"
	"strconv"
)

func Get(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) int {
	value, _ := strconv.Atoi(os.Getenv(key))
	return value
}
