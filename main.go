package main

import (
	"gorm/handlers"
	"gorm/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Mux se usa para la creacion de rutas o endpoints
func main() {

	models.MigrarUser()

	//Rutas
	mux := mux.NewRouter()

	//Endpoints:
	// Obtener todos los users
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	// //Obtener user por Id.
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	// //Registrar un usuario
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	// //Editar
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	// //Eliminar
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))

}
