// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

// routers
package routers

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"middleware-jwt/controllers/movies"
)

func SetMoviesRoutes(router *mux.Router) *mux.Router {
	router.Handle("/v1/movies/best", negroni.New(
		negroni.HandlerFunc(movies.Get),
	)).Methods("GET")

	return router
}