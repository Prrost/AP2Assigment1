package main

import (
	"api-gateway/Storage/Sqlite"
	"api-gateway/config"
	"api-gateway/gateway/handlers"
)

func main() {
	cfg := config.LoadConfig()

	storage := Sqlite.NewSqliteStorage(cfg)

	r := handlers.SetupRouter(cfg, storage)

	err := r.Run(cfg.Port)
	if err != nil {
		panic(err)
	}
}
