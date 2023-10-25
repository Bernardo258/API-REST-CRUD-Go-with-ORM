package main

import (
	"Practica_3/handlers"
	"fmt"
	"net/http"
)

func main() {
	router := handlers.InitializeRoutes()
	http.Handle("/", router)
	fmt.Println("Servidor en ejecución en el puerto 8000")
	http.ListenAndServe(":8000", nil)
}
