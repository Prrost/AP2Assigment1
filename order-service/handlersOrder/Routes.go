package handlersOrder

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"log/slog"
	"net/http"
	"order-service/Storage"
	"order-service/config"
	"order-service/domain"
	"strconv"
	"strings"
	"time"
)

type Object struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int64  `json:"amount"`
	Available bool   `json:"available"`
}

type ObjectResponse struct {
	Object Object `json:"object"`
}

func SetupRoutes(group *gin.RouterGroup, cfg *config.Config, storage Storage.Storage) {

	group.POST("/", func(c *gin.Context) {
		CreateOrder(c, storage)
	})

	group.GET("/:id", func(c *gin.Context) {
		GetOrderByID(c, storage)
	})

	group.PUT("/:id", func(c *gin.Context) {
		UpdateOrderByID(c, storage)
	})

	group.GET("/", func(c *gin.Context) {
		ListAllOrders(c, storage)
	})
}

func CreateOrder(c *gin.Context, storage Storage.Storage) {

	var insertOrder struct {
		ProductId int   `json:"productId" binding:"required"`
		UserId    int   `json:"userId"`
		Amount    int64 `json:"amount" binding:"required"`
	}

	id, ok := ExtractUserID(c)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id extract failed"})
		return
	} else {
		insertOrder.UserId = int(id.(float64))
	}

	err := c.ShouldBindJSON(&insertOrder)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if CheckProductExists(insertOrder.ProductId, c) == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "product doesn't exist"})
		return
	}
	if CheckUserExists(insertOrder.UserId, c) == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "user doesn't exist"})
		return
	}

	var order domain.Order

	if insertOrder.Amount < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "amount can't be less than 1"})
		return
	}

	productAmount := CheckProductAmount(insertOrder.ProductId, c)
	if productAmount == -2 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to check product amount"})
		return
	}

	if productAmount < insertOrder.Amount {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "not enough product amount"})
		return
	}

	order.Amount = insertOrder.Amount

	order.UserID = insertOrder.UserId
	order.ProductID = insertOrder.ProductId
	order.Status = "created"

	createdOrder, err := storage.CreateOrderX(order)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ProductChange(insertOrder.Amount, insertOrder.ProductId, productAmount)

	c.IndentedJSON(http.StatusOK, gin.H{"order": createdOrder})

}

func GetOrderByID(c *gin.Context, storage Storage.Storage) {

	orderID, ok := c.Params.Get("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required in query param"})
		return
	}

	id, err := strconv.Atoi(orderID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if CheckProductExists(id, c) == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "product doesn't exist"})
		return
	}

	order, err := storage.GetOrderByIDX(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"order": order})

}

func UpdateOrderByID(c *gin.Context, storage Storage.Storage) {

	orderId, ok := c.Params.Get("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required in query param"})
		return
	}

	id, err := strconv.Atoi(orderId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedOrder struct {
		UserID    int    `json:"userId"`
		ProductID int    `json:"productId"`
		Amount    int64  `json:"amount"`
		Status    string `json:"status"`
	}

	userId, ok := ExtractUserID(c)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id extract failed"})
		return
	} else {
		updatedOrder.UserID = int(userId.(float64))
	}

	err = c.ShouldBindJSON(&updatedOrder)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	object, err := storage.GetOrderByIDX(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tempAmount := object.Amount

	if CheckProductAmount(updatedOrder.ProductID, c) == -2 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to check product amount"})
		return
	}

	if updatedOrder.Amount > CheckProductAmount(updatedOrder.ProductID, c) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "not enough product amount"})
		return
	}

	switch {
	case updatedOrder.Amount < 1:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "amount can't be less than 1"})
		return
	case updatedOrder.Amount > 0:
		object.Amount = updatedOrder.Amount
	}

	if CheckUserExists(updatedOrder.UserID, c) {
		object.UserID = updatedOrder.UserID
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "user doesn't exist"})
		return
	}

	if CheckProductExists(updatedOrder.ProductID, c) {
		object.ProductID = updatedOrder.ProductID
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "product doesn't exist"})
		return
	}

	if updatedOrder.Status != "" {
		object.Status = updatedOrder.Status
	}

	newObject, err := storage.UpdateOrderByIDX(id, object)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	inStorage := CheckProductAmount(updatedOrder.ProductID, c)

	log.Println(updatedOrder.Amount - tempAmount)

	ProductChange(updatedOrder.Amount-tempAmount, updatedOrder.ProductID, inStorage)

	c.IndentedJSON(http.StatusOK, gin.H{"object": newObject})

}

func ListAllOrders(c *gin.Context, storage Storage.Storage) {

	products, err := storage.ListAllOrdersX()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})

}

func CheckProductAmount(productId int, c *gin.Context) int64 {

	url := fmt.Sprintf("http://localhost:8081/products/%d", productId)

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		slog.Error("ss", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to contact inventory service"})
		return -2
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.IndentedJSON(resp.StatusCode, gin.H{"error": "failed to fetch object"})
		return -2
	}

	var objectResponse ObjectResponse
	err = json.NewDecoder(resp.Body).Decode(&objectResponse)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to parse inventory response"})
		return -2
	}

	return objectResponse.Object.Amount
}

func CheckUserExists(userId int, c *gin.Context) bool {
	return true
}

func CheckProductExists(productId int, c *gin.Context) bool {
	return true
}

func ExtractUserID(c *gin.Context) (interface{}, bool) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, false
	}

	parts := strings.Split(authHeader, "Bearer ")
	if len(parts) != 2 {
		return nil, false
	}

	tokenString := parts[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("supersecretkey"), nil
	})

	if err != nil || !token.Valid {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if id, ok := claims["ID"]; ok {
			return id, true
		}
	}

	return nil, false
}

func ProductChange(amount int64, productId int, inStorage int64) {
	url := fmt.Sprintf("http://localhost:8081/products/%d", productId)

	type Payload struct {
		Amount int64 `json:"amount"`
	}

	data := Payload{Amount: inStorage - amount}
	log.Println(inStorage - amount)
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(err)
		return
	}

	log.Println("product successfully changed")
}
