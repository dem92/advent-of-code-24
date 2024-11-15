package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Arg for task was not provided")
	}

	day := os.Args[1]
	log.Printf("Run requested for %s", day)

	switch day {
	default:
		log.Fatal("Failed to match task arg with function")
	}
}
