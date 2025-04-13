package handlers

import (
	"api-gateway/config"
	"api-gateway/gateway/proxi"
	"github.com/gin-gonic/gin"
)

func SetupProducts(group *gin.RouterGroup, cfg *config.Config) {

	group.GET("/:id", func(c *gin.Context) {
		proxi.ProxiRequest(cfg.InventoryService, c)
	})

	group.GET("/", func(c *gin.Context) {
		proxi.ProxiRequest(cfg.InventoryService, c)
	})

	group.POST("/", func(c *gin.Context) {
		proxi.ProxiRequest(cfg.InventoryService, c)
	})

	group.PUT("/:id", func(c *gin.Context) {
		proxi.ProxiRequest(cfg.InventoryService, c)
	})

	group.DELETE("/:id", func(c *gin.Context) {
		proxi.ProxiRequest(cfg.InventoryService, c)
	})
}
