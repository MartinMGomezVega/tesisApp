package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
	"github.com/MartinMGomezVega/tesisApp/pkg/service"
)

// SavePostulationJob: permite guardar la postulacion al empleo en la bd
func SavePostulationJob(w http.ResponseWriter, r *http.Request) {
	var postulation models.PostulationJob
	err := json.NewDecoder(r.Body).Decode(&postulation)

	if err != nil {
		// error aldecodificar el cuerpo de la solicitud HTTP en el objeto estructurado
		http.Error(w, "Error when decoding the HTTP request body in the structured object. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	registerSavePostulationJob := models.PostulationJob{
		Name:            postulation.Name,
		Surname:         postulation.Surname,
		CountryCode:     postulation.CountryCode,
		MobilePhone:     postulation.MobilePhone,
		Email:           postulation.Email,
		Describe:        postulation.Describe,
		CV:              postulation.CV,
		IdJob:           postulation.IdJob,
		DatePostulation: time.Now(), // Fecha de la postulacion
	}

	_, status, err := bd.InsertPostulationJob(registerSavePostulationJob)
	if err != nil {
		// error al guardar la postulacion
		http.Error(w, "An error occurred while saving the postulation. Please try again. Error: "+err.Error(), 400)
		return
	}

	// Enviar mail al usuario que postulo el empleo
	statusSendEmail, err := service.SendPostulationEmail(registerSavePostulationJob)
	if err != nil {
		// error al enviar la postulacion por email al usuario que publico el empleo
		http.Error(w, "Error sending the application by email to the user who posted the job: "+err.Error(), 400)
		return
	} else {
		if statusSendEmail {
			fmt.Println("Email sent successfully")
		}
	}

	if !status {
		// No se guardo la publicacion en la bd
		http.Error(w, "Failed to save the postulation.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
