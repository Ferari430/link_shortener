package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   Dbconfig
	Auth AuthConfig
}

type AuthConfig struct {
	Secret string
}

type Dbconfig struct {
	DSN string
}

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("Cant read .env file")
	}

	return &Config{Db: Dbconfig{
		DSN: os.Getenv("DSN"),
	},
		Auth: AuthConfig{
			Secret: os.Getenv("Token")}}

}
