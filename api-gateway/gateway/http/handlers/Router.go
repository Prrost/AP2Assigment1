package handlers

import (
	"api-gateway/config"
	"api-gateway/gateway/grpc/clients"
	"api-gateway/gateway/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config, grpcClient *clients.Client) *gin.Engine {
	router := gin.Default()

	productsGroup := router.Group("/products")
	ordersGroup := router.Group("/orders")
	authGroup := router.Group("/auth")

	productsGroup.Use(middleware.AuthMiddleware(cfg))
	ordersGroup.Use(middleware.AuthMiddleware(cfg))

	SetupAuth(authGroup, grpcClient)
	SetupOrders(ordersGroup, cfg)
	SetupProducts(productsGroup, cfg)

	return router
}
