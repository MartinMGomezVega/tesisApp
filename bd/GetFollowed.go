package bd

import (
	"context"
	"log"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetFollowed: obtener las personas a las que sigues
func GetFollowed(userID string) []string {
	// Establecer contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Conectar a la base de datos y seleccionar la colección
	db := MongoConnect.Database("AppThesis")
	col := db.Collection("relation")

	// Definir los criterios de búsqueda
	condition := bson.M{
		"userId": userID,
	}

	var results []string

	cursor, err := col.Find(ctx, condition)
	if err != nil {
		log.Fatal(err.Error())
		return results
	}

	// Contexto nuevo y vacío para no mezclar con el ctx
	for cursor.Next(context.TODO()) {
		// Por cada iteración, crea una nueva variable de registro
		var register models.Relation
		err := cursor.Decode(&register)
		if err != nil {
			return results
		}
		// Si no hay errores, agrega el userRelationId al slice results
		results = append(results, register.UserRelationID)
	}

	return results
}
