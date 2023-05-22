package product

import (
	"github.com/gin-gonic/gin"
)

func ImageProduct(c *gin.Context) {
	id := c.Param("id")

	c.File("upload/" + id + ".jpg")
}
