package main

import (
	"fantlab/api/docs"
	"flag"
	"os"
)

var gendocs = flag.Bool("gendocs", false, "")

func main() {
	_ = docs.Generate(os.Stdout)
	// flag.Parse()

	// if *gendocs {
	// 	_ = docs.Generate(os.Stdout)
	// } else {
	// 	log.SetFlags(0)
	// 	log.SetPrefix("$ ")

	// 	startServer()
	// }
}
