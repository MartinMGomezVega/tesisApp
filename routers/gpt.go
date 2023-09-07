package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
	openai "github.com/sashabaranov/go-openai"
)

// ChatGPT: Realiza las consultas a chat gpt de OpenAI
func ChatGPT(w http.ResponseWriter, r *http.Request) {
	var req models.ChatGPTRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid question: "+err.Error(), 400)
		return
	}

	// fmt.Println("api_key:", req.APIKey)
	fmt.Println("Question:", req.Question)
	fmt.Println("id user:", IDUser)
	fmt.Println("len(IDUser):", len(IDUser))

	if len(IDUser) == 0 {
		http.Error(w, "IDUser is null: "+err.Error(), 400)
		return
	}

	// Obtener la api key de la base de datos con el Id del usuario
	fmt.Println("Obtener la Api Key y el Datetime para el usuario: ", IDUser)
	apiKeyInfo, err := bd.GetLatestAPIKeyAndDatetimeByUserID(IDUser)
	if err != nil {
		http.Error(w, "Error obtaining api key: "+err.Error(), 400)
		return
	}

	fmt.Println("apiKeyInfo.APIKey:", apiKeyInfo.APIKey)
	fmt.Println("apiKeyInfo.DateTime:", apiKeyInfo.DateTime)

	// Si la diferencia de tiempo del DateTime es mayor a 20 minutos, validar nuevamente la Key
	timeDiff := time.Since(apiKeyInfo.DateTime)
	validateKey := false
	if timeDiff.Minutes() > 20 {
		// Validar la api Key
		fmt.Println("Fue creada hace", timeDiff.Minutes(), "minutos")
		// Realizar una llamada de prueba a la API de OpenAI
		req, err := http.NewRequest("GET", "https://api.openai.com/v1/engines", nil)
		if err != nil {
			http.Error(w, "Failed to create test request: "+err.Error(), http.StatusInternalServerError)
			return
		}
		req.Header.Set("Authorization", "Bearer "+apiKeyInfo.APIKey)

		client := http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			// Si hay un error en la llamada, se considera que la clave no es válida
			response := models.APIKeyValidationResponse{Valid: false}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
		defer resp.Body.Close()
		// Leer la respuesta de la llamada de prueba
		// Si el estado de respuesta es 200, se considera que la clave es válida
		validateKey = resp.StatusCode == http.StatusOK
		fmt.Println("Status Api Key (resp.StatusCode == http.StatusOK): ", resp.StatusCode == http.StatusOK)
	} else {
		// Al ser menor la diferencia de tiempo del DateTime a 20 minutos, se considera valida
		validateKey = true
	}

	// Al ser valida, realizar la consulta a chat gpt
	var response string
	var errorText string
	if validateKey {
		fmt.Println("Status Api Key: ", validateKey)

		client := openai.NewClient(apiKeyInfo.APIKey)
		ctx := context.Background()

		openaiReq := openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: req.Question,
				},
			},
		}

		resp, err := client.CreateChatCompletion(ctx, openaiReq)
		if err != nil {
			errorText = "Error when calling the api to create the response to the sent message: " + err.Error()
			if strings.Contains(errorText, "You exceeded your current quota") {
				errorText = "Ha superado su cuota actual, por favor compruebe los detalles de su plan y facturación en OpenAI."
			} else {
				http.Error(w, errorText, 400)
				return
			}
		}

		// openaiReq.Messages[0].Role = openai.ChatMessageRoleUser
		// openaiReq.Messages[0].Content = resp.Choices[0].Message.Content
		// fmt.Println("resp.Choices[0].Message.Content", resp.Choices[0].Message.Content+"\n")
		if strings.Contains(errorText, "Ha superado su cuota actual") {
			response = "Ha superado su cuota actual, por favor compruebe los detalles de su plan y facturación en OpenAI."
		} else {
			response = resp.Choices[0].Message.Content
		}

	} else {
		response = "Key not valid"
	}

	fmt.Println("Respuesta: ", response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
