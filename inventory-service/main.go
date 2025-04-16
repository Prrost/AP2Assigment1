package main

import (
	"inventory-service/Storage"
	"inventory-service/config"
	"inventory-service/grpc/server"
	"inventory-service/useCase"
)

func main() {
	cfg := config.LoadConfig()

	db := Storage.NewSqliteStorage(cfg)

	UseCase := useCase.NewUseCase(db, cfg)

	server.RunGRPCServer(cfg, UseCase)
}
