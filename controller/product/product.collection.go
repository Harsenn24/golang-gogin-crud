package product

import (
	"go-api/config"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var product_collection *mongo.Collection = config.GetCollection(config.DB, "product")
var validate = validator.New()