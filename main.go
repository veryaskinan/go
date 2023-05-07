package main

import (
	"main/_lib/db"
	"main/_lib/env"
	"main/_lib/httpServer"
	"main/_lib/redis"
	"main/represent"
)

func init() {
	env.Init()
	db.Init()
	redis.Init()
}

func main() {
	port := env.Get("APP_PORT")
	endpoints := represent.GetEnpoints()
	httpServer.Start(port, endpoints)
}
