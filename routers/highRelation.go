package routers

import (
	"net/http"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
)

// HighRelation: realiza el registro de la relacion entre usuarios
func HighRelation(w http.ResponseWriter, r *http.Request) {
	// Obtener el id
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id.", http.StatusBadRequest) // Se debe de enviar el id
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.InsertRelation(t)
	if err != nil {
		http.Error(w, "An error occurred when trying to insert the relation.", http.StatusBadRequest) // ocurrio un error al intentar insertar la relacion
		return
	}

	if !status {
		http.Error(w, "Failed to insert the relation.", http.StatusBadRequest) // no se logro insertar la relacion
		return
	}

	w.WriteHeader(http.StatusCreated)
}
