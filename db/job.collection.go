package db

import "go.mongodb.org/mongo-driver/mongo"

func JobCollection() *mongo.Collection {
	return Client.Database("go-job-api").Collection("job")
}
