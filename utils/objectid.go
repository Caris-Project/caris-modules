package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

// GenerateMongoID ...
func GenerateMongoID() primitive.ObjectID {
	return primitive.NewObjectID()
}
