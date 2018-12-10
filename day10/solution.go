package day10

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/piotrplaneta/adventofcode2018/utils"
)

//MovingPoint represents moving point
type MovingPoint struct {
	x, y, xVelocity, yVelocity int
}

//PotentialSolution represents potential solution
type PotentialSolution struct {
	solution    []MovingPoint
	elapsedTime int
}

//SolvePart1 returns the answer for part 1 of day 10
func SolvePart1() {
	SavePossibleMessagesAsImages(adventInput(), 15)
}

//SolvePart2 returns the answer for part 2 of day 10
func SolvePart2() int {
	return PossibleMessages(adventInput(), 15)[0].elapsedTime
}

//SavePossibleMessagesAsImages saves all possible messages as images
func SavePossibleMessagesAsImages(input []string, heightLimit int) {
	for i, v := range PossibleMessages(input, heightLimit) {
		saveAsImage(v.solution, strconv.Itoa(i))
	}
}

//PossibleMessages returns all points configurations constained within a given height limit
func PossibleMessages(input []string, heightLimit int) []PotentialSolution {
	possibleMessages := make([]PotentialSolution, 0)

	movingPoints := parseInput(input)
	minY, maxY := minMaxYCoordinates(movingPoints)

	for i := 1; i <= abs(minY)+maxY; i++ {
		movingPoints = tick(movingPoints)
		minY, maxY := minMaxYCoordinates(movingPoints)

		if maxY-minY <= heightLimit {
			possibleMessages = append(possibleMessages, PotentialSolution{solution: movingPoints, elapsedTime: i})
		}
	}

	return possibleMessages
}

func saveAsImage(points []MovingPoint, filename string) {
	minX, maxX := minMaxXCoordinates(points)
	minY, maxY := minMaxYCoordinates(points)
	imgRect := image.Rect(minX, minY, maxX, maxY)
	img := image.NewNRGBA(imgRect)

	for _, p := range points {
		img.Set(p.x, p.y, color.NRGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		})
	}

	fDst, err := os.Create(fmt.Sprintf("day10/%s.png", filename))
	if err != nil {
		panic(err)
	}
	defer fDst.Close()
	err = png.Encode(fDst, img)
	if err != nil {
		panic(err)
	}
}

func minMaxYCoordinates(points []MovingPoint) (int, int) {
	minY := math.MaxInt32
	maxY := math.MinInt32

	for _, p := range points {
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	return minY, maxY
}

func minMaxXCoordinates(points []MovingPoint) (int, int) {
	minX := math.MaxInt32
	maxX := math.MinInt32

	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
	}

	return minX, maxX
}

func tick(points []MovingPoint) []MovingPoint {
	pointsAfterTick := make([]MovingPoint, len(points))

	for i, p := range points {
		pointAfterTick := MovingPoint{
			x:         p.x + p.xVelocity,
			y:         p.y + p.yVelocity,
			xVelocity: p.xVelocity,
			yVelocity: p.yVelocity,
		}

		pointsAfterTick[i] = pointAfterTick
	}

	return pointsAfterTick
}

func parseInput(input []string) []MovingPoint {
	points := make([]MovingPoint, len(input))

	for i, s := range input {
		position := strings.Split(strings.Split(s, ">")[0], "<")[1]
		velocity := strings.Split(s, "velocity=<")[1]
		velocity = velocity[:len(velocity)-1]

		point := MovingPoint{
			x:         stringToInt(strings.Split(position, ",")[0]),
			y:         stringToInt(strings.Split(position, ",")[1]),
			xVelocity: stringToInt(strings.Split(velocity, ",")[0]),
			yVelocity: stringToInt(strings.Split(velocity, ",")[1]),
		}

		points[i] = point
	}

	return points
}

func stringToInt(input string) int {
	result, err := strconv.Atoi(strings.Trim(input, " "))

	if err != nil {
		panic(err)
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func adventInput() []string {
	lines, _ := utils.ReadLinesFromFile("day10/input")
	return lines
}
