package testing

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	name := c.Request.FormValue("name")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(500, "error")
		return
	}


	tempFilepath := filepath.Join("temp", file.Filename)

	err = c.SaveUploadedFile(file, tempFilepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":   name,
		"status": "single",
	})
}
