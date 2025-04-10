package main

import (
	"order-service/Storage"
	"order-service/config"
	"order-service/handlersOrder"
)

func main() {

	cfg := config.LoadConfig()

	db := Storage.NewSqliteStorage(cfg)

	router := handlersOrder.SetupRouter(cfg, db)

	err := router.Run(cfg.Port)
	if err != nil {
		panic(err)
	}

}
