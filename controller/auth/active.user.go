package auth

import (
	"context"
	"go-api/helper"
	"go-api/intface"
	"go-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

func ActiveUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var user intface.BodyUserUpdate

	var validate = validator.New()

	defer cancel()

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": validationErr.Error()}})
		return
	}

	Object_id, err := helper.ParseObjectID(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	filter := bson.D{
		{Key: "_id", Value: Object_id},
	}

	set_update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "active", Value: true},
		}},
	}

	result_update, err := userCollection.UpdateOne(ctx, filter, set_update)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	if result_update.MatchedCount == 0  {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": "id not found"}})
		return
	}

	if result_update.MatchedCount == 1 && result_update.ModifiedCount == 0 {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": "This account already active"}})
		return
	}

	c.JSON(200, responses.UserResponse{Status: 200, Message: "success", Data: map[string]interface{}{"result": "Acount is active"}})
}
