package dbOps

import (
	"AdvancedNetwork/connections"
	"AdvancedNetwork/pkg/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func DB_GetAll_Videos() ([]models.Video, error) {

	var Result []models.Video

	filterCursor, err := connections.DATABASE.Collection("AN_Videos").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	for filterCursor.Next(context.TODO()) {
		var filteredResult models.Video
		if err := filterCursor.Decode(&filteredResult); err != nil {
			fmt.Println("Error ", "mCursor.Decode(&extract)")
			fmt.Println(err.Error())
			return nil, err
		} else {
			if filteredResult.Title != "" {
				Result = append(Result, filteredResult)
			}
		}
	}
	return Result, nil

}
