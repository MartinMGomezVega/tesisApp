package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertPublicationJob: guardo el anuncio del empleo en la bd
func InsertPublicationJob(t models.SavePublicationJob) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("AppThesis")
	col := db.Collection("jobs")

	// Anuncio/publicacion del empleo
	registerPublicationJob := bson.M{
		"userId":          t.UserID,
		"position":        t.Position,
		"company":         t.Company,
		"typeOfWorkplace": t.TypeOfWorkplace,
		"jobLocation":     t.JobLocation,
		"jobType":         t.JobType,
		"datePublication": t.DatePublication,
		"description":     t.Description,
		"finished":        t.Finished,
	}

	// Insertar la publicacion del empleo
	result, err := col.InsertOne(ctx, registerPublicationJob)
	if err != nil {
		return "", false, err
	}

	// Obtener el id de la publicacion del empleo
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
