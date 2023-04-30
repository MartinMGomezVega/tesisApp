package models

// GptRequest: formato de como es el request para obtener datos desde la API de OpenAI
type GptRequest struct {
	Model    string `json:"model"`
	Question string `json:"question"`
}
