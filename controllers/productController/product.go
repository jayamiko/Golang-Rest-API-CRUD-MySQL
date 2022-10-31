package productcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api-go/helper"
	"rest-api-go/models"

	_ "github.com/go-sql-driver/mysql"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		fmt.Println("Success")
	}

	ResponseJson(w, http.StatusOK, products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, product)
}
