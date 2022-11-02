package productcontroller

import (
	"encoding/json"
	"net/http"
	"rest-api-go/models"

	_ "github.com/go-sql-driver/mysql"
)

// var ResponseJson = helper.ResponseJson

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	
	var products []models.Product

	db := models.DB

	result, err := db.Query("SELECT * FROM products");
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
	 json.NewEncoder(w).Encode(products)
	//  ResponseJson(w, http.StatusOK, products)
}

// func CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	var product models.Product

// 	decoder := json.NewDecoder(r.Body)

// 	if err := decoder.Decode(&product); err != nil {
// 		ResponseError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	defer r.Body.Close()

// 	if err := models.DB.Query(&product).Error; err != nil {
// 		ResponseError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	ResponseJson(w, http.StatusCreated, product)
// }
