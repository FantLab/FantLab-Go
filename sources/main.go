package main

import (
	"fantlab/server"
	"flag"
)

var gendocs = flag.Bool("gendocs", false, "")

func main() {
	flag.Parse()

	if *gendocs {
		server.GenerateDocs()
	} else {
		server.Start()
	}
}
