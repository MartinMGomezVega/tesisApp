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
	err := r.ParseMultipartForm(10 << 20) // 10 MB max size
	fmt.Println("Entro!")

	if err != nil {
		fmt.Println("Error:", err.Error())
		http.Error(w, "Error when parsing the form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Name:", r.FormValue("name"))
	fmt.Println("Surname:", r.FormValue("surname"))
	fmt.Println("countryCode:", r.FormValue("countryCode"))
	fmt.Println("mobilePhone:", r.FormValue("mobilePhone"))
	fmt.Println("email:", r.FormValue("email"))
	fmt.Println("describe:", r.FormValue("describe"))
	fmt.Println("idJob:", r.FormValue("idJob"))

	// Obtener el archivo adjunto (currículum)
	cvFile, header, err := r.FormFile("cv")
	if err != nil {
		fmt.Println("Error:", err.Error())
		http.Error(w, "Error al obtener el currículum: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer cvFile.Close()
	fmt.Println("cv nombre:", header.Filename)
	fmt.Println("cv tamaño:", header.Size)

	// Leer el contenido del archivo en un slice de bytes
	cvData, err := io.ReadAll(cvFile)
	if err != nil {
		fmt.Println("Error al leer el currículum: ", err)
		http.Error(w, "Error al leer el currículum: "+err.Error(), http.StatusBadRequest)
		return
	}

	attachment := models.Attachment{
		Filename: header.Filename,
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
		http.Error(w, "Error al guardar la postulación: "+err.Error(), http.StatusInternalServerError)
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
