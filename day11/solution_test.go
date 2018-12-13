package day11

import "testing"

func TestBiggestPowerCoordinate3x3Square(t *testing.T) {
	square := BiggestPowerCoordinate(42, 3)
	if square != (SquareWithValue{x: 21, y: 61, size: 3, value: 30}) {
		t.Errorf("Result was incorrect, got: (%d, %d, %d, %d), want: %s.", square.x, square.y, square.size, square.value, "21, 61, 2")
	}
}

func TestBiggestPowerCoordinateAnySquare(t *testing.T) {
	square := BiggestPowerCoordinate(42, 300)
	if square != (SquareWithValue{x: 232, y: 251, size: 12, value: 119}) {
		t.Errorf("Result was incorrect, got: (%d, %d, %d, %d), want: %s.", square.x, square.y, square.size, square.value, "232, 251, 12")
	}
}
