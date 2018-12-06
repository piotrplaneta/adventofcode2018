package day6

import "testing"

func TestLargestEmptyAreaSize(t *testing.T) {
	size := LargestEmptyAreaSize([]string{"1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"})
	if size != 17 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", size, 17)
	}
}

func TestSizeOfRegionCloseToAllLocations(t *testing.T) {
	size := SizeOfRegionCloseToAllLocations([]string{"1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"}, 32)
	if size != 16 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", size, 16)
	}
}
