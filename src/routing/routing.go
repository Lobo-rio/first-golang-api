package routing

import (
	"modules/src/routing/routers"

	"github.com/gorilla/mux"
)

func Generate() {
	r := mux.NewRouter()
	return routers.Configure(r)
}