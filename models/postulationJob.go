package models

import (
	"time"
)

type Attachment struct {
	Filename string
	Content  []byte
}

// PostulationJob: captura el body y el mensaje con los campos respectivos de la postulacion al empleo
type PostulationJob struct {
	Name            string     `bson:"name" json:"name"`
	Surname         string     `bson:"surname" json:"surname"`
	CountryCode     string     `bson:"countryCode" json:"countryCode"`
	MobilePhone     string     `bson:"mobilePhone" json:"mobilePhone"`
	Email           string     `bson:"email" json:"email"`
	Describe        string     `bson:"describe" json:"describe"`
	CV              Attachment `bson:"cv" json:"cv"`
	IdJob           string     `bson:"idJob" json:"idJob"`
	DatePostulation time.Time  `bson:"datePostulation" json:"datePostulation"`
}
