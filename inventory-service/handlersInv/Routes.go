package handlersInv

import (
	"errors"
	"github.com/gin-gonic/gin"
	"inventory-service/Storage"
	"inventory-service/config"
	"inventory-service/domain"
	"net/http"
	"strconv"
)

func SetupRoutes(group *gin.RouterGroup, cfg *config.Config, storage Storage.Storage) {

	group.POST("/", func(c *gin.Context) {
		CreateProduct(c, storage)
	})

	group.GET("/:id", func(c *gin.Context) {
		GetProductByID(c, storage)
	})

	group.PUT("/:id", func(c *gin.Context) {
		UpdateProductByID(c, storage)
	})

	group.DELETE("/:id", func(c *gin.Context) {
		DeleteObject(c, storage)
	})

	group.GET("/", func(c *gin.Context) {
		GetAllProducts(c, storage)
	})
}

func CreateProduct(c *gin.Context, storage Storage.Storage) {
	var insertObject struct {
		Name   string `json:"name" binding:"required"`
		Amount int64  `json:"amount" binding:"required"`
	}

	err := c.ShouldBindJSON(&insertObject)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var object domain.Object

	if insertObject.Amount < -1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "amount can't be less than -1"})
		return
	}

	if insertObject.Amount == -1 {
		object.Amount = 0
	} else {
		object.Amount = insertObject.Amount
	}

	object.Name = insertObject.Name

	if object.Amount > 0 {
		object.Available = true
	} else {
		object.Available = false
	}

	createdObject, err := storage.CreateObject(object)
	if err != nil {
		if errors.Is(err, Storage.ErrAlreadyExists) {
			c.IndentedJSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"object": createdObject})

}

func GetProductByID(c *gin.Context, storage Storage.Storage) {
	productID, ok := c.Params.Get("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required in query param"})
		return
	}

	id, err := strconv.Atoi(productID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	object, err := storage.GetObjectByID(id)
	if err != nil {
		if errors.Is(err, Storage.ErrNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"object": object})

}

func DeleteObject(c *gin.Context, storage Storage.Storage) {
	productID, ok := c.Params.Get("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required in query param"})
		return
	}

	id, err := strconv.Atoi(productID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	object, err := storage.DeleteObjectByID(id)
	if err != nil {
		if errors.Is(err, Storage.ErrNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"deletedObject": object})
}

func UpdateProductByID(c *gin.Context, storage Storage.Storage) {
	productID, ok := c.Params.Get("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required in query param"})
		return
	}

	id, err := strconv.Atoi(productID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedObject struct {
		Name   string `json:"name"`
		Amount int64  `json:"amount"`
	}

	err = c.ShouldBindJSON(&updatedObject)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	object, err := storage.GetObjectByID(id)
	if err != nil {
		if errors.Is(err, Storage.ErrNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updatedObject.Name != object.Name {
		ok, err = storage.IsProductExists(updatedObject.Name)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if ok {
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "Product name already exists"})
			return
		}
	}

	switch {
	case updatedObject.Amount < -1:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "amount can't be less than -1"})
		return
	case updatedObject.Amount == -1:
		object.Amount = 0
		object.Available = false
	case updatedObject.Amount > 0:
		object.Amount = updatedObject.Amount
		object.Available = true
	}

	if updatedObject.Name != "" {
		object.Name = updatedObject.Name
	}

	newObject, err := storage.UpdateObjectByID(id, object)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"object": newObject})

}

func GetAllProducts(c *gin.Context, storage Storage.Storage) {

	name := c.Query("name")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	products, err := storage.ListProducts(name, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
