package middleware

import (
	"go-api/controller/auth"
	"go-api/helper"
	"go-api/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authguard(c *gin.Context) {

	author := c.GetHeader("authorization")

	secretKey := []byte("123456789")

	data, err := helper.DecryptJWT(author, secretKey)

	if err != nil {
		c.AbortWithStatusJSON(500, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	email := data.Email

	find_account, err := auth.CheckAccount(c, email)

	if err != nil {
		c.AbortWithStatusJSON(500, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if len(find_account) == 0 {
		c.AbortWithStatusJSON(500, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Authorization Failed"}})
		return
	}

	c.Next()

}
