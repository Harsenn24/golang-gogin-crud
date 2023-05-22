package helper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FilterSearch(key_name string, value_name string) primitive.D {
	if value_name == "" {
		return bson.D{}
	}

	filter := bson.D{
		{Key: key_name, Value: primitive.D{
			{Key: "$regex", Value: value_name},
			{Key: "$options", Value: "i"},
		}},
	}
	return filter
}
