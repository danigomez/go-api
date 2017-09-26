package products

import (
	"encoding/json"
	"fmt"
	"github.com/danigomez/go-api/common"
	"log"
	"net/http"
)

func DoGetOne(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	GetOneById(id)
}

func DoPost(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var product Product
	err := decoder.Decode(&product)

	fmt.Println(product.Name)

	if err != nil {
		log.Fatal("There was an error while processing data " + err.Error())
	} else if status, message := CreateOne(product); status {
		var result common.RequestResult
		result.StatusCode = "200"
		result.Data = message

		encoder := json.NewEncoder(rw)
		encoder.Encode(&result)
	}
}
