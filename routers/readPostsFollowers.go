package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MartinMGomezVega/tesisApp/bd"
)

// ReadPostsFollowers: leer las publicaciones de todos mis seguidores
func ReadPostsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the page number. ", http.StatusBadRequest) // Debe enviar la pagina
		return
	}

	// En page se debe convertir en dato numerico:
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the page number with a value greater than zero. ", http.StatusBadRequest) // Debe enviar el numero de pagina con un valor mayor a cero
		return
	}

	response, correct := bd.ReadPostsFollowers(IDUser, page)
	if !correct {
		// Si no es correcto leer las publicaciones:
		http.Error(w, "Error when reading publications. ", http.StatusBadRequest) // Error al leer las publicaciones
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
