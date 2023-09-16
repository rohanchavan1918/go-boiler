package main

import (
	"log"

	"github.com/rohanchavan1918/go-boiler/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
