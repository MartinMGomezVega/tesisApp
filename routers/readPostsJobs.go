package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MartinMGomezVega/tesisApp/bd"
)

// leer mis publicaciones y la de todos mis seguidores
// ReadPostsJobs: leer los empleos publicados
func ReadPostsJobs(w http.ResponseWriter, r *http.Request) {
	// Paginación
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the page number. ", http.StatusBadRequest) // Debe enviar la pagina
		return
	}

	// En page se debe convertir en dato numerico:
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page == 0 {
		http.Error(w, "You must send the page number with a value greater than zero. ", http.StatusBadRequest) // Debe enviar el numero de pagina con un valor mayor a cero
		return
	}

	// Leer todas los empleos
	pageNumber := int64(page)
	response, correct := bd.ReadAllJobs(pageNumber)
	if !correct {
		// Si no es correcto leer los empleos:
		http.Error(w, "Error when reading jobs. ", http.StatusBadRequest) // Error al leer los empleos
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
