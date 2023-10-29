package service

import (
	"fmt"
	"io"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
	"gopkg.in/gomail.v2"
)

// SendPostulationEmail: Se encarga de enviar el email al reclutador con la información del postulante
func SendPostulationEmail(candidate models.PostulationJob) (bool, error) {
	// Buscar el email del usuario que postuló el empleo con el t.IdJob
	recruiterPostulation, err := bd.SearchRecruiter(candidate.IdJob)
	if err != nil {
		return false, err
	}

	// Configurar el cliente de correo
	d := gomail.NewDialer("smtp.gmail.com", 587, "valkiria.jobs@gmail.com", "zzmp qkxj nmas kubm")

	// Construcción del cuerpo del mensaje
	body := "Información del candidato:\n"
	body += "\t" + "Nombre: " + candidate.Name + " " + candidate.Surname + "\n"
	body += "\t" + "Email: " + candidate.Email + "\n"
	body += "\t" + "Código del teléfono: " + candidate.CountryCode + "\n"
	body += "\t" + "Teléfono: " + candidate.MobilePhone + "\n\n"

	body += "Descripción del candidato:\n"
	body += "\t" + candidate.Describe

	// Asunto del email
	subject := recruiterPostulation.Position + " | " + candidate.Name + " " + candidate.Surname

	m := gomail.NewMessage()
	m.SetHeader("From", "valkiria.jobs@gmail.com")         // Se envía desde el email de ValkirIA
	m.SetHeader("To", recruiterPostulation.EmailRecruiter) // Se le envía al reclutador
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Adjuntar el archivo PDF
	attachmentData := candidate.CV.Content

	// Agregar el archivo adjunto al correo
	fmt.Println("Tamaño del archivo antes de enviarlo por email: ", candidate.CV.Size)
	m.Attach(candidate.CV.Filename, gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(attachmentData)
		if err != nil {
			return err
		}
		return nil
	}))

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error DialAndSend: ", err.Error())
		return false, err
	}

	return true, nil
}
