package database

import (
	model "Practica_3/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "Bernardo2:BradoDSNS23@tcp(localhost:3306)/your-database-name?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error al conectar a la base de datos")
	}
	defer db.Close()

	db.AutoMigrate(&model.Products{})
}
