package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("AppThesis")
	col := db.Collection("users")

	// Crear un mapa para realizar las modificaciones
	register := make(map[string]interface{}) // Permite crear slice o mapas

	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.Surname) > 0 {
		register["surname"] = u.Surname
	}
	register["dateOfBirth"] = u.DateOfBirth
	if len(u.Email) > 0 {
		register["email"] = u.Email
	}
	if len(u.Password) > 0 {
		register["password"] = u.Password
	}
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		register["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		register["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		register["webSite"] = u.WebSite
	}

	// Resigstro para actualizarlo en la bd
	updateRegister := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}} // filtro para la bd

	// UpdateOne para actualizar un dato. UpdateMany para actualizar muchos registros
	_, err := col.UpdateOne(ctx, filter, updateRegister)
	if err != nil {
		return false, err
	}

	return true, nil
}
