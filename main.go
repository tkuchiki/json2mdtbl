package main

import (
	"log"
	"os"
)

func main() {
	converter := NewConverter(os.Stdin, os.Stdout)

	if err := converter.Read(); err != nil {
		log.Fatal(err)
	}

	converter.Write()
}
