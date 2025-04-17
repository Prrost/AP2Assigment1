package clients

import (
	"api-gateway/config"
	"api-gateway/orderpb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func InitOrderClient(ctx context.Context, cfg *config.Config) (orderpb.OrderServiceClient, error) {
	const op = "InitOrderClient"

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	clientConn, err := grpc.NewClient(
		cfg.OrderService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("[%s] gRPC client init error: %s", op, err)
		return nil, err
	}

	OrderClient := orderpb.NewOrderServiceClient(clientConn)
	return OrderClient, nil
}
