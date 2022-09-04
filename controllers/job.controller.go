package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/This-Is-Prince/go-jobs-api/db"
	"github.com/This-Is-Prince/go-jobs-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllJobs(w http.ResponseWriter, r *http.Request) {
	JobCollection := db.JobCollection()
	cursor, err := JobCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var allJobs []primitive.M
	for cursor.Next(context.Background()) {
		var job primitive.M
		if err := cursor.Decode(&job); err != nil {
			log.Fatal(err)

		}
		allJobs = append(allJobs, job)
	}
	json.NewEncoder(w).Encode(allJobs)
}
func CreateJob(w http.ResponseWriter, r *http.Request) {
	JobCollection := db.JobCollection()

	var job models.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		log.Fatal(err)
	}
	job.Status = "pending"
	fmt.Println(job)
	defer r.Body.Close()
	result, err := JobCollection.InsertOne(context.Background(), job)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}
func DeleteAllJobs(w http.ResponseWriter, r *http.Request) {
	JobCollection := db.JobCollection()
	result, err := JobCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}
