package main

import (
	"log"
	"my_project/internal/link"
	"my_project/internal/stat"
	"my_project/internal/user"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	defer log.Println("Migrations succses")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cant migrate data in DB")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic("Cant connect to DB in auto")
	}
	log.Println("Database connected")

	db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})

}
