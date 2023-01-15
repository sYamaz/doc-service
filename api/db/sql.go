package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection() *gorm.DB {
	dsn := "host=localhost user=app password=postgres dbname=postgres port=5432 sslmode=disable search_path=public"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
