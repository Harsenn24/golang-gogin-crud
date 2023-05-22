package intface

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateProductBody struct {
	Name        string `binding:"required"`
	Price       int    `binding:"required,number"`
	Description string `binding:"required"`
	Quantity    int    `binding:"required,number"`
}

type CreateProductDocument struct {
	Name        string `binding:"required"`
	Price       int    `binding:"required,number"`
	Description string `binding:"required"`
	Quantity    int    `binding:"required,number"`
	User_id     primitive.ObjectID
}

type ListProduct struct {
	Product_Name        string `json:"product_name"`
	Product_price       int    `json:"product_price"`
	Product_description string `json:"product_description"`
	Product_quantity    int    `json:"product_quantity"`
	User                string `json:"user"`
	Id                  string `json:"id"`
}
