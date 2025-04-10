package handlers

import (
	"api-gateway/config"
	"api-gateway/gateway/proxi"
	"github.com/gin-gonic/gin"
)

func SetupOrders(group *gin.RouterGroup, cfg *config.Config) {

	group.GET("/", func(c *gin.Context) {
		proxi.ProxiRequest(cfg.OrderService, c)
	})

	group.GET("/:id", func(c *gin.Context) {
		proxi.ProxiRequest(cfg.OrderService, c)
	})

	group.POST("/", func(c *gin.Context) {
		proxi.ProxiRequest(cfg.OrderService, c)
	})

	group.PUT("/:id", func(c *gin.Context) {
		proxi.ProxiRequest(cfg.OrderService, c)
	})
}
