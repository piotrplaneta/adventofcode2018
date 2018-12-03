package day3

import "testing"

func TestOverlappingRectanglesArea(t *testing.T) {
	input := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	area := OverlappingRectanglesArea(input)
	if area != 4 {
		t.Errorf("Area was incorrect, got: %d, want: %d.", area, 4)
	}
}

func TestNonOverlappingRectangle(t *testing.T) {
	input := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	rectangleID := NonOverlappingRectangle(input)
	if rectangleID != "#3" {
		t.Errorf("Id was incorrect, got: %s, want: %s.", rectangleID, "#3")
	}
}
