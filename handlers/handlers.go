package handlers

import (
	model "Practica_3/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "User:Password@tcp(localhost:3306)/tienda?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error al conectar a la base de datos")
	}

	db.AutoMigrate(&model.Products{})
}

// Función para obtener todos los productos
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []model.Products
	db.Find(&products)
	respondWithJSON(w, http.StatusOK, products)
}

// Función para obtener un producto por ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product model.Products
	db.First(&product, params["id"])
	if product.ID == 0 {
		respondWithError(w, http.StatusNotFound, "Producto no encontrado")
		return
	}
	respondWithJSON(w, http.StatusOK, product)
}

// Función para crear un nuevo producto
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Products
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Datos de entrada no válidos")
		return
	}
	db.Create(&product)
	respondWithJSON(w, http.StatusCreated, product)
}

// Función para actualizar un producto
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product model.Products
	db.First(&product, params["id"])
	if product.ID == 0 {
		respondWithError(w, http.StatusNotFound, "Producto no encontrado")
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Datos de entrada no válidos")
		return
	}
	db.Save(&product)
	respondWithJSON(w, http.StatusOK, product)
}

// Función para eliminar un producto
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product model.Products
	db.First(&product, params["id"])
	if product.ID == 0 {
		respondWithError(w, http.StatusNotFound, "Producto no encontrado")
		return
	}
	db.Delete(&product)
	respondWithJSON(w, http.StatusNoContent, nil)
}

// Función para responder con JSON y código de estado HTTP
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Función para responder con un error y código de estado HTTP
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
