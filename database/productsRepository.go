package database

import (
	"github.com/danigomez/go-api/database/model"
	"strconv"
)

type productTable struct {
	data []model.Product
}


func (p *productTable) CreateOne(product model.Product) (status bool, err error) {

	// Assign a new id according the data currently loaded
	product.Id = strconv.Itoa(len(p.data))

	p.data = append(p.data, product)

	status = true
	err = nil
	return
}

func (p productTable) GetAll() []model.Product {
	return p.data
}


var ProductsTable productTable




