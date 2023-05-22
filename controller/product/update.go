package product

import (
	"context"
	"go-api/intface"
	"go-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdatePRoduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var product intface.CreateProductBody

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if vaidationErr := validate.Struct(&product); vaidationErr != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": vaidationErr.Error()}})
		return
	}

	id_product := c.Param("id")

	//change id_product to object id
	id, err := primitive.ObjectIDFromHex(id_product)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	filter := bson.D{
		{Key: "_id", Value: id},
	}

	set_update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: product.Name},
			{Key: "price", Value: product.Price},
			{Key: "quantity", Value: product.Quantity},
			{Key: "description", Value: product.Description},
		}},
	}

	result_update, err := product_collection.UpdateOne(ctx, filter, set_update)	
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if result_update.MatchedCount == 0 {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "id not found"}})
		return
	}

	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "success update product"}})

}
