package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
)

// Register: Funcion para crear en la base de datos el registro de un usuario nuevo
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	// Decodificar lo que viene del body y lo decodifica en el modelo t
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		// Error en los datos recibidos
		http.Error(w, "Error in data received: "+err.Error(), 400)
		return
	}

	// Si el mail viene vacio es un error
	if len(t.Email) == 0 {
		// Error no se ingreso un email
		http.Error(w, "Error email not entered", 400)
		return
	}

	// Si la contraseña es menor a 6 caracteres
	if len(t.Password) < 6 {
		// Error, debe de ingresar una contraseña de al menos 6 caracteres
		http.Error(w, "Error, you must enter a password of at least 6 characters.", 400)
		return
	}

	// Validacion contra los datos recibidos por el registro del usuario
	_, encontrado, _ := bd.CheckTheUserAlreadyExists(t.Email) // Chequeo de que ya existe el creado registrado
	if encontrado {
		// Error ya existe un usuario registrado con el email ingresado
		http.Error(w, "Error already exists a registered user with the entered email address", 400)
		return
	}

	_, status, err := bd.InsertRegister(t)
	if err != nil {
		// ocurrio un error al intentar realizar el registro del usuario
		http.Error(w, "An error occurred while trying to register the user: "+err.Error(), 400)
		return
	}

	if !status {
		// no se logro insertar el registro del usuario
		http.Error(w, "failed to insert user registration: "+err.Error(), 400)
		return
	}

	// Devoler el status
	w.WriteHeader(http.StatusCreated)
}
