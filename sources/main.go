package main

import (
	"fantlab/apiserver"
	"flag"
	"fmt"
	"os"

	"github.com/FantLab/go-kit/crypto/signed"
)

var gendocs = flag.Bool("gendocs", false, "")
var genkeys = flag.Bool("genkeys", false, "")

func main() {
	flag.Parse()

	if *gendocs {
		apiserver.GenerateDocs()
		return
	}

	if *genkeys {
		coder, _ := signed.Generate()
		fmt.Fprintln(os.Stdout, "Public key:", coder.PublicKey())
		fmt.Fprintln(os.Stdout, "Private key:", coder.PrivateKey())
		return
	}

	apiserver.Start()
}
