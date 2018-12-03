package day3

import (
	"strconv"
	"strings"

	"github.com/piotrplaneta/adventofcode2018/utils"
)

type oneInchSquare struct {
	x0, x1, y0, y1 int
}
type rectangle struct {
	id             string
	x0, x1, y0, y1 int
}
type area map[oneInchSquare]int

//SolvePart1 returns the answer for part 1 of day 3
func SolvePart1() int {
	return OverlappingRectanglesArea(adventInput())
}

//SolvePart2 returns the answer for part 2 of day 3
func SolvePart2() string {
	return NonOverlappingRectangle(adventInput())
}

//OverlappingRectanglesArea returns area of overlapping rectangles
func OverlappingRectanglesArea(input []string) int {
	rectangles := parseRectangles(input)
	rectanglesOnArea := areaWithRectangles(rectangles)

	return squaresWithAtLeastNRectangles(rectanglesOnArea, 2)
}

//NonOverlappingRectangle returns id of nonoverlapping rectangle
func NonOverlappingRectangle(input []string) string {
	rectangles := parseRectangles(input)
	rectanglesOnArea := areaWithRectangles(rectangles)

	for _, rectangle := range rectangles {
		if onlyOneRectangle(rectangle, rectanglesOnArea) {
			return rectangle.id
		}
	}

	return ""
}

func parseRectangles(input []string) []rectangle {
	result := make([]rectangle, len(input))

	for i, line := range input {
		idAndCoordinates := strings.Split(line, "@")

		id := strings.Trim(idAndCoordinates[0], " ")
		baseAndSize := strings.Split(strings.Trim(idAndCoordinates[1], " "), ":")

		x0 := stringToInt(strings.Split(strings.Trim(baseAndSize[0], " "), ",")[0])
		y0 := stringToInt(strings.Split(strings.Trim(baseAndSize[0], " "), ",")[1])

		x1 := x0 + stringToInt(strings.Split(strings.Trim(baseAndSize[1], " "), "x")[0])
		y1 := y0 + stringToInt(strings.Split(strings.Trim(baseAndSize[1], " "), "x")[1])

		result[i] = rectangle{id: id, x0: x0, x1: x1, y0: y0, y1: y1}
	}

	return result
}

func areaWithRectangles(rectangles []rectangle) area {
	result := make(area)

	for _, rectangle := range rectangles {
		for x := rectangle.x0; x < rectangle.x1; x++ {
			for y := rectangle.y0; y < rectangle.y1; y++ {
				result[oneInchSquare{x0: x, x1: x + 1, y0: y, y1: y + 1}]++
			}
		}
	}

	return result
}

func squaresWithAtLeastNRectangles(rectanglesOnArea area, atLeast int) int {
	result := 0

	for _, v := range rectanglesOnArea {
		if v >= atLeast {
			result++
		}
	}

	return result
}

func onlyOneRectangle(rectangle rectangle, rectanglesOnArea area) bool {
	for x := rectangle.x0; x < rectangle.x1; x++ {
		for y := rectangle.y0; y < rectangle.y1; y++ {
			if rectanglesOnArea[oneInchSquare{x0: x, x1: x + 1, y0: y, y1: y + 1}] > 1 {
				return false
			}
		}
	}

	return true
}

func stringToInt(input string) int {
	result, err := strconv.Atoi(input)

	if err != nil {
		panic(err)
	}

	return result
}

func adventInput() []string {
	lines, _ := utils.ReadLinesFromFile("day3/input")
	return lines
}
