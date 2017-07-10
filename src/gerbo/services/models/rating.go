// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

package models

type Rating struct {
	ID      int    `json:"id"`
	Feeling string `json:"feeling"`
	Score   int    `json:"score"`
	User    User   `json:"user"`
}