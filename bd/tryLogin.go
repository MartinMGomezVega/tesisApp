package bd

import (
	"github.com/MartinMGomezVega/tesisApp/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin: realiza el chequeo de login a la bd
func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := CheckUserAlreadyExists(email)

	// Si no fue encontrado
	if !found {
		return user, false
	}

	passwordBytes := []byte(password)   // Pass no encriptada
	passwordBd := []byte(user.Password) // pass encriptada de la bd
	err := bcrypt.CompareHashAndPassword(passwordBd, passwordBytes)

	// Si falla es porque no coinciden
	if err != nil {
		return user, false
	}
	return user, true
}
