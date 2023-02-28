package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
)

// DeleteRelation: borra la relacion en la base de datos
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Si supera 15 segundos devuelve un timeout
	// Cancel: Cancela el timeout
	defer cancel()

	// Conectarse a la bd
	db := MongoConnect.Database("AppThesis")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
