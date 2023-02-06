package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertRegister: es la parada final con la bd para insertar los datos del usuario
// El string de devolucion es un id
func InsertRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Si supera 15 segundos devuelve un timeout
	// Cancel: Cancela el timeout
	defer cancel()

	// Conectarse a la bd
	db := MongoConnect.Database("AppThesis")
	col := db.Collection("users")

	// Grabar la password encriptada en la bd
	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil
}
