package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Video struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
}
