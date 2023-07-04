package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadPostsFollowers: leer las publicaciones de mi seguidores
func ReadPostsFollowers(ID string, page int) ([]models.ReturnPostsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnect.Database("AppThesis")
	col := db.Collection("relation")

	// Paginacion de a 20 resultados
	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userId": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "publication",
			"localField":   "userRelationId",
			"foreignField": "userId",
			"as":           "publication",
		}})

	// Obtener los resultados para poderlos procesar: $unwind -> permite que todos los documentos vengan iguales
	conditions = append(conditions, bson.M{"$unwind": "$publication"})
	// Obtener los datos ordenados de forma descendente por fecha
	conditions = append(conditions, bson.M{"$sort": bson.M{"date": -1}})
	// Cuantos registros tiene que saltear la consulta
	conditions = append(conditions, bson.M{"$skip": skip})
	// limite de pagina
	conditions = append(conditions, bson.M{"$limit": 20})

	var result []models.ReturnPostsFollowers
	// Framework de MongoDB: aggregate
	cursor, err := col.Aggregate(ctx, conditions)
	if err != nil {
		return result, false
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true
}
