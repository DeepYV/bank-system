package delivery

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetId(r *http.Request) string {

	vars := mux.Vars(r)
	id := vars["id"]

	return id

}

