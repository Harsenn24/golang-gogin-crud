package testing

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func HelloQuery(c *gin.Context) {
	id := c.Param("id")
	query := c.Query("ini_query")
	c.JSON(http.StatusOK, gin.H{
		"name":   "harsenn",
		"status": "single",
		"id":     id,
		"query":  query,
	})
}




