package day5

import "testing"

func TestPolymerLength(t *testing.T) {
	length := PolymerLength("dabAcCaCBAcCcaDA")
	if length != 10 {
		t.Errorf("Length was incorrect, got: %d, want: %d.", length, 10)
	}
}

func TestShortestPolymerWithoutOneUnitLength(t *testing.T) {
	length := ShortestPolymerWithoutOneUnitLength("dabAcCaCBAcCcaDA")
	if length != 4 {
		t.Errorf("Length was incorrect, got: %d, want: %d.", length, 4)
	}
}
