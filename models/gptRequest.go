package models

// ChatGPTRequest: formato de como es el request para obtener datos desde la API de OpenAI
type ChatGPTRequest struct {
	APIKey   string `json:"apikey"`
	Question string `json:"question"`
}
