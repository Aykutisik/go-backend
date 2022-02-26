package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoElements struct {
	Id     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Text   string             `json:"text"`
	Status int                `json:"status"`
}
