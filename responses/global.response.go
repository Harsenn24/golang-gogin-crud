package responses

import "github.com/gin-gonic/gin"

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewResponses(status int, message string, data interface{}) (response gin.H) {
	return gin.H{
		"status":  status,
		"data":    data,
		"message": message,
	}
}
