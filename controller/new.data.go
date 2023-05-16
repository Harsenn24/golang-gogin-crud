package controller

import (
	"fmt"
	"net/http"

	"go-api/intface"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)



func NewData(c *gin.Context) {
	var inputData intface.I_NewData

	err := c.ShouldBindJSON(&inputData)

	if err != nil {

		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": errorMessages,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":         inputData.Name,
		"age":          inputData.Age,
		"tempat_lahir": inputData.BirthPlace,
	})

}