package bd

import (
	"context"
	"log"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadPosts: se leen publicaciones paginadas de un perfil
func ReadPosts(ID string, page int64) ([]*models.ReturnPublications, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnect.Database("AppThesis")
	col := db.Collection("publication")

	var results []*models.ReturnPublications
	condition := bson.M{
		"userId": ID,
	}

	// En MongoDB existen las opciones que son para obtener documentos
	options := options.Find()
	options.SetLimit(20)                              // Cuantos documentos trae
	options.SetSort(bson.D{{Key: "date", Value: -1}}) // Traer por fecha en orden descendente
	options.SetSkip((page - 1) * 20)                  // Cuantos documentos hay que saltear (es el limite que se debe de saltear)

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	// Contexto nuevo y vacio para no mezclar con el ctx
	for cursor.Next(context.TODO()) {
		// Por cada iteracion crea una nueva variable de registro
		var register models.ReturnPublications
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		// Si no hay errores agrega el registro al slice results
		results = append(results, &register)
	}

	return results, true
}
