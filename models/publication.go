package models

// Publication: captura el body y el mensaje que recibe
type Publication struct {
	Message string `bson:"message" json:"message"`
}
