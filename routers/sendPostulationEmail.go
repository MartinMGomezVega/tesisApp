package routers

import (
	"fmt"
	"io"
	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
	"gopkg.in/gomail.v2"
)

func SendPostulationEmail(t models.PostulationJob) (bool, error) {
	// Buscar el email del usuario que postuló el empleo con el t.IdJob
	toEmail := bd.SearchRecruiter(t.IdJob)

	// Configurar el cliente de correo valkiria.jobs@gmail.com
	d := gomail.NewDialer("smtp.your-email-provider.com", 587, "valkiria.jobs@gmail.com", "Tesis1999")

	body := t.Describe
	subject := "titulo del puesto " + " | " + t.Name + " " + t.Surname

	m := gomail.NewMessage()
	m.SetHeader("From", "valkiria.jobs@gmail.com")
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Adjuntar el archivo CV
	if err := m.Attach(t.CV.Filename, gomail.SetCopyFunc(func(w *gomail.Writer))){
		file, err := t.CV.Open()
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = fmt.Fprintf(w, "Content-Type: %s\nContent-Disposition: attachment; filename=\"%s\"\r\n\r\n", t.CV.Header.Get("Content-Type"), t.CV.Filename)
		if err != nil {
			return false, err
		}
		_, err = io.Copy(w, file)
		return false, err
	}

	// Manejar el error si ocurrió al adjuntar el archivo
	if err != nil {
		return false, err
	}

	if err := d.DialAndSend(m); err != nil {
		return false, err
	}

	return true, nil
}
