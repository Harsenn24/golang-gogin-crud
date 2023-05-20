package product

import (

	"github.com/gin-gonic/gin"
)

func Createproduct(c *gin.Context){
	user := c.MustGet("user")
	c.JSON(200, user)
}