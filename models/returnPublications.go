package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReturnPublications: es la estructura con la que se devuelven las publicaciones
type ReturnPublications struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userId" json:"userId,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
