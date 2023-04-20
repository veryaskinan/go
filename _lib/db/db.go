package db

import (
	"main/_lib/db/types"
	"main/_lib/env"
)

var Client types.DbClient

func Init() {
	Client.Init(types.DbConfig{
		Host:   env.Get("DATABASE_HOST"),
		Port:   env.Get("DATABASE_PORT"),
		User:   env.Get("DATABASE_USER"),
		Pass:   env.Get("DATABASE_PASS"),
		DbName: env.Get("DATABASE_NAME"),
	})
}
