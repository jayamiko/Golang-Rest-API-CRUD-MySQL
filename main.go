package main

import (
	"log"
	"net/http"
	productcontroller "rest-api-go/controllers/productController"
	"rest-api-go/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/products", productcontroller.GetAllProducts).Methods("GET")
	// r.HandleFunc("/product/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/product", productcontroller.AddProduct).Methods("POST")
	// r.HandleFunc("/product/{id}", productcontroller.Update).Methods("PUT")
	// r.HandleFunc("/product", productcontroller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
