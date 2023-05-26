package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
)

// SavePublication: permite guardar la publicacion en la bd
func SavePublication(w http.ResponseWriter, r *http.Request) {
	var message models.Publication
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		// error aldecodificar el cuerpo de la solicitud HTTP en el objeto estructurado
		http.Error(w, "Error when decoding the HTTP request body in the structured object. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	registerSavePublication := models.SavePublication{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertPublication(registerSavePublication)
	if err != nil {
		// error al guardar la publicacion
		http.Error(w, "An error occurred while saving the publication. Please try again. Error: "+err.Error(), 400)
		return
	}

	if !status {
		// No se guardo la publicacion en la bd
		http.Error(w, "Failed to save the publication.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
