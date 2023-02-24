package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeletePublication: borra una publicacion determinada
func DeletePublication(ID string, UserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("AppThesis")
	col := db.Collection("publication")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id":    objID,
		"userId": UserId,
	}

	_, err := col.DeleteOne(ctx, condition) // Borrar publicacion
	return err
}
