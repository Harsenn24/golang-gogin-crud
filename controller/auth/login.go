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

func LoginUser(c *gin.Context) {

	var validate = validator.New()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var userlogin intface.I_Login

	defer cancel()

	if err := c.BindJSON(&userlogin); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&userlogin); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	pipeline := bson.A{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "email", Value: userlogin.Email},
			}},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "email", Value: "$email"},
				{Key: "password", Value: "$password"},
				{Key: "id", Value: bson.D{
					{Key: "$toString", Value: "$_id"},
				}},
			}},
		},
	}

	result, err := userCollection.Aggregate(ctx, pipeline)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	var results []intface.I_LoginResult
	if err := result.All(ctx, &results); err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	matchPassword := helper.DecryptPassword(results[0].Password, userlogin.Password)

	if matchPassword == "password match" {

		payload := intface.I_LoginResult{
			Email:    results[0].Email,
			Password: results[0].Password,
			Id:       results[0].Id,
		}
		token_data, err := helper.JwtSign(&payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
	
		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"token": token_data}})
		return
	}

	c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"result": matchPassword}})
	
}
