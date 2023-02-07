package jwt

import (
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT: genera el encriptado con JWT
func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("MastersdelDesarrollo_grupodeFacebook")

	// Lista de privilegios
	payload := jwt.MapClaims{
		"email":      t.Email,
		"name":       t.Name,
		"surname":    t.Surname,
		"dateOfBith": t.DateOfBirth,
		"biography":  t.Biography,
		"location":   t.Location,
		"webSite":    t.WebSite,
		"_id":        t.ID.Hex(),
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
