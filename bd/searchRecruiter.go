package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchRecruiter: Busca una postulación de empleo en la bd y devuelve el documento completo
func SearchRecruiter(IdJob string) (models.PublicationJob, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("AppThesis")
	col := db.Collection("jobs")

	var job models.PublicationJob
	objID, err := primitive.ObjectIDFromHex(IdJob)
	if err != nil {
		return job, err
	}

	condition := bson.M{
		"_id": objID,
	}

	err = col.FindOne(ctx, condition).Decode(&job)
	if err != nil {
		return job, err
	}

	return job, nil
}
