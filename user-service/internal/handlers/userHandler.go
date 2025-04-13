package handlers

import (
	"context"
	userpb "github.com/Prrost/assignment1proto/proto/user"
	"strconv"
	"user-service/config"
	"user-service/domain"
	"user-service/useCase"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
	cfg *config.Config
	uc  *useCase.UseCase
}

func NewUserServer(cfg *config.Config, uc *useCase.UseCase) *UserServer {
	return &UserServer{
		cfg: cfg,
		uc:  uc,
	}
}

func (s *UserServer) RegisterUser(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	var user domain.User

	user.Email = req.GetEmail()
	user.Password = req.GetPassword()

	userOut, err := s.uc.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &userpb.RegisterResponse{
		Id:      int64(userOut.ID),
		Message: "User created successfully",
	}, nil
}

func (s *UserServer) AuthenticateUser(ctx context.Context, req *userpb.AuthRequest) (*userpb.AuthResponse, error) {
	var user domain.User

	user.Email = req.GetEmail()
	user.Password = req.GetPassword()

	token, err := s.uc.LoginUser(user)
	if err != nil {
		return nil, err
	}

	return &userpb.AuthResponse{
		Token:   token,
		Message: "Authentication successful",
	}, nil
}

func (s *UserServer) GetUserProfile(ctx context.Context, req *userpb.UserID) (*userpb.UserProfile, error) {
	id := req.GetId()

	user, err := s.uc.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	idString := strconv.Itoa(int(user.ID))

	return &userpb.UserProfile{
		Id:    idString,
		Email: user.Email,
	}, nil
}
