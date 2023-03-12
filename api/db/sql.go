package db

import (
	"doc-api/api/gateway/model"
	"doc-api/env"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(
	host env.DB_HOST,
	port env.DB_PORT,
	dbname env.DB_NAME,
	user env.DB_USER,
	password env.DB_PASS,
	schema env.DB_SCHEMA) *gorm.DB {
	dsn := strings.Join([]string{
		"host=" + string(host),
		"port=" + string(port),
		"dbname=" + string(dbname),
		"user=" + string(user),
		"password=" + string(password),
		"search_path=" + string(schema),
		"sslmode=disable",
	}, " ")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(model.User{}, model.Document{}); err != nil {
		panic(err)
	}

	return db
}
