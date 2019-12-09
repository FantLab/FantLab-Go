package main

import (
	"fantlab/server"
	"flag"
	"log"
)

var gendocs = flag.Bool("gendocs", false, "")

func main() {
	flag.Parse()

	if *gendocs {
		server.GenerateDocs()
	} else {
		log.SetFlags(0)
		log.SetPrefix("$ ")

		server.Start()
	}
}
