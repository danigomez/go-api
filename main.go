package main

import (
	"github.com/danigomez/go-api/modules/products"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products", products.RootRoutes)
	router.HandleFunc("/products/{id}", products.IdRoutes)
	log.Fatal(http.ListenAndServe(":9001", router))
}
