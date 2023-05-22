package product

import (
	"go-api/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func QueryProductList(query string) primitive.A {
	pipeLine := bson.A{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "price", Value: bson.D{
					{Key: "$gt", Value: 46000},
				}},
			}},
		},
		bson.D{
			{Key: "$match", Value: helper.FilterSearch("name", query)},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "id", Value: bson.D{
					{Key: "$toString", Value: "$_id"},
				}},
				{Key: "product_name", Value: "$name"},
				{Key: "product_price", Value: "$price"},
			}},
		},
	}

	return pipeLine
}
