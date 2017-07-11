// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

// routers
package routers

import (
	"github.com/gorilla/mux"
)

// Initialize the routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetMoviesRoutes(router)
	return router
}