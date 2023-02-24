package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/MartinMGomezVega/tesisApp/bd"
)

// GetBanner: envia el avatar al HTTP
func GetBanner(w http.ResponseWriter, r *http.Request) {
	// Obtener el id
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id.", http.StatusBadRequest) // Se debe de enviar el id
		return
	}

	// Obtener el perfil con el id
	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User not found.", http.StatusBadRequest) // Se debe de enviar el id
		return
	}

	// abrir el archivo/banner
	OpenFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Image not found.", http.StatusBadRequest) // Se debe de enviar el id
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Image not found.", http.StatusBadRequest) // Se debe de enviar el id
		return
	}

}
