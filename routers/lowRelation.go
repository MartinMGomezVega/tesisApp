package routers

import (
	"net/http"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
)

// LowRelation: realiza el borrado de la relacion entre usuarios
func LowRelation(w http.ResponseWriter, r *http.Request) {
	// Obtener el id
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.DeleteRelation(t)
	if err != nil {
		http.Error(w, "An error occurred when trying to delete the relation.", http.StatusBadRequest) // ocurrio un error al intentar insertar la relacion
		return
	}

	if !status {
		http.Error(w, "Failed to delete the relation.", http.StatusBadRequest) // no se logro insertar la relacion
		return
	}

	w.WriteHeader(http.StatusCreated)
}
