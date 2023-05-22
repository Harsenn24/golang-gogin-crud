package product

import (
	"context"
	"go-api/intface"
	"go-api/responses"
	"go-api/query/product"

	"time"

	"github.com/gin-gonic/gin"
)

func ListProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	pipeLine := product.QueryProductList()

	result, err := product_collection.Aggregate(ctx, pipeLine)
	if err != nil {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	var results []intface.ListProduct
	if err := result.All(ctx, &results); err != nil {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	c.JSON(200, responses.UserResponse{Status: 200, Message: "success", Data: map[string]interface{}{"data": results}})

}
