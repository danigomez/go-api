package products

import (
	"fmt"
	"github.com/danigomez/go-api/database"
	"github.com/danigomez/go-api/database/model"
)


func CreateOne(data model.Product) (status bool, message string) {
	fmt.Println("Processing data => " + data.String())

	database.ProductsRepository.CreateOne(data)
	status = true
	message = "Created successfully!"
	return
}

func GetAll() []model.Product {
	return database.ProductsRepository.GetAll()
}

func GetOneById(id string) model.Product {
	product, err := database.ProductsRepository.GetOneById(id)
	if err != nil {
		panic(err)
	}
	return product
}

func DeleteOneById(id string) bool {
	fmt.Println("Processing data => " + id)

	return true
}
