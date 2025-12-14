package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MichelFortes/httpit/internal/config"
	"github.com/MichelFortes/httpit/internal/runner"
)

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: httpit <config-file>")
		os.Exit(1)
	}

	scheme, err := config.GetTestScheme(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading scheme: %v\n", err)
		os.Exit(1)
	}

	runner := runner.NewRunner()
	if err := runner.Run(scheme); err != nil {
		panic(err)
	}
}
