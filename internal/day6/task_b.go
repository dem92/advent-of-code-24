package day6

import (
	"aoc24/internal/utils"
	"fmt"
	"log"
)

func Day6b() {
	fileContent := utils.ReadFile("6", false)
	areaMap := [][]rune{}

	for _, line := range fileContent {
		characters := []rune(line)
		areaMap = append(areaMap, characters)
	}

	guardPosition := findGuard(areaMap)
	successfulObstacles := moveGuardAndAttemptLoop(areaMap, guardPosition, utils.DirNorth, 0)

	log.Printf("Total: %d", successfulObstacles)
}

func moveGuardAndAttemptLoop(areaMap [][]rune, guardPosition position, direction utils.Direction, successfulObstacles int) int {
	// utils.PrintCharMap(areaMap)

	nextX := guardPosition.X + direction.X
	nextY := guardPosition.Y + direction.Y

	if nextX < 0 || nextY < 0 || nextX > len(areaMap)-1 || nextY > len(areaMap[0])-1 {
		return successfulObstacles
	}

	nextSquare := areaMap[nextY][nextX]
	nextPosition := position{X: nextX, Y: nextY}

	switch nextSquare {
	case '^':
		return moveGuardAndAttemptLoop(areaMap, nextPosition, direction, successfulObstacles)
	case '#':
		return moveGuardAndAttemptLoop(areaMap, guardPosition, getNextDirection(direction), successfulObstacles)
	default:
		if checkForSuccessfulObstacle(areaMap, guardPosition, nextPosition, direction) {
			successfulObstacles++
		}

		areaMap[nextY][nextX] = '^'
		return moveGuardAndAttemptLoop(areaMap, nextPosition, direction, successfulObstacles)
	}
}

func checkForSuccessfulObstacle(areaMap [][]rune, guardPosition, obstaclePosition position, direction utils.Direction) bool {
	areaMap[obstaclePosition.Y][obstaclePosition.X] = '#'
	// utils.PrintCharMap(areaMap)
	squaresEntered := map[string][]utils.Direction{}
	isLoop := checkForLoop(areaMap, guardPosition, direction, squaresEntered)

	areaMap[obstaclePosition.Y][obstaclePosition.X] = '.'

	return isLoop
}

func checkForLoop(areaMap [][]rune, guardPosition position, direction utils.Direction, squaresEntered map[string][]utils.Direction) bool {
	if checkForPreviousVisit(guardPosition, direction, squaresEntered) {
		return true
	}

	nextX := guardPosition.X + direction.X
	nextY := guardPosition.Y + direction.Y

	if nextX < 0 || nextY < 0 || nextX > len(areaMap)-1 || nextY > len(areaMap[0])-1 {
		return false
	}

	nextSquare := areaMap[nextY][nextX]

	switch nextSquare {
	case '#':
		return checkForLoop(areaMap, guardPosition, getNextDirection(direction), squaresEntered)
	default:
		addVisit(guardPosition.Y, guardPosition.X, direction, squaresEntered)
		return checkForLoop(areaMap, position{X: nextX, Y: nextY}, direction, squaresEntered)
	}
}

func addVisit(y, x int, direction utils.Direction, squaresEntered map[string][]utils.Direction) {
	key := fmt.Sprintf("%d,%d", y, x)
	previousVisits := squaresEntered[key]
	previousVisits = append(previousVisits, direction)
	squaresEntered[key] = previousVisits
}

func checkForPreviousVisit(guardPosition position, direction utils.Direction, squaresEntered map[string][]utils.Direction) bool {
	key := fmt.Sprintf("%d,%d", guardPosition.Y, guardPosition.X)
	previousVisits := squaresEntered[key]

	for _, dir := range previousVisits {
		if dir == direction {
			return true
		}
	}

	return false
}
