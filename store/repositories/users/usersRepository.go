package auth

import (
	"github.com/google/uuid"
	"main/_lib/db"
)

func Create(phone string) {
	userId := uuid.New().String()
	db.Client.Exec("insert into users (id, phone) values (?, ?)", userId, phone)
}
