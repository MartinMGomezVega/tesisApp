package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
)

// SavePublicationJob: permite guardar la publicacion del empleo en la bd
func SavePublicationJob(w http.ResponseWriter, r *http.Request) {
	var publication models.PublicationJob
	err := json.NewDecoder(r.Body).Decode(&publication)

	if err != nil {
		// error aldecodificar el cuerpo de la solicitud HTTP en el objeto estructurado
		http.Error(w, "Error when decoding the HTTP request body in the structured object. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	registerSavePublicationJob := models.SavePublicationJob{
		UserID:          IDUser, // ID del reclutador (usuario que anuncia el empleo)
		Position:        publication.Position,
		Company:         publication.Company,
		TypeOfWorkplace: publication.TypeOfWorkplace,
		JobLocation:     publication.JobLocation,
		JobType:         publication.JobType,
		DatePublication: time.Now(), // Fecha del anuncio
	}

	_, status, err := bd.InsertPublicationJob(registerSavePublicationJob)
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
