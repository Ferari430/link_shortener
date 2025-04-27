package db

import (
	"log"
	"my_project/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(config.Db.DSN))

	if err != nil {
		panic("Cant connect to DB")
	}
	log.Println("Database connected")
	return &Db{db}
}
