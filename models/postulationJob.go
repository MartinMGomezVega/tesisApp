package models

import (
	"mime/multipart"
	"time"
)

// PostulationJob: captura el body y el mensaje con los campos respectivos de la postulacion al empleo
type PostulationJob struct {
	Name            string                `bson:"name" json:"name"`
	Surname         string                `bson:"surname" json:"surname"`
	CountryCode     string                `bson:"countryCode" json:"countryCode"`
	MobilePhone     string                `bson:"mobilePhone" json:"mobilePhone"`
	Email           string                `bson:"email" json:"email"`
	Describe        string                `bson:"describe" json:"describe"`
	CV              *multipart.FileHeader `bson:"cv" json:"cv"`
	IdJob           string                `bson:"idJob" json:"idJob"`
	DatePostulation time.Time             `bson:"datePostulation" json:"datePostulation"`
}
