package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID     primitive.ObjectID `bson: "_id, omitempty" json: "_id, omitempty"`
	Task   string             `json: "task, omitempty"`
	Status bool               `json: "status, omitempty"`
}
