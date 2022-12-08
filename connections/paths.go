package connections

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
)

var DATABASE *mongo.Database

const DATABASE_URL = "mongodb+srv://dami:dami123@cluster0.otujs.mongodb.net/?retryWrites=true&w=majority"

const DATABASE_NAME = "evocodeui" // Do not change DB name
