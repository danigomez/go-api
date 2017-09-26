package products

import (
	"bytes"
	"fmt"
	"strconv"
)

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       int    `json:"value"`
}

func (p Product) String() string {
	var buf bytes.Buffer

	buf.WriteString("[id]: " + p.Id + "\n")
	buf.WriteString("[Name]: " + p.Name + "\n")
	buf.WriteString("[Description]: " + p.Description + "\n")
	buf.WriteString("[Value]: " + strconv.Itoa(p.Value) + "\n")

	return buf.String()
}

func CreateOne(data Product) (status bool, message string) {
	fmt.Println("Processing data => " + data.String())
	status = true
	message = "This is a test"
	return
}

func GetOneById(id string) Product {
	fmt.Println("Processing data => " + id)

	var product Product

	return product
}

func DeleteOneById(id string) bool {
	fmt.Println("Processing data => " + id)

	return true
}
