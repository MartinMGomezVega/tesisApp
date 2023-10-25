package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchRecruiter: Busca un empleo en la bd
func SearchRecruiter(IdJob string) (models.PublicationJob, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnect.Database("AppThesis")
	col := db.Collection("jobs")

	var profile models.PublicationJob
	objID, _ := primitive.ObjectIDFromHex(IdJob)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	if err != nil {
		fmt.Println("Register not found " + err.Error()) // No se encontro el registro
		return profile, err
	}
	return profile, nil
}
