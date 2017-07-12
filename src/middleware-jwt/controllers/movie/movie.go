// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

package movie

import (
	"net/http"
	liberr "gerbo/lib/error"
	"middleware-jwt/core/response"
	"middleware-jwt/services/movie"
	"github.com/gorilla/mux"
	"strconv"
	"gopkg.in/mgo.v2/bson"
)

// Test HTTP benchmarking
// see more in https://github.com/wg/wrk
// wrk -t12 -c200 -d30s http://127.0.0.1:6060/v1/movies/best/page/10
func BestMovies(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var responseStatus = http.StatusOK
	var errors = []liberr.Errors{}
	var result = []string{}

	vars := mux.Vars(r)
	page, _ := strconv.ParseInt(vars["page"], 0, 64)

	data, err := movie.GetBestMovies(page)
	envelope := response.Envelope(liberr.ErrorsAPI{
		Errors: errors,
		Url:    string(r.URL.Path),
		Method: string(r.Method),
	}, responseStatus, data)
	envelope.Pagination.Page = page

	if err != nil {
		errors = []liberr.Errors{
			0: {
				Message: "An unexpected error occurred while searching movies",
				Type:    "internal_error",
			},
		}
		responseStatus = 500
		envelope = response.Envelope(liberr.ErrorsAPI{
			Errors: errors,
			Url:    string(r.URL.Path),
			Method: string(r.Method),

		}, responseStatus, result)
		envelope.Pagination.Page = page
	}
	response.Header(w, responseStatus, envelope, response.Headers{})
}

func ProcessRequest(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ProcessResponse(w, r, next)
}

func ProcessResponse(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var responseStatus = http.StatusOK
	var errors = []liberr.Errors{}
	var result = []string{}
	var err error = nil
	var data = []bson.M{}

	switch r.URL.Path {
		case "/v1/movies/genre/best":
			data, err = movie.GetBestMoviesGenres()
		case "/v1/movies/genre/winner":
			data, err = movie.GetMoreMoviesGenre()
		case "/v1/movies/rating/genre":
			data, err = movie.GetMoviesRatingPerGenre()
		case "/v1/movies/rating/year":
			data, err = movie.GetMoviesRatingPerYear()
		case "/v1/movies/distribution/year":
			data, err = movie.GetMoviesDistributionPerYear()
	}

	envelope := response.Envelope(liberr.ErrorsAPI{
		Errors: errors,
		Url:    string(r.URL.Path),
		Method: string(r.Method),
	}, responseStatus, data)
	if err != nil {
		errors = []liberr.Errors{
			0: {
				Message: "An unexpected error occurred while searching movies",
				Type:    "internal_error",
			},
		}
		responseStatus = 500
		envelope = response.Envelope(liberr.ErrorsAPI{
			Errors: errors,
			Url:    string(r.URL.Path),
			Method: string(r.Method),

		}, responseStatus, result)
	}
	response.Header(w, responseStatus, envelope, response.Headers{})
}

func MoviesDistributionPerDecade(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var responseStatus = http.StatusOK
	var errors = []liberr.Errors{}
	var result = []string{}

	data, err := movie.GetMoviesDistributionPerDecade()
	envelope := response.Envelope(liberr.ErrorsAPI{
		Errors: errors,
		Url:    string(r.URL.Path),
		Method: string(r.Method),
	}, responseStatus, data)
	if err != nil {
		errors = []liberr.Errors{
			0: {
				Message: "An unexpected error occurred while searching movies",
				Type:    "internal_error",
			},
		}
		responseStatus = 500
		envelope = response.Envelope(liberr.ErrorsAPI{
			Errors: errors,
			Url:    string(r.URL.Path),
			Method: string(r.Method),

		}, responseStatus, result)
	}
	response.Header(w, responseStatus, envelope, response.Headers{})
}