package bd

import (
	"context"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadUsers: lee los usuarios registrados en el sistema, si se recibe "R" en quienes trae solo los que se relacionan conmigo
func ReadUsers(ID string, page int64, search string, typeSearch string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Si supera 15 segundos devuelve un timeout
	// Cancel: Cancela el timeout
	defer cancel()

	// Conectarse a la bd
	db := MongoConnect.Database("AppThesis")
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	// Si el skip es 0 me quedo donde estoy y si es 20 se posiciona en 20 y luego se setea el limite.
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	// Consulta
	// `(?i)`: no importa si son mayusculas o minusculas
	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		// fmt.Println(err.Error())
		return results, false
	}

	var found, include bool

	for cursor.Next(ctx) {
		var s models.User
		err := cursor.Decode(&s)
		if err != nil {
			return results, false
		}
		// Consultar la relacion del usuario
		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false
		found, _ = ConsultationRelation(r) // Devuelve true o false

		if typeSearch == "new" && !found {
			include = true // Es un usuario que no sigo
		}

		if typeSearch == "follow" && found {
			include = true // Es un usuario que sigo
		}

		if r.UserRelationID == ID {
			include = false // No son los mismos
		}

		if include {
			// los campos que no son necesarios los blanqueo
			s.Password = ""
			s.Biography = ""
			s.WebSite = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s) // Grabo al slice con el puntero de memoria &s
		}
	}

	err = cursor.Err()
	if err != nil {
		return results, false
	}

	cursor.Close(ctx)
	return results, true
}
