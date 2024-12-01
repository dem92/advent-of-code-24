package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(taskName string, useExample bool) []string {
	var filePath string

	if useExample {
		filePath = "/internal/day%s/example.txt"
	} else {
		filePath = "/internal/day%s/input.txt"
	}

	file, err := os.Open(fmt.Sprintf("."+filePath, taskName))

	if err != nil {
		// For debug mode
		file, err = os.Open(fmt.Sprintf(".."+filePath, taskName))

		if err != nil {
			log.Fatal(err)
		}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	fileContent := []string{}

	for scanner.Scan() {
		text := scanner.Text()
		fileContent = append(fileContent, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fileContent
}
