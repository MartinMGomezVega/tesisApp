package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertAPIKeys: Insertar en la bd la Key (validada) de la API de Open AI ingresada por el usuario
func InsertAPIKeys(t models.SaveAPIKey) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("AppThesis")
	col := db.Collection("apikeys")

	// Verificar si ya existe una API Key con el mismo valor
	existingKeyFilter := bson.M{"apikey": t.Apikey}
	existingKey := col.FindOne(ctx, existingKeyFilter)
	if existingKey.Err() == nil {
		var existingDoc bson.M
		if err := existingKey.Decode(&existingDoc); err != nil {
			return "Error decoding api key.", false, err
		}
		objID := existingDoc["_id"].(primitive.ObjectID)
		return objID.String(), true, nil
	} else if existingKey.Err() != mongo.ErrNoDocuments {
		return "No api key found.", false, existingKey.Err()
	}

	// Key
	registerAPIKey := bson.M{
		"userId":   t.UserId,
		"apikey":   t.Apikey,
		"datetime": t.DateTime,
	}

	// Insertar la Key de la API de Open AI
	result, err := col.InsertOne(ctx, registerAPIKey)
	if err != nil {
		return "Error inserting api key.", false, err
	}

	// Obtener el id de la key
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
