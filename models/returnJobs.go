package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReturnJobs: es la estructura con la que se devuelven las publicaciones
type ReturnJobs struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID          string             `bson:"userId" json:"userId,omitempty"`
	Position        string             `bson:"position" json:"position"`
	Company         string             `bson:"company" json:"company"`
	TypeOfWorkplace string             `bson:"typeOfWorkplace" json:"typeOfWorkplace"`
	JobLocation     string             `bson:"jobLocation" json:"jobLocation"`
	JobType         string             `bson:"jobType" json:"jobType"`
	Description     string             `bson:"description" json:"description"`
	DatePublication time.Time          `bson:"datePublication" json:"datePublication"`
	Finished        bool               `bson:"finished" json:"finished"`
}
