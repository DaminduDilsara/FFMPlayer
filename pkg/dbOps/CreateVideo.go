package dbOps

import (
	"AdvancedNetwork/connections"
	"AdvancedNetwork/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_Create_Video(application models.Video) (*mongo.InsertOneResult, error) {

	id, err := connections.DATABASE.Collection("AN_Videos").InsertOne(context.Background(), application)

	if err != nil {
		return nil, err
	}
	return id, nil
}
