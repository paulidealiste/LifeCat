package main

import (
	"flag"
)

func main() {

	var t1 string
	var t2 string

	flag.StringVar(&t1, "taxa1", "Rupicapra", "string")
	flag.StringVar(&t2, "taxa2", "Rupicapra", "string")

	flag.Parse()

	// todex := catalogueoflife.ReadAndUnmarsh(t1, t2)
	// catalogueoflife.PrintTaxon(&todex)
}
