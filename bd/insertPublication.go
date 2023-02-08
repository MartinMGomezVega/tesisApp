package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertPublication: guardo la publicacion en la bd
func InsertPublication(t models.SavePublication) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("AppThesis")
	col := db.Collection("publication")

	// Publicacion
	registerPublication := bson.M{
		"userId":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	// Insertar la publicacion
	result, err := col.InsertOne(ctx, registerPublication)
	if err != nil {
		return "", false, err
	}

	// Obtener el id de la publicacion
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
