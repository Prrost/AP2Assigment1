package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port             string
	InventoryService string
	OrderService     string
	JWTSecret        string
	DBPath           string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s" + err.Error())
	}

	if os.Getenv("ENV") == "doc" {
		log.Println("Running in docker")
		return &Config{
			Port:             os.Getenv("PORT"),
			InventoryService: os.Getenv("DOCKER_INVENTORY_SERVICE"),
			DBPath:           os.Getenv("DB_PATH"),
		}
	} else {
		return &Config{
			Port:             os.Getenv("PORT"),
			InventoryService: os.Getenv("INVENTORY_SERVICE"),
			DBPath:           os.Getenv("DB_PATH"),
		}
	}

}
