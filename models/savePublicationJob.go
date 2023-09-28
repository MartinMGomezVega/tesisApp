package models

import (
	"time"
)

// SavePublicationJob: es el formato o estructura que tendrá la publicación del empleo en la bd
type SavePublicationJob struct {
	UserID          string    `bson:"userID" json:"userID"`
	Position        string    `bson:"position" json:"position"`
	Company         string    `bson:"company" json:"company"`
	TypeOfWorkplace string    `bson:"typeOfWorkplace" json:"typeOfWorkplace"`
	JobLocation     string    `bson:"jobLocation" json:"jobLocation"`
	JobType         string    `bson:"jobType" json:"jobType"`
	DatePublication time.Time `bson:"datePublication" json:"datePublication"`
}
