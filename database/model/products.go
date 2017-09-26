package model

import (
	"bytes"
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