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
		keyPair, _ := edsign.GenerateNewKeyPair()
		fmt.Fprintln(os.Stdout, "Public key:", keyPair.PublicKey)
		fmt.Fprintln(os.Stdout, "Private key:", keyPair.PrivateKey)
		return
	}

	server.Start()
}
