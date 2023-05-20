package product

import (
	"context"
	"go-api/helper"
	"go-api/intface"
	"go-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	payload_jwt := c.MustGet("user").(*intface.JwtClaim)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var productbody intface.CreateProductBody

	defer cancel()

	if err := c.BindJSON(&productbody); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&productbody); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": validationErr.Error()}})
		return
	}

	user_id, err := helper.ParseObjectID(payload_jwt.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	new_product := intface.CreateProductDocument{
		Name:        productbody.Name,
		Price:       productbody.Price,
		Description: productbody.Description,
		Quantity:    productbody.Quantity,
		User_id:     user_id,
	}

	result_insert, err := product_collection.InsertOne(ctx, new_product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	c.JSON(201, responses.UserResponse{Status: 201, Message: "success", Data: map[string]interface{}{"result": result_insert}})


}
