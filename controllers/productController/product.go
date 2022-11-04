package productcontroller

import (
	"encoding/json"
	"net/http"
	"rest-api-go/helper"
	"rest-api-go/models"

	_ "github.com/go-sql-driver/mysql"
)

// var ResponseJson = helper.ResponseJson
var db = models.DB

var ResponseJson = helper.ResponseJson

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	var products []models.Product

	result, err := db.Query("SELECT * FROM products")
	if err != nil {
		db.Close()
		panic(err.Error())
	}

	for result.Next() {
		var data models.Product
		err := result.Scan(&data.Id, &data.Name, &data.Stock, &data.Price)
		if err != nil {
			db.Close()
			panic(err.Error())
		}

		products = append(products, data)
	}

	defer result.Close()
	//  json.NewEncoder(w).Encode(products)
	ResponseJson(w, http.StatusOK, products) // Response JSON
}

func GetDetailProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var data models.Product
	result, err := db.Query("SELECT * from products WHERE id = ? ", id)
	if err != nil {
		db.Close()
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&data.Id, &data.Name, &data.Stock, &data.Price)
		if err != nil {
			db.Close()
			panic(err.Error())
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(data)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	var id int

	result, err := db.Query("SELECT MAX(id) FROM products")
	if err != nil {
		db.Close()
		panic(err)
	}

	for result.Next() {
		err := result.Scan(&id)
		if err != nil {
			db.Close()
			panic(err)
		}
	}

	result.Close()

	productId := id + 1

	stmt, err := db.Prepare("INSERT INTO products(id, name, stock, price) VALUES (?,?,?,?)")

	if err != nil {
		db.Close()
		panic(err.Error())
	}

	name := r.Form.Get("name")
	stock := r.Form.Get("stock")
	price := r.Form.Get("price")

	_, err = stmt.Exec(productId, name, stock, price)

	if err != nil {
		db.Close()
		panic(err.Error())
	}

	stmt.Close()

	json.NewEncoder(w).Encode("Added Product Successfully")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE products SET name = ? WHERE id = ?")
	if err != nil {
		db.Close()
		panic(err.Error())
	}

	name := r.Form.Get("name")
	// stock := r.Form.Get("stock")
	// price := r.Form.Get("price")

	_, err = stmt.Exec(name, id)
	if err != nil {
		db.Close()
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		db.Close()
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		db.Close()
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}
