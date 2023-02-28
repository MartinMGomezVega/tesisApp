package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ConsultationRelation: consulta la relacion entre dos usuarios
func ConsultationRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Si supera 15 segundos devuelve un timeout
	// Cancel: Cancela el timeout
	defer cancel()

	// Conectarse a la bd
	db := MongoConnect.Database("AppThesis")
	col := db.Collection("relation")

	condition := bson.M{
		"userId":         t.UserID,
		"userRelationId": t.UserRelationID,
	}

	var result models.Relation
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
