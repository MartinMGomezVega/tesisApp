package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
)

// InsertRelation: Graba la relacion en la base de datos
func InsertRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Si supera 15 segundos devuelve un timeout
	// Cancel: Cancela el timeout
	defer cancel()

	// Conectarse a la bd
	db := MongoConnect.Database("AppThesis")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
