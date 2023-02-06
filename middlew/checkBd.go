package middlew

import (
	"net/http"

	"github.com/MartinMGomezVega/tesisApp/bd"
)

// Los middlewares tienen que recibir y dar datos del mismo tipo, sino no es un pasa manos.
// CheckBD: es el middleware que permite conocer el estado de la base de datos
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.CheckConnection() {
			http.Error(w, "Lost connection to the database", 500) // Falla en la conexion con la base de datos
			return
		}
		next.ServeHTTP(w, r) // Si no falla le paso todos los datos
	}
}
