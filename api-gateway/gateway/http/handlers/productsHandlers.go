package handlers

import (
	"api-gateway/gateway/grpc/clients"
	"api-gateway/mapping"
	inventorypb "github.com/Prrost/assignment1proto/proto/inventory"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"strconv"
)

func SetupProducts(group *gin.RouterGroup, grpcClient *clients.Client) {

	group.GET("/:id", func(c *gin.Context) {
		GetByID(c, grpcClient)
	})

	group.GET("/", func(c *gin.Context) {
		GetAll(c, grpcClient)
	})

	group.POST("/", func(c *gin.Context) {
		CreateProduct(c, grpcClient)
	})

	group.PUT("/:id", func(c *gin.Context) {
		UpdateProduct(c, grpcClient)
	})

	group.DELETE("/:id", func(c *gin.Context) {
		DeleteProduct(c, grpcClient)
	})
}

func CreateProduct(c *gin.Context, grpcClient *clients.Client) {
	const op = "CreateProduct"

	var product mapping.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := grpcClient.InventoryClient.CreateProduct(c.Request.Context(), &inventorypb.CreateRequest{
		Name:   product.Name,
		Amount: product.Amount,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("[%s] grpc_status:%s/ error creating products: %s", op, st.Code(), err)
			switch st.Code() {
			case codes.AlreadyExists:
				c.IndentedJSON(http.StatusConflict, gin.H{"error": err.Error()})
			case codes.Internal:
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			case codes.InvalidArgument:
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		} else {
			log.Printf("[%s] error creating products: %s", op, err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"product": mapping.ToProduct(res.Product),
		"message": res.Message,
	})

}

func GetAll(c *gin.Context, grpcClient *clients.Client) {
	const op = "GetAll"

	name := c.Query("name")
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		log.Printf("[%s] limit atoi err: %v", op, err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	intOffset, err := strconv.Atoi(offset)
	if err != nil {
		log.Printf("[%s] offset atoi err: %v", op, err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := grpcClient.InventoryClient.GetAllProducts(c.Request.Context(), &inventorypb.GetAllRequest{
		Name:   name,
		Limit:  int32(intLimit),
		Offset: int32(intOffset),
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("[%s] grpc_status:%s/ error getting products: %s", op, st.Code(), err)
			switch st.Code() {
			case codes.NotFound:
				c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			case codes.Internal:
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			case codes.InvalidArgument:
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		} else {
			log.Printf("[%s] error getting products: %s", op, err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"products": mapping.ToProducts(res.Products),
		"message":  res.Message,
	})

}

func GetByID(c *gin.Context, grpcClient *clients.Client) {
	const op = "GetByID"
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[%s] error converting id to int: %s", op, err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := grpcClient.InventoryClient.GetProductById(c.Request.Context(), &inventorypb.GetByIdRequest{Id: int64(idInt)})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("[%s] grpc_status:%s/ error getting product by id: %s", op, st.Code(), err)
			switch st.Code() {
			case codes.NotFound:

				c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			case codes.Internal:
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			case codes.InvalidArgument:
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		} else {
			log.Printf("[%s] error getting product by id: %s", op, err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"product": mapping.ToProduct(res.Product),
		"message": res.Message,
	})
}

func UpdateProduct(c *gin.Context, grpcClient *clients.Client) {
	const op = "UpdateProduct"

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[%s] error converting id to int: %s", op, err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var product mapping.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := grpcClient.InventoryClient.UpdateProduct(c.Request.Context(), &inventorypb.UpdateRequest{
		Id:     int64(idInt),
		Name:   product.Name,
		Amount: product.Amount,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("[%s] grpc_status:%s/ error updating product by id: %s", op, st.Code(), err)
			switch st.Code() {
			case codes.NotFound:
				c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			case codes.Internal:
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			case codes.InvalidArgument:
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			case codes.AlreadyExists:
				c.IndentedJSON(http.StatusConflict, gin.H{"error": err.Error()})
			}
		} else {
			log.Printf("[%s] error updating product by id: %s", op, err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"product": mapping.ToProduct(res.Product),
		"message": res.Message,
	})
}

func DeleteProduct(c *gin.Context, grpcClient *clients.Client) {
	const op = "DeleteProduct"

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[%s] error converting id to int: %s", op, err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := grpcClient.InventoryClient.DeleteProduct(c.Request.Context(), &inventorypb.DeleteRequest{
		Id: int64(idInt),
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("[%s] grpc_status:%s/ error deleting product by id: %s", op, st.Code(), err)
			switch st.Code() {
			case codes.NotFound:

				c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			case codes.Internal:
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			case codes.InvalidArgument:
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		} else {
			log.Printf("[%s] error deleting product by id: %s", op, err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"product": mapping.ToProduct(res.Product),
		"message": res.Message,
	})

}
