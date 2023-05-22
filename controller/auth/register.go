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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(c *gin.Context) {
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

	find_account, err := CheckAccount(c, user.Email)

	if err != nil {
		c.AbortWithStatusJSON(500, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if len(find_account) > 0 {
		c.AbortWithStatusJSON(500, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Email already in used"}})
		return
	}

	epoch, err := helper.ConvertToEpoch(user.Birthday)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	newUser := intface.I_userMongo{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
		Birthday: int(epoch),
		Active:   false,
	}

	result_insert, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	
	object_id := result_insert.InsertedID.(primitive.ObjectID).Hex()

	message_email := "Thank you for registering your account!\n\nYour user ID: " + object_id

	result_send, err := helper.SendEmail(newUser.Email, "REGISTRATION SUCCESS", message_email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "sucess", Data: map[string]interface{}{"result": result_send}})




}
