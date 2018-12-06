package day6

import (
	"math"
	"strconv"
	"strings"

	"github.com/piotrplaneta/adventofcode2018/utils"
)

type point struct {
	x, y int
}

type location struct {
	x, y int
}

type area struct {
	minX, maxX, minY, maxY int
}

//SolvePart1 returns the answer for part 1 of day 6
func SolvePart1() int {
	return LargestEmptyAreaSize(adventInput())
}

//SolvePart2 returns the answer for part 2 of day 6
func SolvePart2() int {
	return SizeOfRegionCloseToAllLocations(adventInput(), 10000)
}

//LargestEmptyAreaSize return the size of area with points cloesest to singular location
func LargestEmptyAreaSize(lines []string) int {
	locations := linesToLocations(lines)
	surface := calculateSurface(locations)

	closestLocationSizes := calculateClosestLocationSizes(surface, locations)
	locationsWithFiniteClosestPoints := generateClosedLocations(locations)
	validLocationSizes := selectValidLocationSizes(closestLocationSizes, locationsWithFiniteClosestPoints)
	return findLargestEmptyAreaSize(validLocationSizes)
}

//SizeOfRegionCloseToAllLocations returns the size of region which is close to all points
func SizeOfRegionCloseToAllLocations(lines []string, maxDistance int) int {
	locations := linesToLocations(lines)
	surface := calculateSurface(locations)
	sumOfdistancesToLocationsFromEachSurfacePoint := calcuteDistanceSums(surface, locations)

	areaSize := 0
	for _, v := range sumOfdistancesToLocationsFromEachSurfacePoint {
		if v < maxDistance {
			areaSize++
		}
	}

	return areaSize
}

func linesToLocations(lines []string) []location {
	locations := make([]location, len(lines))

	for i, v := range lines {
		x := stringToInt(strings.Trim(strings.Split(v, ",")[0], " "))
		y := stringToInt(strings.Trim(strings.Split(v, ",")[1], " "))

		locations[i] = location{x: x, y: y}
	}

	return locations
}

func calculateSurface(locations []location) area {
	minX := math.MaxInt32
	maxX := 0
	minY := math.MaxInt32
	maxY := 0

	for _, l := range locations {
		if l.x < minX {
			minX = l.x
		}
		if l.x > maxX {
			maxX = l.x
		}
		if l.y < minY {
			minY = l.y
		}
		if l.y > maxY {
			maxY = l.y
		}
	}

	return area{minX: minX, minY: minY, maxX: maxX, maxY: maxY}
}

func calculateClosestLocationSizes(surface area, locations []location) map[location]int {
	pointsClosestToLocation := make(map[location]int)

	for i := surface.minX; i <= surface.maxX; i++ {
		for j := surface.minY; j <= surface.maxY; j++ {
			closestLocations := findClosestLocations(i, j, locations)
			if len(closestLocations) == 1 {
				pointsClosestToLocation[closestLocations[0]]++
			}
		}
	}

	return pointsClosestToLocation
}

func findClosestLocations(x, y int, locations []location) []location {
	closestLocations := make([]location, 0)

	for _, l := range locations {
		if len(closestLocations) == 0 {
			closestLocations = append(closestLocations, l)
		} else if distance(x, y, l.x, l.y) < distance(x, y, closestLocations[0].x, closestLocations[0].y) {
			closestLocations = []location{l}
		} else if distance(x, y, l.x, l.y) == distance(x, y, closestLocations[0].x, closestLocations[0].y) {
			closestLocations = append(closestLocations, l)
		}
	}

	return closestLocations
}

func generateClosedLocations(locations []location) []location {
	closedLocations := make([]location, 0)

	for _, checkedLocation := range locations {
		isClosed := []bool{false, false, false, false}

		for _, l := range locations {
			if l != checkedLocation && l.x-checkedLocation.x >= abs(l.y-checkedLocation.y) {
				isClosed[0] = true
			}
			if l != checkedLocation && checkedLocation.x-l.x >= abs(l.y-checkedLocation.y) {
				isClosed[1] = true
			}
			if l != checkedLocation && l.y-checkedLocation.y >= abs(l.x-checkedLocation.x) {
				isClosed[2] = true
			}
			if l != checkedLocation && checkedLocation.y-l.y >= abs(l.x-checkedLocation.x) {
				isClosed[3] = true
			}
		}

		if allTrue(isClosed) {
			closedLocations = append(closedLocations, checkedLocation)
		}
	}

	return closedLocations
}

func selectValidLocationSizes(closestLocationSizes map[location]int, validLocations []location) map[location]int {
	validLocationSizes := make(map[location]int)

	for k, v := range closestLocationSizes {
		if contains(validLocations, k) {
			validLocationSizes[k] = v
		}
	}

	return validLocationSizes
}

func findLargestEmptyAreaSize(validLocationSizes map[location]int) int {
	max := 0

	for _, v := range validLocationSizes {
		if v > max {
			max = v
		}
	}

	return max
}

func calcuteDistanceSums(surface area, locations []location) map[point]int {
	distanceSums := make(map[point]int)

	for i := surface.minX; i <= surface.maxX; i++ {
		for j := surface.minY; j <= surface.maxY; j++ {
			for _, l := range locations {
				distanceSums[point{x: i, y: j}] += distance(i, j, l.x, l.y)
			}
		}
	}

	return distanceSums
}

func distance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func contains(a []location, l location) bool {
	for _, v := range a {
		if v == l {
			return true
		}
	}

	return false
}

func allTrue(slice []bool) bool {
	allTrue := true

	for _, v := range slice {
		allTrue = allTrue && v
	}

	return allTrue
}

func stringToInt(input string) int {
	result, err := strconv.Atoi(input)

	if err != nil {
		panic(err)
	}

	return result
}

func adventInput() []string {
	lines, _ := utils.ReadLinesFromFile("day6/input")
	return lines
}
