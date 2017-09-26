package database

import (
	"github.com/danigomez/go-api/database/model"
	"strconv"
	"errors"
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

func (p productTable) GetOneById(id string) (ret model.Product, err error) {

	for _, product := range p.data {
		if product.Id == id {
			ret = product
			err = nil
			return
		}
	}

	err = errors.New("the requested id doesn't exists")
	return
}


var ProductsRepository productTable




