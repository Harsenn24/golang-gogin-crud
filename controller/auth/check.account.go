package auth

import (
	"context"
	"go-api/intface"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func CheckAccount(c *gin.Context, email string) ([]intface.I_user, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	pipeline := bson.A{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "email", Value: email},
			}},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "email", Value: "$email"},
			}},
		},
	}

	result, err := userCollection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var results []intface.I_user
	if err := result.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil

}
