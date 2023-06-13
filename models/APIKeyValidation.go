package models

// APIKeyValidationRequest: Key de la API de Open AI
type APIKeyValidationRequest struct {
	Apikey string `bson:"apikey" json:"apikey"`
}

// APIKeyValidationResponse: True o False, si es valida o no la key
type APIKeyValidationResponse struct {
	Valid bool `json:"valid"`
}
