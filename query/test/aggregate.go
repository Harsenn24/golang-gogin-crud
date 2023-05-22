package test

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func QueryTest() primitive.A {
	pipeLine := bson.A{
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "detail", Value: bson.D{
					{Key: "name", Value: "$detail.name"},
					{Key: "image", Value: "$detail.image"},
					{Key: "phone", Value: "$detail.phone"},
				}},
				{Key: "UserAccess", Value: bson.D{
					{Key: "$map", Value: bson.D{
						{Key: "input", Value: "$userAccess"},
						{Key: "in", Value: bson.D{
							{Key: "email", Value: "$$this.email"},
							{Key: "username", Value: "$$this.username"},
						}},
					}},
				}},
			}},
		},
	}

	return pipeLine
}
