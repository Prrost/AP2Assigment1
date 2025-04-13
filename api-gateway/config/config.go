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
	UserService      string
	JWTSecret        string
	DBPath           string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s" + err.Error())
	}

	return &Config{
		Port:             os.Getenv("PORT"),
		InventoryService: os.Getenv("INVENTORY_SERVICE"),
		OrderService:     os.Getenv("ORDER_SERVICE"),
		UserService:      os.Getenv("USER_SERVICE"),
		JWTSecret:        os.Getenv("JWT_SECRET"),
		DBPath:           os.Getenv("DB_PATH"),
	}
}
