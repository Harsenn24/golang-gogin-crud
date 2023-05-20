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
