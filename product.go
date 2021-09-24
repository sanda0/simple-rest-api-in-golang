package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/database"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func InitMigrate() {
	database.Connect()
	database.DB.AutoMigrate(&Product{})

}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product
	database.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	params := mux.Vars(r)
	database.DB.First(&product, params["id"])
	json.NewEncoder(w).Encode(product)
}
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("created")
	w.Header().Set("Content-Type", "application/json")
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	params := mux.Vars(r)
	database.DB.First(&product, params["id"])

	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Save(&product)

	json.NewEncoder(w).Encode(product)

}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	params := mux.Vars(r)
	database.DB.Delete(&product, params["id"])
	json.NewEncoder(w).Encode("product is deleted!")

}
