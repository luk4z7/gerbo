// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

// server.go
package main

import (
	"github.com/codegangsta/negroni"
	"middleware-jwt/routers"
	"middleware-jwt/settings"
	"net/http"
	"github.com/rs/cors"
)

func main() {
	settings.Init()

	// Get all routes defined
	router := routers.InitRoutes()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"accept", "authorization", "content-type", "*"},
		AllowedMethods: []string{"PUT", "POST", "DELETE", "OPTIONS", "GET"},
	})
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)
	http.ListenAndServe(":6060", n)
}