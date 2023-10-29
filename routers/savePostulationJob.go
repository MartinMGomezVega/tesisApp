package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
	"github.com/MartinMGomezVega/tesisApp/pkg/service"
)

// SavePostulationJob: permite guardar la postulación al empleo en la bd
func SavePostulationJob(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(20 << 20) // Aumentar el límite a 20 MB (20 * 1024 * 1024 bytes)

	if err != nil {
		http.Error(w, "Error when parsing the form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener el archivo adjunto (currículum)
	cvFile, header, err := r.FormFile("cv")
	if err != nil {
		http.Error(w, "Error al obtener el currículum: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer cvFile.Close()
	// fmt.Println("cv nombre:", header.Filename)
	// fmt.Println("cv tamaño:", header.Size)

	// Leer el contenido del archivo en un slice de bytes
	cvData, err := io.ReadAll(cvFile)
	if err != nil {
		http.Error(w, "Error when reading the curriculum vitae: "+err.Error(), http.StatusBadRequest)
		return
	}

	attachment := models.Attachment{
		Filename: header.Filename,
		Size:     header.Size,
		Content:  cvData,
	}

	postulation := models.PostulationJob{
		Name:            r.FormValue("name"),
		Surname:         r.FormValue("surname"),
		CountryCode:     r.FormValue("countryCode"),
		MobilePhone:     r.FormValue("mobilePhone"),
		Email:           r.FormValue("email"),
		Describe:        r.FormValue("describe"),
		CV:              attachment, // Asignar el objeto Attachment
		IdJob:           r.FormValue("idJob"),
		DatePostulation: time.Now(), // Fecha de la postulación
	}

	// Guardar la postulación en la base de datos
	_, err = bd.InsertPostulationJob(postulation)
	if err != nil {
		http.Error(w, "Error saving the application:"+err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar mail al usuario que postuló el empleo
	statusSendEmail, err := service.SendPostulationEmail(postulation)
	if err != nil {
		http.Error(w, "Error sending the application by email to the user who posted the job: "+err.Error(), http.StatusBadRequest)
		return
	} else {
		if statusSendEmail {
			fmt.Println("Email sent successfully")
		}
	}

	// Enviar una respuesta JSON de éxito
	response := map[string]string{"message": "Postulation saved successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
