package clients

import (
	"api-gateway/config"
	"context"
	userpb "github.com/Prrost/assignment1proto/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func InitUserClient(ctx context.Context, cfg *config.Config) (userpb.UserServiceClient, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	clientConn, err := grpc.NewClient(
		cfg.UserService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	userClient := userpb.NewUserServiceClient(clientConn)
	return userClient, nil
}
