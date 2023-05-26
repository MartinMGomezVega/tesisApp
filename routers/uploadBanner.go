package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
)

// UploadBanner: Subir el banner a la bd y servidor
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "Error getting the banner: "+err.Error(), http.StatusBadRequest) // error al subir el banner
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1] // El elemento 1 es un string
	var archivo string = "uploads/banners/" + IDUser + "." + extension

	// Obtener el archivo/banner
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error when uploading banner: "+err.Error(), http.StatusBadRequest) // error al subir el banner
		return
	}

	// Copiar el archivo/banner
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error when copying banner: "+err.Error(), http.StatusBadRequest) // error al copiar el banner
		return
	}

	var user models.User
	var status bool

	// Guardar el banner en la base de datos
	user.Banner = IDUser + "." + extension
	status, err = bd.ModifyRegister(user, IDUser)
	if err != nil || !status {
		http.Error(w, "Error when saving the banner in the database: "+err.Error(), http.StatusBadRequest) // error al guardar el banner en la bd
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
