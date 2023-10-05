package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadAllJobs: leer todos los empleos
func ReadAllJobs(page int64) ([]*models.ReturnJobs, bool) {
	// Establecer contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Conectar a la base de datos y seleccionar la colección
	db := MongoConnect.Database("AppThesis")
	col := db.Collection("jobs")

	// Definir los criterios de búsqueda
	var results []*models.ReturnJobs
	// Posteos de empleos que no esten finalizados
	condition := bson.M{"finished": false}

	// En MongoDB existen las opciones que son para obtener documentos
	findOptions := options.Find()
	findOptions.SetLimit(20)                              // Cuantos documentos trae
	findOptions.SetSort(bson.D{{Key: "date", Value: -1}}) // Traer por fecha en orden descendente
	findOptions.SetSkip((page - 1) * 20)                  // Cuantos documentos hay que saltear (es el limite que se debe de saltear)

	cursor, err := col.Find(ctx, condition, findOptions)
	if err != nil {
		// log.Fatal(err.Error())
		fmt.Println(err.Error())
		return results, false
	}

	// Contexto nuevo y vacio para no mezclar con el ctx
	for cursor.Next(context.TODO()) {
		// Por cada iteracion crea una nueva variable de registro
		var register models.ReturnJobs
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		// Si no hay errores agrega el registro al slice results
		results = append(results, &register)
	}

	return results, true
}
