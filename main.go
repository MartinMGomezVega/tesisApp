package main

import (
	"log"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/handlers"
)

func main() {
	// Chequear la conexión si es false no hay conexión
	if !bd.CheckConnection() {
		log.Fatal("Error, sin conexión a la base de datos")
		return
	}

	// Manejadores
	handlers.Drivers()
}
