package types

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/_lib/logger"
)

type DbClient struct {
	connection *gorm.DB
}

func (dbc DbClient) Init(config DbConfig) {
	dsn := "host=" + config.Host + " " + "port=" + config.Port + " " +
		"user=" + config.User + " " + "password=" + config.Pass + " " +
		"dbname=" + config.DbName + " " + "sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		logger.Info("Error connecting to database")
	}
	dbc.connection = connection
}

func (dbc DbClient) Select(query string, bindings map[string]interface{}) map[string]interface{} {
	var result map[string]interface{}
	dbc.connection.Raw(query, bindings).Scan(&result)
	return result
}

func (dbc DbClient) Exec(query string, bindings map[string]interface{}) {
	dbc.connection.Exec(query, bindings)
}
