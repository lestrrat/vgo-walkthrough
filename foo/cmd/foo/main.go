package main

import (
	"log"
	"os"

	"github.com/lestrrat/vgo-walkthrough/foo"
)

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

func _main() error {
	return foo.Main()
}
