package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MartinMGomezVega/tesisApp/bd"
)

// ReadPosts: leer publicaciones
func ReadPosts(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the Id: ", http.StatusBadRequest) // Debe enviar el Id
		return
	}

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
	pageNumber := int64(page)
	response, correct := bd.ReadPosts(ID, pageNumber)
	if !correct {
		// Si no es correcto leer las publicaciones:
		http.Error(w, "Error when reading publications. ", http.StatusBadRequest) // Error al leer las publicaciones
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
