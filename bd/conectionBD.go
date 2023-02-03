package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnect = conectBD()

// clientOptions: URL de la base de datos
var clientOptions = options.Client().ApplyURI("mongodb://MartinGomezVega:6hd3TlXoIbJpolfg@ac-jkru0ds-shard-00-00.3ph7abi.mongodb.net:27017,ac-jkru0ds-shard-00-01.3ph7abi.mongodb.net:27017,ac-jkru0ds-shard-00-02.3ph7abi.mongodb.net:27017/test?replicaSet=atlas-n1jjt3-shard-0&ssl=true&authSource=Cluster-Tesis&authMechanism=SCRAM-SHA-1")

// conectBD: Conexion a la base de datos
func conectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Si hay error en la conexion:
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	// Saber si la base de datos está ON
	err = client.Ping(context.TODO(), nil) // Hace otro tipo de comprobaciones
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("¡Conexión exitosa a la base de datos de MongoDB!")
	return client
}

// CheckConnection: checkeo de la conexión a la base de datos
func CheckConnection() bool {
	err := MongoConnect.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		// con error
		return false
	}
	// Sin error
	return true
}
