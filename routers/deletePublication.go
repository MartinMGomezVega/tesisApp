package routers

import (
	"net/http"

	"github.com/MartinMGomezVega/tesisApp/bd"
)

// DeletePublication: permite borrar una publicacion determinada
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the Id. ", http.StatusBadRequest) // Debe enviar el Id
		return
	}

	err := bd.DeletePublication(ID, IDUser)
	if err != nil {
		http.Error(w, "Error while trying to delete the publication. ", http.StatusBadRequest) // error al intentar borrar la publicacion
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
