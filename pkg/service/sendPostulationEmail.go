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
	fmt.Println("Email del reclutador: ", recruiterPostulation.EmailRecruiter)
	if err != nil {
		return false, err
	}

	// Configurar el cliente de correo
	d := gomail.NewDialer("smtp.your-email-provider.com", 587, "valkiria.jobs@gmail.com", "Tesis1999")

	// Construccion del cuerpo del mensaje
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
	m.SetHeader("From", "valkiria.jobs@gmail.com")         // Se envia desde el email de ValkirIA
	m.SetHeader("To", recruiterPostulation.EmailRecruiter) // Se le envia al reclutador
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Adjuntar el archivo PDF
	file, err := candidate.CV.Open()
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Leer los datos del archivo adjunto
	attachmentData, err := io.ReadAll(file)
	if err != nil {
		return false, err
	}

	// Agregar el archivo adjunto al correo
	m.Attach(candidate.CV.Filename, gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(attachmentData)
		return err
	}))

	if err := d.DialAndSend(m); err != nil {
		return false, err
	}

	return true, nil
}
