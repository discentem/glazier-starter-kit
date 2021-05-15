package main

import (
	"log"

	"github.com/discentem/glazier-config/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
