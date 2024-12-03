package main

import (
	"aoc24/internal/day1"
	"aoc24/internal/day2"
	"aoc24/internal/day3"
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
	case "2a":
		day2.Day2a()
	case "2b":
		day2.Day2b()
	case "3a":
		day3.Day3a()
	default:
		log.Fatal("Failed to match task arg with function")
	}
}
