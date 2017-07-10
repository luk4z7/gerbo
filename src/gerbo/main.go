// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

package main

import (
	"gerbo/lib/logs"
	"gerbo/controllers/operation"
)

func main() {
	logs.Start()
	logs.INFO.Println("Running!")

	operation.Start()
}