package models

import (
	"time"
)

// SaveAPIKey: es el formato o estructura que tendrá el guardado y manejo de la key de la API de Open AI
type SaveAPIKey struct {
	UserId   string    `bson:"userId" json:"userId,omitempty"`
	Apikey   string    `bson:"apikey" json:"apikey,omitempty"`
	DateTime time.Time `bson:"datetime" json:"date,omitempty"`
}
