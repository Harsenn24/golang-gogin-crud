package product

import (
	"context"
	"go-api/helper"
	"go-api/intface"
	"go-api/responses"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(c *gin.Context) {
	
	payload_jwt := c.MustGet("user").(*intface.JwtClaim)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	name := c.Request.FormValue("name")
	price, err:= helper.StringToInT(c.Request.FormValue("price"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": "parse string to number"}})
		return
	}
	description := c.Request.FormValue("description")
	quantity, err:= helper.StringToInT(c.Request.FormValue("quantity"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": "parse string to number"}})
		return
	}


	user_id, err := helper.ParseObjectID(payload_jwt.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	new_product := intface.CreateProductDocument{
		Name:        name,
		Price:       price,
		Description: description,
		Quantity:    quantity,
		User_id:     user_id,
	}

	result_insert, err := product_collection.InsertOne(ctx, new_product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	id_insert_string := result_insert.InsertedID.(primitive.ObjectID).Hex()
	

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": "error getting file"}})
		return
	}

	tempFilepath := filepath.Join("upload", id_insert_string + ".jpg")

	err = c.SaveUploadedFile(file, tempFilepath)
	if err != nil {
		c.JSON(400, responses.UserResponse{Status: 400, Message: "error", Data: map[string]interface{}{"error": "error upload"}})
		return
	}

	c.JSON(201, responses.UserResponse{Status: 201, Message: "success", Data: map[string]interface{}{"result": "success create product"}})


}
