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
				{Key: "$and", Value: bson.A{
					bson.D{
						{Key: "email", Value: userlogin.Email},
						{Key: "active", Value: true},
					},
				}},
			}},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "email", Value: "$email"},
				{Key: "password", Value: "$password"},
				{Key: "active", Value: "$active"},
				{Key: "id", Value: bson.D{
					{Key: "$toString", Value: "$_id"},
				}},
			}},
		},
	}

	result, err := userCollection.Aggregate(ctx, pipeline)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var results []intface.I_LoginResult
	
	if err := result.All(ctx, &results); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(results) == 0 {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": "user not found"}})
		return
	}

	if !results[0].Active {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": "This account is not already active"}})
		return
	}

	matchPassword := helper.DecryptPassword(results[0].Password, userlogin.Password)

	if matchPassword == "password match" {

		payload := intface.CheckAccount{
			Email: results[0].Email,
			Id:    results[0].Id,
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
