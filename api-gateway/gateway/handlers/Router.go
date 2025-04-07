package handlers

import (
	"api-gateway/Storage"
	"api-gateway/config"
	"api-gateway/gateway/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config, storage Storage.Repo) *gin.Engine {
	router := gin.Default()

	productsGroup := router.Group("/products")
	ordersGroup := router.Group("/orders")
	authGroup := router.Group("/auth")

	productsGroup.Use(middleware.AuthMiddleware(cfg))
	ordersGroup.Use(middleware.AuthMiddleware(cfg))

	SetupAuth(authGroup, cfg, storage)
	SetupOrders(ordersGroup, cfg)
	SetupProducts(productsGroup, cfg)

	return router
}
