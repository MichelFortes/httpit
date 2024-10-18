package main

import (
	"fmt"

	"github.com/MichelFortes/httpit/internal/cli"
)

func main() {

	scheme, err := cli.GetTestScheme()
	if err != nil {
		panic(err)
	}

	fmt.Print(scheme)

}
