package handlersInv

import (
	"github.com/gin-gonic/gin"
	"inventory-service/Storage"
	"inventory-service/config"
)

func SetupRouter(cfg *config.Config, storage Storage.Storage) *gin.Engine {
	router := gin.Default()

	mainGroup := router.Group("/products")

	SetupRoutes(mainGroup, cfg, storage)

	return router
}
