package server

import (
	inventorypb "github.com/Prrost/assignment1proto/proto/inventory"
	"google.golang.org/grpc"
	"inventory-service/config"
	"inventory-service/grpc/handlers"
	"inventory-service/grpc/middleware"
	"inventory-service/useCase"
	"log"
	"net"
)

func RunGRPCServer(cfg *config.Config, uc *useCase.UseCase) *grpc.Server {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.LoggingInterceptor))

	inventoryServer := handlers.NewInventoryServer(cfg, uc)

	inventorypb.RegisterInventoryServiceServer(grpcServer, inventoryServer)

	log.Printf("starting gRPC server on port %s", cfg.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return grpcServer
}
