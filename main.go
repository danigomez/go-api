package main

import (
	"github.com/danigomez/go-api/modules/products"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", products.Routes)

	log.Fatal(http.ListenAndServe(":9001", mux))
}
