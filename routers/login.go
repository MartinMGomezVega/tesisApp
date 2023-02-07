package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/jwt"
	"github.com/MartinMGomezVega/tesisApp/models"
)

// Login: realiza el login del usuario
func Login(w http.ResponseWriter, r *http.Request) {
	// el contenido es de tipo json:
	w.Header().Add("content-type", "aplication/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid user or password: "+err.Error(), 400)
		return
	}

	// Si el mail viene vacio es un error
	if len(t.Email) == 0 {
		// Error no se ingreso un email
		http.Error(w, "Error email not entered", 400)
		return
	}

	document, exists := bd.TryLogin(t.Email, t.Password)
	// Si no existe el usuario
	if !exists {
		http.Error(w, "Invalid user or password.", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "An error occurred while generating the corresponding token: "+err.Error(), 400)
		return
	}

	// Si se genero el token se le devuelve al navegador
	response := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // devolver un valor
	json.NewEncoder(w).Encode(response)

	// Grabar el token en la cookie del usuario
	expirationTime := time.Now().Add(24 * time.Hour) // expiracion del token en 24hs
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
