package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	mongoURI := os.Getenv("MONGO_URI")

	clientOptions := options.Client().ApplyURI(mongoURI)

	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error connecting to mongodb")
	}
	log.Println("MongoDB Connection Successful....")
}
