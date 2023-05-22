package testing

import (
	"context"
	"go-api/config"
	"go-api/intface"
	"go-api/responses"
	"go-api/query/test"


	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var brand_collection *mongo.Collection = config.GetCollection(config.DB, "brand")

func AggragateExample(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	pipeline := test.QueryTest()

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
