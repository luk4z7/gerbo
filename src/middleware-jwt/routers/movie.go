// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

// routers
package routers

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"middleware-jwt/controllers/movie"
)

func SetMovieRoutes(router *mux.Router) *mux.Router {
	// Quais os filmes com melhor avaliação média?
	router.Handle("/v1/movies/best/page/{page:[0-9]+}", negroni.New(
		negroni.HandlerFunc(movie.BestMovies),
	)).Methods("GET")

	// Quais os gêneros com melhor avaliação média?
	router.Handle("/v1/movies/genre/best", negroni.New(
		negroni.HandlerFunc(movie.ProcessRequest),
	)).Methods("GET")

	// Quais os gêneros com mais filmes?
	router.Handle("/v1/movies/genre/winner", negroni.New(
		negroni.HandlerFunc(movie.ProcessRequest),
	)).Methods("GET")

	// Qual a avaliação média por gênero?
	router.Handle("/v1/movies/rating/genre", negroni.New(
		negroni.HandlerFunc(movie.ProcessRequest),
	)).Methods("GET")

	// Qual a avaliação média por ano?
	router.Handle("/v1/movies/rating/year", negroni.New(
		negroni.HandlerFunc(movie.ProcessRequest),
	)).Methods("GET")

	// Qual a distribuição do número de filmes produzidos por ano?
	router.Handle("/v1/movies/distribution/year", negroni.New(
		negroni.HandlerFunc(movie.ProcessRequest),
	)).Methods("GET")

	// Qual a distribuição do número de filmes produzidos por década?
	router.Handle("/v1/movies/distribution/decade", negroni.New(
		negroni.HandlerFunc(movie.MoviesDistributionPerDecade),
	)).Methods("GET")

	return router
}