package main

import (
	"context"
	"fmt"
	inventorypb "github.com/Prrost/assignment1proto/proto/inventory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"order-service/Routes"
	"order-service/Storage"
	"order-service/config"
	"order-service/orderpb"
	"time"
)

func main() {

	cfg := config.LoadConfig()
	db := Storage.NewSqliteStorage(cfg)

	fmt.Println("Server starting on port:", cfg.Port)
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		panic(fmt.Sprintf("Failed to listen on port %s: %v", cfg.Port, err))
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	clientConn, err := grpc.NewClient(
		cfg.InventoryService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("gRPC client init error: %s", err)
	}

	inventoryClient := inventorypb.NewInventoryServiceClient(clientConn)

	grpcServer := grpc.NewServer()
	orderService := &Routes.OrderServiceServer{
		Storage:   db,
		InvClient: inventoryClient,
	}
	orderpb.RegisterOrderServiceServer(grpcServer, orderService)
	fmt.Println("OrderServiceServer registered successfully")

	reflection.Register(grpcServer)
	fmt.Println("Reflection registered successfully")

	if err := grpcServer.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to start gRPC server: %v", err))
	}

}
