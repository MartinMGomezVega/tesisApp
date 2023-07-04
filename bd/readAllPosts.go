package bd

import (
	"context"
	"log"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadAllPosts: leer todas las publicaciones
func ReadAllPosts(userID string, page int64) ([]*models.ReturnPublications, bool) {
	// Establecer contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Conectar a la base de datos y seleccionar la colección
	db := MongoConnect.Database("AppThesis")
	col := db.Collection("publication")

	// Obtener los seguidores del usuario
	idFollowers := GetFollowed(userID)
	idFollowers = append(idFollowers, userID) // Incluir al usuario en la lista de seguidores
	// fmt.Println("Seguidores: ", idFollowers)

	// Definir los criterios de búsqueda para las publicaciones de los seguidores
	var results []*models.ReturnPublications
	condition := bson.M{
		"userId": bson.M{"$in": idFollowers}, // Buscar por los idFollowers
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
