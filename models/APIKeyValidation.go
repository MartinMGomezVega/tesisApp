package models

// APIKeyValidationRequest: Key de la API de Open AI
type APIKeyValidationRequest struct {
	APIKey string `json:"api_key"`
}

// APIKeyValidationResponse: True o False, si es valida o no la key
type APIKeyValidationResponse struct {
	Valid bool `json:"valid"`
}
