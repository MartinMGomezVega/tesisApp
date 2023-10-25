package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertPostulationJob: guardo la postulacion al empleo en la bd
func InsertPostulationJob(t models.PostulationJob) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("AppThesis")
	col := db.Collection("postulations")

	// Postulacion del usuario al empleo
	registerPostulationJob := bson.M{
		"name":            t.Name,
		"surname":         t.Surname,
		"countryCode":     t.CountryCode,
		"mobilePhone":     t.MobilePhone,
		"email":           t.Email,
		"describe":        t.Describe,
		"cv":              t.CV,
		"idJob":           t.IdJob,
		"datePostulation": t.DatePostulation,
	}

	// Insertar la postulacion del usuario
	result, err := col.InsertOne(ctx, registerPostulationJob)
	if err != nil {
		return "", false, err
	}

	// Obtener el id de la postulacion del usuario
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
