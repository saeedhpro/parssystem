package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Automobile struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Manufacture string             `bson:"manufacture,omitempty"`
	Type        string             `bson:"type,omitempty"`
	Model       int                `bson:"model,omitempty"`
	Parts       []Part             `bson:"parts,omitempty"`
}
