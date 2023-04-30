package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MartinMGomezVega/tesisApp/models"
	openai "github.com/sashabaranov/go-openai"
)

// ChatGPT: Realiza las consultas a chat gpt de OpenAI
func ChatGPT(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Estoy en Chat GPT...")
	var req models.GptRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid question or model: "+err.Error(), 400)
		return
	}

	client := openai.NewClient("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")
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
		fmt.Println("Error 2")
		http.Error(w, "Error when calling the api to create the response to the sent message: "+err.Error(), 400)
		return
	}
	// // Guardar respuesta en una variable
	// var responseText string = resp.Choices[0].Message.Content
	// fmt.Println("Response from Chat-GPT: ", responseText)

	openaiReq.Messages[0].Role = openai.ChatMessageRoleUser
	openaiReq.Messages[0].Content = resp.Choices[0].Message.Content

	fmt.Println("resp.Choices[0].Message.Content", resp.Choices[0].Message.Content+"\n")

	fmt.Println("resp: ", resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp.Choices[0].Message.Content)

}
