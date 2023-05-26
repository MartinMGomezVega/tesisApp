package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertAPIKeys: Insertar en la bd la Key (validada) de la API de Open AI ingresada por el usuario
func InsertAPIKeys(t models.SaveAPIKey) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("AppThesis")
	col := db.Collection("apikeys")

	// Key
	registerAPIKey := bson.M{
		"userId":   t.UserId,
		"apikey":   t.Apikey,
		"datetime": t.DateTime,
	}

	// Insertar la Key de la API de Open AI
	result, err := col.InsertOne(ctx, registerAPIKey)
	if err != nil {
		return "", false, err
	}

	// Obtener el id de la publicacion
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
