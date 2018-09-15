package main

import (
	"fmt"
	"os"

	"github.com/rigado/edge-connect/examples/api/command"
)

func main() {

	r := command.InitCommands()

	if err := r.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
