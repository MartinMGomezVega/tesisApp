package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserAlreadyExists: recibe un email de parametro y chequea si ya existe en la base de datos
func CheckUserAlreadyExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // Cancelar cuando termina la busqueda en la bd

	db := MongoConnect.Database("AppThesis")
	col := db.Collection("users")

	// Find en MongoDB
	condition := bson.M{"email": email} // Devuelve un formato json y se le tiene que enviar clave valor

	var result models.User

	// Busqueda de un solo registro
	err := col.FindOne(ctx, condition).Decode(&result) // lo convierto en json en la variable result
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
