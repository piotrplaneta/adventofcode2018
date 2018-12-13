package day11

import (
	"strconv"
	"strings"
)

//SquareWithValue represents 2d square with value
type SquareWithValue struct {
	x, y  int
	size  int
	value int
}

const gridSize = 300

//SolvePart1 returns the answer for part 1 of day 11
func SolvePart1() SquareWithValue {
	return BiggestPowerCoordinate(9424, 3)
}

//SolvePart2 returns the answer for part 2 of day 11
func SolvePart2() SquareWithValue {
	return BiggestPowerCoordinate(9424, gridSize)
}

//BiggestPowerCoordinate returns top left corner of square with given max size with the biggest power value
func BiggestPowerCoordinate(serialNumber, maxSquareSize int) SquareWithValue {
	pointValues := generatePointValues(serialNumber)
	prefixSums := generatePrefixSums(pointValues)

	maxSquare := SquareWithValue{-1, -1, -1, -1}

	for squareSize := 1; squareSize <= maxSquareSize; squareSize++ {
		for i := -1; i < gridSize-squareSize; i++ {
			for j := -1; j < gridSize-squareSize; j++ {
				sum := calculateSum(i, j, squareSize, prefixSums)

				if sum > maxSquare.value {
					maxSquare = SquareWithValue{i + 2, j + 2, squareSize, sum}
				}
			}
		}
	}

	return maxSquare
}

func generatePointValues(serialNumber int) [][]int {
	pointsWithValues := make([][]int, gridSize)

	for i := 0; i < gridSize; i++ {
		pointsWithValues[i] = make([]int, gridSize)
		for j := 0; j < gridSize; j++ {
			pointsWithValues[i][j] = calculatePower(i, j, serialNumber)
		}
	}

	return pointsWithValues
}

func generatePrefixSums(pointValues [][]int) [][]int {
	prefixSums := make([][]int, gridSize)
	prefixSums[0] = make([]int, gridSize)
	copy(prefixSums[0], pointValues[0])

	for j := 1; j < gridSize; j++ {
		prefixSums[0][j] += prefixSums[0][j-1]
	}

	for i := 1; i < gridSize; i++ {
		prefixSums[i] = make([]int, gridSize)
		copy(prefixSums[i], pointValues[i])

		prefixSums[i][0] += prefixSums[i-1][0]

		for j := 1; j < gridSize; j++ {
			prefixSums[i][j] += (prefixSums[i-1][j] + prefixSums[i][j-1] - prefixSums[i-1][j-1])
		}
	}

	return prefixSums
}

func calculatePower(x, y, serialNumber int) int {
	pointPower := strconv.Itoa(((x+11)*(y+1) + serialNumber) * (x + 11))

	if len(pointPower) > 2 {
		thirdDigit := stringToInt(pointPower[len(pointPower)-3 : len(pointPower)-2])
		return thirdDigit - 5
	}

	return 0
}

func calculateSum(i, j, squareSize int, prefixSums [][]int) int {
	if i > 0 && j > 0 {
		return prefixSums[i+squareSize][j+squareSize] - prefixSums[i][j+squareSize] - prefixSums[i+squareSize][j] + prefixSums[i][j]
	} else if i > 0 {
		return prefixSums[i+squareSize][j+squareSize] - prefixSums[i][j+squareSize]
	} else if j > 0 {
		return prefixSums[i+squareSize][j+squareSize] - prefixSums[i+squareSize][j]
	} else {
		return prefixSums[i+squareSize][j+squareSize]
	}
}

func stringToInt(input string) int {
	result, err := strconv.Atoi(strings.Trim(input, " "))

	if err != nil {
		panic(err)
	}

	return result
}
