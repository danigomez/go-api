package products

import (
	"net/http"
)

func Routes(rw http.ResponseWriter, r *http.Request) {

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
