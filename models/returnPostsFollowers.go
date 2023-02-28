package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReturnPostsFollowers: tiene el token que se devuelve con el login
type ReturnPostsFollowers struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userId" json:"userId,omitempty"`
	UserRelationID string             `bson:"userRelationId" json:"userRelationId,omitempty"`
	Posts          struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
