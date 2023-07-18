package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/MartinMGomezVega/tesisApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetLatestAPIKeyAndDatetimeByUserID: Obtener la última API key y datetime asociados a un ID de usuario
func GetLatestAPIKeyAndDatetimeByUserID(userID string) (models.APIKeyInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var apiKeyInfo models.APIKeyInfo
	db := MongoConnect.Database("AppThesis")
	col := db.Collection("apikeys")

	// Configurar las opciones de ordenamiento para obtener la última API key cargada
	opts := options.Find().SetSort(bson.M{"datetime": -1}).SetLimit(1)

	// Realizar la consulta para obtener la última API key asociada al IDUser
	filter := bson.M{"userId": userID}
	cursor, err := col.Find(ctx, filter, opts)
	if err != nil {
		return apiKeyInfo, err
	}
	defer cursor.Close(ctx)

	// Obtener el primer documento del cursor (el último por fecha debido al ordenamiento)
	var apiKeyDoc bson.M
	if cursor.Next(ctx) {
		if err := cursor.Decode(&apiKeyInfo); err != nil {
			return apiKeyInfo, err
		}
	}

	if apiKeyDoc == nil {
		// No se encontró ninguna API key para el IDUser dado
		fmt.Println("The api key is not found with the user id: " + userID)
		return apiKeyInfo, err
	}

	// Obtener la API Key y el datetime
	apiKey := apiKeyDoc["apikey"].(string)
	datetime := apiKeyDoc["datetime"].(primitive.DateTime).Time()

	// Crear y retornar el objeto APIKeyInfo
	apiKeyInfo = models.APIKeyInfo{
		APIKey:   apiKey,
		DateTime: datetime,
	}

	fmt.Println("API Key obtenida: ", apiKey)
	fmt.Println("DateTime: ", datetime)
	return apiKeyInfo, nil
}
