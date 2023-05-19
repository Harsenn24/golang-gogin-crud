package testing

import (
	"context"
	"go-api/config"
	"go-api/intface"
	"go-api/responses"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var brand_collection *mongo.Collection = config.GetCollection(config.DB, "brand")

func AggragateExample(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	pipeline := bson.A{
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "detail", Value: bson.D{
					{Key: "name", Value: "$detail.name"},
					{Key: "image", Value: "$detail.image"},
					{Key: "phone", Value: "$detail.phone"},
				}},
				{Key: "UserAccess", Value: bson.D{
					{Key: "$map", Value: bson.D{
						{Key: "input", Value: "$userAccess"},
						{Key: "in", Value: bson.D{
							{Key: "email", Value: "$$this.email"},
							{Key: "username", Value: "$$this.username"},
						}},
					}},
				}},
			}},
		},
	}

	result, err := brand_collection.Aggregate(ctx, pipeline)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	var results []intface.I_Brand
	if err := result.All(ctx, &results); err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": results}})

}
