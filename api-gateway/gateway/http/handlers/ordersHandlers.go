package handlers

import (
	"api-gateway/gateway/grpc/clients"
	"api-gateway/mapping"
	"api-gateway/orderpb"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func SetupOrders(group *gin.RouterGroup, grpcClient *clients.Client) {

	group.GET("/", func(c *gin.Context) {
		GetAllOrders(c, grpcClient)
	})

	group.GET("/:id", func(c *gin.Context) {
		GetOrderById(c, grpcClient)
	})

	group.POST("/", func(c *gin.Context) {
		CreateOrder(c, grpcClient)
	})

	group.PUT("/:id", func(c *gin.Context) {
		UpdateOrder(c, grpcClient)
	})
}

func CreateOrder(c *gin.Context, client *clients.Client) {
	const op = "CreateOrder"

	var order mapping.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, exist := c.Get("id")
	if !exist {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "user ID not found"})
		return
	}

	var id int
	switch v := userId.(type) {
	case int:
		id = v
	case float64:
		id = int(v)
	default:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "user ID must be an integer"})
		return
	}

	order.UserId = int32(id)

	res, err := client.OrderClient.CreateOrder(c.Request.Context(), &orderpb.CreateOrderRequest{
		ProductId: order.ProductID,
		UserId:    int32(id),
		Amount:    order.Amount,
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"order":   mapping.ToOrder(res.Order),
		"message": res.Message,
	})
}

func GetAllOrders(c *gin.Context, client *clients.Client) {
	const op = "GetAllOrders"

	res, err := client.OrderClient.ListAllOrders(c.Request.Context(), &orderpb.ListAllOrdersRequest{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"orders": mapping.ToOrders(res.Orders),
	})
}

func GetOrderById(c *gin.Context, client *clients.Client) {
	const op = "GetOrderByID"

	orderId := c.Param("id")
	id, err := strconv.ParseInt(orderId, 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	res, err := client.OrderClient.GetOrder(c.Request.Context(), &orderpb.GetOrderRequest{
		OrderId: int32(id),
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"order":   mapping.ToOrder(res.Order),
		"message": res.Message,
	})
}

func UpdateOrder(c *gin.Context, client *clients.Client) {
	const op = "UpdateOrder"

	orderId := c.Param("id")
	id, err := strconv.ParseInt(orderId, 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	var order mapping.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(id)

	order.OrderID = int32(id)

	res, err := client.OrderClient.UpdateOrder(c.Request.Context(), &orderpb.UpdateOrderRequest{
		OrderId: order.OrderID,
		Amount:  order.Amount,
		Status:  order.Status,
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"order":   mapping.ToOrder(res.Order),
		"message": res.Message,
	})
}
