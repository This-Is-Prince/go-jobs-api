package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company  string             `json:"company" bson:"company"`
	Position string             `json:"position" bson:"position"`
	Status   string             `json:"status" bson:"status" default:"pending"`
}
