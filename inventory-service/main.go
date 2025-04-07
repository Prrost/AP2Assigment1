package main

import (
	"inventory-service/Storage"
	"inventory-service/config"
	"inventory-service/handlersInv"
)

func main() {
	cfg := config.LoadConfig()

	db := Storage.NewSqliteStorage(cfg)

	router := handlersInv.SetupRouter(cfg, db)

	err := router.Run(cfg.Port)
	if err != nil {
		panic(err)
	}
}
