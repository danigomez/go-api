package products

import (
	"net/http"
)

func RootRoutes(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		DoGetAll(rw, r)
		break
	case "POST":
		DoPost(rw, r)
		break
	}
}

func IdRoutes(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		DoGetOne(rw, r)
		break
	case "DELETE":
		DoDeleteOne(rw, r)
		break
	}
}
