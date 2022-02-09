package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Part struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Files []string           `bson:"files,omitempty"`
}
