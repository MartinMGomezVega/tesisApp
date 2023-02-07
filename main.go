package main

import (
	"log"

	"github.com/MartinMGomezVega/tesisApp/bd"
	"github.com/MartinMGomezVega/tesisApp/handlers"
)

func main() {
	// Chequear la conexión si es false no hay conexión
	if !bd.CheckConnection() {
		log.Fatal("No database connection error")
		return
	}

	// Manejadores
	handlers.Drivers()
}
