package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
)

func ConsultationRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	var response models.RelationQueryResponse

	status, err := bd.ConsultationRelation(t)
	if err != nil || !status {
		response.Status = false // No hay relacion
	} else {
		response.Status = true // Hay relacion
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
