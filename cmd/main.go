package main

import (
	"aoc24/internal/day1"
	"aoc24/internal/day2"
	"aoc24/internal/day3"
	"aoc24/internal/day4"
	"aoc24/internal/day5"
	"aoc24/internal/day6"
	"aoc24/internal/day7"
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
	case "3b":
		day3.Day3b()
	case "4a":
		day4.Day4a()
	case "4b":
		day4.Day4b()
	case "5a":
		day5.Day5a()
	case "5b":
		day5.Day5b()
	case "6a":
		day6.Day6a()
	case "6b":
		day6.Day6b()
	case "7a":
		day7.Day7a()
	case "7b":
		day7.Day7b()
	default:
		log.Fatal("Failed to match task arg with function")
	}
}
