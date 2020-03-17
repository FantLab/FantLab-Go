package main

import (
	"fantlab/base/edsign"
	"fantlab/server"
	"flag"
	"fmt"
	"os"
)

var gendocs = flag.Bool("gendocs", false, "")
var genkeys = flag.Bool("genkeys", false, "")

func main() {
	flag.Parse()

	if *gendocs {
		server.GenerateDocs()
		return
	}

	if *genkeys {
		pub, priv, _ := edsign.GenerateNewKeyPair()
		fmt.Fprintln(os.Stdout, "Public key:", pub)
		fmt.Fprintln(os.Stdout, "Private key:", priv)
		return
	}

	server.Start()
}
