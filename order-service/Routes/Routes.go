package Routes

import (
	"context"
	"fmt"
	inventorypb "github.com/Prrost/assignment1proto/proto/inventory"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"order-service/Storage"
	"order-service/domain"
	"order-service/orderpb"
)

type OrderServiceServer struct {
	orderpb.UnimplementedOrderServiceServer
	InvClient inventorypb.InventoryServiceClient
	Storage   Storage.Storage
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	if req.Amount < 1 {
		return nil, fmt.Errorf("amount can't be less than 1")
	}

	order := domain.Order{
		ProductID: int(req.ProductId),
		UserID:    int(req.UserId),
		Amount:    req.Amount,
		Status:    "Created",
	}

	availableAmount, err := CheckProductAmount(int(req.ProductId), s.InvClient)
	if err != nil {
		return nil, status.Error(codes.Internal, "error getting product amount")
	}
	log.Println("1:", availableAmount, "2:", order.Amount)

	if availableAmount < order.Amount {
		return nil, status.Error(codes.InvalidArgument, "not enough amount in inventory")
	}

	createdOrder, err := s.Storage.CreateOrderX(order)
	if err != nil {
		return nil, status.Error(codes.Internal, "error creating order")
	}

	err = ProductChange(order.Amount, order.ProductID, availableAmount, s.InvClient)
	if err != nil {
		return nil, err
	}

	return &orderpb.CreateOrderResponse{
		Message: "Order created successfully",
		Order: &orderpb.Order{
			OrderId:   int32(createdOrder.ID),
			UserId:    int32(createdOrder.UserID),
			ProductId: int32(createdOrder.ProductID),
			Amount:    createdOrder.Amount,
			Status:    createdOrder.Status,
		},
	}, nil
}

func (s *OrderServiceServer) GetOrder(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.GetOrderResponse, error) {
	order, err := s.Storage.GetOrderByIDX(int(req.OrderId))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &orderpb.GetOrderResponse{
		Message: "Order retrieved successfully",
		Order: &orderpb.Order{
			OrderId:   int32(order.ID),
			UserId:    int32(order.UserID),
			ProductId: int32(order.ProductID),
			Amount:    int64(order.Amount),
			Status:    order.Status,
		},
	}, nil
}

func (s *OrderServiceServer) UpdateOrder(ctx context.Context, req *orderpb.UpdateOrderRequest) (*orderpb.UpdateOrderResponse, error) {
	// Проверяем, существует ли заказ
	order, err := s.Storage.GetOrderByIDX(int(req.OrderId))
	if err != nil {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	tempAmount := order.Amount

	// Проверяем, что количество товара >= 1
	if req.Amount < 1 {
		return nil, status.Error(codes.InvalidArgument, "amount can't be less than 1")
	}

	// Получаем текущий объем доступного продукта
	availableAmount, err := CheckProductAmount(order.ProductID, s.InvClient)
	if err != nil {
		return nil, status.Error(codes.Internal, "error getting product amount")
	}

	// Проверяем, хватает ли доступного товара
	log.Println("av:", availableAmount+tempAmount, "req:", req.Amount)
	if req.Amount > availableAmount+tempAmount {
		return nil, status.Error(codes.InvalidArgument, "not enough product amount in inventory")
	}

	// Подготовка обновленных данных заказа
	order.UserID = int(req.OrderId)
	order.Amount = req.Amount
	order.Status = req.Status

	// Обновление в базе данных
	updatedOrder, err := s.Storage.UpdateOrderByIDX(int(req.OrderId), order)
	if err != nil {
		return nil, status.Error(codes.Internal, "error updating order")
	}

	// Изменяем количество товара в наличии
	err = ProductChange(req.Amount-tempAmount, updatedOrder.ProductID, availableAmount, s.InvClient)
	if err != nil {
		return nil, status.Error(codes.Internal, "error updating product amount in inventory")
	}

	// Возвращаем обновленный заказ
	return &orderpb.UpdateOrderResponse{
		Message: "Order updated successfully",
		Order: &orderpb.Order{
			OrderId:   int32(updatedOrder.ID),
			UserId:    int32(updatedOrder.UserID),
			ProductId: int32(updatedOrder.ProductID),
			Amount:    updatedOrder.Amount,
			Status:    updatedOrder.Status,
		},
	}, nil
}

func (s *OrderServiceServer) ListAllOrders(ctx context.Context, req *orderpb.ListAllOrdersRequest) (*orderpb.ListAllOrdersResponse, error) {
	orders, err := s.Storage.ListAllOrdersX()
	if err != nil {
		return nil, err
	}

	var grpcOrders []*orderpb.Order
	for _, o := range orders {
		grpcOrders = append(grpcOrders, &orderpb.Order{
			OrderId:   int32(o.ID),
			UserId:    int32(o.UserID),
			ProductId: int32(o.ProductID),
			Amount:    int64(o.Amount),
			Status:    o.Status,
		})
	}

	return &orderpb.ListAllOrdersResponse{
		Orders: grpcOrders,
	}, nil
}

func CheckProductAmount(productId int, InvClient inventorypb.InventoryServiceClient) (int64, error) {
	const op = "CheckProductAmount"

	ctx := context.Background()

	res, err := InvClient.GetProductById(ctx, &inventorypb.GetByIdRequest{Id: int64(productId)})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("[%s] grpc_status:%s/ error getting product by id: %s", op, st.Code(), err)
			switch st.Code() {
			case codes.NotFound:
				log.Printf("[%s] product not found", op)
			case codes.Internal:
				log.Printf("[%s] internal error", op)
			case codes.InvalidArgument:
				log.Printf("[%s] invalid argument", op)
			}
		} else {
			log.Printf("[%s] error getting product by id: %s", op, err)
		}
		return 0, nil
	}

	return res.Product.Amount, nil
}

func ProductChange(amount int64, productId int, inStorage int64, InvClient inventorypb.InventoryServiceClient) error {
	const op = "ProductChange"

	newAmount := inStorage - amount
	if newAmount == 0 {
		newAmount = -1
	}

	ctx := context.Background()
	_, err := InvClient.UpdateProduct(ctx, &inventorypb.UpdateRequest{
		Id:     int64(productId),
		Name:   "",
		Amount: newAmount,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("[%s] grpc_status:%s/ error updating product by id: %s", op, st.Code(), err)
			switch st.Code() {
			case codes.NotFound:
				log.Printf("[%s] product not found", op)
			case codes.Internal:
				log.Printf("[%s] internal error", op)
			case codes.InvalidArgument:
				log.Printf("[%s] invalid argument", op)
			case codes.AlreadyExists:
				log.Printf("[%s] product already exists", op)
			}
		} else {
			log.Printf("[%s] error updating product by id: %s", op, err)
		}
		return err
	}

	log.Printf("[%s] product updated", op)
	return nil

}
