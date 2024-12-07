package router

import (
	"devbook-api/pkg/router/routes"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Configurate(r)
}
