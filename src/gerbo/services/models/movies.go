// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

package models

type MovieRequest struct {
	ID            int    `json:"id"`
	Filme         string `json:"filme"`
	Ano           int    `json:"ano"`
	GeneroID      int    `json:"genero_id"`
	Genero        string `json:"genero"`
	AvaliacaoID   int    `json:"avaliacao_id"`
	Avaliacao     string `json:"avaliacao"`
	AvaliacaoNota int    `json:"avaliacao_nota"`
	TwitterID     int    `json:"twitter_id"`
}

type MoviesResponse struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Genre  []Genre  `json:"genre"`
	Rating []Rating `json:"rating"`
}