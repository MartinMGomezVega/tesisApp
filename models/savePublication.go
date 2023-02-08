package models

import (
	"time"
)

// SavePublication: es el formato o estructura que tendrá la publicación en la bd
type SavePublication struct {
	UserID  string    `bson:"userID" json:"userID,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
