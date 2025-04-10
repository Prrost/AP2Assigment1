package handlersOrder

import (
	"github.com/gin-gonic/gin"
	"order-service/Storage"
	"order-service/config"
)

func SetupRouter(cfg *config.Config, storage Storage.Storage) *gin.Engine {
	router := gin.Default()

	mainGroup := router.Group("/orders")

	SetupRoutes(mainGroup, cfg, storage)

	return router
}
