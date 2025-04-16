package clients

import (
	"api-gateway/config"
	"context"

	inventorypb "github.com/Prrost/assignment1proto/proto/inventory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func InitInventoryClient(ctx context.Context, cfg *config.Config) (inventorypb.InventoryServiceClient, error) {
	const op = "InitInventoryClient"

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	clientConn, err := grpc.NewClient(
		cfg.InventoryService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("[%s] gRPC client init error: %s", op, err)
		return nil, err
	}

	inventoryClient := inventorypb.NewInventoryServiceClient(clientConn)
	return inventoryClient, nil
}
