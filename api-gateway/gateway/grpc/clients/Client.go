package clients

import (
	"api-gateway/config"
	"api-gateway/orderpb"
	"context"
	"fmt"
	inventorypb "github.com/Prrost/assignment1proto/proto/inventory"
	userpb "github.com/Prrost/assignment1proto/proto/user"
	"log"
)

type Client struct {
	UserClient      userpb.UserServiceClient
	InventoryClient inventorypb.InventoryServiceClient
	OrderClient     orderpb.OrderServiceClient
}

func NewMainClient(ctx context.Context, cfg *config.Config) (*Client, []error) {
	const op = "NewClient"
	var errArr []error

	userClient, err := InitUserClient(ctx, cfg)
	if err != nil {
		log.Printf("%s: gRPC user client error: %s", op, err)
		errArr = append(errArr, fmt.Errorf("InitUserClient error: %w", err))
	}

	inventoryClient, err := InitInventoryClient(ctx, cfg)
	if err != nil {
		log.Printf("%s: gRPC inventory client error: %s", op, err)
		errArr = append(errArr, fmt.Errorf("InitInventoryClient error: %w", err))
	}

	orderClient, err := InitOrderClient(ctx, cfg)
	if err != nil {
		log.Printf("%s: gRPC order client error: %s", op, err)
		errArr = append(errArr, fmt.Errorf("InitOrderClient error: %w", err))
	}

	return &Client{
		UserClient:      userClient,
		InventoryClient: inventoryClient,
		OrderClient:     orderClient,
	}, errArr
}
