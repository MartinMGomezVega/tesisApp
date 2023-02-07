package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/MartinMGomezVega/tesisApp/middlew"
	"github.com/MartinMGomezVega/tesisApp/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// drivers: Setear el puerto y escuchar el servidor
func Drivers() {
	router := mux.NewRouter() // Devuelve informacion del router

	router.HandleFunc("/register", middlew.CheckBD(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	//router.HandleFunc("/viewProfile", middlew.CheckBD(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")

	// abrir el puerto
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// Creando un handler/importacion
	// Los cors son quienes otorgan los permisos
	handler := cors.AllowAll().Handler(router) // Todos pueden acceder

	log.Fatal(http.ListenAndServe(":"+PORT, handler)) // Escucha el puerto
}
