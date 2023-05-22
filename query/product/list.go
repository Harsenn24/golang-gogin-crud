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
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "user"},
				{Key: "localField", Value: "user_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "user"},
				{Key: "pipeline", Value: bson.A{
					bson.D{
						{Key: "$project", Value: bson.D{
							{Key: "name", Value: "$name"},
						}},
					},
				}},
			}},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "id", Value: bson.D{
					{Key: "$toString", Value: "$_id"},
				}},
				{Key: "product_name", Value: "$name"},
				{Key: "product_price", Value: "$price"},
				{Key: "product_quantity", Value: "$quantity"},
				{Key: "product_description", Value: "$description"},
				{Key: "user", Value: bson.D{
					{Key: "$ifNull", Value: bson.A{
						bson.D{
							{Key: "$first", Value: "$user.name"},
						},
						"no user",
					}},
				}},
			}},
		},
	}

	return pipeLine
}
