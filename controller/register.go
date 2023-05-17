package controller

import (
	"context"
	"go-api/config"
	"go-api/helper"
	"go-api/intface"
	"go-api/responses"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)


func CreateUser(c *gin.Context) {
	var userCollection *mongo.Collection = config.GetCollection(config.DB, "user")
	var validate = validator.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user intface.I_user
	defer cancel()

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	epoch, err := helper.ConvertToEpoch(user.Birthday, "2006-01-02")
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	newUser := intface.I_userMongo{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
		Birthday: int(epoch),
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"result": result}})

}
