package mapping

import "api-gateway/orderpb"

type Order struct {
	OrderID   int32  `json:"order_id"`
	ProductID int32  `json:"product_id"`
	UserId    int32  `json:"user_id"`
	Amount    int64  `json:"amount"'`
	Status    string `json:"status"`
}

func ToOrder(order *orderpb.Order) *Order {
	return &Order{
		OrderID:   order.GetOrderId(),
		ProductID: order.GetProductId(),
		UserId:    order.GetUserId(),
		Amount:    order.GetAmount(),
		Status:    order.GetStatus(),
	}
}

func ToOrders(orders []*orderpb.Order) []*Order {
	result := make([]*Order, 0, len(orders))
	for _, order := range orders {
		result = append(result, ToOrder(order))
	}
	return result
}
