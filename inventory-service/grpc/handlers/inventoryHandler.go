package handlers

import (
	"context"
	inventorypb "github.com/Prrost/assignment1proto/proto/inventory"
	"inventory-service/config"
	"inventory-service/domain"
	"inventory-service/mapping"
	"inventory-service/useCase"
)

type InventoryServer struct {
	inventorypb.UnimplementedInventoryServiceServer
	cfg *config.Config
	uc  *useCase.UseCase
}

func NewInventoryServer(cfg *config.Config, uc *useCase.UseCase) *InventoryServer {
	return &InventoryServer{
		cfg: cfg,
		uc:  uc,
	}
}

func (i *InventoryServer) CreateProduct(ctx context.Context, req *inventorypb.CreateRequest) (*inventorypb.CreateResponse, error) {
	var object domain.Object

	object.Name = req.GetName()
	object.Amount = req.GetAmount()

	outObject, err := i.uc.CreateProduct(object)
	if err != nil {
		return nil, err
	}

	return &inventorypb.CreateResponse{
		Product: mapping.ToProtoProduct(outObject),
		Message: "Object created successfully",
	}, nil
}

func (i *InventoryServer) GetProductById(ctx context.Context, req *inventorypb.GetByIdRequest) (*inventorypb.GetByIdResponse, error) {
	id := req.GetId()

	object, err := i.uc.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return &inventorypb.GetByIdResponse{
		Product: mapping.ToProtoProduct(object),
		Message: "User retrieved successfully",
	}, nil
}

func (i *InventoryServer) GetAllProducts(ctx context.Context, req *inventorypb.GetAllRequest) (*inventorypb.GetAllResponse, error) {
	name := req.GetName()
	limit := req.GetLimit()
	offset := req.GetOffset()

	objects, err := i.uc.GetAllProducts(name, limit, offset)
	if err != nil {
		return nil, err
	}

	var protoObjects []*inventorypb.Product

	for _, object := range objects {
		protoObjects = append(protoObjects, mapping.ToProtoProduct(object))
	}

	return &inventorypb.GetAllResponse{
		Products: protoObjects,
		Message:  "All matching objects retrieved",
	}, nil
}

func (i *InventoryServer) UpdateProduct(ctx context.Context, req *inventorypb.UpdateRequest) (*inventorypb.UpdateResponse, error) {
	inputObject := domain.Object{
		ID:     int(req.GetId()),
		Name:   req.GetName(),
		Amount: req.GetAmount(),
	}

	updatedObject, err := i.uc.UpdateProductById(inputObject)
	if err != nil {
		return nil, err
	}

	return &inventorypb.UpdateResponse{
		Product: mapping.ToProtoProduct(updatedObject),
		Message: "Object updated",
	}, nil
}

func (i *InventoryServer) DeleteProduct(ctx context.Context, req *inventorypb.DeleteRequest) (*inventorypb.DeleteResponse, error) {
	id := req.GetId()

	object, err := i.uc.DeleteProductById(id)
	if err != nil {
		return nil, err
	}

	return &inventorypb.DeleteResponse{
		Product: mapping.ToProtoProduct(object),
		Message: "Deleted object",
	}, nil
}
