package main

import (
	"aoc24/internal/day1"
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
	case "1a":
		day1.Day1a()
	case "1b":
		day1.Day1b()
	default:
		log.Fatal("Failed to match task arg with function")
	}
}
