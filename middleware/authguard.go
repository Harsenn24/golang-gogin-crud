package middleware

import (
	"go-api/controller/auth"
	"go-api/environment"
	"go-api/helper"
	"go-api/responses"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Authguard(c *gin.Context) {

	author := c.GetHeader("authorization")

	environment.ExportEnv()
	keyJwt := os.Getenv("KEYJWT")

	secretKey := []byte(keyJwt)

	data, err := helper.DecryptJWT(author, secretKey)

	if err != nil {
		c.AbortWithStatusJSON(500, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": "you dont have authorization"}})
		return
	}

	email := data.Email

	find_account, err := auth.CheckAccount(c, email)

	if err != nil {
		c.AbortWithStatusJSON(500, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": "This JWT does not match the account"}})
		return
	}

	if len(find_account) == 0 {
		c.AbortWithStatusJSON(500, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Authorization Failed"}})
		return
	}

	c.Set("user", data)

	c.Next()

}
