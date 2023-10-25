package handlers

import (
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// Rutas para el CRUD
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")

	return router
}
