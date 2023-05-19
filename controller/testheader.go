package controller

import (
	"go-api/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestCode(c *gin.Context) {

	author := c.GetHeader("authorization")

	if author == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "You have to login first"}})

		return
	}

	c.Next()


}