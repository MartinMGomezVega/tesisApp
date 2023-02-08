package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MartinMGomezVega/tesisApp/bd"
)

// ViewProfile: permite extraer los valores del perfil
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the parameter", http.StatusBadRequest) // Debe de enviar el parametro
		return
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "An error occurred while trying to search for the record "+err.Error(), 400) // ocurrio un error al intentar buscar el registro
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
