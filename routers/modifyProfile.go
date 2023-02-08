package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
)

// ModifyProfile: modifica el perfil de usuario
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	// Crear modelo usuario
	var t models.User

	// Grabo el body y se decodifica en el modelo usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect data: "+err.Error(), 400)
		return
	}

	var status bool
	// Resultado de la modificacion del registro
	status, err = bd.ModifyRegister(t, IDUser)
	if err != nil {
		http.Error(w, "An error occurred while trying to modify the registry. Please try again. Error: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Failed to modify the user's registration. Error: "+err.Error(), 400) // no se ha logrado modificar el registro del usuario
		return
	}

	w.WriteHeader(http.StatusCreated)
}
