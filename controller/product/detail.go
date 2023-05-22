package product

import (
	"context"
	"go-api/helper"
	"go-api/intface"
	"go-api/query/product"
	"go-api/responses"
	"time"

	"github.com/gin-gonic/gin"
)

func DetailProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	id := c.Param("id")

	convert_id, err := helper.ParseObjectID(id)
	if err != nil {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	pipeLine := product.QueryProductDetail(convert_id)

	result, err := product_collection.Aggregate(ctx, pipeLine)
	if err != nil {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	var results []intface.DetailProduct
	if err := result.All(ctx, &results); err != nil {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	if len(results) == 0 {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": "data not found"}})
		return
	}

	c.JSON(200, responses.NewResponses(200, "success", results[0]))

}
