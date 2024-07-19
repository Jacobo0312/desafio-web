package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

func Load() (*ConfigAppDefault, error) {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}, nil

}
