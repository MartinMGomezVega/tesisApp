package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
)

// UploadAvatar: Subir el avatar a la bd y servidor
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1] // El elemento 1 es un string
	var archivo string = "uploads/avatars/" + IDUser + "." + extension

	// Obtener el archivo/avatar
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error when uploading avatar: "+err.Error(), http.StatusBadRequest) // error al subir el avatar
		return
	}

	// Copiar el archivo/avatar
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error when copying avatar: "+err.Error(), http.StatusBadRequest) // error al copiar el avatar
		return
	}

	var user models.User
	var status bool

	// Guardar el avatar en la base de datos
	user.Avatar = IDUser + "." + extension
	status, err = bd.ModifyRegister(user, IDUser)
	if err != nil || !status {
		http.Error(w, "Error when saving the avatar in the database: "+err.Error(), http.StatusBadRequest) // error al guardar el avatar en la bd
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
