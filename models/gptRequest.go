package models

// ChatGPTRequest: formato de como es el request para obtener datos desde la API de OpenAI
type ChatGPTRequest struct {
	Question string `json:"question"`
}
