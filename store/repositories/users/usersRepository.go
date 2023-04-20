package auth

import "main/_lib/db"

func Create(phone string) {
	db.Client.Exec("insert into users (phone) values (?)", map[string]interface{}{phone: phone})
}
