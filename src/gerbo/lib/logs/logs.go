// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

// lib
package logs

import (
	"github.com/agtorre/gocolorize"
	"log"
	"os"
)

var (
	INFO     *log.Logger
	WARNING  *log.Logger
	CRITICAL *log.Logger
)

func Start() {
	info := gocolorize.NewColor("green")
	warning := gocolorize.NewColor("yellow")
	critical := gocolorize.NewColor("back+u:red")

	//helper functions to shorten code
	i := info.Paint
	w := warning.Paint
	c := critical.Paint

	INFO = log.New(os.Stdout, i("INFO "), log.Ldate|log.Lmicroseconds|log.Lshortfile)
	WARNING = log.New(os.Stdout, w("WARNING "), log.Ldate|log.Lmicroseconds|log.Lshortfile)
	CRITICAL = log.New(os.Stdout, c("CRITICAL "), log.Ldate|log.Lmicroseconds|log.Lshortfile)
}
