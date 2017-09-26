package products

import (
	"encoding/json"
	"fmt"
	"github.com/danigomez/go-api/common"
	"log"
	"net/http"
	"github.com/danigomez/go-api/database/model"
	"github.com/gorilla/mux"
)

func DoGetOne(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	productData := GetOneById(id)

	// Create Result data
	var result common.RequestResult
	result.StatusCode = "200"
	result.Data = productData

	json.NewEncoder(rw).Encode(&result)
}

func DoDeleteOne(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	DeleteOneById(id)
}

func DoGetAll(rw http.ResponseWriter, r *http.Request) {
	allData := GetAll()

	// Create Result data
	var result common.RequestResult
	result.StatusCode = "200"
	result.Data = allData

	json.NewEncoder(rw).Encode(&result)
}

func DoPost(rw http.ResponseWriter, r *http.Request) {
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	fmt.Println(product.Name)

	if err != nil {
		log.Fatal("There was an error while processing data " + err.Error())
	} else if status, message := CreateOne(product); status {
		var result common.RequestResult
		result.StatusCode = "200"
		result.Data = message

		json.NewEncoder(rw).Encode(&result)
	}
}
