package Storage

import "order-service/domain"

type Storage interface {
	CreateOrderX(order domain.Order) (domain.Order, error)
	GetOrderByIDX(id int) (domain.Order, error)
	UpdateOrderByIDX(id int, order domain.Order) (domain.Order, error)
	ListAllOrdersX() ([]domain.Order, error)
}
