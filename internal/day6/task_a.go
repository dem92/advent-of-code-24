package day6

import (
	"aoc24/internal/utils"
	"log"
)

type position struct {
	X int
	Y int
}

func Day6a() {
	fileContent := utils.ReadFile("6", false)
	areaMap := [][]rune{}

	for _, line := range fileContent {
		characters := []rune(line)
		areaMap = append(areaMap, characters)
	}

	guardPosition := findGuard(areaMap)
	squaresEntered := moveGuard(areaMap, guardPosition, utils.DirNorth, 1)

	log.Printf("Total: %d", squaresEntered)
}

func findGuard(areaMap [][]rune) position {
	for y, column := range areaMap {
		for x, square := range column {
			if square == '^' {
				return position{X: x, Y: y}
			}
		}
	}

	panic("Guard not found! D:")
}

func moveGuard(areaMap [][]rune, guardPosition position, direction utils.Direction, squaresEntered int) int {
	// utils.PrintCharMap(areaMap)

	nextX := guardPosition.X + direction.X
	nextY := guardPosition.Y + direction.Y

	if nextX < 0 || nextY < 0 || nextX > len(areaMap)-1 || nextY > len(areaMap[0])-1 {
		return squaresEntered
	}

	nextSquare := areaMap[nextY][nextX]

	switch nextSquare {
	case '^':
		return moveGuard(areaMap, position{X: nextX, Y: nextY}, direction, squaresEntered)
	case '#':
		return moveGuard(areaMap, guardPosition, getNextDirection(direction), squaresEntered)
	default:
		areaMap[nextY][nextX] = '^'
		return moveGuard(areaMap, position{X: nextX, Y: nextY}, direction, squaresEntered+1)
	}
}

func getNextDirection(direction utils.Direction) utils.Direction {
	switch direction.Dir {
	case "N":
		return utils.DirEast
	case "E":
		return utils.DirSouth
	case "S":
		return utils.DirWest
	default:
		return utils.DirNorth
	}
}
