package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
)

// ValidateAPIKey: validar si la API key de OpenAI ingresada por el usuario es correcta
func ValidateAPIKey(w http.ResponseWriter, r *http.Request) {
	var request models.APIKeyValidationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid API key: "+err.Error(), http.StatusBadRequest)
		return
	}

	valid := false
	if len(request.Apikey) != 0 {
		// Realizar una llamada de prueba a la API de OpenAI
		req, err := http.NewRequest("GET", "https://api.openai.com/v1/engines", nil)
		if err != nil {
			http.Error(w, "Failed to create test request: "+err.Error(), http.StatusInternalServerError)
			return
		}
		req.Header.Set("Authorization", "Bearer "+request.Apikey)

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
		valid = resp.StatusCode == http.StatusOK
	} else {
		// Error
		if len(request.Apikey) == 0 {
			http.Error(w, "The API key entered is null.", http.StatusInternalServerError)
			return
		} else {
			http.Error(w, "The user id entered is null.", http.StatusInternalServerError)
			return
		}
	}

	response := models.APIKeyValidationResponse{Valid: valid}

	if valid {
		registerSaveAPIKey := models.SaveAPIKey{
			UserId:   IDUser,
			Apikey:   request.Apikey,
			DateTime: time.Now(),
		}

		// Guardar la clave en la BD
		_, status, err := bd.InsertAPIKeys(registerSaveAPIKey)
		if err != nil {
			// error al guardar la key
			http.Error(w, "An error occurred while saving the key. Please try again. Error: "+err.Error(), 400)
			return
		}
		if !status {
			// No se guardo la key en la bd
			http.Error(w, "Failed to save the Key.", 400)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
