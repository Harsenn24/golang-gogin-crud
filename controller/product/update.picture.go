package product

import (
	"go-api/responses"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UpdateImageProduct(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": "error getting file"}})
		return
	}

	id := c.Param("id")

	save_image := filepath.Join("upload", id +".jpg")

	err = c.SaveUploadedFile(file, save_image)
	if err != nil {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": "error saving file"}})
		return
	}

	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "success update image"}})
}
