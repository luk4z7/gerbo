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
				Message: "An unexpected error occurred while searching the best movies",
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

func BestMoviesGenres(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}

func MoreMoviesGenre(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}

func MoviesRatingPerGenre(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}

func MoviesRatingPerYear(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}

func MoviesDistributionPerYear(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}

func MoviesDistributionPerDecade(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}