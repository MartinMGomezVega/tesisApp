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

// drivers: Setear el puerto y escuchar el servidor (se encarga de manejar las solicitudes HTTP que llegan al servidor)
func Drivers() {
	router := mux.NewRouter() // Devuelve informacion del router
	// Registro
	router.HandleFunc("/register", middlew.CheckBD(routers.Register)).Methods("POST")
	// Login
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	// Ver perfil
	router.HandleFunc("/viewProfile", middlew.CheckBD(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	// Modificar perfil
	router.HandleFunc("/modifyProfile", middlew.CheckBD(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	// Guardar publicacion
	router.HandleFunc("/savePublication", middlew.CheckBD(middlew.ValidateJWT(routers.SavePublication))).Methods("POST")
	// Leer publicaciones
	router.HandleFunc("/readPosts", middlew.CheckBD(middlew.ValidateJWT(routers.ReadPosts))).Methods("GET")
	// Borrar publicaciones
	router.HandleFunc("/deletePublication", middlew.CheckBD(middlew.ValidateJWT(routers.DeletePublication))).Methods("DELETE")
	// Subir avatar
	router.HandleFunc("/uploadAvatar", middlew.CheckBD(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	// Obtener avatar
	router.HandleFunc("/getAvatar", middlew.CheckBD(routers.GetAvatar)).Methods("GET")
	// Subir banner
	router.HandleFunc("/uploadBanner", middlew.CheckBD(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	// Obtener banner
	router.HandleFunc("/getBanner", middlew.CheckBD(routers.GetBanner)).Methods("GET")
	// Alta de la relacion (seguir)
	router.HandleFunc("/highRelation", middlew.CheckBD(middlew.ValidateJWT(routers.HighRelation))).Methods("POST")
	// Baja de la relacion (dejar de seguir)
	router.HandleFunc("/lowRelation", middlew.CheckBD(middlew.ValidateJWT(routers.LowRelation))).Methods("DELETE")
	// Consultar sobre la relacion entre dos personas
	router.HandleFunc("/consultationRelation", middlew.CheckBD(middlew.ValidateJWT(routers.ConsultationRelation))).Methods("GET")
	// Listar usuarios
	router.HandleFunc("/listUsers", middlew.CheckBD(middlew.ValidateJWT(routers.ListUsers))).Methods("GET")
	// Leer publicaciones de mis seguidores
	router.HandleFunc("/readPostsFollowers", middlew.CheckBD(middlew.ValidateJWT(routers.ReadPostsFollowers))).Methods("GET")

	// INTELIGENCIAS ARTIFICIALES
	router.HandleFunc("/gpt", middlew.CheckBD(routers.ChatGPT)).Methods("POST")
	// Validar la Key de la API de Open AI
	router.HandleFunc("/validateAPIKey", middlew.CheckBD(routers.ValidateAPIKey)).Methods("POST")

	// EMPLEOS
	// Guardar publicacion del empleo
	router.HandleFunc("/savePublicationJob", middlew.CheckBD(middlew.ValidateJWT(routers.SavePublicationJob))).Methods("POST")

	// abrir el puerto
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// Creando un handler/importacion
	// Los cors son quienes otorgan los permisos
	handler := cors.AllowAll().Handler(router)        // Todos pueden acceder
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) // Escucha el puerto
}
