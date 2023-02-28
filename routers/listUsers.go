package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MartinMGomezVega/tesisApp/bd"
)

// ListUsers: lee la lista de los usuarios
func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemporal, err := strconv.Atoi(page) // Convierte string a int
	if err != nil {
		http.Error(w, "You must send the page parameter as an integer greater than zero.", http.StatusBadRequest) // Debe enviar el parametro pagina como entero mayor a cero.
		return
	}

	pag := int64(pageTemporal)

	// Leer los usuarios
	result, status := bd.ReadUsers(IDUser, pag, search, typeUser)
	if !status {
		http.Error(w, "Error when reading users.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
