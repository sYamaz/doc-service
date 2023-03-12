package db

import (
	"doc-api/api/gateway/model"
	"doc-api/env"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDBConnection(
	host env.DB_HOST,
	port env.DB_PORT,
	dbname env.DB_NAME,
	user env.DB_USER,
	password env.DB_PASS) *gorm.DB {

	schemaName := "user_schema"

	dsn := strings.Join([]string{
		"host=" + string(host),
		"port=" + string(port),
		"dbname=" + string(dbname),
		"user=" + string(user),
		"password=" + string(password),
		"search_path=" + schemaName,
		"sslmode=disable",
	}, " ")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: schemaName + ".",
		},
	})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(model.User{}, model.Document{}); err != nil {
		panic(err)
	}

	return db
}
