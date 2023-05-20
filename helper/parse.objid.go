package helper

import "go.mongodb.org/mongo-driver/bson/primitive"

func ParseObjectID(idString string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return objectID, nil
}
