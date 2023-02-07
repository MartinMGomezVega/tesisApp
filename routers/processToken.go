package routers

import (
	"errors"
	"strings"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Variables que se usaran en otros archivos
// Email: es el valor de Email usado en todos los EndPoints
var Email string

// IDUser: es el ID devuelto del modelo, que se usa en todos los EndPoints
var IDUser string

// ProcessToken: procesa el token para extraer sus valores
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("TesisUNDAV_IngenieriaInformatica")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	// Sacar los espacios
	tk = strings.TrimSpace(splitToken[1])

	// Validacion del token
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := bd.CheckUserAlreadyExists(claims.Email)
		if found {
			// El usuario existe y el token es valido
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}

	if !tkn.Valid {
		// Token invalido
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
